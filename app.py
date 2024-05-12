#!/usr/bin/env python3
import os

import aws_cdk as cdk

from cdk.greeting_stack import GreetingStack


app = cdk.App()
GreetingStack(app, "GreetingStack",
    env=cdk.Environment(account='480763985297', region='us-east-1'),
)

app.synth()
