# ECManager

A terminal UI to manage your AWS ElasticContainerService-Clusters

## Note

> This tool is currently under development and therefore not released yet.

## Local Development

If you want to use a local AWS ECS-Mock, we have prepared a docker-compose file for you.
Just run the following command:

```bash
docker compose up -d
```

After that you can start the ECManager with the following environment variables:

```bash
AWS_ENDPOINT_URL=http://localhost:5000
go run cmd/aws/main.go
```

## Setting up Moto

If you want to use Moto as a local AWS-Mock, you can start the docker-compose file:

```bash
docker compose up -d
```

After that you can interact with the Moto-Server via aws cli:

```bash
# create a cluster
aws --endpoint-url=http://localhost:5000 ecs create-cluster --cluster-name test-cluster
aws --endpoint-url=http://localhost:5000 ecs ...
```
