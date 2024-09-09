---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_bandwidth_package_attachment"
sidebar_current: "docs-alicloud-resource-cen-bandwidth-package-attachment"
description: |-
  Provides a Alicloud CEN bandwidth package attachment resource.
---

# alicloud_cen_bandwidth_package_attachment

Provides a CEN bandwidth package attachment resource. The resource can be used to bind a bandwidth package to a specified CEN instance.

-> **NOTE:** Available since v1.18.0.

## Example Usage
<div class="oics-button" style="float: right;margin: 0 0 -40px 0;">
  <a href="https://api.aliyun.com/api-tools/terraform?resource=alicloud_cen_bandwidth_package_attachment&exampleId=39c912f8-fb25-6bdb-4fa2-81d45687b48ba60f1a62&activeTab=example&spm=docs.r.cen_bandwidth_package_attachment.0.39c912f8fb" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>

Basic Usage

```terraform
resource "alicloud_cen_instance" "example" {
  cen_instance_name = "tf_example"
  description       = "an example for cen"
}

resource "alicloud_cen_bandwidth_package" "example" {
  bandwidth                  = 5
  cen_bandwidth_package_name = "tf_example"
  geographic_region_a_id     = "China"
  geographic_region_b_id     = "China"
}

resource "alicloud_cen_bandwidth_package_attachment" "example" {
  instance_id          = alicloud_cen_instance.example.id
  bandwidth_package_id = alicloud_cen_bandwidth_package.example.id
}
```
## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, ForceNew) The ID of the CEN.
* `bandwidth_package_id` - (Required, ForceNew) The ID of the bandwidth package.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the resource, the same as bandwidth_package_id.

## Timeouts

-> **NOTE:** Available in 1.206.0+.

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 5 mins) Used when create the CEN bandwidth package attachment.
* `delete` - (Defaults to 5 mins) Used when delete the CEN bandwidth package attachment.

## Import

CEN bandwidth package attachment resource can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_bandwidth_package_attachment.example bwp-abc123456
```
