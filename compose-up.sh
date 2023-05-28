#!/bin/sh

echo "removing mysql-data volumes"

docker volume rm mysql-data

echo "giving execute permission"

chmod a+x ./mysql/init/create_table.sh

echo "running command [docker compose up -d --build]"

docker compose up -d --build
