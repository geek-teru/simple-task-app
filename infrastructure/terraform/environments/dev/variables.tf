variable "env" {
  type    = string
  default = "dev"
}

variable "sys_name" {
  type    = string
  default = "task-app"
}

data "aws_caller_identity" "current" {}

data "terraform_remote_state" "cmn_vpc" {
  backend = "s3"

  config = {
    bucket = "dev-terraform-aws"
    key    = "cmn-vpc/terraform.tfstate"
    region = "ap-northeast-1"
  }
}
