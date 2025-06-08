output "alb_backend_api_sg_id" {
  value = aws_security_group.alb_backend_api_sg.id
}

output "ecs_backend_api_sg_id" {
  value = aws_security_group.ecs_backend_api_sg.id
}

output "rds_sg_id" {
  value = aws_security_group.rds_sg.id
}
