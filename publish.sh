#/bin/bash

docker build -f min.Dockerfile . -t action-cm-generator
docker tag action-cm-generator:latest public.ecr.aws/n0a0a3c3/action-cm-generator:latest
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/n0a0a3c3
docker push public.ecr.aws/n0a0a3c3/action-cm-generator:latest