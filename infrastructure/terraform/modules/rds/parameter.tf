resource "aws_rds_cluster_parameter_group" "cluster_parameter" {
  name   = "aurora-pg16-cluster-logging"
  family = "aurora-postgresql16"

  parameter {
    name         = "log_min_duration_statement"
    value        = "500"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "log_statement"
    value        = "ddl"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "log_connections"
    value        = "1"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "log_disconnections"
    value        = "1"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "log_line_prefix"
    value        = "%t:%r:%u@%d:[%p]:"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "client_encoding"
    value        = "UTF8"
    apply_method = "immediate"
  }

  tags = {
    Environment = "dev"
    Project     = "aurora-pg16"
  }
}


resource "aws_db_parameter_group" "db_parameter" {
  name   = "aurora-pg16-instance-params"
  family = "aurora-postgresql16"

  parameter {
    name         = "max_connections"
    value        = "200"
    apply_method = "pending-reboot"
  }

  tags = {
    Environment = "${var.env}"
    Project     = "aurora-pg16"
  }
}

