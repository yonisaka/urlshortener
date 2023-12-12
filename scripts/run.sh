#!/bin/bash

# app config
export APP_ENV=dev
export SERVER_PORT=8080
export BASE_URL=localhost
export CODE_SIZE=8

# master db config
export POSTGRES_USER_MASTER=test
export POSTGRES_PASSWORD_MASTER=test
export POSTGRES_HOST_MASTER=localhost
export POSTGRES_PORT_MASTER=5432
export POSTGRES_DB_MASTER=test

# slave db config
export POSTGRES_USER_SLAVE=test
export POSTGRES_PASSWORD_SLAVE=test
export POSTGRES_HOST_SLAVE=localhost
export POSTGRES_PORT_SLAVE=5433
export POSTGRES_DB_SLAVE=test

# use replica config
export IS_REPLICA=false

# tracing config
export OTEL_AGENT=http://localhost:14268/api/traces

# cache config
#export REDIS_HOST=localhost:6379

go build -o main ./cmd/main.go && ./main
