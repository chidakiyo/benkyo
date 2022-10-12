resource "google_dns_managed_zone" "a-zone" {
  name        = "a-zone"
  dns_name    = var.api_domain
  description = "a zone."
}

resource "google_dns_managed_zone" "b-zone" {
  name        = "b-zone"
  dns_name    = "sub.${var.api_domain}"
  description = "b zone."
}

resource "google_dns_record_set" "zone_subzone_ns" {
  name = "sub.${google_dns_managed_zone.a-zone.dns_name}"
  type = "NS"
  ttl  = 300

  managed_zone = google_dns_managed_zone.a-zone.name

  rrdatas = google_dns_managed_zone.b-zone.name_servers
}

#resource "google_dns_record_set" "hoge_root_ns" {
#  name         = "hoge-ns"
#  managed_zone = google_dns_managed_zone.a-zone.name
#  type         = "NS"
#  ttl          = 300
#
#  rrdatas = var.name_servers
#}