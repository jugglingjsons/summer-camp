resource "kubernetes_service" "summer-service" {
  metadata {
    name = "summer-service"

    labels {
      app  = "summer-service"
      role = "master"
      tier = "backend"
    }
  }

  spec {
    selector {
      app  = "summer-service"
    }

    port {
      port        = 80
      target_port = 8081
    }
    type = "LoadBalancer"
  }
}

output "lb_ip" {
  value = "${kubernetes_service.summer-service.load_balancer_ingress.0.ip}"
}