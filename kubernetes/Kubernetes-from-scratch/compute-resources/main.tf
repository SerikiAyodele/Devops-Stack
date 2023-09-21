locals {
    project = var.project_name
}

#-------
# VPC
#-------
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "kthw-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1"]
  private_subnets = ["10.0.1.0/24"]
  public_subnets  = ["10.0.101.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true
  enable_dns_hostnames = true 
  enable_dns_support = true

  tags = {
    Terraform = "true"
    Environment = "dev"
  }
}
#----------------
# EC2 INSTANCE
#----------------
module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"

  name = "kthw-instance"

  instance_type          = "t2.micro"
  key_name               = "kthw-user"
  monitoring             = true
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = "subnet-eddcdzz4"

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}