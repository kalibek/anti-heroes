#!/bin/bash

docker run -p 5432:5432 -d --rm \
  -e POSTGRES_USER=test \
  -e POSTGRES_DB=test \
  -e POSTGRES_PASSWORD=test \
  postgres:alpine
