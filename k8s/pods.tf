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
      container {
        // TODO: push to k8s during the deployment - based on the provisioned app credentials
        image = "gcr.io/summer-camp-244710/summer_app:lastest"
        name  = "master"

        port {
          container_port = 8081
        }

        resources {
          requests {
            cpu    = "100m"
            memory = "100Mi"
          }
        }
      }
    }
  }
}