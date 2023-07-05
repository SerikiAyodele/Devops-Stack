provider "aws" {
    region = "eu-west-2"
}

module "ec2module" {
    source = "./ec2"
    name = "Name from module"
}

