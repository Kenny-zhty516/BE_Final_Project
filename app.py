#!/usr/bin/env python3
import os

import aws_cdk as cdk

from cdk.greeting_stack import GreetingStack


app = cdk.App()
GreetingStack(app, "GreetingStack",
    env=cdk.Environment(account='891377079819', region='us-east-1'),
)

app.synth()
