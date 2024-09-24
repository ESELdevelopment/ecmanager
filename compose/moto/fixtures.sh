#!/bin/sh

echo $PWD

aws ecs create-cluster --cluster-name moto
aws ecs register-task-definition --cli-input-json file://nginx.json
aws ecs create-service \
    --cluster moto \
    --service-name nginx-service \
    --task-definition nginx-fargate-task \
    --desired-count 1 \
    --launch-type FARGATE
