#!/bin/bash
export GOAPP=WhereIsMyDriver
export GOENV=local
export PORT=3001
export DB_NAME=where_is_my_driver
export DB_HOST=localhost
export DB_USER=root
export DB_PASSWORD=root
export DB_PORT=3306

cd $GOPATH/src/WhereIsMyDriver/models && go test -v --cover
cd $GOPATH/src/WhereIsMyDriver/routers && go test -v --cover