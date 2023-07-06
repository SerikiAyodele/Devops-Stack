provider "aws" {
    region = "eu-west-2"
}

module "db" {
    source = "./db"
}

module "web" {
    source = "./web"
}

output "privateIP" {
    value = module.db.PrivateIP
}

output "publicIP" {
    value = module.web.pub_ip
}

