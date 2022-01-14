#!/bin/bash
sudo service codedeploy-agent start
cd test/
sudo go build main.go
sudo service test restart