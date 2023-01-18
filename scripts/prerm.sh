#!/bin/bash


if [ -f "/etc/systemd/system/mycloudberry.service" ]; then
    systemctl stop mycloudberry
    systemctl disable mycloudberry
    systemctl daemon-reload
fi
