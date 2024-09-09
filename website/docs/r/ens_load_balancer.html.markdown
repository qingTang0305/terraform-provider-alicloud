---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_load_balancer"
description: |-
  Provides a Alicloud ENS Load Balancer resource.
---

# alicloud_ens_load_balancer

Provides a ENS Load Balancer resource. Load balancing. When you use it for the first time, please contact the product classmates to add a resource whitelist.

For information about ENS Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createloadbalancer).

-> **NOTE:** Available since v1.213.0.

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_ens_load_balancer&exampleId=cd617ca6-e9d9-748a-b19c-a142ef0dd8cb59a8a80d&activeTab=example&spm=docs.r.ens_load_balancer.0.cd617ca6e9&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_ens_network" "network" {
  network_name  = var.name
  description   = var.name
  cidr_block    = "192.168.2.0/24"
  ens_region_id = "cn-chenzhou-telecom_unicom_cmcc"
}

resource "alicloud_ens_vswitch" "switch" {
  description   = var.name
  cidr_block    = "192.168.2.0/24"
  vswitch_name  = var.name
  ens_region_id = "cn-chenzhou-telecom_unicom_cmcc"
  network_id    = alicloud_ens_network.network.id
}


resource "alicloud_ens_load_balancer" "default" {
  load_balancer_name = var.name
  payment_type       = "PayAsYouGo"
  ens_region_id      = "cn-chenzhou-telecom_unicom_cmcc"
  load_balancer_spec = "elb.s1.small"
  vswitch_id         = alicloud_ens_vswitch.switch.id
  network_id         = alicloud_ens_network.network.id
}
```

## Argument Reference

The following arguments are supported:
* `ens_region_id` - (Required, ForceNew) The ID of the ENS node.
* `load_balancer_name` - (Optional) Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
* `load_balancer_spec` - (Required, ForceNew) Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
* `network_id` - (Required, ForceNew) The network ID of the created edge load balancing (ELB) instance.
* `payment_type` - (Required, ForceNew) Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
* `vswitch_id` - (Required, ForceNew) The ID of the vSwitch to which the VPC instance belongs.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation Time (UTC) of the load balancing instance.
* `status` - The status of the SLB instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Load Balancer.
* `delete` - (Defaults to 5 mins) Used when delete the Load Balancer.
* `update` - (Defaults to 5 mins) Used when update the Load Balancer.

## Import

ENS Load Balancer can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_load_balancer.example <id>
```