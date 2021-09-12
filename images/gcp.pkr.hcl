variables {
  gcp_build_region       = "us-west1-a"
  gcp_build_instance     = "custom-8-16"
  gcp_build_ubuntu_image = "ubuntu-1804-lts"
}

variables {
  gcp_project = "cml-gcp-test"
}

locals {
  tags = {
    environment = var.test ? "test" : "production"
  }
  release_regions = [
    "asia",
    "eu",
    "us",
    "australia-southeast1",
    "southamerica-east1"
  ]
}

source "googlecompute" "source" {
  zone       = var.gcp_build_region
  project_id = var.project

  source_image_family = "ubuntu-os-cloud/${var.gcp_build_ubuntu_image}"

  accelerator_count = 1
  accelerator_type  = "projects/${var.gcp_project}/zones/${var.gcp_build_region}/acceleratorTypes/nvidia-tesla-k80"

  on_host_maintenance = "TERMINATE"

  image_family      = var.test ? var.image_name : "${var.image_name}-test"
  image_labels      = local.tags
  image_name        = "${local.image_name}-${uuid()}"
  image_description = local.image_description

  image_storage_locations = local.release_regions
}
