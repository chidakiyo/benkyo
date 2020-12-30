resource "google_bigquery_dataset" "dataset" {
  dataset_id                  = "auditlog"
  friendly_name               = "auditlog"
  description                 = "your app's audit log"
  location                    = "asia-northeast1"
//  default_table_expiration_ms = 3600000 // Optional

//  labels = {
//    env = "default"
//  }
//
//  access {
//    role          = "OWNER"
//    user_by_email = google_service_account.bqowner.email
//  }
//
//  access {
//    role   = "READER"
//    domain = "hashicorp.com"
//  }

}

//resource "google_service_account" "bqowner" {
//  account_id = "bqowner"
//}