#!/bin/bash
set -e

REPO_NAME=my-projectname
PACKAGE_NAME=my-projectname
BINTRAY_OWNER=wolfeidau

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that dir because we expect that
cd $DIR

# Get the version from the command line
VERSION=$1
if [ -z $VERSION ]; then
    echo "Please specify a version."
    exit 1
fi

# Make sure we have a bintray API key
if [ -z $BINTRAY_API_KEY ]; then
    echo "Please set your bintray API key in the BINTRAY_API_KEY env var."
    exit 1
fi

# Make sure we have a bintray API key
if [ -z $BINTRAY_USER ]; then
    echo "Please set your bintray API user in the BINTRAY_USER env var."
    exit 1
fi

for ARCHIVE in ./pkg/${VERSION}/dist/*; do
    ARCHIVE_NAME=$(basename ${ARCHIVE})

    echo Uploading: $ARCHIVE_NAME
    curl \
        -T ${ARCHIVE} \
        -u${BINTRAY_USER}:${BINTRAY_API_KEY} \
        "https://api.bintray.com/content/${BINTRAY_OWNER}/${REPO_NAME}/${PACKAGE_NAME}/${VERSION}/${ARCHIVE_NAME}"
done
