output "api_endpoint" {
  value       = "${aws_apigatewayv2_stage.lambda.invoke_url}/hello"
  description = "The API endpoint"
}

output "quality-trace_url" {
  value       = "http://${aws_lb.quality-trace-alb.dns_name}:11633"
  description = "Qualitytrace public URL"
}

output "jaeger_ui_url" {
  value       = "http://${aws_lb.quality-trace-alb.dns_name}:16686"
  description = "Jaeger public URL"
}

output "internal_jaeger_api_url" {
  value       = "${aws_lb.internal_quality-trace_alb.dns_name}:16685"
  description = "Jaeger internal API URL"
}
