#!/bin/sh
while ! nc -z mongo 27017;
do
    sleep 3;
done