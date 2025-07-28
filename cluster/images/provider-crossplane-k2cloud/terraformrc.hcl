provider_installation {
  network_mirror {
  url     = "file:///terraform/provider-mirror"
  include = ["registry.terraform.io/hashicorp/aws"]
  }
  dev_overrides {
    "hashicorp/aws" = "/terraform/provider-mirror/registry.terraform.io/C2Devel/rockitcloud/25.1.1/linux_amd64/"
  }
  filesystem_mirror {
    path    = "/terraform/provider-mirror"
    include = ["*/*"]
    readonly = true 
  }
  direct {
    exclude = ["*/*"]
  }
}
