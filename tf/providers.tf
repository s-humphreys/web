terraform {
  required_version = "1.13.4"

  backend "gcs" {
    bucket = "shumphreys-tf-state"
    prefix = "terraform/state"
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 7.9.0"
    }
  }
}

provider "google" {
  project = var.gcp_project
}
