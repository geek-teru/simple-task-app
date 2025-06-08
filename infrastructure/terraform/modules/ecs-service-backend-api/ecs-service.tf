locals {
  full_service_name = "${var.env}-${var.sys_name}-${var.service_name}"
}

data "aws_secretsmanager_secret" "rds_master_password_secret" {
  arn = var.rds_secret_arn
}

data "aws_secretsmanager_secret_version" "rds_master_password_secret_version" {
  secret_id = data.aws_secretsmanager_secret.rds_master_password_secret.id
}

data "aws_ssm_parameter" "ecs-service-release" {
  name = "/ecs/${local.full_service_name}/release"
}

# Logs Group
resource "aws_cloudwatch_log_group" "logs" {
  name              = "/ecs/${local.full_service_name}"
  retention_in_days = 14
}

# ---------------------------------------
# ECS Task Definition
# ---------------------------------------
data "template_file" "ecs_task_template" {
  template = file("${path.module}/task-definitions/${var.service_name}.json")
  vars = {
    service_name                          = local.full_service_name
    ecr_uri                               = "${var.aws_account_id}.dkr.ecr.ap-northeast-1.amazonaws.com/${local.full_service_name}"
    ecr_image_tag                         = data.aws_ssm_parameter.ecs-service-release.value
    system_env                            = var.env
    log_level                             = var.ecs_task_definition.environment.log_level
    db_addr                               = var.ecs_task_definition.environment.db_addr
    db_user                               = var.ecs_task_definition.environment.db_user
    db_name                               = var.ecs_task_definition.environment.db_name
    db_port                               = var.ecs_task_definition.environment.db_port
    aws_secretsmanager_secret_version_arn = data.aws_secretsmanager_secret_version.rds_master_password_secret_version.arn
    log_group_name                        = aws_cloudwatch_log_group.logs.name
  }
  depends_on = [
    aws_cloudwatch_log_group.logs
  ]
}

resource "aws_ecs_task_definition" "ecs_task" {
  family                   = local.full_service_name
  container_definitions    = data.template_file.ecs_task_template.rendered
  requires_compatibilities = ["FARGATE"]
  skip_destroy             = false
  task_role_arn            = var.ecs_task_definition.ecs_task_role_arn
  execution_role_arn       = var.ecs_task_definition.ecs_execution_role_arn
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition.cpu
  memory                   = var.ecs_task_definition.memory
  track_latest             = true
  depends_on = [
    aws_cloudwatch_log_group.logs
  ]
}

# ---------------------------------------
# ECS Service
# ---------------------------------------
data "aws_ecs_task_definition" "ecs_task" {
  task_definition = aws_ecs_task_definition.ecs_task.family
}

resource "aws_ecs_service" "ecs_service" {
  name                               = local.full_service_name
  cluster                            = var.ecs_service.ecs_cluster_id
  task_definition                    = data.aws_ecs_task_definition.ecs_task.arn
  desired_count                      = var.ecs_service.desired_count
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200
  launch_type                        = "FARGATE"
  platform_version                   = "LATEST"
  enable_execute_command             = var.ecs_service.enable_execute_command

  network_configuration {
    assign_public_ip = false
    subnets          = var.ecs_service.subnet_ids
    security_groups  = var.ecs_service.security_group_ids
  }

  load_balancer {
    target_group_arn = var.ecs_service.lb_target_group_arn
    container_name   = local.full_service_name
    container_port   = 8080
  }

  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }

  tags = {
    Name = local.full_service_name
  }
  depends_on = [
    aws_cloudwatch_log_group.logs,
    aws_ecs_task_definition.ecs_task
  ]
}
