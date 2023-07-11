provider "aws" {
    region = "eu-west-2"
}

variable "ingressrules" {
    type = list(number)
    default = [80,443]
}

variable "egressrules" {
    type = list(number)
    default = [80,443, 25, 3306, 53, 8080]
}

resource "aws_instance" "db" {
    ami = "ami-032598fcc7e9d4ce7c"
    instance_type = "t2.micro"
    security_groups = [aws_security_group.webtraffic.name]
    tags = {
        Name = "db-server"
    }
}

resource "aws_instance" "web" {
    ami = "ami-032598fcc7e9d4ce7c"
    instance_type = "t2.micro"
    security_groups = [aws_security_group.webtraffic.name]
    user_data = file("server-script.sh")
    tags = {
        Name = "web-server"
    }
}

resource "aws_security_group" "webtraffic" {
    name  = "Allow HTTPS"

    dynamic "ingress" {
        iterator = port
        for_each = var.ingressrules
        content {
          from_port = port.value
          to_port = port.value
          protocol = "TCP"
          cidr_blocks = ["0.0.0.0/0"]
        }
    }

    dynamic "egress" {
        iterator = port
        for_each = var.egressrules
        content {
          from_port = port.value
          to_port = port.value
          protocol = "TCP"
          cidr_blocks = ["0.0.0.0/0"]
        }
    }
}

output "privateIP" {
    value = aws_instance.db.private_ip
}

output "publicIP" {
    value = aws_instance.web.public_ip
}