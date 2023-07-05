provider "aws" {
    region = "eu-west-2"
}

module "ec2" {
    source = "../modules/modules/ec2"
    name = "ec2 name"
}