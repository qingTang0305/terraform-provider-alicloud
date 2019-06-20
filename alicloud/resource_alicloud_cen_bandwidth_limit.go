package alicloud

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAlicloudCenBandwidthLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudCenBandwidthLimitCreate,
		Read:   resourceAlicloudCenBandwidthLimitRead,
		Update: resourceAlicloudCenBandwidthLimitUpdate,
		Delete: resourceAlicloudCenBandwidthLimitDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MaxItems: 2,
				MinItems: 2,
			},
			"bandwidth_limit": {
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value < 1 {
						errors = append(errors, fmt.Errorf("%s should be at least than 1 Mbps", k))
					}

					return
				},
			},
		},
	}
}

func resourceAlicloudCenBandwidthLimitCreate(d *schema.ResourceData, meta interface{}) error {
	cenId := d.Get("instance_id").(string)
	regionIds := d.Get("region_ids").(*schema.Set).List()
	if len(regionIds) != 2 {
		return WrapError(Error("Two different region ids should be set for bandwidth limit. "))
	}

	localRegionId := regionIds[0].(string)
	oppositeRegionId := regionIds[1].(string)

	if strings.Compare(localRegionId, oppositeRegionId) <= 0 {
		d.SetId(cenId + COLON_SEPARATED + localRegionId + COLON_SEPARATED + oppositeRegionId)
	} else {
		d.SetId(cenId + COLON_SEPARATED + oppositeRegionId + COLON_SEPARATED + localRegionId)
	}

	return resourceAlicloudCenBandwidthLimitUpdate(d, meta)
}

func resourceAlicloudCenBandwidthLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cenService := CenService{client}

	object, err := cenService.DescribeCenBandwidthLimit(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	respRegionIds := make([]string, 0)
	respRegionIds = append(respRegionIds, object.LocalRegionId, object.OppositeRegionId)

	d.Set("region_ids", respRegionIds)
	d.Set("instance_id", object.CenId)
	d.Set("bandwidth_limit", object.BandwidthLimit)

	return nil
}

func resourceAlicloudCenBandwidthLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cenService := CenService{client}
	cenId := d.Get("instance_id").(string)

	regionIds := d.Get("region_ids").(*schema.Set).List()
	if len(regionIds) != 2 {
		return WrapError(Error("Two different region ids should be set for bandwidth limit. "))
	}

	localRegionId := regionIds[0].(string)
	oppositeRegionId := regionIds[1].(string)
	var bandwidthLimit int

	if d.HasChange("bandwidth_limit") {
		bandwidthLimit = d.Get("bandwidth_limit").(int)
		if bandwidthLimit == 0 {
			return WrapError(Error("the bandwidth limit should be at least than 1 Mbps"))
		}
		err := resource.Retry(5*time.Minute, func() *resource.RetryError {
			err := cenService.SetCenInterRegionBandwidthLimit(cenId, localRegionId, oppositeRegionId, bandwidthLimit)
			if err != nil {
				if IsExceptedError(err, InvalidCenInstanceStatus) {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if err != nil {
			return WrapError(err)
		}

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			stateConf := &resource.StateChangeConf{
				Pending:                   []string{"Modifying"},
				Target:                    []string{"Active"},
				Refresh:                   cenService.CenBandwidthLimitStateRefreshFunc(d.Id(), []string{}),
				Timeout:                   d.Timeout(schema.TimeoutUpdate),
				Delay:                     3 * time.Second,
				MinTimeout:                3 * time.Second,
				PollInterval:              3 * time.Second,
				ContinuousTargetOccurence: 1,
			}

			if _, err = stateConf.WaitForState(); err != nil {
				if IsExceptedError(err, PvtzThrottlingUser) {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if err != nil {
			return WrapError(err)
		}
	}
	return resourceAlicloudCenBandwidthLimitRead(d, meta)
}

func resourceAlicloudCenBandwidthLimitDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cenService := CenService{client}
	cenId := d.Get("instance_id").(string)

	regionIds := d.Get("region_ids").(*schema.Set).List()
	if len(regionIds) != 2 {
		return fmt.Errorf("Two different region ids should be set for bandwidth limit")
	}

	localRegionId := regionIds[0].(string)
	oppositeRegionId := regionIds[1].(string)

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		err := cenService.SetCenInterRegionBandwidthLimit(cenId, localRegionId, oppositeRegionId, 0)
		if err != nil {
			if IsExceptedError(err, InvalidCenInstanceStatus) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return WrapError(err)
	}

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		stateConf := &resource.StateChangeConf{
			Pending:                   []string{"Active", "Modifying"},
			Target:                    []string{},
			Refresh:                   cenService.CenBandwidthLimitStateRefreshFunc(d.Id(), []string{}),
			Timeout:                   d.Timeout(schema.TimeoutDelete),
			Delay:                     3 * time.Second,
			MinTimeout:                3 * time.Second,
			PollInterval:              3 * time.Second,
			ContinuousTargetOccurence: 1,
		}

		_, err = stateConf.WaitForState()
		if IsExceptedError(err, PvtzThrottlingUser) {
			return resource.RetryableError(err)
		}
		return resource.NonRetryableError(err)
	})

	return WrapError(err)
}
