# Security Group for ECS Service

resource "aws_security_group" "ecs_backend_api_sg" {
  name   = "${var.env}-${var.sys_name}-ecs-sg"
  vpc_id = var.vpc_id

  tags = {
    Name        = "${var.env}-${var.sys_name}-ecs-sg"
    Environment = var.env
  }
}

resource "aws_vpc_security_group_ingress_rule" "ecs_backend_api_ingress" {
  from_port                    = 8080
  to_port                      = 8080
  ip_protocol                  = "tcp"
  referenced_security_group_id = aws_security_group.ecs_backend_api_sg.id
  security_group_id            = aws_security_group.ecs_backend_api_sg.id
}

resource "aws_vpc_security_group_egress_rule" "ecs_backend_api_egress" {
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
  security_group_id = aws_security_group.ecs_backend_api_sg.id
}
