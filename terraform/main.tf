data "google_container_engine_versions" "gke_versions" {
  location = var.zone
}

resource "google_container_cluster" "primary" {
  name               = "monorepo-base-staging"
  location           = var.zone
  initial_node_count = 2
  min_master_version = data.google_container_engine_versions.gke_versions.latest_master_version

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/servicecontrol",
      "https://www.googleapis.com/auth/service.management.readonly",
      "https://www.googleapis.com/auth/trace.append",
    ]

    machine_type = "n1-standard-1"
    image_type   = "COS"
  }

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}
