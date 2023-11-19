resource "aws_security_group" "echo_playground_sg" {
  name        = "echo-playground-sg"
  description = "Security group for Echo Playground"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port   = 1323
    to_port     = 1323
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_ecs_cluster" "echo_playground" {
  name = "echo-playground"
}

resource "aws_ecs_task_definition" "echo_task" {
  family                   = "echo-playground"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = "arn:aws:iam::614442032670:role/ecsTaskExecutionRole"
  cpu                      = "256"
  memory                   = "512"

  container_definitions = jsonencode([{
    name  = "echo-playground",
    image = "614442032670.dkr.ecr.ap-southeast-2.amazonaws.com/echo-playground:latest",
    cpu   = 128,
    memory = 256,
    essential = true,
    portMappings = [{
      containerPort = 1323
    }]
  }])
}

resource "aws_ecs_service" "echo_service" {
  name            = "echo-playground"
  cluster         = aws_ecs_cluster.echo_playground.id
  task_definition = aws_ecs_task_definition.echo_task.arn
  launch_type     = "FARGATE"
  desired_count   = 1

  network_configuration {
    subnets         = data.aws_subnets.default.ids
    security_groups = [aws_security_group.echo_playground_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.echo_tg.arn
    container_name   = "echo-playground"
    container_port   = 1323
  }

  depends_on = [
    aws_lb_listener.echo_listener,
  ]
}

