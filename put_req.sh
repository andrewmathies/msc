#!/bin/bash

curl -v localhost:9090/erds/2 -XPUT -d '{"name": "Erd_CameraStatusNew", "address":"0x0504", "fields": ["Erd_CameraStatusNew"], "actions": ["Read", "Write"] }'
