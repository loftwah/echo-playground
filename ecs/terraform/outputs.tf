output "ecs_cluster_id" {
  description = "The ID of the ECS Cluster"
  value       = aws_ecs_cluster.echo_playground.id
}

output "ecs_cluster_arn" {
  description = "The ARN of the ECS Cluster"
  value       = aws_ecs_cluster.echo_playground.arn
}

output "ecs_task_definition_arn" {
  description = "The ARN of the ECS Task Definition"
  value       = aws_ecs_task_definition.echo_task.arn
}

output "ecs_service_name" {
  description = "The name of the ECS Service"
  value       = aws_ecs_service.echo_service.name
}

output "ecs_service_id" {
  description = "The ID of the ECS Service"
  value       = aws_ecs_service.echo_service.id
}

output "alb_dns_name" {
  description = "The DNS name of the Application Load Balancer"
  value       = aws_lb.echo_alb.dns_name
}

output "alb_arn" {
  description = "The ARN of the Application Load Balancer"
  value       = aws_lb.echo_alb.arn
}

output "target_group_arn" {
  description = "The ARN of the Target Group"
  value       = aws_lb_target_group.echo_tg.arn
}

output "security_group_id" {
  description = "The ID of the ECS Security Group"
  value       = aws_security_group.echo_playground_sg.id
}

output "alb_security_group_id" {
  description = "The ID of the ALB Security Group"
  value       = aws_security_group.alb_sg.id
}
