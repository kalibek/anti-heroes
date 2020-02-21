#!/bin/bash

psql -h localhost -p 5432 -U test -W -d test -f ./create_tables.sql
