variable "env" {
  type    = string
  default = "prd"
}

variable "sys_name" {
  type    = string
  default = "cmn"
}

data "aws_caller_identity" "current" {}


# vpc
variable "cmn_vpc_config" {
  type = map(any)
  default = {
    cidr       = "10.100.0.0/16",
    managed_by = "terraform"
  }
}

# public subnet
variable "cmn_pub_subnet_config" {
  type = map(any)
  default = {
    az-a = { az = "ap-northeast-1a", cidr = "10.100.10.0/24", managed_by = "terraform" }
  }
}

# private subnet
variable "cmn_priv_subnet_config" {
  type = map(any)
  default = {
    az-a = { az = "ap-northeast-1a", cidr = "10.100.110.0/24", managed_by = "terraform" }
  }
}

# securitygroup
variable "sg_config" {
  type = list(any)
  default = [
    { name = "linux_sg", port = [22], src_ip = ["0.0.0.0/0"], managed_by = "terraform" },
    { name = "windows_sg", port = [3389], src_ip = ["0.0.0.0/0"], managed_by = "terraform" }
  ]
}
