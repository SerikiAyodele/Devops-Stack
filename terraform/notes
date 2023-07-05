TERRAFORM

--------------
TERRAFORM 101
--------------
resource “aws_vpc” “my_vpc” {
	cidr_block = “10.0.0.0/12”
} 

aws_vpc: type of the resource
my_vpc: name we are giving the resource within terraform

Difference between list and tupleIn a list we can only specify one data type, in a tupule we can have multiple data types
Tuple is similar to list
Object is similar to Map
 
Tfswitch allows you to easily switch between versions of terraform


--------------------------
TERRAFORMM VERSION CHANGE
--------------------------
Example of using count
To for example decide wether to deploy something to prod or test environment
count = var.environment == "prod" ? 1:0

-var-file for targeting a filr when running terraform commands


----
EC2
----
- Imagine we have to create 50 ports on a security group, we could use dynamic blocks instead on ingress, egress blocks.
  We create a list of the ports we want open and it takes care of them for us 

- Dynamic blocks is a fast way of setting up security groups if we have multiple groups

- Elastic IPs remain the same4 after redeploying unlike public IPs


--------
MODULES
--------
- To access any atribute or value for anything we've created inside a module we have
  to create an output


-----------
OTHER NOTES
-----------
- A local is only accessible within the local module, a terraform variable is       globally scoped

- Data sources allow terraform to use information defined outside of terraform, defined by another seperate terraform configuration

- <cidr_blocks = [for s in data.aws_subnet.redshift_ingress : s.cidr_block]> - uses
 a list comprehension to extract the CIDR blocks from  'data.aws_subnet.redshift_ingress' 
 data source and assign them to the 'cidr_blocks' variable

