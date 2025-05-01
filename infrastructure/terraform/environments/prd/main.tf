module "vpc" {
  source = "../../modules/vpc"

  env_name       = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id

  # vpc
  vpc_config = var.cmn_vpc_config


  # public subnet
  pub_subnet_config = var.cmn_pub_subnet_config

  # private subnet
  priv_subnet_config = var.cmn_priv_subnet_config
}

module "security-group" {
  source = "../../modules/security-group"

  vpc_id    = module.vpc.cmn_vpc.id
  sg_config = var.sg_config
}

output "vpc" {
  value = module.vpc
}

output "security_group" {
  value = module.security-group
}