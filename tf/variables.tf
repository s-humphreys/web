variable "gcp_project" {
  description = "The GCP project ID"
  type        = string
  default     = "samhumphreys-website"
}

variable "gcp_billing_account" {
  description = "The GCP billing account ID"
  type        = string
}

variable "gcp_region" {
  description = "The GCP region"
  type        = string
  default     = "europe-west1"
}

variable "gcp_region_appengine" {
  description = "The GCP AppEngine region"
  type        = string
  default     = "europe-west"
}

variable "gcp_apis" {
  description = "The GCP APIs to enable within the project"
  type        = set(string)
  default     = ["appengine"]
}

variable "force_destroy" {
  description = "Whether to force destroy resources or not"
  type        = bool
  default     = false
}
