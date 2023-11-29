// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product VPN
// ---------------------------------------------------------------

package vpn

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func ResourceGateway() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGatewayCreate,
		UpdateContext: resourceGatewayUpdate,
		ReadContext:   resourceGatewayRead,
		DeleteContext: resourceGatewayDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the VPN gateway. Only letters, digits, underscores(_) and hypens(-) are supported.`,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`^[\-_A-Za-z0-9]+$`),
						"the input is invalid"),
					validation.StringLenBetween(1, 64),
				),
			},
			"availability_zones": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				ForceNew:    true,
				Description: `The availability zone IDs.`,
			},
			"flavor": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The flavor of the VPN gateway.`,
				ValidateFunc: validation.StringInSlice([]string{
					"V1G", "V300", "Basic", "Professional1", "Professional2",
				}, false),
			},
			"attachment_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "vpc",
				ForceNew:    true,
				Description: `The attachment type.`,
				ValidateFunc: validation.StringInSlice([]string{
					"vpc", "er",
				}, false),
			},
			"network_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The network type of the VPN gateway.`,
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The ID of the VPC to which the VPN gateway is connected.`,
			},
			"local_subnets": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `The local subnets.`,
			},
			"connect_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The Network ID of the VPC subnet used by the VPN gateway.`,
			},
			"er_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The enterprise router ID to attach with to VPN gateway.`,
			},
			"ha_mode": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				Description:   `The HA mode of the VPN gateway.`,
				ValidateFunc:  validation.StringInSlice([]string{"active-active", "active-standby"}, false),
				ConflictsWith: []string{"master_eip", "slave_eip"},
			},
			"master_eip": {
				Type:         schema.TypeList,
				MaxItems:     1,
				Elem:         GatewayEipSchema(),
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				Description:  utils.SchemaDesc("", utils.SchemaDescInput{Internal: true}),
				RequiredWith: []string{"slave_eip"},
			},
			"eip1": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Elem:          GatewayEipSchema(),
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"master_eip", "slave_eip"},
				RequiredWith:  []string{"eip2"},
			},
			"slave_eip": {
				Type:         schema.TypeList,
				MaxItems:     1,
				Elem:         GatewayEipSchema(),
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				Description:  utils.SchemaDesc("", utils.SchemaDescInput{Internal: true}),
				RequiredWith: []string{"master_eip"},
			},
			"eip2": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Elem:          GatewayEipSchema(),
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"master_eip", "slave_eip"},
				RequiredWith:  []string{"eip1"},
			},
			"access_vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The access VPC ID of the VPN gateway.`,
			},
			"access_subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The access subnet ID of the VPN gateway.`,
			},
			"asn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     64512,
				ForceNew:    true,
				Description: `The ASN number of BGP`,
			},
			"enterprise_project_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  `The enterprise project ID`,
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The status of VPN gateway.`,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The create time.`,
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The update time.`,
			},
			"used_connection_group": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of used connection groups.`,
			},
			"used_connection_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of used connections.`,
			},
		},
	}
}

func GatewayEipSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The public IP ID.`,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}`),
						"the input is invalid"),
					validation.StringLenBetween(36, 36),
				),
			},

			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The EIP type. The value can be **5_bgp** and **5_sbgp**.`,
			},
			"bandwidth_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The bandwidth name.`,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`^[\-_A-Za-z0-9]+$`),
						"the input is invalid"),
					validation.StringLenBetween(1, 64),
				),
			},
			"bandwidth_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Description: `Bandwidth size in Mbit/s. When the flavor is **V300**, the value cannot be greater than **300**.
`,
				ValidateFunc: validation.IntBetween(1, 1024),
			},
			"charge_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The charge mode of the bandwidth. The value can be **bandwidth** and **traffic**.`,
				ValidateFunc: validation.StringInSlice([]string{
					"bandwidth", "traffic",
				}, false),
			},

			"bandwidth_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The bandwidth ID.`,
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The public IP address.`,
			},
			"ip_version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The public IP version.`,
			},
		},
	}
	return &sc
}

func resourceGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createGateway: Create a VPN Gateway.
	var (
		createGatewayHttpUrl = "v5/{project_id}/vpn-gateways"
		createGatewayProduct = "vpn"
	)
	createGatewayClient, err := cfg.NewServiceClient(createGatewayProduct, region)
	if err != nil {
		return diag.Errorf("error creating Gateway Client: %s", err)
	}

	createGatewayPath := createGatewayClient.Endpoint + createGatewayHttpUrl
	createGatewayPath = strings.ReplaceAll(createGatewayPath, "{project_id}", createGatewayClient.ProjectID)

	createGatewayOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			201,
		},
	}
	createGatewayOpt.JSONBody = utils.RemoveNil(buildCreateGatewayBodyParams(d, cfg))
	createGatewayResp, err := createGatewayClient.Request("POST", createGatewayPath, &createGatewayOpt)
	if err != nil {
		return diag.Errorf("error creating Gateway: %s", err)
	}

	createGatewayRespBody, err := utils.FlattenResponse(createGatewayResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("vpn_gateway.id", createGatewayRespBody)
	if err != nil {
		return diag.Errorf("error creating Gateway: ID is not found in API response")
	}
	d.SetId(id.(string))

	err = createGatewayWaitingForStateCompleted(ctx, d, meta, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.Errorf("error waiting for the Create of Gateway (%s) to complete: %s", d.Id(), err)
	}
	return resourceGatewayRead(ctx, d, meta)
}

func buildCreateGatewayBodyParams(d *schema.ResourceData, cfg *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"vpn_gateway": buildCreateGatewayVpnGatewayChildBody(d, cfg),
	}
	return bodyParams
}

func buildCreateGatewayVpnGatewayChildBody(d *schema.ResourceData, cfg *config.Config) map[string]interface{} {
	haMode := d.Get("ha_mode").(string)
	masterEIP := buildCreateGatewayEIPChildBody(d, "master_eip")
	slaveEIP := buildCreateGatewayEIPChildBody(d, "slave_eip")

	// default use "active-standby" ha_mode type when declare master_eip and slave_eip
	if haMode == "" && masterEIP != nil && slaveEIP != nil {
		haMode = "active-standby"
	}
	params := map[string]interface{}{
		"attachment_type":       utils.ValueIngoreEmpty(d.Get("attachment_type")),
		"availability_zone_ids": utils.ValueIngoreEmpty(d.Get("availability_zones")),
		"bgp_asn":               utils.ValueIngoreEmpty(d.Get("asn")),
		"connect_subnet":        utils.ValueIngoreEmpty(d.Get("connect_subnet")),
		"enterprise_project_id": utils.ValueIngoreEmpty(common.GetEnterpriseProjectID(d, cfg)),
		"flavor":                utils.ValueIngoreEmpty(d.Get("flavor")),
		"local_subnets":         utils.ValueIngoreEmpty(d.Get("local_subnets")),
		"name":                  utils.ValueIngoreEmpty(d.Get("name")),
		"vpc_id":                utils.ValueIngoreEmpty(d.Get("vpc_id")),
		"ha_mode":               utils.ValueIngoreEmpty(haMode),
		"eip1":                  buildCreateGatewayEIPChildBody(d, "eip1"),
		"master_eip":            masterEIP,
		"eip2":                  buildCreateGatewayEIPChildBody(d, "eip2"),
		"slave_eip":             slaveEIP,
		"access_vpc_id":         utils.ValueIngoreEmpty(d.Get("access_vpc_id")),
		"access_subnet_id":      utils.ValueIngoreEmpty(d.Get("access_subnet_id")),
		"er_id":                 utils.ValueIngoreEmpty(d.Get("er_id")),
		"network_type":          utils.ValueIngoreEmpty(d.Get("network_type")),
	}

	return params
}

