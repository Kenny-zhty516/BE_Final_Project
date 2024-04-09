from aws_cdk import (
    aws_ec2 as ec2,
    aws_lambda as _lambda,
    aws_apigateway as api_gw,
)
from constructs import Construct


class APIGConstructProps:
    def __init__(self, lambda_func: _lambda.Function):
        self.lambda_func = lambda_func


class APIGConstruct(Construct):
    def __init__(
        self,
        scope: Construct,
        construct_id: str,
        props: APIGConstructProps,
        **kwargs
    ) -> None:
        super().__init__(scope, construct_id, **kwargs)
        api = api_gw.LambdaRestApi(
            self, 'Greetings',
            handler=props.lambda_func,
            api_key_source_type=api_gw.ApiKeySourceType.HEADER,
        )
        api.root.add_method("GET", api_key_required=True)
        proxy = api.root.get_resource("{proxy+}")
        proxy.add_method("GET", api_key_required=True)

        apiKey = api_gw.ApiKey(self, "GreetingApiKey")

        plan = api.add_usage_plan("UsagePlan",
            name="Easy",
            throttle=api_gw.ThrottleSettings(
                rate_limit=10,
                burst_limit=2
            )
        )
        plan.add_api_stage(api=api, stage=api.deployment_stage)
        plan.add_api_key(apiKey)

        self.api = api
