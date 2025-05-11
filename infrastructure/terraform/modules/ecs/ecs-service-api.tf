## ECS Service Definition
data "aws_ssm_parameter" "ecs_backend_release" {
  name = "/${var.env}/ecs/${var.sys_name}-backend/release"
}

data "template_file" "ecs_backend_task_template" {
  template = file("${path.module}/task-definitions/backend.json")
  vars = {
    ecr_uri                               = aws_ecr_repository.ecr_repository.repository_url
    ecr_image_tag                         = data.aws_ssm_parameter.ecs_backend_release.value
    system_env                            = var.env
    log_level                             = var.log_level
    db_addr                               = var.db_config.addr
    db_port                               = var.db_config.port
    db_user                               = var.db_config.user
    db_name                               = var.db_config.database_name
    aws_secretsmanager_secret_version_arn = var.db_config.aws_secretsmanager_secret_version_arn
  }
}

resource "aws_ecs_task_definition" "ecs_backend_task" {
  family                   = "${var.env}-${var.sys_name}-backend"
  container_definitions    = data.template_file.ecs_backend_task_template.rendered
  requires_compatibilities = ["FARGATE"]
  skip_destroy             = false
  # task_role_arn            = var.task_role_arn #optional
  execution_role_arn = var.execution_role_arn
  network_mode       = "awsvpc"
  cpu                = var.backend.cpu
  memory             = var.backend.memory

  tags = {
    Environment = var.env
    Name        = "${var.env}-${var.sys_name}-backend"
  }
}

# ECS Service
# data "aws_ecs_task_definition" "ecs_backend_task" {
#   task_definition = aws_ecs_task_definition.ecs_backend_task.family
# }

# resource "aws_ecs_service" "ecs_backend_service" {
#   name                               = "${var.env}-${var.sys_name}-backend"
#   cluster                            = var.api_service.cluster_id
#   task_definition                    = data.aws_ecs_task_definition.ecs_backend_task.arn
#   desired_count                      = var.api_service.desired_count
#   deployment_minimum_healthy_percent = 100
#   deployment_maximum_percent         = 200
#   launch_type                        = "FARGATE"
#   platform_version                   = "LATEST"
#   enable_execute_command             = var.api_service.enable_execute_command

#   network_configuration {
#     assign_public_ip = false
#     subnets          = var.api_service.subnet_ids
#     security_groups  = var.api_service.security_group_ids
#   }

#   load_balancer {
#     target_group_arn = var.api_service.lb_target_group_arn
#     container_name   = "backend"
#     container_port   = 8080
#   }

#   deployment_circuit_breaker {
#     enable   = true
#     rollback = true
#   }

#   tags = {
#     Name = "${var.env}-${var.sys_name}-ecs-service-backend"
#   }
# }

# resource "aws_cloudwatch_log_group" "logs_backend" {
#   name              = "/app/backend"
#   retention_in_days = 14

#   tags = {
#     Name = "logs_backend"
#   }
# }

