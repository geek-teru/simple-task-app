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

module "rds" {
  source         = "../../modules/rds"
  env            = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id
  db_subnet_group_ids = (
    var.env == "dev"
    ? data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc_pub_subnet_ids # devはパブリックサブネットを使用
    : data.terraform_remote_state.cmn_vpc.outputs.cmn_vpc_priv_subnet_ids
  )

  rds_config = {
    database_name           = "taskapp"
    master_username         = "postgres"
    port                    = 5432
    instance_count          = 0
    instance_class          = "db.serverless"
    min_capacity            = 1
    max_capacity            = 1
    vpc_security_group_ids  = [module.security-groups.rds_sg_id]
    performance_insights    = true
    enhanced_monitoring     = true
    backup_retention_period = 7
  }
}

output "vpc" {
  value = data.terraform_remote_state.cmn_vpc.outputs.vpc
}

output "security_group" {
  value = data.terraform_remote_state.cmn_vpc.outputs.security_group
}
