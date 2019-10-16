#!/bin/sh

curl -d '{"instances": [[1,2,3]]}' -X OPTIONS http://localhost:81/v1/models/goat_herd:predict -I
