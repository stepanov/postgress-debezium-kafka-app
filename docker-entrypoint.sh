#!/bin/sh
set -e

# run migrations if enabled
if [ "${MIGRATE_ON_START}" = "true" ]; then
  echo "running migrations..."
  /app/migrate
fi

exec "$@"
