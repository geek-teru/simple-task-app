variable "env" {
  type = string
}

variable "sys_name" {
  type = string
}

variable "service_name" {
  type = string
}

variable "certificate_arn" {
  type = string
}

variable "alb" {
  type = object({
    subnet_group_ids       = list(string)
    vpc_security_group_ids = list(string)
    system_logs_bucket     = string
    vpc_id                 = string
    enable_access_log      = bool
  })
}
