#!/bin/bash


if ! [ -d /var/lib/mycloudberry/ ]; then
    mkdir /var/lib/mycloudberry
fi

if [ -f "/etc/systemd/system/mycloudberry.service" ]; then
    systemctl stop mycloudberry
    systemctl disable mycloudberry
    systemctl daemon-reload
fi
