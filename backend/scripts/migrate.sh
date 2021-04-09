#!/bin/bash

# Enable golang to be used in this script
GO111MODULE=on go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# TODO make a check to the PG DB metadata versions table and check if we need to migrate
# 