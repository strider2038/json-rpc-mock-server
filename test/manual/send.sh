#!/usr/bin/env bash

scriptDirectory="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

request=$(cat "${scriptDirectory}/$1.json")
request="${request//[$'\t\r\n ']}"

echo "JSON RPC request: ${request}"

(echo ${request}; sleep 1) | nc localhost 4000
