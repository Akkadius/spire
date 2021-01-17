#!/bin/bash

eval $(egrep -v '^#' .env | xargs); mysql -u$MYSQL_USERNAME -p$MYSQL_PASSWORD -h $MYSQL_HOST $MYSQL_DATABASE