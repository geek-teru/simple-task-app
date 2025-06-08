# Security Groups for RDS (Aurora PostgreSQL)

resource "aws_security_group" "rds_sg" {
  name   = "${var.env}-${var.sys_name}-rds-sg"
  vpc_id = var.vpc_id

  tags = {
    Name        = "${var.env}-${var.sys_name}-rds-sg"
    Environment = var.env
  }
}

resource "aws_vpc_security_group_ingress_rule" "redshift_sg_ingress_from_ecs" {
  from_port                    = 5432
  to_port                      = 5432
  ip_protocol                  = "tcp"
  referenced_security_group_id = aws_security_group.rds_sg.id
  security_group_id            = aws_security_group.rds_sg.id
}

resource "aws_vpc_security_group_egress_rule" "rds_sg_egress" {
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
  security_group_id = aws_security_group.rds_sg.id
}
