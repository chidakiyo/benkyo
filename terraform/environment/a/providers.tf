terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.23.0"
    }

    postgresql = {
      source  = "cyrilgdn/postgresql"
      version = "1.16.0"
    }
  }

  required_version = "~> 1.2.2"
}

provider "google" {
  project = var.gcp_project
  region  = var.region
}

provider "google-beta" {
  project = var.gcp_project
  region  = var.region
}
