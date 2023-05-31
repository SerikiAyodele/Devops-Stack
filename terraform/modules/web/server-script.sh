#!/bin/bash
sudo yum update
sudo yum install -y httpd
sudo sytemctl start httpd
sudo systemctl enable httpd
echo "<h1>Hello, badass DevOps engineer</h1>" | sudo tee /var/www/html/index.html 