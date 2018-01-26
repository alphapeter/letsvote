#!/bin/sh
TRAVIS_TAG=${TRAVIS_TAG-latest}
BASE_DIR=$PWD
chmod 755 $npm_package_config_output_path/linux/amd64/$npm_package_config_binary_name

cp $npm_package_config_settings_path/settings.json $npm_package_config_output_path/windows/x86/
cp $npm_package_config_settings_path/settings.json $npm_package_config_output_path/windows/amd64/

cp $npm_package_config_settings_path/settings.json $npm_package_config_output_path/linux/amd64/

mkdir -p $npm_package_config_artifact_path

cd $BASE_DIR/$npm_package_config_output_path/windows/x86/
zip $BASE_DIR/$npm_package_config_artifact_path/$npm_package_config_binary_name-$TRAVIS_TAG-windows-32bit.zip ./*

cd $BASE_DIR/$npm_package_config_output_path/windows/amd64
zip $BASE_DIR/$npm_package_config_artifact_path/$npm_package_config_binary_name-$TRAVIS_TAG-windows-64bit.zip ./*

cd $BASE_DIR/$npm_package_config_output_path/linux/amd64
tar -cvzf $BASE_DIR/$npm_package_config_artifact_path/$npm_package_config_binary_name-$TRAVIS_TAG-linux-amd64.tar.gz ./*
