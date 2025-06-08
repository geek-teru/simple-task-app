# HTTP
resource "aws_alb_listener_rule" "http_listener_rule" {
  listener_arn = aws_alb_listener.public_alb_listener_http.arn
  priority     = 100

  action {
    type = "redirect"
    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }

  condition {
    path_pattern {
      values = ["/*"]
    }
  }

  condition {
    host_header {
      values = ["*.geek-teru.com"]
    }
  }

  tags = {
    Name = "HTTP"
  }
}

# HTTPS
resource "aws_alb_listener_rule" "https_listener_rule" {
  listener_arn = aws_alb_listener.public_alb_listener_https.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.target_group.arn
  }

  condition {
    path_pattern {
      values = ["/*"]
    }
  }

  condition {
    host_header {
      values = ["*.geek-teru.com"]
    }
  }

  tags = {
    Name = "HTTPS"
  }
}
