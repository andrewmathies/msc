#!/bin/bash

curl localhost:9090/erds -d '{"name": "Erd_CameraStatus", "address":"0x0503", "fields": ["Erd_CameraStatus"], "actions": ["Read", "Write"] }'
