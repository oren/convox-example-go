#!/bin/sh

docker-compose build
docker-compose up -d
sleep 10 # give the database server enough time to start!
cd integration-test && npm start && cd ..
docker-compose down
