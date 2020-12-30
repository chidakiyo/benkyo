resource "google_project_service" "service" {
  for_each = toset([
  ])

  service = each.key

  project            = var.PROJECT_ID
  disable_on_destroy = false
}