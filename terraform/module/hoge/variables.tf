variable "gcp_project" {}
variable "region" {}

variable "api_domain" {}

variable "name_servers" {
  type = list(string)
  default = []
}
