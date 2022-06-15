// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),

// Description: Interface configuration.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceCreate,
		Read:   resourceSystemInterfaceRead,
		Update: resourceSystemInterfaceUpdate,
		Delete: resourceSystemInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"autogenerated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"ip6_autoconf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"lacp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacp_speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_up_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"min_links": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_links_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource while getting object: %v", err)
	}

	v, _ := d.GetOk("type")
	ag, _ := d.GetOk("autogenerated")
	if v == "physical" || ag == "auto" {
		_, err = c.UpdateSystemInterface(obj, adomv, (*obj)["name"].(string), nil)
	} else {
		_, err = c.CreateSystemInterface(obj, adomv, nil)
	}

	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemInterfaceRead(d, m)
}

func resourceSystemInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterface(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemInterfaceRead(d, m)
}

func resourceSystemInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	v, _ := d.GetOk("type")
	ag, _ := d.GetOk("autogenerated")
	if v == "physical" || ag == "auto" {
		d.SetId("")
		return nil
	}

	err = c.DeleteSystemInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenSystemInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenSystemInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_autoconf"
	if _, ok := i["ip6-autoconf"]; ok {
		result["ip6_autoconf"] = flattenSystemInterfaceIpv6Ip6Autoconf(i["ip6-autoconf"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLacpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLacpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLinkUpDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			v := flattenSystemInterfaceMemberInterfaceName(i["interface-name"], d, pre_append)
			tmp["interface_name"] = fortiAPISubPartPatch(v, "SystemInterface-Member-InterfaceName")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceMemberInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMinLinks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMinLinksDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aggregate", flattenSystemInterfaceAggregate(o["aggregate"], d, "aggregate")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregate"], "SystemInterface-Aggregate"); ok {
			if err = d.Set("aggregate", vv); err != nil {
				return fmt.Errorf("Error reading aggregate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregate: %v", err)
		}
	}

	if err = d.Set("alias", flattenSystemInterfaceAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "SystemInterface-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "SystemInterface-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemInterfaceDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemInterface-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemInterface-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv6"], "SystemInterface-Ipv6"); ok {
				if err = d.Set("ipv6", vv); err != nil {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6"); ok {
			if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv6"], "SystemInterface-Ipv6"); ok {
					if err = d.Set("ipv6", vv); err != nil {
						return fmt.Errorf("Error reading ipv6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			}
		}
	}

	if err = d.Set("lacp_mode", flattenSystemInterfaceLacpMode(o["lacp-mode"], d, "lacp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-mode"], "SystemInterface-LacpMode"); ok {
			if err = d.Set("lacp_mode", vv); err != nil {
				return fmt.Errorf("Error reading lacp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_mode: %v", err)
		}
	}

	if err = d.Set("lacp_speed", flattenSystemInterfaceLacpSpeed(o["lacp-speed"], d, "lacp_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-speed"], "SystemInterface-LacpSpeed"); ok {
			if err = d.Set("lacp_speed", vv); err != nil {
				return fmt.Errorf("Error reading lacp_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_speed: %v", err)
		}
	}

	if err = d.Set("link_up_delay", flattenSystemInterfaceLinkUpDelay(o["link-up-delay"], d, "link_up_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-up-delay"], "SystemInterface-LinkUpDelay"); ok {
			if err = d.Set("link_up_delay", vv); err != nil {
				return fmt.Errorf("Error reading link_up_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_up_delay: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenSystemInterfaceMember(o["member"], d, "member")); err != nil {
			if vv, ok := fortiAPIPatch(o["member"], "SystemInterface-Member"); ok {
				if err = d.Set("member", vv); err != nil {
					return fmt.Errorf("Error reading member: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenSystemInterfaceMember(o["member"], d, "member")); err != nil {
				if vv, ok := fortiAPIPatch(o["member"], "SystemInterface-Member"); ok {
					if err = d.Set("member", vv); err != nil {
						return fmt.Errorf("Error reading member: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	if err = d.Set("min_links", flattenSystemInterfaceMinLinks(o["min-links"], d, "min_links")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-links"], "SystemInterface-MinLinks"); ok {
			if err = d.Set("min_links", vv); err != nil {
				return fmt.Errorf("Error reading min_links: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_links: %v", err)
		}
	}

	if err = d.Set("min_links_down", flattenSystemInterfaceMinLinksDown(o["min-links-down"], d, "min_links_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-links-down"], "SystemInterface-MinLinksDown"); ok {
			if err = d.Set("min_links_down", vv); err != nil {
				return fmt.Errorf("Error reading min_links_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_links_down: %v", err)
		}
	}

	if err = d.Set("mtu", flattenSystemInterfaceMtu(o["mtu"], d, "mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu"], "SystemInterface-Mtu"); ok {
			if err = d.Set("mtu", vv); err != nil {
				return fmt.Errorf("Error reading mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("speed", flattenSystemInterfaceSpeed(o["speed"], d, "speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["speed"], "SystemInterface-Speed"); ok {
			if err = d.Set("speed", vv); err != nil {
				return fmt.Errorf("Error reading speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemInterfaceStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemInterface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemInterfaceType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemInterface-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceAggregate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-address"], _ = expandSystemInterfaceIpv6Ip6Address(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-allowaccess"], _ = expandSystemInterfaceIpv6Ip6Allowaccess(d, i["ip6_allowaccess"], pre_append)
	} else {
		result["ip6-allowaccess"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_autoconf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-autoconf"], _ = expandSystemInterfaceIpv6Ip6Autoconf(d, i["ip6_autoconf"], pre_append)
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6Autoconf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLacpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLacpSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLinkUpDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-name"], _ = expandSystemInterfaceMemberInterfaceName(d, i["interface_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceMemberInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMinLinks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMinLinksDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggregate"); ok || d.HasChange("aggregate") {
		t, err := expandSystemInterfaceAggregate(d, v, "aggregate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandSystemInterfaceAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandSystemInterfaceAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemInterfaceDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemInterfaceIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandSystemInterfaceIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("lacp_mode"); ok || d.HasChange("lacp_mode") {
		t, err := expandSystemInterfaceLacpMode(d, v, "lacp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-mode"] = t
		}
	}

	if v, ok := d.GetOk("lacp_speed"); ok || d.HasChange("lacp_speed") {
		t, err := expandSystemInterfaceLacpSpeed(d, v, "lacp_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-speed"] = t
		}
	}

	if v, ok := d.GetOk("link_up_delay"); ok || d.HasChange("link_up_delay") {
		t, err := expandSystemInterfaceLinkUpDelay(d, v, "link_up_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-up-delay"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandSystemInterfaceMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("min_links"); ok || d.HasChange("min_links") {
		t, err := expandSystemInterfaceMinLinks(d, v, "min_links")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-links"] = t
		}
	}

	if v, ok := d.GetOk("min_links_down"); ok || d.HasChange("min_links_down") {
		t, err := expandSystemInterfaceMinLinksDown(d, v, "min_links_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-links-down"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok || d.HasChange("mtu") {
		t, err := expandSystemInterfaceMtu(d, v, "mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("speed"); ok || d.HasChange("speed") {
		t, err := expandSystemInterfaceSpeed(d, v, "speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speed"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemInterfaceStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemInterfaceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
