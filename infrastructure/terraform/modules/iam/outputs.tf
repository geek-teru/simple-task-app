output "ecs_execution_role_arn" {
  value = aws_iam_role.ecs_execution_role.arn
}

output "ecs_ecs_backend_api_role_arn" {
  value = aws_iam_role.ecs_backend_api_role.arn
}
