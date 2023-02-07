variable "ec2name" {
    type = string
}

resource "aws_instance" "ec2" {
    ami = "ami-032598fcc7e9d4ce7c"
    instance_type = "t2.micro"
    security_groups = [aws_security_group.webtraffic.name]
    tags = {
        Name = var.ec2name
    }
}

output "instance_id" {
    value = aws_instance.ec2.id
} 