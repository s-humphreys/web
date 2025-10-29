resource "google_app_engine_application" "app" {
  project     = google_project.project.project_id
  location_id = var.gcp_region_appengine
}
