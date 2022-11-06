#!/bin/bash
# 
# Description
#   Get the latest apk file from the artifacts
#   and copy it to the deploy folder.
#   * This file is executed at 0 sharp in Japan time(by cron).
#
# Usage:
#   bash scripts/deploy_cgi.sh
#

# Working dir
WORKING_DIR="/home/ubuntu/work/tmp/$(TZ='Japan' date +%Y_%m%d_%H%M%S)/"
LOG_OUT="log"
LOG_ERR="error"
mkdir -p "${WORKING_DIR}"

# Set true if u wanna delete the working directory at the end of this file.
DELETE_DIR=false

# Log settings
exec 1> >(
    TZ='Japan' awk '{print strftime("[%Y-%m-%d %H:%M:%S]"),$0 } { fflush() } ' \
    >> "${WORKING_DIR}${LOG_OUT}"
)
exec 2> >(
    TZ='Japan' awk '{print strftime("[%Y-%m-%d %H:%M:%S]"),$0 } { fflush() } ' \
    >> "${WORKING_DIR}${LOG_ERR}"
)

# Constants for GitHub api
GITHUB_URL="https://api.github.com"
MY_NAME="android-project-46group"
REPO_NAME="android"

# Get id of the latest apk file.
# Assume that id comes first, then name comes.
ACTION_ID="$(curl -s \
    -H "Accept: application/vnd.github.v3+json" \
    $GITHUB_URL/repos/$MY_NAME/$REPO_NAME/actions/artifacts | \
    awk '{if($1 == "\"id\":"){id = $2}; 
    if($1 == "\"name\":" && $2 == "\"apk\","){print id; exit}}' | \
    tr -d ,)"

# Log: action id
echo "action_id ${ACTION_ID}"

# GitHub credential (name and token)
PATH_TO_CREDENTIAL="/home/ubuntu/work/.secret/_netrc"
APK_NAME="app-debug.apk"

# Save apk.zip file to working_dir.
curl -s \
    --netrc-file $PATH_TO_CREDENTIAL \
    -L \
    -o "${WORKING_DIR}apk.zip" \
    $GITHUB_URL/repos/$MY_NAME/$REPO_NAME/actions/artifacts/$ACTION_ID/zip -f

unzip "${WORKING_DIR}apk.zip" -d "$WORKING_DIR"

# Copy the downloaded apk file to the deploy folder.
PATH_TO_DEPLOY_DIR="/var/www/html/"
cp "$WORKING_DIR$APK_NAME" "$PATH_TO_DEPLOY_DIR$APK_NAME"

# By default, the directory will not be deleted.
if [ "$DELETE_DIR" = true ]; then
    rm -rf "${WORKING_DIR}"
fi
