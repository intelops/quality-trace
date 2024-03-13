output "quality-trace_url" {
  value       = "http://${aws_lb.quality-trace_alb.dns_name}:11633"
  description = Qualitytrace public URL"
}
