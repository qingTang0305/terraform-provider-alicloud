---
subcategory: "Api Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_api_gateway_group"
sidebar_current: "docs-alicloud-resource-api-gateway-group"
description: |-
  Provides a Alicloud Api Gateway Group Resource.
---

# alicloud_api_gateway_group

Provides an api group resource.To create an API, you must firstly create a group which is a basic attribute of the API.

For information about Api Gateway Group and how to use it, see [Create An Api Group](https://www.alibabacloud.com/help/en/api-gateway/latest/api-cloudapi-2016-07-14-createapigroup)

-> **NOTE:** Available since v1.19.0.

-> **NOTE:** Terraform will auto build api group while it uses `alicloud_api_gateway_group` to build api group.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_api_gateway_group&exampleId=ff7858f8-9d2d-c61b-0ca0-ee35e84a1d6bcf60cf3a&activeTab=example&spm=docs.r.api_gateway_group.0.ff7858f89d" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
resource "alicloud_api_gateway_group" "default" {
  name        = "tf_example"
  description = "tf_example"
  base_path   = "/"
}
```
## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the api gateway group. Defaults to null.
* `description` - (Optional) The description of the api gateway group. Defaults to null.
* `instance_id` - (Optional, ForceNew, Available in 1.179.0+)	The id of the api gateway.
* `base_path` - (Optional, Computed, Available since v1.228.0) The base path of the api gateway group. Defaults to `/`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the api group of api gateway.
* `sub_domain` - (Available in 1.69.0+)	Second-level domain name automatically assigned to the API group.
* `vpc_domain` - (Available in 1.69.0+)	Second-level VPC domain name automatically assigned to the API group.

## Import

Api gateway group can be imported using the id, e.g.

```shell
$ terraform import alicloud_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91"
```
