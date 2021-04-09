#!/bin/bash
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root this script is attempting to start the database service."
    exit
fi

# This script is intended to simulate the init script if postgres was in a separate container and docker needed to do inital postgres setup (since containers always are init)
# For local dev this is a quick tool to get going without having to setup the DB manually.

# Grab env vars from file and print during run
set -o allexport
source systemproperties.env
echo "--- Environment variables ---"
for i in `cat systemproperties.env`
do
    echo $i
done
echo -e '--- Environment variables ---\n'

which psql

if [ "$?" -gt "0" ]; then
    echo -e 'Postgres not installed please install before running this script, will stop execution now.'
    exit
else
    echo -e 'Installed and now checking to see if database already exists or running.'
    status=$(service postgresql status)
    echo -e "postgres status: $status\n"
fi

if [[ "$status" =~ "down"  ]]; then
    service postgresql start
else
    echo -e "The database is already running.\n"
fi

# Check if the database exists if not create the local database in postgres
sudo -u postgres psql -tAc "SELECT 1 FROM pg_database WHERE datname='${DBNAME}'" | grep -q 1 || sudo -u postgres psql -c "CREATE DATABASE ${DBNAME}"

# Check if default dev user is configured then create and grant privileges
if [ "$( sudo -u postgres psql -tAc "SELECT 1 FROM pg_roles WHERE rolname='${DEVUSER}'" )" = '1' ]
then
    echo "User ${DEVUSER} already exists and database ${DBNAME} is present, exiting startdb.sh"
    exit
else
    # Check pg for the default admin role and create with privilege
    echo -e "User ${DEVUSER} does not exist...creating now\n"
    sudo -u postgres psql -c "create user ${DEVUSER}"
    sudo -u postgres psql -c "grant all privileges on database ${DBNAME} to ${DEVUSER}"
    sudo -u postgres psql -c "alter user ${DEVUSER} with encrypted password '${DEVPASS}'"
    echo -e "--- ${DEVUSER} role is created with rights on database ${DBNAME} ---\n"
fi

# disables env export after work is done
set +o allexport

echo "startdb.sh execution finished - local database '${DBNAME}' exists with default user of ${DEVUSER} password of ${DEVPASS}"