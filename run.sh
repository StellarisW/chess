#!/bin/bash

echo '----------build go----------'
go build -o ./app/service/api/chess-api -v ./app/service/api/chess.go
go build -o ./app/service/rpc/register/register-rpc -v ./app/service/rpc/register/register.go

echo '----------docker compose----------'
docker-compose down --rmi all
docker-compose up