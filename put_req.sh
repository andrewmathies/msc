#!/bin/bash

curl -v localhost:9090/2 -XPUT -d '{"name": "Erd_CameraStatusNew", "fields": ["CameraStatusNew"], "actions": ["Read", "Write"] }' | python -m json.tool
