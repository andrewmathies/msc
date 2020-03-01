#!/bin/bash

curl localhost:9090/erds -d '{"name": "Erd_CameraStatus", "fields": ["CameraStatus"], "actions": ["Read", "Write"] }' | python -m json.tool
