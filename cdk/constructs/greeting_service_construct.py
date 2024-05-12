from aws_cdk import (
    Duration,
    aws_ec2 as ec2,
    aws_ecs as ecs,
    aws_ecs_patterns as ecsp,
)
from constructs import Construct
VPC_ID = "vpc-05a95664281543ea3"
# VPC_ID = "vpc-0ddeab84c11dcca89"

class GreetingServiceConstruct(Construct):
    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        vpc = self._prepare_vpc()
        private_subnet = self._prepare_private_subnet(vpc=vpc)
        self._prepare_vpc_endpoints(vpc=vpc, private_subnet=private_subnet)
        cluster = self._prepare_ecs_cluster(vpc=vpc)
        task_definition = self._prepare_task_definition()
        service = ecsp.ApplicationLoadBalancedFargateService(self, "GreetingService",
            cluster=cluster,
            memory_limit_mib=2048,
            desired_count=1,
            cpu=1024,
            task_subnets=ec2.SubnetSelection(
                subnets=[private_subnet]
            ),
            load_balancer_name="application-lb-name",
            task_definition=task_definition,
        )
        self.service = service

    def _prepare_vpc(self) -> ec2.IVpc:
        """
        Prepares VPC with endpoints for private connection from Fargate
        """
        vpc = ec2.Vpc.from_lookup(self, "vpc", vpc_id=VPC_ID)
        return vpc

    def _prepare_vpc_endpoints(self, vpc: ec2.IVpc, private_subnet: ec2.PrivateSubnet):
        # access ECR
        ec2.InterfaceVpcEndpoint(self, 'ECRVpcEndpoint', 
            vpc=vpc,
            service=ec2.InterfaceVpcEndpointAwsService.ECR,
            private_dns_enabled=True
        )
        ec2.InterfaceVpcEndpoint(self, 'ECRDockerVpcEndpoint',
            vpc=vpc,
            service=ec2.InterfaceVpcEndpointAwsService.ECR_DOCKER,
            private_dns_enabled=True
        )
        ec2.GatewayVpcEndpoint(self, 'S3GatewayEndpoint', 
            vpc=vpc,
            service=ec2.GatewayVpcEndpointAwsService.S3,
            subnets=[ec2.SubnetSelection(
                subnets=[private_subnet]
            )]
        )

        # access Cloudwatch logging
        ec2.InterfaceVpcEndpoint(self, 'CloudWatchLogsVpcEndpoint', 
            vpc=vpc,
            service=ec2.InterfaceVpcEndpointAwsService.CLOUDWATCH_LOGS,
            private_dns_enabled=True
        )

    def _prepare_private_subnet(self, vpc: ec2.IVpc) -> ec2.Subnet:
        private_subnet = ec2.PrivateSubnet(self, "MyPrivateSubnet",
            availability_zone="us-east-1a",
            cidr_block="172.31.96.0/20",
            vpc_id=vpc.vpc_id,
        )
        return private_subnet

    def _prepare_ecs_cluster(self, vpc: ec2.IVpc) -> ecs.ICluster:
        """
        Prepares cluster with Fargate capacity provider
        """
        cluster = ecs.Cluster(self, "FargateCPCluster",
            vpc=vpc,
            enable_fargate_capacity_providers=True
        )
        return cluster

    def _prepare_task_definition(self) -> ecs.FargateTaskDefinition:
        """
        Prepares Fargate task definition for container under execution
        """
        task_definition = ecs.FargateTaskDefinition(self, "GreetingTaskDef",
            memory_limit_mib=1024,
            cpu=512
        )
        task_definition.add_container(
            "greeting-container",
            image=ecs.ContainerImage.from_asset(directory="greeting_http_server"),
            port_mappings=[ecs.PortMapping(
                container_port=8080,
                protocol=ecs.Protocol.TCP,
            )],
            logging=ecs.LogDrivers.aws_logs(stream_prefix="greeting-web"),
            health_check=ecs.HealthCheck(
                # Health check options
                # 1.  curl -f http://localhost:8080/
                # 2.  wget --no-verbose --tries=1 --spider http://localhost:8080/
                # Base image for greeting-servers is Alpine and does not have `curl` installed
                command=["CMD-SHELL", "wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1"],
                # the properties below are optional
                interval=Duration.minutes(1),
                retries=3,
                start_period=Duration.seconds(30),
                timeout=Duration.seconds(10)
            )
        )
        return task_definition