#!/bin/bash

docker build -t local/backend -f Dockerfile.backend .
docker build -t local/api -f Dockerfile.api .

kubectl create -f api.yaml
kubectl create -f backend.yaml


