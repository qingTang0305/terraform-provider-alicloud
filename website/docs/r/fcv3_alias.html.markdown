---
subcategory: "Function Compute Service V3 (FCV3)"
layout: "alicloud"
page_title: "Alicloud: alicloud_fcv3_alias"
description: |-
  Provides a Alicloud FCV3 Alias resource.
---

# alicloud_fcv3_alias

Provides a FCV3 Alias resource.

Alias for functions.

For information about FCV3 Alias and how to use it, see [What is Alias](https://www.alibabacloud.com/help/en/functioncompute/developer-reference/api-fc-2023-03-30-createalias).

-> **NOTE:** Available since v1.228.0.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_fcv3_alias&exampleId=964a205d-f3d9-3bda-b718-c780deee06a81557d045&activeTab=example&spm=docs.r.fcv3_alias.0.964a205df3" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "function_name" {
  default = "flask-3xdg"
}


resource "alicloud_fcv3_alias" "default" {
  version_id    = "1"
  function_name = var.function_name
  description   = "create alias"
  alias_name    = var.name
  additional_version_weight = {
    "2" = "0.5"
  }
}
```

## Argument Reference

The following arguments are supported:
* `additional_version_weight` - (Optional, Map) Grayscale version
* `alias_name` - (Optional, ForceNew, Computed) Function Alias
* `description` - (Optional) Description
* `function_name` - (Required, ForceNew) Function Name
* `version_id` - (Optional) The version that the alias points

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<function_name>:<alias_name>`.
* `create_time` - The creation time of the resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Alias.
* `delete` - (Defaults to 5 mins) Used when delete the Alias.
* `update` - (Defaults to 5 mins) Used when update the Alias.

## Import

FCV3 Alias can be imported using the id, e.g.

```shell
$ terraform import alicloud_fcv3_alias.example <function_name>:<alias_name>
```