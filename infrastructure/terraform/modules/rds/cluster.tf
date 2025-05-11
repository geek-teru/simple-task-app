locals {
  rds_cluster_id = "${var.env}-${var.sys_name}"
}

# Log groups for MySQL
# resource "aws_cloudwatch_log_group" "rds_log_group_error" {
#   name              = "/aws/rds/cluster/${local.rds_cluster_id}/error"
#   retention_in_days = 7
# }

# resource "aws_cloudwatch_log_group" "rds_log_group_audit" {
#   name              = "/aws/rds/cluster/${local.rds_cluster_id}/audit"
#   retention_in_days = 7
# }

# resource "aws_cloudwatch_log_group" "rds_log_group_slowquery" {
#   name              = "/aws/rds/cluster/${local.rds_cluster_id}/slowquery"
#   retention_in_days = 7
# }

# DB subnet group
resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "${var.env}-${var.sys_name}-db-subnet-group"
  subnet_ids = var.db_subnet_group_ids
}

# RDS cluster and instances
resource "aws_rds_cluster" "cluster" {
  cluster_identifier = local.rds_cluster_id
  engine_mode        = "provisioned"
  engine             = "aurora-postgresql"
  engine_version     = "16.6"
  database_name      = var.rds_config.database_name
  port               = var.rds_config.port

  ## general configurations
  master_username             = var.rds_config.master_username
  manage_master_user_password = true
  db_subnet_group_name        = aws_db_subnet_group.db_subnet_group.id
  vpc_security_group_ids      = var.rds_config.vpc_security_group_ids
  #iam_roles                           = var.rds_config.iam_roles
  iam_database_authentication_enabled = true
  db_cluster_parameter_group_name     = aws_rds_cluster_parameter_group.cluster_parameter.name
  db_instance_parameter_group_name    = aws_db_parameter_group.db_parameter.name

  serverlessv2_scaling_configuration {
    max_capacity = var.rds_config.max_capacity
    min_capacity = var.rds_config.min_capacity
  }

  ## backup and maintenance configurations
  preferred_backup_window      = "16:00-17:00"
  preferred_maintenance_window = "sun:17:30-sun:18:30"
  backup_retention_period      = var.rds_config.backup_retention_period
  final_snapshot_identifier    = "${var.env}-${var.sys_name}-snapshot"
  skip_final_snapshot          = true

  # log configurations
  allow_major_version_upgrade = false
  storage_encrypted           = true

  tags = {
    Environment = var.env
  }
}

resource "aws_rds_cluster_instance" "instance" {
  identifier              = "${var.env}-${var.sys_name}-instance"
  cluster_identifier      = aws_rds_cluster.cluster.id
  instance_class          = var.rds_config.instance_class
  engine                  = "aurora-postgresql"
  db_parameter_group_name = aws_db_parameter_group.db_parameter.name
  publicly_accessible     = false
}
