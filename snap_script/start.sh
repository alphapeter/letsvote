#!/bin/bash

# Migrate config if necessary
if [ ! -f $SNAP_DATA/settings.json ]; then
    cp $SNAP/conf/settings.json $SNAP_DATA/settings.json
fi
cd $SNAP_DATA
$SNAP/bin/letsvote
