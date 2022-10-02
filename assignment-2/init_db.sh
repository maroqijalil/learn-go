#!/bin/bash

docker -v
if [ $? -ne 128 ]; then
  docker run --rm -e POSTGRES_PASSWORD=postgres -dp 5432:5432 postgres
fi
