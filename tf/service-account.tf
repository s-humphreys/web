locals {
  gsa_name = "cicd-app-engine-deployer"
}

resource "google_service_account" "cicd_app_enginer_deployer" {
  project      = google_project.project.project_id
  account_id   = local.gsa_name
  display_name = local.gsa_name
}

resource "google_service_account_key" "cicd_app_enginer_deployer_key" {
  service_account_id = google_service_account.cicd_app_enginer_deployer.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}
