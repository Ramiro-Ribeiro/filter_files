#!/bin/sh

./myapp &

sleep 5

exec nginx -g 'daemon off;'