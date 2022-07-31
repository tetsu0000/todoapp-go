#!/bin/bash

RESPONSE=$(curl "${ENDPOINT}/todos" \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"message":"post todo"}' \
  -w"\n")
echo ${RESPONSE}
TODO_ID=$(echo ${RESPONSE} | jq .data.id -r)
echo $TODO_ID

curl "${ENDPOINT}/todos" \
  -w"\n"

curl "${ENDPOINT}/todos/${TODO_ID}" \
  -X PUT \
  -H "Content-Type: application/json" \
  -d '{"message":"put todo"}' \
  -w"\n"

curl "${ENDPOINT}/todos/${TODO_ID}" \
  -X DELETE \
  -w"\n"
