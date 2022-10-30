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
set -eu

dir_name="/home/ubuntu/work/go/web"

if [ `pwd` != $dir_name ]; then
    echo "You should use this file at the following directory."
    echo "$dir_name"
    echo ""
    echo "Please change directory and try again."
    exit 0
fi

sed -i -E "s@IS_CGI=false@IS_CGI=true@g" app.env

go build -ldflags "-s -w" -o server.cgi main.go && \
    sudo mv server.cgi /usr/lib/cgi-bin/ && \
    sudo systemctl restart apache2

tmp=$?
if [ $tmp -eq 0 ]; then
    echo "Deploy finished successfully"
else
    echo "Something wrong happend."
    echo "The status code is $tmp"
fi
