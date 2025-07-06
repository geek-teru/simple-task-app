resource "aws_alb_target_group" "target_group" {
  name                 = "${local.full_service_name}-tg"
  port                 = 8080
  deregistration_delay = 15
  protocol             = "HTTP"
  protocol_version     = "HTTP1"
  target_type          = "ip"
  vpc_id               = var.alb.vpc_id

  health_check {
    port = 8080
    path = "/healthcheck"
  }

  depends_on = [
    aws_alb.public_alb
  ]

  tags = {
    Name = "${local.full_service_name}-tg"
  }
}
