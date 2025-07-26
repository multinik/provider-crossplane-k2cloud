provider_installation {
  filesystem_mirror {
    path    = "/terraform/provider-mirror"
    include = ["*/*"]
    readonly = true 
  }
  dev_overrides {
    "hashicorp/aws" = "/terraform/provider-mirror/registry.terraform.io/C2Devel/rockitcloud/25.1.1/linux_amd64"
  }
  direct {
    exclude = ["*/*"]
  }
}
