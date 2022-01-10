#!/bin/bash
cd test/
sudo go build main.go
sudo systemctl enable test
sudo service test restart