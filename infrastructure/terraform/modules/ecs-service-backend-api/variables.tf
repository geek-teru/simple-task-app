variable "env" {
  type = string
}
variable "sys_name" {
  type = string
}

variable "aws_account_id" {
  type = string
}

variable "service_name" {
  type = string
}

variable "rds_secret_arn" {
  type = string
}

variable "ecs_service" {
  type = object({
    ecs_cluster_id         = string
    desired_count          = number
    subnet_ids             = list(string)
    security_group_ids     = list(string)
    enable_execute_command = bool
    lb_target_group_arn    = string
  })
}

variable "ecs_task_definition" {
  type = object({
    cpu                    = number
    memory                 = number
    environment            = map(string)
    ecs_execution_role_arn = string
    ecs_task_role_arn      = string
  })
}
