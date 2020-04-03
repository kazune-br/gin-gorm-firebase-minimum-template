#!/bin/bash

until mysqladmin ping -h 127.0.0.1 -uroot -P 3390 --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 1
done