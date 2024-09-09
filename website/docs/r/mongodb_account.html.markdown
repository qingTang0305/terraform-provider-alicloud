---
subcategory: "MongoDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_mongodb_account"
sidebar_current: "docs-alicloud-resource-mongodb-account"
description: |-
  Provides a Alicloud MongoDB Account resource.
---

# alicloud_mongodb_account

Provides a MongoDB Account resource.

For information about MongoDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/doc-detail/62154.html).

-> **NOTE:** Available since v1.148.0.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_mongodb_account&exampleId=6b410fdc-06af-0ce6-22f7-149630f8a0982c790f03&activeTab=example&spm=docs.r.mongodb_account.0.6b410fdc06" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}
data "alicloud_mongodb_zones" "default" {}
locals {
  index   = length(data.alicloud_mongodb_zones.default.zones) - 1
  zone_id = data.alicloud_mongodb_zones.default.zones[local.index].id
}
resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.17.3.0/24"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "172.17.3.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = local.zone_id
}

resource "alicloud_mongodb_instance" "default" {
  engine_version      = "4.2"
  db_instance_class   = "dds.mongo.mid"
  db_instance_storage = 10
  vswitch_id          = alicloud_vswitch.default.id
  security_ip_list    = ["10.168.1.12", "100.69.7.112"]
  name                = var.name
  tags = {
    Created = "TF"
    For     = "example"
  }
}

resource "alicloud_mongodb_account" "default" {
  account_name        = "root"
  account_password    = "Example_123"
  instance_id         = alicloud_mongodb_instance.default.id
  account_description = var.name
}
```

## Argument Reference

The following arguments are supported:

* `account_description` - (Optional) The description of the account.
  * The description must start with a letter, and cannot start with `http://` or `https://`.
  * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
* `account_name` - (Required) The name of the account. Valid values: `root`.
* `account_password` - (Required, Sensitive) The Password of the Account.
  * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
  * The password must be `8` to `32` characters in length.
* `instance_id` - (Required, ForceNew) The ID of the instance.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID of Account. The value formats as `<instance_id>:<account_name>`.
* `status` - The status of the account. Valid values: `Unavailable`, `Available`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 10 mins) Used when create the Account.
* `update` - (Defaults to 10 mins) Used when update the Account.

## Import

MongoDB Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_mongodb_account.example <instance_id>:<account_name>
```