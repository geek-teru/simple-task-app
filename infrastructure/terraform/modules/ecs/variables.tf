variable "env" {
  type = string
}

variable "sys_name" {
  type = string
}

variable "aws_account_id" {
  type = string
}

variable "backend" {
  type = object({
    cpu    = number
    memory = number
  })
}

variable "db_config" {
  type = object({
    addr                                  = string
    port                                  = string
    user                                  = string
    database_name                         = string
    aws_secretsmanager_secret_version_arn = string
  })
}

variable "log_level" {
  type = string
}

variable "execution_role_arn" {
  type = string
}
