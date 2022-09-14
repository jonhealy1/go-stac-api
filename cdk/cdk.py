from aws_cdk import App
from fastapi_stack import FastAPIStack
 
app = App()
FastAPIStack(app, "FastAPIStack")
app.synth()