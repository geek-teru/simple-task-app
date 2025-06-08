# iam_role
resource "aws_iam_role" "ecs_backend_api_role" {
  name = "${var.env}-${var.sys_name}-ecs-backend-api-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}
