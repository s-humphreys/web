resource "google_project_iam_member" "cicd" {
  for_each = toset([
    "roles/appengine.deployer",
    "roles/appengine.serviceAdmin",
    "roles/cloudbuild.builds.editor",
    "roles/storage.objectAdmin",
  ])

  project = google_project.project.project_id
  role    = each.key
  member  = "serviceAccount:${google_service_account.cicd_app_enginer_deployer.email}"
}
