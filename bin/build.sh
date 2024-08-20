#!/usr/bin/env bash

# set -e

docker build -t apiserver server
docker build -t apiclient client

docker network create gogo

docker run --name express-client -dp 8080:8080 --rm --network gogo apiserver
docker run --name go-server -dp 3000:3000 --rm --network gogo -e PORT=3000 -e API_SERVER=express-client:8080 apiclient

