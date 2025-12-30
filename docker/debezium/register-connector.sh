#!/bin/sh
set -e

CONNECT_URL="http://connect:8083"
CONFIG_FILE="/docker/debezium/postgres-connector.json"

# wait for connect to be up
until curl -s ${CONNECT_URL}/connectors >/dev/null 2>&1; do
  echo "waiting for connect to be available..."
  sleep 2
done

# register connector (idempotent)
if curl -s ${CONNECT_URL}/connectors/postgres-connector >/dev/null 2>&1; then
  echo "connector already registered"
else
  echo "registering connector"
  curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" ${CONNECT_URL}/connectors -d @${CONFIG_FILE}
fi
