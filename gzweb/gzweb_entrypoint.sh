#!/bin/bash

# Link models from external source as frontend assets
SOURCE="/data/models"
TARGET="/gzweb/static/assets"

if [ ! -d ${SOURCE} ]; then
    echo "Models volume not mounted to: ${SOURCE}."
    exit 1
fi

if [ -d ${TARGET} ] && [ ! "$(ls -A ${TARGET})" ]; then
    rmdir $TARGET
fi

if [ ! -d ${TARGET} ]; then
    echo "Linking ${SOURCE} -> ${TARGET}"
    ln -s ${SOURCE} ${TARGET}
fi

# Start gzweb
cd /gzweb || exit
GAZEBO_MASTER_URI="${1}" ./server.js
