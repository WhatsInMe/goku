#!/bin/sh
. ./.env
export HOST_PORT=${HOST_PORT:-80}
docker build -t yw/goku:dev . || exit 1
docker rmi -f $(docker images -f "dangling=true" -q)
docker run \
        -e DOCKER_PORT=${DOCKER_PORT} \
        -e POSTGRES_DBNAME=${POSTGRES_DBNAME} \
        -e POSTGRES_HOST=${POSTGRES_HOST} \
        -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
        -e POSTGRES_PORT=${POSTGRES_PORT} \
        -e POSTGRES_USER=${POSTGRES_USER} \
        --rm \
        -p ${HOST_PORT}:${DOCKER_PORT} \
        --name goku \
        yw/goku:dev
