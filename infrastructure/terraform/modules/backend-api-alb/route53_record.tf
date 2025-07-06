resource "aws_route53_record" "route53_record" {
  zone_id = "Z00821753QKFELRQIFTW2"
  name    = "${local.full_service_name}.geek-teru.com"
  type    = "A"

  alias {
    name                   = aws_alb.public_alb.dns_name
    zone_id                = aws_alb.public_alb.zone_id
    evaluate_target_health = true
  }
}
