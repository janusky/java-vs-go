#!/bin/env bash

# Creación de usuario
psql -U postgres -c "DROP USER IF EXISTS $POSTGRES_USER;"
psql -U postgres -c "CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';"
# Creación de base de datos
psql -U postgres -c "DROP DATABASE IF EXISTS $POSTGRES_DB;"
psql -U postgres -c "CREATE DATABASE $POSTGRES_DB OWNER $POSTGRES_USER;"
# Creación de las tablas
psql -U postgres -c "CREATE TABLE client(id SERIAL PRIMARY KEY NOT NULL, name VARCHAR(20), email VARCHAR(20), phone VARCHAR(20));"
psql -U postgres -c "CREATE TABLE transaction(id SERIAL PRIMARY KEY NOT NULL, from_client_id INTEGER, to_client_id INTEGER, amount INTEGER);"