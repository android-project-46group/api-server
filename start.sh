#!/bin/sh

set -e

echo "run db migration"
if [ -z "$DB_SOURCE" ]; then
    # docekr build etc
    echo "environment variable DB_SOURCE is empty"
    source /app/app.env
fi
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

/app/populate

echo "start the app"
exec "$@"
