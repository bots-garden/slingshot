#!/bin/bash

sudo chmod 666 /var/run/docker.sock
ls -l /var/run/docker.sock
sudo systemctl restart docker.service

