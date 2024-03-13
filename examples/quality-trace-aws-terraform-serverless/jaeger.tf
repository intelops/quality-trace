resource "aws_ecs_task_definition" "jaeger" {
  family                   = "${local.name}_jaeger"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  execution_role_arn       = aws_iam_role.quality-trace_task_execution_role.arn
  container_definitions = jsonencode([
    {
      "name" : "jaeger",
      "image" : "jaegertracing/all-in-one:1.42",
      "cpu" : 1024,
      "memory" : 2048,
      "essential" : true,
      "environment" : [{
        "name" : "COLLECTOR_OTLP_ENABLED",
        "value" : "true"
      }],
      "portMappings" : [
        {
          "hostPort" : 14269,
          "protocol" : "tcp",
          "containerPort" : 14269
        },
        {
          "hostPort" : 14268,
          "protocol" : "tcp",
          "containerPort" : 14268
        },
        {
          "hostPort" : 6832,
          "protocol" : "udp",
          "containerPort" : 6832
        },
        {
          "hostPort" : 6831,
          "protocol" : "udp",
          "containerPort" : 6831
        },
        {
          "hostPort" : 5775,
          "protocol" : "udp",
          "containerPort" : 5775
        },
        {
          "hostPort" : 14250,
          "protocol" : "tcp",
          "containerPort" : 14250
        },
        {
          "hostPort" : 16685,
          "protocol" : "tcp",
          "containerPort" : 16685
        },
        {
          "hostPort" : 5778,
          "protocol" : "tcp",
          "containerPort" : 5778
        },
        {
          "hostPort" : 16686,
          "protocol" : "tcp",
          "containerPort" : 16686
        },
        {
          "hostPort" : 9411,
          "protocol" : "tcp",
          "containerPort" : 9411
        },
        {
          "hostPort" : 4318,
          "protocol" : "tcp",
          "containerPort" : 4318
        },
        {
          "hostPort" : 4317,
          "protocol" : "tcp",
          "containerPort" : 4317
        }
      ],
      "logConfiguration" : {
        "logDriver" : "awslogs",
        "options" : {
          "awslogs-create-group" : "true",
          "awslogs-group" : "/ecs/jaeger",
          "awslogs-region" : "us-west-2",
          "awslogs-stream-prefix" : "ecs"
        }
      },
    }
  ])
}

resource "aws_ecs_service" "jaeger_service" {
  name            = "jaeger-service"
  cluster         = aws_ecs_cluster.quality-trace-cluster.id
  task_definition = aws_ecs_task_definition.jaeger.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  load_balancer {
    target_group_arn = aws_lb_target_group.quality-trace-jaeger-tg.arn
    container_name   = "jaeger"
    container_port   = 16686
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.quality-trace-jaeger-api-tg.arn
    container_name   = "jaeger"
    container_port   = 16685
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.quality-trace-jaeger-collector-tg.arn
    container_name   = "jaeger"
    container_port   = 4318
  }

  network_configuration {
    subnets          = module.network.private_subnets_ids
    security_groups  = [module.quality-trace_ecs_service_security_group.security_group_id]
    assign_public_ip = false
  }
}

resource "aws_lb_target_group" "quality-trace-jaeger-api-tg" {
  name             = "quality-trace-jaeger-api-tg"
  port             = 16685
  protocol         = "HTTP"
  protocol_version = "GRPC"
  vpc_id           = module.network.vpc_id
  target_type      = "ip"
}

resource "aws_lb_target_group" "quality-trace-jaeger-collector-tg" {
  name        = "quality-trace-jaeger-collector-tg"
  port        = 4318
  protocol    = "HTTP"
  vpc_id      = module.network.vpc_id
  target_type = "ip"

  health_check {
    path              = "/"
    port              = "4318"
    protocol          = "HTTP"
    healthy_threshold = 2
    matcher           = "200-499"
  }
}

resource "aws_lb_target_group" "quality-trace-jaeger-tg" {
  name        = "quality-trace-jaeger-tg"
  port        = 16686
  protocol    = "HTTP"
  vpc_id      = module.network.vpc_id
  target_type = "ip"
}

resource "aws_lb_listener" "quality-trace-jaeger-alb-listener" {
  load_balancer_arn = aws_lb.quality-trace-alb.arn
  port              = "16686"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.quality-trace-jaeger-tg.arn
  }
}

resource "aws_lb_listener" "quality-trace-jaeger-collector-alb-listener" {
  load_balancer_arn = aws_lb.internal_quality-trace_alb.arn
  port              = "4318"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.quality-trace-jaeger-collector-tg.arn
  }
}

resource "aws_lb_listener" "quality-trace-jaeger-api-alb-listener" {
  load_balancer_arn = aws_lb.internal_quality-trace_alb.arn
  port              = "16685"
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.cert.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.quality-trace-jaeger-api-tg.arn
  }
}

resource "tls_private_key" "quality-trace_private_key" {
  algorithm = "RSA"
}

resource "tls_self_signed_cert" "quality-trace_self_signed_cert" {
  private_key_pem = tls_private_key.quality-trace_private_key.private_key_pem

  subject {
    common_name  = "quality-trace.com"
    organization = "Qualitytrace"
  }

  validity_period_hours = 720

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "aws_acm_certificate" "cert" {
  private_key      = tls_private_key.quality-trace_private_key.private_key_pem
  certificate_body = tls_self_signed_cert.quality-traceself_signed_cert.cert_pem
}
