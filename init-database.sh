#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE chats;
    CREATE DATABASE messages;
    CREATE DATABASE users;
    GRANT ALL PRIVILEGES ON DATABASE chats TO user;
EOSQL