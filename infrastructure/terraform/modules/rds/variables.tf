variable "env" {
  type = string
}

variable "sys_name" {
  type = string
}

variable "aws_account_id" {
  type = string
}


variable "db_subnet_group_ids" {
  type = list(string)
}

variable "rds_config" {
  type = object({
    database_name          = string
    master_username        = string
    port                   = number
    min_capacity           = number
    max_capacity           = number
    instance_count         = number
    instance_class         = string
    vpc_security_group_ids = list(string)
    #iam_roles               = list(string)
    performance_insights    = bool
    enhanced_monitoring     = bool
    backup_retention_period = number
  })
}
