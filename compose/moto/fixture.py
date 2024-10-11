#!/usr/bin/env python

import boto3
import os

client = boto3.client("ecs", endpoint_url=os.getenv("AWS_ENDPOINT_URL"))

cluster_name = "default"

client.create_cluster(clusterName=cluster_name)

definition = client.register_task_definition(
  family="hello_world",
  containerDefinitions=[
    {
      "name": "hello_world",
      "image": "docker/hello-world:latest",
      "cpu": 1024,
      "memory": 400,
    }
  ],
)

task_definition_arn = definition["taskDefinition"]["taskDefinitionArn"]

client.create_service(
  cluster=cluster_name,
  serviceName="hello_service",
  taskDefinition=task_definition_arn,
  desiredCount=2,
  platformVersion="2",
)

client.run_task(
  launchType="FARGATE",
  cluster=cluster_name,
  overrides={},
  taskDefinition=task_definition_arn,
  count=2,
  startedBy="moto",
)
