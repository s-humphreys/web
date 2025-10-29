resource "google_storage_bucket" "shumphreys-tf-state" {
  project                     = google_project.project.project_id
  name                        = "shumphreys-tf-state"
  location                    = var.gcp_region
  force_destroy               = var.force_destroy
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
}
