# iam_policy
resource "aws_iam_policy" "ecs_execution_policy" {
  name   = "${var.env}-${var.sys_name}-ecs-execution-policy"
  policy = file("${path.module}/policies/ecs-execution.json")
}

# iam_role
resource "aws_iam_role" "ecs_execution_role" {
  name = "${var.env}-${var.sys_name}-ecs-execution-role"
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

resource "aws_iam_role_policy_attachment" "ecs_execution_policy_attachment" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = aws_iam_policy.ecs_execution_policy.arn
}
