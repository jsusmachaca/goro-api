#!/usr/bin/env bash

# set -e

docker build -t apiserver server
docker build -t apiclient client

docker run --name go-server -dp 8080:8080 --rm --network bridge -e PORT=8080 apiserver
docker run --name express-client -dp 3000:3000 --rm --network bridge -e PORT=3000 -e API_SERVER=go-server:8080 apiclient

docker cp go-server:/app/goroapi .

printf "\e[H\e[2J\n"

echo "Container created and listen on ports: "
ports=$(docker ps | grep -o '0.0.0.0:[0-9]*')
for port in $ports; do
    echo "=> $port"
done