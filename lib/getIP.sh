#!/bin/bash
echo "Current using IP:"
# Now curl package is neccesary for this script.
curl -x 127.0.0.1:8118 http://icanhazip.com/
