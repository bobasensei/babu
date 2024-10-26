#!/bin/sh

USERNAME=""
PASSWORD=""

curl -L https://auth.enterprise.wikimedia.com/v1/login \
	-H "Content-Type: application/json" \
	-d "{\"username\":\"$USERNAME\", \"password\":\"$PASSWORD\"}"
