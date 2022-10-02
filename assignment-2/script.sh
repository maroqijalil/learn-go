#!/bin/bash

container_id=""

docker -v
if [ $? -ne 128 ]; then
  container_id=$(docker run --rm -e POSTGRES_PASSWORD=postgres -dp 5432:5432 postgres)
fi

sleep 2s

go run main.go

docker stop $container_id
