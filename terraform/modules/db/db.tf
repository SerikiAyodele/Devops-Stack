resource "aws_instance" "db" {
    ami = "ami-032598fcc7e9d4ce7c"
    instance_type = "t2.micro"
    tags = {
        Name = "db-server"
    }
}

output "PrivateIP" {
    value = aws_instance.db.private_ip
}