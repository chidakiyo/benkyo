resource "google_logging_project_sink" "audit-log-sink" {
  name = "audit-logging-sink"
  destination = "bigquery.googleapis.com/projects/${var.PROJECT_ID}/datasets/${google_bigquery_dataset.dataset.dataset_id}"
  filter = "resource.type = cloud_run_revision AND jsonPayload.mark = Target"
//  unique_writer_identity = true
  depends_on = [
    google_bigquery_dataset.dataset
  ]
}