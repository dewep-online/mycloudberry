#!/bin/bash


if [ -f "/etc/systemd/system/mycloudberry.service" ]; then
    systemctl start mycloudberry
    systemctl enable mycloudberry
    systemctl daemon-reload
fi
