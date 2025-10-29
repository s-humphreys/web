resource "google_project" "project" {
  name                = "Website for shumphreys"
  project_id          = var.gcp_project
  billing_account     = var.gcp_billing_account
  auto_create_network = false
}

resource "google_project_service" "service" {
  for_each = var.gcp_apis

  project            = google_project.project.project_id
  service            = "${each.value}.googleapis.com"
  disable_on_destroy = false
}
