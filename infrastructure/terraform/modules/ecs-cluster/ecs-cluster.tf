resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${var.env}-${var.sys_name}"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
  tags = {
    Environment = var.env
  }
}


