#!/bin/bash


curl -s -X POST -F file=@"$1" https://ipfs.forta.network/api/v0/add | jq -r .Hash
