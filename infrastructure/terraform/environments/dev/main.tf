# ------------------------------------------
# Basic Resources
# ------------------------------------------
module "iam" {
  source         = "../../modules/iam"
  env            = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id
}

module "security_groups" {
  source         = "../../modules/security-groups"
  env            = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id
  vpc_id         = data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc.id
}

# ------------------------------------------
# Infrastructure Resources
# ------------------------------------------
module "rds" {
  source              = "../../modules/rds"
  env                 = var.env
  sys_name            = var.sys_name
  aws_account_id      = data.aws_caller_identity.current.account_id
  db_subnet_group_ids = data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc_priv_subnet_ids

  rds_config = {
    database_name           = "taskapp"
    master_username         = "postgres"
    port                    = 5432
    instance_count          = 0
    instance_class          = "db.serverless"
    min_capacity            = 1
    max_capacity            = 1
    vpc_security_group_ids  = [module.security_groups.rds_sg_id]
    performance_insights    = true
    enhanced_monitoring     = true
    backup_retention_period = 7
  }

  depends_on = [
    module.iam,
    module.security_groups
  ]
}

module "ecs_cluster" {
  source   = "../../modules/ecs-cluster"
  env      = var.env
  sys_name = var.sys_name
}

# ------------------------------------------
# backend API Resources
# ------------------------------------------
module "alb_backend_api" {
  source          = "../../modules/alb-backend-api"
  env             = var.env
  sys_name        = var.sys_name
  service_name    = "backend-api"
  certificate_arn = "arn:aws:acm:ap-northeast-1:775538353788:certificate/f0dfab13-4e9f-443c-bae9-67770b742683"

  alb = {
    subnet_group_ids       = data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc_pub_subnet_ids
    vpc_security_group_ids = [module.security_groups.alb_backend_api_sg_id]
    system_logs_bucket     = data.terraform_remote_state.log_ops.outputs.s3.logs_bucket.bucket
    vpc_id                 = data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc.id
    enable_access_log      = true
  }

  depends_on = [
    module.iam,
    module.security_groups,
    module.rds,
    module.ecs_cluster
  ]
}

module "ecs_backend_api" {
  source         = "../../modules/ecs-service-backend-api"
  env            = var.env
  sys_name       = var.sys_name
  aws_account_id = data.aws_caller_identity.current.account_id
  service_name   = "backend-api"
  rds_secret_arn = module.rds.rds_master_password_secret_arn

  ecs_service = {
    ecs_cluster_id         = module.ecs_cluster.ecs_cluster_id
    desired_count          = 2
    subnet_ids             = data.terraform_remote_state.cmn_vpc.outputs.vpc.cmn_vpc_priv_subnet_ids
    security_group_ids     = [module.security_groups.ecs_backend_api_sg_id]
    enable_execute_command = true
    lb_target_group_arn    = module.alb_backend_api.target_group_arn
  }

  ecs_task_definition = {
    cpu    = 512
    memory = 2048
    environment = {
      log_level = "debug"
      db_addr   = module.rds.rds_endpoint
      db_user   = module.rds.rds_master_username
      db_name   = module.rds.rds_database_name
      db_port   = module.rds.rds_port

    }
    ecs_execution_role_arn = module.iam.ecs_execution_role_arn
    ecs_task_role_arn      = module.iam.ecs_ecs_backend_api_role_arn
  }

  depends_on = [
    module.iam,
    module.security_groups,
    module.rds,
    module.ecs_cluster
  ]
}

# ------------------------------------------
# frontend Resources
# ------------------------------------------

# ToDo
