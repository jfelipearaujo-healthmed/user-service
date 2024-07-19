#!/bin/sh

echo "Initializing Secrets Manager..."

awslocal secretsmanager create-secret \
    --name db-secret-url \
    --description "DB URL" \
    --secret-string "postgres://user:user@localhost:5432/user_db?sslmode=disable"

awslocal secretsmanager create-secret \
    --name cache-host-secret \
    --description "CACHE HOST" \
    --secret-string "127.0.0.1:6379"

echo "Secrets Manager initialized!"