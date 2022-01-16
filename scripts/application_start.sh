#!/bin/bash
sudo service codedeploy-agent start
cd /home/ubuntu/test
sudo go build main.go
sudo service test restart