func buildCreateGatewayEIPChildBody(d *schema.ResourceData, param string) map[string]interface{} {
	if rawArray, ok := d.Get(param).([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		raw, ok := rawArray[0].(map[string]interface{})
		if !ok {
			return nil
		}

		params := map[string]interface{}{
			"bandwidth_name": utils.ValueIngoreEmpty(raw["bandwidth_name"]),
			"bandwidth_size": utils.ValueIngoreEmpty(raw["bandwidth_size"]),
			"charge_mode":    utils.ValueIngoreEmpty(raw["charge_mode"]),
			"id":             utils.ValueIngoreEmpty(raw["id"]),
			"type":           utils.ValueIngoreEmpty(raw["type"]),
		}
		return params
	}
	return nil
}

func createGatewayWaitingForStateCompleted(ctx context.Context, d *schema.ResourceData, meta interface{}, t time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED"},
		Refresh: func() (interface{}, string, error) {
			cfg := meta.(*config.Config)
			region := cfg.GetRegion(d)
			// createGatewayWaiting: missing operation notes
			var (
				createGatewayWaitingHttpUrl = "v5/{project_id}/vpn-gateways/{id}"
				createGatewayWaitingProduct = "vpn"
			)
			createGatewayWaitingClient, err := cfg.NewServiceClient(createGatewayWaitingProduct, region)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error creating Gateway Client: %s", err)
			}

			createGatewayWaitingPath := createGatewayWaitingClient.Endpoint + createGatewayWaitingHttpUrl
			createGatewayWaitingPath = strings.ReplaceAll(createGatewayWaitingPath, "{project_id}", createGatewayWaitingClient.ProjectID)
			createGatewayWaitingPath = strings.ReplaceAll(createGatewayWaitingPath, "{id}", d.Id())

			createGatewayWaitingOpt := golangsdk.RequestOpts{
				KeepResponseBody: true,
				OkCodes: []int{
					200,
				},
			}
			createGatewayWaitingResp, err := createGatewayWaitingClient.Request("GET", createGatewayWaitingPath, &createGatewayWaitingOpt)
			if err != nil {
				return nil, "ERROR", err
			}

			createGatewayWaitingRespBody, err := utils.FlattenResponse(createGatewayWaitingResp)
			if err != nil {
				return nil, "ERROR", err
			}
			statusRaw, err := jmespath.Search(`vpn_gateway.status`, createGatewayWaitingRespBody)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error parse %s from response body", `vpn_gateway.status`)
			}

			status := fmt.Sprintf("%v", statusRaw)

			targetStatus := []string{
				"ACTIVE",
			}
			if utils.StrSliceContains(targetStatus, status) {
				return createGatewayWaitingRespBody, "COMPLETED", nil
			}

			pendingStatus := []string{
				"PENDING_CREATE",
			}
			if utils.StrSliceContains(pendingStatus, status) {
				return createGatewayWaitingRespBody, "PENDING", nil
			}

			return createGatewayWaitingRespBody, status, nil

		},
		Timeout:      t,
		Delay:        10 * time.Second,
		PollInterval: 5 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func resourceGatewayRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	// getGateway: Query the VPN gateway detail
	var (
		getGatewayHttpUrl = "v5/{project_id}/vpn-gateways/{id}"
		getGatewayProduct = "vpn"
	)
	getGatewayClient, err := cfg.NewServiceClient(getGatewayProduct, region)
	if err != nil {
		return diag.Errorf("error creating Gateway Client: %s", err)
	}

	getGatewayPath := getGatewayClient.Endpoint + getGatewayHttpUrl
	getGatewayPath = strings.ReplaceAll(getGatewayPath, "{project_id}", getGatewayClient.ProjectID)
	getGatewayPath = strings.ReplaceAll(getGatewayPath, "{id}", d.Id())

	getGatewayOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getGatewayResp, err := getGatewayClient.Request("GET", getGatewayPath, &getGatewayOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving Gateway")
	}

	getGatewayRespBody, err := utils.FlattenResponse(getGatewayResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("attachment_type", utils.PathSearch("vpn_gateway.attachment_type", getGatewayRespBody, nil)),
		d.Set("availability_zones", utils.PathSearch("vpn_gateway.availability_zone_ids", getGatewayRespBody, nil)),
		d.Set("asn", utils.PathSearch("vpn_gateway.bgp_asn", getGatewayRespBody, nil)),
		d.Set("connect_subnet", utils.PathSearch("vpn_gateway.connect_subnet", getGatewayRespBody, nil)),
		d.Set("created_at", utils.PathSearch("vpn_gateway.created_at", getGatewayRespBody, nil)),
		d.Set("enterprise_project_id", utils.PathSearch("vpn_gateway.enterprise_project_id", getGatewayRespBody, nil)),
		d.Set("flavor", utils.PathSearch("vpn_gateway.flavor", getGatewayRespBody, nil)),
		d.Set("local_subnets", utils.PathSearch("vpn_gateway.local_subnets", getGatewayRespBody, nil)),
		d.Set("ha_mode", utils.PathSearch("vpn_gateway.ha_mode", getGatewayRespBody, nil)),
		d.Set("eip1", flattenGetGatewayResponseBodyVPNGatewayBody(getGatewayRespBody, "eip1")),
		d.Set("master_eip", flattenGetGatewayResponseBodyVPNGatewayBody(getGatewayRespBody, "master_eip")),
		d.Set("name", utils.PathSearch("vpn_gateway.name", getGatewayRespBody, nil)),
		d.Set("eip2", flattenGetGatewayResponseBodyVPNGatewayBody(getGatewayRespBody, "eip2")),
		d.Set("slave_eip", flattenGetGatewayResponseBodyVPNGatewayBody(getGatewayRespBody, "slave_eip")),
		d.Set("status", utils.PathSearch("vpn_gateway.status", getGatewayRespBody, nil)),
		d.Set("updated_at", utils.PathSearch("vpn_gateway.updated_at", getGatewayRespBody, nil)),
		d.Set("used_connection_group", utils.PathSearch("vpn_gateway.used_connection_group", getGatewayRespBody, nil)),
		d.Set("used_connection_number", utils.PathSearch("vpn_gateway.used_connection_number", getGatewayRespBody, nil)),
		d.Set("vpc_id", utils.PathSearch("vpn_gateway.vpc_id", getGatewayRespBody, nil)),
		d.Set("access_vpc_id", utils.PathSearch("vpn_gateway.access_vpc_id", getGatewayRespBody, nil)),
		d.Set("access_subnet_id", utils.PathSearch("vpn_gateway.access_subnet_id", getGatewayRespBody, nil)),
		d.Set("er_id", utils.PathSearch("vpn_gateway.er_id", getGatewayRespBody, nil)),
		d.Set("network_type", utils.PathSearch("vpn_gateway.network_type", getGatewayRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenGetGatewayResponseBodyVPNGatewayBody(resp interface{}, paramName string) []interface{} {
	var rst []interface{}
	curJson, err := jmespath.Search(fmt.Sprintf("vpn_gateway.%s", paramName), resp)
	if err != nil {
		log.Printf("[ERROR] error parsing vpn_gateway.%s from response= %#v", paramName, resp)
		return rst
	}

	rst = []interface{}{
		map[string]interface{}{
			"bandwidth_id":   utils.PathSearch("bandwidth_id", curJson, nil),
			"bandwidth_name": utils.PathSearch("bandwidth_name", curJson, nil),
			"bandwidth_size": utils.PathSearch("bandwidth_size", curJson, nil),
			"charge_mode":    utils.PathSearch("charge_mode", curJson, nil),
			"id":             utils.PathSearch("id", curJson, nil),
			"ip_address":     utils.PathSearch("ip_address", curJson, nil),
			"ip_version":     utils.PathSearch("ip_version", curJson, nil),
			"type":           utils.PathSearch("type", curJson, nil),
		},
	}
	return rst
}

func resourceGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	updateGatewayHasChanges := []string{
		"local_subnets",
		"name",
	}

	if d.HasChanges(updateGatewayHasChanges...) {
		// updateGateway: Update the configuration of VPN gateway
		var (
			updateGatewayHttpUrl = "v5/{project_id}/vpn-gateways/{id}"
			updateGatewayProduct = "vpn"
		)
		updateGatewayClient, err := cfg.NewServiceClient(updateGatewayProduct, region)
		if err != nil {
			return diag.Errorf("error creating Gateway Client: %s", err)
		}

		updateGatewayPath := updateGatewayClient.Endpoint + updateGatewayHttpUrl
		updateGatewayPath = strings.ReplaceAll(updateGatewayPath, "{project_id}", updateGatewayClient.ProjectID)
		updateGatewayPath = strings.ReplaceAll(updateGatewayPath, "{id}", d.Id())

		updateGatewayOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				200,
			},
		}
		updateGatewayOpt.JSONBody = utils.RemoveNil(buildUpdateGatewayBodyParams(d))
		_, err = updateGatewayClient.Request("PUT", updateGatewayPath, &updateGatewayOpt)
		if err != nil {
			return diag.Errorf("error updating Gateway: %s", err)
		}
		err = updateGatewayWaitingForStateCompleted(ctx, d, meta, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return diag.Errorf("error waiting for the Update of Gateway (%s) to complete: %s", d.Id(), err)
		}
	}
	return resourceGatewayRead(ctx, d, meta)
}

func buildUpdateGatewayBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"vpn_gateway": buildUpdateGatewayVpnGatewayChildBody(d),
	}
	return bodyParams
}

func buildUpdateGatewayVpnGatewayChildBody(d *schema.ResourceData) map[string]interface{} {
	params := map[string]interface{}{
		"local_subnets": utils.ValueIngoreEmpty(d.Get("local_subnets")),
		"name":          utils.ValueIngoreEmpty(d.Get("name")),
	}
	return params
}

func updateGatewayWaitingForStateCompleted(ctx context.Context, d *schema.ResourceData, meta interface{}, t time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED"},
		Refresh: func() (interface{}, string, error) {
			cfg := meta.(*config.Config)
			region := cfg.GetRegion(d)
			// updateGatewayWaiting: missing operation notes
			var (
				updateGatewayWaitingHttpUrl = "v5/{project_id}/vpn-gateways/{id}"
				updateGatewayWaitingProduct = "vpn"
			)
			updateGatewayWaitingClient, err := cfg.NewServiceClient(updateGatewayWaitingProduct, region)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error creating Gateway Client: %s", err)
			}

			updateGatewayWaitingPath := updateGatewayWaitingClient.Endpoint + updateGatewayWaitingHttpUrl
			updateGatewayWaitingPath = strings.ReplaceAll(updateGatewayWaitingPath, "{project_id}", updateGatewayWaitingClient.ProjectID)
			updateGatewayWaitingPath = strings.ReplaceAll(updateGatewayWaitingPath, "{id}", d.Id())

			updateGatewayWaitingOpt := golangsdk.RequestOpts{
				KeepResponseBody: true,
				OkCodes: []int{
					200,
				},
			}
			updateGatewayWaitingResp, err := updateGatewayWaitingClient.Request("GET", updateGatewayWaitingPath, &updateGatewayWaitingOpt)
			if err != nil {
				return nil, "ERROR", err
			}

			updateGatewayWaitingRespBody, err := utils.FlattenResponse(updateGatewayWaitingResp)
			if err != nil {
				return nil, "ERROR", err
			}
			statusRaw, err := jmespath.Search(`vpn_gateway.status`, updateGatewayWaitingRespBody)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error parse %s from response body", `vpn_gateway.status`)
			}

			status := fmt.Sprintf("%v", statusRaw)

			targetStatus := []string{
				"ACTIVE",
			}
			if utils.StrSliceContains(targetStatus, status) {
				return updateGatewayWaitingRespBody, "COMPLETED", nil
			}

			pendingStatus := []string{
				"PENDING_UPDATE",
			}
			if utils.StrSliceContains(pendingStatus, status) {
				return updateGatewayWaitingRespBody, "PENDING", nil
			}

			return updateGatewayWaitingRespBody, status, nil
		},
		Timeout:      t,
		Delay:        10 * time.Second,
		PollInterval: 5 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func resourceGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// deleteGateway: Delete an existing VPN Gateway
	var (
		deleteGatewayHttpUrl = "v5/{project_id}/vpn-gateways/{id}"
		deleteGatewayProduct = "vpn"
	)
	deleteGatewayClient, err := cfg.NewServiceClient(deleteGatewayProduct, region)
	if err != nil {
		return diag.Errorf("error creating Gateway Client: %s", err)
	}

	deleteGatewayPath := deleteGatewayClient.Endpoint + deleteGatewayHttpUrl
	deleteGatewayPath = strings.ReplaceAll(deleteGatewayPath, "{project_id}", deleteGatewayClient.ProjectID)
	deleteGatewayPath = strings.ReplaceAll(deleteGatewayPath, "{id}", d.Id())

	deleteGatewayOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			204,
		},
	}
	_, err = deleteGatewayClient.Request("DELETE", deleteGatewayPath, &deleteGatewayOpt)
	if err != nil {
		return diag.Errorf("error deleting Gateway: %s", err)
	}

	err = deleteGatewayWaitingForStateCompleted(ctx, d, meta, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.Errorf("error waiting for the Delete of Gateway (%s) to complete: %s", d.Id(), err)
	}
	return nil
}

