packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.1"
      source  = "github.com/hashicorp/amazon"
    }
    googlecompute = {
      version = ">= 0.0.1"
      source  = "github.com/hashicorp/googlecompute"
    }
  }
}

variable "test" {
  type    = bool
  default = false
}

variables {
  image_name        = "iterative-cml"
  image_description = "CML (Continuous Machine Learning) Ubuntu 18.04"
}

build {
  sources = ["source.amazon-ebs.source", "sources.googlecompute.source"]
  provisioner "shell" {
    script = "./setup.sh"
  }
}
