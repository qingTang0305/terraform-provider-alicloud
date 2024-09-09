package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAliCloudGovernanceAccount_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_governance_account.default"
	ra := resourceAttrInit(resourceId, AlicloudGovernanceAccountMap7372)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GovernanceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGovernanceAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sgovernanceaccount%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudGovernanceAccountBasicDependence7372)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"account_id":          "1493822914031335",
					"baseline_id":         "${data.alicloud_governance_baselines.default.ids.0}",
					"payer_account_id":    "${data.alicloud_account.default.id}",
					"display_name":        name,
					"account_name_prefix": name,
					"folder_id":           "${data.alicloud_resource_manager_folders.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"baseline_id":         CHECKSET,
						"payer_account_id":    CHECKSET,
						"display_name":        CHECKSET,
						"account_name_prefix": CHECKSET,
						"folder_id":           CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_name_prefix", "display_name", "folder_id", "payer_account_id", "baseline_id"},
			},
		},
	})
}

func TestAccAliCloudGovernanceAccount_basic7372(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_governance_account.default"
	ra := resourceAttrInit(resourceId, AlicloudGovernanceAccountMap7372)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GovernanceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGovernanceAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sgovernanceaccount%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudGovernanceAccountBasicDependence7372)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"baseline_id":         "${data.alicloud_governance_baselines.default.ids.0}",
					"payer_account_id":    "${data.alicloud_account.default.id}",
					"display_name":        name,
					"account_name_prefix": name,
					"folder_id":           "${data.alicloud_resource_manager_folders.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"baseline_id":         CHECKSET,
						"payer_account_id":    CHECKSET,
						"display_name":        CHECKSET,
						"account_name_prefix": CHECKSET,
						"folder_id":           CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"baseline_id": "${data.alicloud_governance_baselines.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"baseline_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_name_prefix", "display_name", "folder_id", "payer_account_id", "baseline_id"},
			},
		},
	})
}

var AlicloudGovernanceAccountMap7372 = map[string]string{
	"status":     CHECKSET,
	"account_id": CHECKSET,
}

func AlicloudGovernanceAccountBasicDependence7372(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_account" "default" {
}

data "alicloud_governance_baselines" "default" {
}

data "alicloud_resource_manager_folders" "default" {
}


`, name)
}
