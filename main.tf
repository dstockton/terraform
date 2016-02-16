provider "rightscale" {
  cloud_name = "EC2 eu-west-1"
}

resource "rightscale_network" "terraform.test" {
  cidr_block = "192.168.0.0/24" 
}
