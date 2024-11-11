#!/bin/sh

if [ "$APP_ENV" = "development" ]; then
  echo "Starting in development mode with hot-reloading..."
  air -c .air.toml
else
  echo "Starting in production mode..."
  go build -o main ./main.go
  ./main
fi
