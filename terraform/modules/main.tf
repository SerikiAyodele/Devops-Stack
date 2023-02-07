provider "aws" {
    region = "eu-west-2"
}

module "ec2module" {
    source = "./ec2"
    ec2name = "Name from module"
}

