resource "kubernetes_replication_controller" "summer-service" {
  metadata {
    name = "summer-service"
    labels {
      app  = "summer-service"
      role = "master"
      tier = "backend"
    }
  }

  spec {
    replicas = 2
    selector = {
      app  = "summer-service"
      role = "master"
      tier = "backend"
    }
    template {
      metadata {
        labels {
          summer-service = "summer-service"
          app = "summer-service"
          role = "master"
          tier = "backend"
        }
      }
      spec {
        container {
          image = "gcr.io/summer-camp-244710/summer_app:latest"
          name  = "master"
          port {
            container_port = 8081
          }
          liveness_probe {
            http_get {
              path = "/health"
              port = 8081
            }
            initial_delay_seconds = 5
            period_seconds        = 10
          }
          resources {

            requests {
              cpu    = "50m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}