resource "aws_ssm_parameter" "ecs_backend_release" {
  name        = "/${var.env}/ecs/${var.sys_name}/release"
  description = "${var.env}-${var.sys_name} current version"
  type        = "String"
  value       = ""
}
