provider_installation {
  dev_overrides {
    "hashicorp/aws" = "/terraform/provider-mirror/registry.terraform.io/C2Devel/rockitcloud/25.1.1/linux_amd64"
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
