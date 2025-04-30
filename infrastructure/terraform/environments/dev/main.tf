module "iam" {
  source         = "../../modules/iam"
  env            = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id
}

module "security-groups" {
  source         = "../../modules/security-groups"
  env            = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id
  vpc_id         = data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc.id
}

output "vpc" {
  value = data.terraform_remote_state.cmn_vpc.outputs.vpc
}

output "security_group" {
  value = data.terraform_remote_state.cmn_vpc.outputs.security_group
}
