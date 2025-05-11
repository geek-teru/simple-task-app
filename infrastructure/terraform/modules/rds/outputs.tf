output "rds_cluster" {
  value = aws_rds_cluster.cluster
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
