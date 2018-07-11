#!/bin/sh

# Migrate config if necessary
if [ ! -d $SNAP_DATA/settings.json ]; then
    cp $SNAP/conf/settings.json $SNAP_DATA/settings.json
fi
cd $SNAP_DATA
$SNAP/bin/letsvote

