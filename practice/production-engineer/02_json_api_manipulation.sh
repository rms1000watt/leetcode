#!/usr/bin/env bash

# https://programminghistorian.org/en/lessons/json-and-jq
# https://jqplay.org/

curl -s https://jsonplaceholder.typicode.com/todos/1 | jq '. | {NewID: .id, NewTitle: .title}'
