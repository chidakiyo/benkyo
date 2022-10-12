module "hoge" {
  source = "../../module/hoge"

  gcp_project = var.gcp_project
  region      = var.region

  api_domain = "chida.dev."

}
