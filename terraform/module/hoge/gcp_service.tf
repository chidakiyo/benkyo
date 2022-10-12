locals {
  services = [
    "dns.googleapis.com",
  ]
}
resource "google_project_service" "project" {
  count = length(local.services)

  project = var.gcp_project
  service = local.services[count.index]
}
