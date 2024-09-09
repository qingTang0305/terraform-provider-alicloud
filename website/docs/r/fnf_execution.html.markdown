---
subcategory: "Serverless Workflow (FnF)"
layout: "alicloud"
page_title: "Alicloud: alicloud_fnf_execution"
sidebar_current: "docs-alicloud-resource-fnf-execution"
description: |-
  Provides a Alicloud Serverless Workflow Execution resource.
---

# alicloud_fnf_execution

Provides a Serverless Workflow Execution resource.

For information about Serverless Workflow Execution and how to use it, see [What is Execution](https://www.alibabacloud.com/help/en/doc-detail/122628.html).

-> **NOTE:** Available since v1.149.0+.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_fnf_execution&exampleId=6141d824-e267-2cd5-942e-7b613c3b93af47e5c8cf&activeTab=example&spm=docs.r.fnf_execution.0.6141d824e2" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-shanghai"
}

variable "name" {
  default = "tf-example-fnfflow"
}

resource "alicloud_ram_role" "default" {
  name     = var.name
  document = <<EOF
  {
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
          "Service": [
            "fnf.aliyuncs.com"
          ]
        }
      }
    ],
    "Version": "1"
  }
  EOF
}

resource "alicloud_fnf_flow" "default" {
  definition  = <<EOF
  version: v1beta1
  type: flow
  steps:
    - type: wait
      name: custom_wait
      duration: $.wait
  EOF
  role_arn    = alicloud_ram_role.default.arn
  description = "Test for terraform fnf_flow."
  name        = var.name
  type        = "FDL"
}

resource "alicloud_fnf_execution" "default" {
  execution_name = var.name
  flow_name      = alicloud_fnf_flow.default.name
  input          = "{\"wait\": 600}"
}
```

## Argument Reference

The following arguments are supported:

* `execution_name` - (Required, ForceNew) The name of the execution.
* `flow_name` - (Required, ForceNew) The name of the flow.
* `input` - (Optional, ForceNew) The Input information for this execution.
* `status` - (Optional, Computed) The status of the resource. Valid values: `Stopped`.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID of Execution. The value formats as `<flow_name>:<execution_name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 5 mins) Used when create the Execution.
* `update` - (Defaults to 5 mins) Used when update the Execution.

## Import

Serverless Workflow Execution can be imported using the id, e.g.

```shell
$ terraform import alicloud_fnf_execution.example <flow_name>:<execution_name>
```