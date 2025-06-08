# Security Groups for ALB

resource "aws_security_group" "alb_backend_api_sg" {
  name   = "${var.env}-${var.sys_name}-alb-backend-api-sg"
  vpc_id = var.vpc_id

  tags = {
    Name        = "${var.env}-${var.sys_name}-alb-backend-api-sg"
    Environment = var.env
  }
}

resource "aws_vpc_security_group_ingress_rule" "alb_backend_api_sg_ingress_https" {
  from_port         = 443
  to_port           = 443
  ip_protocol       = "tcp"
  cidr_ipv4         = "0.0.0.0/0"
  security_group_id = aws_security_group.alb_backend_api_sg.id
}

# Allow HTTP traffic from anywhere for redirecting to HTTPS
resource "aws_vpc_security_group_ingress_rule" "alb_backend_api_sg_ingress_http" {
  from_port         = 80
  to_port           = 80
  ip_protocol       = "tcp"
  cidr_ipv4         = "0.0.0.0/0"
  security_group_id = aws_security_group.alb_backend_api_sg.id
}

resource "aws_vpc_security_group_egress_rule" "alb_backend_api_sg_egress" {
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
  security_group_id = aws_security_group.alb_backend_api_sg.id
}
