---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_image_event_operation"
description: |-
  Provides a Alicloud Threat Detection Image Event Operation resource.
---

# alicloud_threat_detection_image_event_operation

Provides a Threat Detection Image Event Operation resource. Image Event Operation.

For information about Threat Detection Image Event Operation and how to use it, see [What is Image Event Operation](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-addimageeventoperation).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_image_event_operation" "default" {
  event_type     = "maliciousFile"
  operation_code = "whitelist"
  event_key      = "alibabacloud_ak"
  scenarios      = <<EOF
{
  "type":"default",
  "value":""
}
EOF
  event_name     = "阿里云AK"
  conditions     = <<EOF
[
  {
      "condition":"MD5",
      "type":"equals",
      "value":"0083a31cc0083a31ccf7c10367a6e783e"
  }
]
EOF
}
```

## Argument Reference

The following arguments are supported:
* `conditions` - (Optional, ForceNew) Event Conditions.
* `event_key` - (Optional, ForceNew) Image Event Key.
* `event_name` - (Optional, ForceNew) Image Event Name.
* `event_type` - (Optional, ForceNew, Computed) Image Event Type.
* `operation_code` - (Optional, ForceNew) Event Operation Code.
* `scenarios` - (Optional) Event Scenarios.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Image Event Operation.
* `delete` - (Defaults to 5 mins) Used when delete the Image Event Operation.
* `update` - (Defaults to 5 mins) Used when update the Image Event Operation.

## Import

Threat Detection Image Event Operation can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_image_event_operation.example <id>
```