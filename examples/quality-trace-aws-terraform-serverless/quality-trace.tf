resource "aws_ecs_task_definition" "quality-trace" {
  family                   = "quality-trace"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  execution_role_arn       = aws_iam_role.quality-trace_task_execution_role.arn
  container_definitions = jsonencode([
    {
      "name" : "${local.name}",
      "image" : "${local.quality-trace_image}",
      "cpu" : 1024,
      "memory" : 2048,
      "essential" : true,
      "portMappings" : [
        {
          "containerPort" : 11633,
          "hostPort" : 11633,
          "protocol" : "tcp"
        }
      ],
      "environment" : [
        {
          "name" : "QUALITYTRACE_POSTGRES_HOST",
          "value" : "${module.db.db_instance_address}"
        },
        {
          "name" : "QUALITYTRACE_POSTGRES_PORT",
          "value" : "${tostring(module.db.db_instance_port)}"
        },
        {
          "name" : "QUALITYTRACE_POSTGRES_DBNAME",
          "value" : "${module.db.db_instance_name}"
        },
        {
          "name" : "QUALITYTRACE_POSTGRES_USER",
          "value" : "${module.db.db_instance_username}"
        },
        {
          "name" : "QUALITYTRACE_POSTGRES_PASSWORD",
          "value" : "${module.db.db_instance_password}"
        },
        {
          "name" : "QUALITYTRACE_PROVISIONING",
          "value" : base64encode(local.provisioning),
        }
      ],
      "logConfiguration" : {
        "logDriver" : "awslogs",
        "options" : {
          "awslogs-create-group" : "true",
          "awslogs-group" : "/ecs/quality-trace",
          "awslogs-region" : "us-west-2",
          "awslogs-stream-prefix" : "ecs"
        }
      },
    }
  ])
}

resource "aws_ecs_service" "quality-trace-service" {
  name            = "${local.name}-service"
  cluster         = aws_ecs_cluster.quality-trace-cluster.id
  task_definition = aws_ecs_task_definition.quality-trace.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  load_balancer {
    target_group_arn = aws_lb_target_group.quality-trace-tg.arn
    container_name   = "quality-trace"
    container_port   = 11633
  }

  network_configuration {
    subnets          = module.network.private_subnets_ids
    security_groups  = [module.quality-trace_ecs_service_security_group.security_group_id]
    assign_public_ip = false
  }
}

// DATABASE
module "db" {
  source = "terraform-aws-modules/rds/aws"

  identifier = local.name

  engine               = "postgres"
  engine_version       = "14"
  family               = "postgres14"
  major_engine_version = "14"
  instance_class       = "db.t4g.micro"

  allocated_storage     = 20
  max_allocated_storage = 100

  db_name  = local.db_name
  username = local.db_username
  port     = 5432

  create_db_subnet_group = true
  subnet_ids             = module.network.private_subnets_ids

  vpc_security_group_ids = [module.db_security_group.security_group_id]
  deletion_protection    = false

  tags = local.tags
}

module "db_security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = local.name
  description = "PostgreSQL security group"
  vpc_id      = module.network.vpc_id

  ingress_with_cidr_blocks = [
    {
      from_port   = 5432
      to_port     = 5432
      protocol    = "tcp"
      description = "PostgreSQL access from within VPC"
      cidr_blocks = local.vpc_cidr
    },
  ]

  tags = local.tags
}

resource "aws_lb_target_group" "quality-trace-tg" {
  name        = "quality-trace-tg"
  port        = 11633
  protocol    = "HTTP"
  vpc_id      = module.network.vpc_id
  target_type = "ip"
}

resource "aws_lb_listener" "quality-trace-alb-listener" {
  load_balancer_arn = aws_lb.quality-trace-alb.arn
  port              = "11633"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.quality-trace-tg.arn
  }
}
