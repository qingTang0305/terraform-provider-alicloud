---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_instance"
description: |-
  Provides a Alicloud ENS Instance resource.
---

# alicloud_ens_instance

Provides a ENS Instance resource. 

For information about ENS Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.216.0.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_ens_instance&exampleId=9a65ef92-61f4-4b21-e60d-0731269974d7689744af&activeTab=example&spm=docs.r.ens_instance.0.9a65ef9261" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}


resource "alicloud_ens_instance" "default" {
  period = 1
  data_disk {
    size     = 20
    category = "cloud_efficiency"
  }
  data_disk {
    size     = 30
    category = "cloud_efficiency"
  }
  data_disk {
    size     = 40
    category = "cloud_efficiency"
  }
  public_ip_identification   = true
  period_unit                = "Month"
  scheduling_strategy        = "Concentrate"
  schedule_area_level        = "Region"
  image_id                   = "centos_6_08_64_20G_alibase_20171208"
  carrier                    = "cmcc"
  instance_type              = "ens.sn1.tiny"
  host_name                  = "exampleHost80"
  password                   = "Example123456@@"
  net_district_code          = "100102"
  internet_charge_type       = "95BandwidthByMonth"
  instance_name              = var.name
  internet_max_bandwidth_out = 100
  ens_region_id              = "cn-wuxi-telecom_unicom_cmcc-2"
  system_disk {
    size = 20
  }
  scheduling_price_strategy = "PriceHighPriority"
  user_data                 = "IyEvYmluL3NoCmVjaG8gIkhlbGxvIFdvcmxkLiAgVGhlIHRpbWUgaXMgbm93ICQoZGF0ZSAtUikhIiB8IHRlZSAvcm9vdC9vdXRwdXQudHh0"
  instance_charge_strategy  = "user"
  payment_type              = "Subscription"
}
```

### Deleting `alicloud_ens_instance` or removing it from your configuration

The `alicloud_ens_instance` resource allows you to manage  `payment_type = "Subscription"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `amount` - (Optional) The number of instances created, with a minimum of 1 and a maximum of 100.
* `auto_renew` - (Optional, Available since v1.208.0) Whether to automatically renew the logo. The default value is false. This parameter is invalid when you pay by volume.
* `auto_use_coupon` - (Optional) Whether to use vouchers. The default is to use. Value:
  - true (used)
  - false (not used).
* `billing_cycle` - (Optional) The billing cycle for instance computing resources. Only instance-level pay-as-you-go is supported. Value
  - Hour: hourly billing
  - Day: Daily billing
  - Month: monthly billing.
* `carrier` - (Optional, Available since v1.208.0) Operator, required for regional scheduling. Optional values:
  - cmcc (mobile)
  - unicom
  - telecom.
