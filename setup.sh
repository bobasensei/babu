#!/bin/sh

export BABU_DATABASE=postgres://USER:PASSWORD@HOST:PORT/babu

USERNAME="USERNAME"
PASSWORD="PASSWORD"

export BABU_WIKIMEDIA=$(curl -s -L https://auth.enterprise.wikimedia.com/v1/login \
	-H "Content-Type: application/json" \
	-d "{\"username\":\"$USERNAME\", \"password\":\"$PASSWORD\"}" | jq .access_token -r)

