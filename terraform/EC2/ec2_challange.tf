1. Create a DB server and output the private IP

resource "aws_instance" "DB-server" {
    ami = "ami-xxxxxxxxxxxxx"
    instance_type = "t2.micro"
    security_groups = 
}