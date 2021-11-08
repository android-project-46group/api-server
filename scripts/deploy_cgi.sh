#!/bin/bash
# 
# Description
#   1. make cgi file 
#   2. move the cgi file to the contents folder
#   3. restart the apache server
#
# Usage:
#   make this file exutable
#   sh scripts/deploy_cgi.sh

dir_name="/home/ubuntu/work/go/web"

if [ `pwd` != $dir_name ]; then
    echo "You should use this file at the following directory."
    echo "$dir_name"
    echo ""
    echo "Please change directory and try again."
    exit 0
fi

go build -ldflags "-s -w" -o server.cgi server.go && \
    sudo mv server.cgi /usr/lib/cgi-bin/ && \
    sudo systemctl restart apache2

tmp=$?
if [ $tmp -eq 0 ]; then
    echo "Deploy finished successfully"
else
    echo "Something wrong happend."
    echo "The status code is $tmp"
fi
