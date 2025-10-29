resource "google_project_iam_member" "app_engine_deployer" {
  project = google_project.project.project_id
  role    = "roles/appengine.deployer"
  member  = "serviceAccount:${google_service_account.cicd_app_enginer_deployer.email}"
}

resource "google_project_iam_member" "app_engine_service_admin" {
  project = google_project.project.project_id
  role    = "roles/appengine.serviceAdmin"
  member  = "serviceAccount:${google_service_account.cicd_app_enginer_deployer.email}"
}

resource "google_project_iam_member" "cloud_build_editor" {
  project = google_project.project.project_id
  role    = "roles/cloudbuild.builds.editor"
  member  = "serviceAccount:${google_service_account.cicd_app_enginer_deployer.email}"
}

resource "google_project_iam_member" "storage_object_admin" {
  project = google_project.project.project_id
  role    = "roles/storage.objectAdmin"
  member  = "serviceAccount:${google_service_account.cicd_app_enginer_deployer.email}"
}

resource "google_service_account_iam_member" "app_enginer_gsa_user" {
  service_account_id = "projects/${google_project.project.project_id}/serviceAccounts/${google_project.project.project_id}@appspot.gserviceaccount.com"
  role               = "roles/iam.serviceAccountUser"
  member             = "serviceAccount:${google_service_account.cicd_app_enginer_deployer.email}"
}
