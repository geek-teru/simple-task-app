# ----------------------------------------
# Provider Settings
# ----------------------------------------
provider "aws" {
  region = "ap-northeast-1"
}

####################################
# Terraform tfstate
####################################
#resource "aws_dynamodb_table" "cmn_ci_tfstate_lock" {
#  count = 1
#  name           = "cmn_ci_tfstate_lock"
#  read_capacity  = 1
#  write_capacity = 1
#  hash_key       = "LockID"
#  attribute {
#    name = "LockID"
#    type = "S"
#  }
#  lifecycle {
#    ignore_changes = ["read_capacity", "write_capacity"]
#  }
#}

terraform {
  required_version = ">= 1.5.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.74.0"
    }
  }
  backend "s3" {
    bucket  = "prd-terraform-aws"
    region  = "ap-northeast-1"
    key     = "cmn-vpc/terraform.tfstate"
    acl     = "bucket-owner-read"
    encrypt = true
  }
}