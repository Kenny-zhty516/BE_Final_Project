from aws_cdk import aws_lambda as _lambda
from constructs import Construct


class GreetingLambdaConstruct(Construct):
    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)
        greeting_lambda = _lambda.DockerImageFunction(
            scope=self,
            id="GreetingLambda",
            # Function name on AWS
            function_name="GolangGreetingLambda",
            # Use aws_cdk.aws_lambda.DockerImageCode.from_image_asset to build
            # a docker image on deployment
            code=_lambda.DockerImageCode.from_image_asset(
                # Directory relative to where you execute cdk deploy
                # contains a Dockerfile with build instructions
                directory="greeting_lambda"
            ),
        )
        self.lambda_func = greeting_lambda