* `data_disk` - (Optional, ForceNew, Available since v1.208.0) Data disk specifications. See [`data_disk`](#data_disk) below.
* `ens_region_id` - (Optional, ForceNew, Available since v1.208.0) The node ID. When ScheduleAreaLevel is Region, EnsRegionId is required. When ScheduleAreaLevel is Big,Middle,Small, EnsRegionId is invalid.
* `force_stop` - (Optional) Whether to force the identity when operating the instance. Optional values:
  - true: Force
  - false (default): non-mandatory
* `host_name` - (Optional, Computed, Available since v1.208.0) The host name of the instance. Example value: test-HostName.
* `image_id` - (Optional, ForceNew, Available since v1.208.0) The image ID of the instance. The arm version card cannot be filled in. Other specifications are required. Example value: m-5si16wo6simkt267p8b7h * * * *.
* `include_data_disks` - (Optional) Whether the Payment type of the disk created with the instance is converted.
* `instance_charge_strategy` - (Optional, Available since v1.208.0) The instance billing policy. Optional values:
  - instance: instance granularity (the subscription method does not support instance)
  - user: user Dimension (user is not transmitted or supported in the prepaid mode).
* `instance_name` - (Optional, Computed, Available since v1.208.0) The instance name. Example value: test-InstanceName. It must be 2 to 128 characters in length and must start with an uppercase or lowercase letter or a Chinese character. It cannot start with http:// or https. Can contain Chinese, English, numbers, half-width colons (:), underscores (_), periods (.), or hyphens (-) The default value is the InstanceId of the instance. .
* `instance_type` - (Required, Available since v1.208.0) The specification of the instance. Example value: ens.sn1.small.
* `internet_charge_type` - (Optional, Available since v1.208.0) Instance bandwidth billing method. If the billing method can be selected for the first purchase, the subsequent value of this field will be processed by default according to the billing method selected for the first time. Optional values:
  - BandwidthByDay: Daily peak bandwidth
  - 95bandwidthbymonth: 95 peak bandwidth.
* `internet_max_bandwidth_out` - (Optional, ForceNew, Computed, Available since v1.208.0) Maximum public network bandwidth. The field type is Long, and the precision may be lost during serialization/deserialization. Please note that the value must not be greater than 9007199254740991.
* `ip_type` - (Optional) The IP type. Value:
  - ipv4 (default):IPv4
  - ipv6:IPv6
  - ipv4Andipv6:IPv4 and IPv6.
* `net_district_code` - (Optional, Available since v1.208.0) The area code. Example value: 350000. Required for regional-level scheduling, invalid for node-level scheduling.
* `net_work_id` - (Optional, ForceNew) The network ID of the instance. Can only be used in node-level scheduling.
* `password` - (Optional, Available since v1.208.0) The instance password. At least one of Password, KeyPairName, and PasswordInherit.
* `password_inherit` - (Optional, Available since v1.208.0) Whether to use image preset password prompt: Password and KeyPairNamePasswordInherit must be passed.
* `payment_type` - (Required, ForceNew, Available since v1.208.0) Instance payment method. Optional values:
  - Subscription: prepaid, annual and monthly
  - PayAsYouGo: Pay by volume.
* `period` - (Optional, Available since v1.208.0) The duration of the resource purchase. Value method:
  - If PeriodUnit is set to Day, Period can only be set to 3.
  - If PeriodUnit is set to Month, Period can be set to 1-9,12.
* `period_unit` - (Optional, Available since v1.208.0) The unit of time for purchasing resources. Value:
  - Month (default): purchase by Month
  - Day: buy by Day.
* `private_ip_address` - (Optional, ForceNew) The private IP address. Can only be used for node-level scheduling. If a private IP address is specified, the number of instances can only be one, and both the private IP address and the vSwitch ID are not empty, the private IP address takes effect.
* `public_ip_identification` - (Optional, Available since v1.208.0) Whether to assign a public IP identifier. Value:
  - true (default): Assign
  - false: do not assign.
* `schedule_area_level` - (Required, Available since v1.208.0) Scheduling level, through which node-level scheduling or area scheduling is performed. Optional values:
  - Node-level scheduling: Region
  - Regional scheduling: Big (region),Middle (province),Small (city).
* `scheduling_price_strategy` - (Optional, Available since v1.208.0) Scheduling price policy. If it is not filled in, the default priority is low price. Value:
  - PriceLowPriority
  - PriceLowPriority (priority low price).
* `scheduling_strategy` - (Optional, Available since v1.208.0) Scheduling policy. Optional values:
  - Concentrate for node-level scheduling
  - For regional scheduling, Concentrate, Disperse.
* `security_id` - (Optional, ForceNew) ID of the security group to which the instance belongs.
* `status` - (Optional, Computed) Status of the instance.
* `system_disk` - (Optional, ForceNew, Available since v1.208.0) System Disk Specification. SystemDisk is a non-required parameter when InstanceType is x86_pm,x86_bmi,x86_bm,pc_bmi, or arm_bmi. SystemDisk is a required parameter when instanceType is other specification families. See [`system_disk`](#system_disk) below.
* `unique_suffix` - (Optional, Available since v1.208.0) Indicates whether to add an ordered suffix to HostName and InstanceName. The ordered suffix starts from 001 and cannot exceed 999.
* `user_data` - (Optional, Available since v1.208.0) User-defined data, maximum support 16KB. You can pass in the UserData information. The UserData is encoded in Base64 format.
* `vswitch_id` - (Optional, ForceNew) The ID of the vSwitch to which the instance belongs. Can only be used in node-level scheduling.

### `data_disk`

The data_disk supports the following:
* `category` - (Optional, ForceNew) Data disk type. Optional values:
  - cloud_efficiency: Ultra cloud disk
  - cloud_ssd: Full Flash cloud disk
  - local_hdd: local hdd disk
  - local_ssd: local disk ssd.
* `size` - (Optional, ForceNew) Data disk size, unit: GB.

### `system_disk`

The system_disk supports the following:
* `category` - (Optional, ForceNew) System disk type. Optional values:
  - cloud_efficiency: Ultra cloud disk
  - cloud_ssd: Full Flash cloud disk
  - local_hdd: local hdd disk
  - local_ssd: local disk ssd.
* `size` - (Optional, ForceNew) System disk size, unit: GB.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `system_disk` - System Disk Specification. SystemDisk is a non-required parameter when InstanceType is x86_pm,x86_bmi,x86_bm,pc_bmi, or arm_bmi. SystemDisk is a required parameter when instanceType is other specification families.
  * `category` - System disk type. Value
  - cloud_efficiency: Ultra cloud disk
  - cloud_ssd: Full Flash cloud disk
  - local_hdd: local hdd disk
  - local_ssd: local disk ssd.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Instance.
* `delete` - (Defaults to 10 mins) Used when delete the Instance.
* `update` - (Defaults to 10 mins) Used when update the Instance.

## Import

ENS Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_instance.example <id>
```