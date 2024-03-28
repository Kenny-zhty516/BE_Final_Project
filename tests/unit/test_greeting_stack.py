import aws_cdk as core
import aws_cdk.assertions as assertions

from cdk.greeting_stack import GreetingStack

# example tests. To run these tests, uncomment this file along with the example
# resource in cdk/greeting_stack.py
def test_sqs_queue_created():
    app = core.App()
    stack = GreetingStack(app, "cdk-docker-lambda")
    template = assertions.Template.from_stack(stack)

#     template.has_resource_properties("AWS::SQS::Queue", {
#         "VisibilityTimeout": 300
#     })
