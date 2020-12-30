variable "GOOGLE_CREDENTIALS" {}
variable "PROJECT_ID" {
  default="chida-201018-bm-test"
}
variable "REGION_ID" {}

provider "google" {
  credentials = var.GOOGLE_CREDENTIALS
  project     = var.PROJECT_ID
  region      = var.REGION_ID
}