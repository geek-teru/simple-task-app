output "rds_cluster_id" {
  value = aws_rds_cluster.cluster.cluster_identifier
}

output "rds_endpoint" {
  value = aws_rds_cluster.cluster.endpoint
}

output "rds_port" {
  value = aws_rds_cluster.cluster.port
}

output "rds_database_name" {
  value = aws_rds_cluster.cluster.database_name
}

output "rds_master_username" {
  value = aws_rds_cluster.cluster.master_username
}

output "rds_master_password_secret_arn" {
  value = aws_rds_cluster.cluster.master_user_secret[0].secret_arn
}

# Loggroups for RDS
# output "rds_log_group_error" {
#   value = aws_cloudwatch_log_group.rds_log_group_error.name
# }

# output "rds_log_group_audit" {
#   value = aws_cloudwatch_log_group.rds_log_group_audit.name
# }

# output "rds_log_group_slowquery" {
#   value = aws_cloudwatch_log_group.rds_log_group_slowquery.name
# }
