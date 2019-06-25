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
      role = "master"
      tier = "backend"
    }

    port {
      port        = 8080
      target_port = 8081
    }
  }
}