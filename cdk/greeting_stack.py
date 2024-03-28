from aws_cdk import (
    Stack,
)
from cdk.constructs import (
    GreetingLambdaConstruct,
    APIGConstructProps,
    APIGConstruct,
    GreetingServiceConstruct,
)
from constructs import Construct


class GreetingStack(Stack):
    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # Lambda solution
        greeting_lambda = GreetingLambdaConstruct(self, "HelloLambda")
        APIGConstruct(
            self, "APIG",
            props=APIGConstructProps(
                lambda_func=greeting_lambda.lambda_func
            )
        )

        # ECS Solution
        greeting_service = GreetingServiceConstruct(self, "GreetingService")

        


