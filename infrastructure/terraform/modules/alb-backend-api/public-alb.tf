locals {
  full_service_name = "${var.env}-${var.sys_name}-${var.service_name}"
}

resource "aws_alb" "public_alb" {
  name                       = "${local.full_service_name}-alb"
  subnets                    = var.alb.subnet_group_ids
  security_groups            = var.alb.vpc_security_group_ids
  internal                   = false
  enable_deletion_protection = false
  load_balancer_type         = "application"

  access_logs {
    bucket  = var.alb.system_logs_bucket
    prefix  = "${var.env}/ALB/${local.full_service_name}-public-alb"
    enabled = var.alb.enable_access_log
  }

  tags = {
    Name = "${local.full_service_name}-public-alb"
  }

}

# HTTP
resource "aws_alb_listener" "public_alb_listener_http" {
  port              = 80
  protocol          = "HTTP"
  load_balancer_arn = aws_alb.public_alb.arn

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      status_code  = "503"
      message_body = "Service Unavailable"
    }
  }
}

# HTTPS 
resource "aws_alb_listener" "public_alb_listener_https" {
  port              = 443
  protocol          = "HTTPS"
  load_balancer_arn = aws_alb.public_alb.arn

  ssl_policy      = "ELBSecurityPolicy-TLS13-1-2-2021-06"
  certificate_arn = var.certificate_arn

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      status_code  = "503"
      message_body = "Service Unavailable"
    }
  }
}
