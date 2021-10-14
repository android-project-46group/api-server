#!/bin/bash

# Usage:
#   make this file exutable
#   ./db_migrate.sh up/down <number>

echo "----------------------------"
CMDNAME=`basename $0`
if [ $# -ne 2 ]; then
  echo "Usage: $CMDNAME up/down <number>" 1>&2
  echo "----------------------------"
  exit 1
fi

# == は bash の性質？ /bin/sh だとエラー
if [ $1 == 'up' ]; then
    echo 'Migratino up'
elif [ $1 == 'down' ]; then
    echo 'Migration down'
else
    echo "Usage: $CMDNAME up/down <number>" 1>&2
    echo "----------------------------"
    exit 1
fi

# use ../.env file as a variable related db
source ../.env

# sslmode=disableは、通信が暗号化されてないという意味
POSTGRESQL_URL="postgres://ubuntu:${DBPASSWORD}@localhost:5432/${DBNAME}?sslmode=disable"

migrate -path ../migrations -database "${POSTGRESQL_URL}" $1 $2
echo "----------------------------"

# 今作り直したDBから、SQLBoilerで使う構造体にも反映させる
cd ../db
sqlboiler psql