func deleteGatewayWaitingForStateCompleted(ctx context.Context, d *schema.ResourceData, meta interface{}, t time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED"},
		Refresh: func() (interface{}, string, error) {
			cfg := meta.(*config.Config)
			region := cfg.GetRegion(d)
			// deleteGatewayWaiting: missing operation notes
			var (
				deleteGatewayWaitingHttpUrl = "v5/{project_id}/vpn-gateways/{id}"
				deleteGatewayWaitingProduct = "vpn"
			)
			deleteGatewayWaitingClient, err := cfg.NewServiceClient(deleteGatewayWaitingProduct, region)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error creating Gateway Client: %s", err)
			}

			deleteGatewayWaitingPath := deleteGatewayWaitingClient.Endpoint + deleteGatewayWaitingHttpUrl
			deleteGatewayWaitingPath = strings.ReplaceAll(deleteGatewayWaitingPath, "{project_id}", deleteGatewayWaitingClient.ProjectID)
			deleteGatewayWaitingPath = strings.ReplaceAll(deleteGatewayWaitingPath, "{id}", d.Id())

			deleteGatewayWaitingOpt := golangsdk.RequestOpts{
				KeepResponseBody: true,
				OkCodes: []int{
					200,
				},
			}
			deleteGatewayWaitingResp, err := deleteGatewayWaitingClient.Request("GET", deleteGatewayWaitingPath, &deleteGatewayWaitingOpt)
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault404); ok {
					return deleteGatewayWaitingResp, "COMPLETED", nil
				}

				return nil, "ERROR", err
			}

			deleteGatewayWaitingRespBody, err := utils.FlattenResponse(deleteGatewayWaitingResp)
			if err != nil {
				return nil, "ERROR", err
			}
			statusRaw, err := jmespath.Search(`vpn_gateway.status`, deleteGatewayWaitingRespBody)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error parse %s from response body", `vpn_gateway.status`)
			}

			status := fmt.Sprintf("%v", statusRaw)

			var targetStatus []string
			if utils.StrSliceContains(targetStatus, status) {
				return deleteGatewayWaitingRespBody, "COMPLETED", nil
			}

			pendingStatus := []string{
				"PENDING_DELETE",
			}
			if utils.StrSliceContains(pendingStatus, status) {
				return deleteGatewayWaitingRespBody, "PENDING", nil
			}

			return deleteGatewayWaitingRespBody, status, nil

		},
		Timeout:      t,
		Delay:        10 * time.Second,
		PollInterval: 5 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
