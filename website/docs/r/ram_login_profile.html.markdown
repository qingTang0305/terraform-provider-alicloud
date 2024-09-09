---
subcategory: "RAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_ram_login_profile"
sidebar_current: "docs-alicloud-resource-ram-login-profile"
description: |-
  Provides a Alicloud RAM User Login Profile resource.
---

# alicloud_ram_login_profile

Provides a RAM User Login Profile resource.

For information about RAM User Login Profile and how to use it, see [What is Login Profile](https://www.alibabacloud.com/help/en/ram/developer-reference/api-ram-2015-05-01-createloginprofile).

-> **NOTE:** Available since v1.0.0.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_ram_login_profile&exampleId=442ec83a-6fe5-90e2-17a2-351fa3abb6a2738d91e8&activeTab=example&spm=docs.r.ram_login_profile.0.442ec83a6f" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
resource "alicloud_ram_user" "user" {
  name         = "terraform_example"
  display_name = "terraform_example"
  mobile       = "86-18688888888"
  email        = "hello.uuu@aaa.com"
  comments     = "terraform_example"
  force        = true
}

resource "alicloud_ram_login_profile" "profile" {
  user_name = alicloud_ram_user.user.name
  password  = "Example_1234"
}
```

## Argument Reference

The following arguments are supported:

* `user_name` - (Required, ForceNew) The name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
* `password` - (Required, Sensitive) The logon password of the RAM user. The password must meet the password strength requirements.
* `password_reset_required` - (Optional, Bool) Specifies whether the RAM user must change the password upon logon. Default value: `false`. Valid values: `true`, `false`.
* `mfa_bind_required` - (Optional, Bool) Specifies whether an MFA device must be attached to the RAM user upon logon. Valid values: `true`, `false`. [To enhance the security of your resources and data, the default value has been changed to `true`](https://www.alibabacloud.com/en/notice/mfa20240524?_p_lc=1) .

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Login Profile. It same as the `user_name`.

## Import

RAM login profile can be imported using the id, e.g.

```shell
$ terraform import alicloud_ram_login_profile.example <id>
```
