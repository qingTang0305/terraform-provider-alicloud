---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_command"
sidebar_current: "docs-alicloud-resource-ecs-command"
description: |-
  Provides a Alicloud ECS Command resource.
---

# alicloud\_ecs\_command

Provides a ECS Command resource.

For information about ECS Command and how to use it, see [What is Command](https://www.alibabacloud.com/help/en/doc-detail/64844.htm).

-> **NOTE:** Available in v1.116.0+.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_ecs_command&exampleId=f456c891-22dc-232b-f26b-53addf27c3ce88600f71&activeTab=example&spm=docs.r.ecs_command.0.f456c89122" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
resource "alicloud_ecs_command" "example" {
  name            = "tf-testAcc"
  command_content = "bHMK"
  description     = "For Terraform Test"
  type            = "RunShellScript"
  working_dir     = "/root"
}

```

## Argument Reference

The following arguments are supported:

* `command_content` - (Required, ForceNew) The Base64-encoded content of the command.
* `description` - (Optional, ForceNew) The description of command.
* `enable_parameter` - (Optional, ForceNew) Specifies whether to use custom parameters in the command to be created. Default to: false.                                                                                                                  
* `name` - (Required, ForceNew) The name of the command, which supports all character sets. It can be up to 128 characters in length.
* `timeout` - (Optional, ForceNew) The timeout period that is specified for the command to be run on ECS instances. Unit: seconds. Default to: `60`.
* `type` - (Required, ForceNew) The command type. Valid Values: `RunBatScript`, `RunPowerShellScript` and `RunShellScript`.
* `working_dir` - (Optional, ForceNew) The execution path of the command in the ECS instance.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Command.

## Import

ECS Command can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_command.example <id>
```
