#!/bin/bash

buf generate
buf generate --template buf.gen-apis.yaml --path api/proto/service.proto

echo '### Generating Proto Files Successfully!! ###'
