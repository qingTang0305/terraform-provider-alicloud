---
subcategory: "DAS"
layout: "alicloud"
page_title: "Alicloud: alicloud_das_switch_das_pro"
sidebar_current: "docs-alicloud-resource-das-switch-das-pro"
description: |-
  Provides a Alicloud DAS Switch Das Pro resource.
---

# alicloud_das_switch_das_pro

Provides a DAS Switch Das Pro resource.

For information about DAS Switch Das Pro and how to use it, see [What is Switch Das Pro](https://www.alibabacloud.com/help/en/database-autonomy-service/latest/enabledaspro).

-> **NOTE:** Available since v1.193.0.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_das_switch_das_pro&exampleId=d75c4507-3124-fb34-5824-953cd7a41aaea0876805&activeTab=example&spm=docs.r.das_switch_das_pro.0.d75c450731" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
variable "name" {
  default = "tfexample"
}

data "alicloud_account" "default" {}
data "alicloud_polardb_node_classes" "default" {
  db_type    = "MySQL"
  db_version = "8.0"
  pay_type   = "PostPaid"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_polardb_node_classes.default.classes[0].zone_id
  vswitch_name = var.name
}

resource "alicloud_polardb_cluster" "default" {
  db_type       = "MySQL"
  db_version    = "8.0"
  db_node_class = "polar.mysql.x4.large"
  pay_type      = "PostPaid"
  vswitch_id    = alicloud_vswitch.default.id
  description   = var.name
  db_cluster_ip_array {
    db_cluster_ip_array_name = "default"
    security_ips             = ["1.2.3.4", "1.2.3.5"]
  }
}

resource "alicloud_das_switch_das_pro" "default" {
  instance_id   = alicloud_polardb_cluster.default.id
  sql_retention = 30
  user_id       = data.alicloud_account.default.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, ForceNew) The ID of the database instance.
* `sql_retention` - (Optional, ForceNew) The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
* `user_id` - (Optional, ForceNew) The ID of the Alibaba Cloud account that is used to create the database instance.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Switch Das Pro. Its value is same as `instance_id`.
* `status` - Whether the database instance has DAS professional.

## Import

DAS Switch Das Pro can be imported using the id, e.g.

```shell
$ terraform import alicloud_das_switch_das_pro.example <id>
```
