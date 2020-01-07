#!/bin/bash

set -xe

apt-get update
apt-get install -y -f ./third_party/*.deb

cp -rf ./third_party/mosquitto.conf /etc/mosquitto/
/usr/sbin/mosquitto -c /etc/mosquitto/mosquitto.conf &
