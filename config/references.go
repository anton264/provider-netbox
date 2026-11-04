package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// globalRefs maps a Terraform argument name to the resource it points to,
// wherever that argument appears. This gives every managed resource fields
// like tenantIdRef / tenantIdSelector, so you can reference other managed
// resources instead of hard-coding NetBox IDs.
var globalRefs = map[string]string{
	"tenant_id":                    "netbox_tenant",
	"site_id":                      "netbox_site",
	"region_id":                    "netbox_region",
	"site_group_id":                "netbox_site_group",
	"location_id":                  "netbox_location",
	"rack_id":                      "netbox_rack",
	"rack_type_id":                 "netbox_rack_type",
	"device_id":                    "netbox_device",
	"installed_device_id":          "netbox_device",
	"device_type_id":               "netbox_device_type",
	"module_id":                    "netbox_module",
	"module_type_id":               "netbox_module_type",
	"module_bay_id":                "netbox_device_module_bay",
	"manufacturer_id":              "netbox_manufacturer",
	"platform_id":                  "netbox_platform",
	"cluster_id":                   "netbox_cluster",
	"cluster_type_id":              "netbox_cluster_type",
	"cluster_group_id":             "netbox_cluster_group",
	"virtual_machine_id":           "netbox_virtual_machine",
	"virtual_chassis_id":           "netbox_virtual_chassis",
	"config_template_id":           "netbox_config_template",
	"vrf_id":                       "netbox_vrf",
	"vlan_id":                      "netbox_vlan",
	"rir_id":                       "netbox_rir",
	"prefix_id":                    "netbox_prefix",
	"ip_range_id":                  "netbox_ip_range",
	"ip_address_id":                "netbox_ip_address",
	"device_interface_id":          "netbox_device_interface",
	"virtual_machine_interface_id": "netbox_interface",
	"mac_address_id":               "netbox_mac_address",
	"power_panel_id":               "netbox_power_panel",
	"circuit_id":                   "netbox_circuit",
	"provider_network_id":          "netbox_circuit_provider_network",
	"contact_id":                   "netbox_contact",
	"choice_set_id":                "netbox_custom_field_choice_set",
	"user_id":                      "netbox_user",
	"tunnel_id":                    "netbox_vpn_tunnel",
	"tunnel_group_id":              "netbox_vpn_tunnel_group",
}

// resourceRefs covers arguments whose target depends on the resource, such as
// role_id or group_id. These take precedence over globalRefs.
var resourceRefs = map[string]map[string]string{
	"netbox_device":                   {"role_id": "netbox_device_role"},
	"netbox_virtual_machine":          {"role_id": "netbox_device_role"},
	"netbox_rack":                     {"role_id": "netbox_rack_role"},
	"netbox_inventory_item":           {"role_id": "netbox_inventory_item_role"},
	"netbox_contact_assignment":       {"role_id": "netbox_contact_role"},
	"netbox_prefix":                   {"role_id": "netbox_ipam_role"},
	"netbox_available_prefix":         {"role_id": "netbox_ipam_role", "parent_prefix_id": "netbox_prefix"},
	"netbox_ip_range":                 {"role_id": "netbox_ipam_role"},
	"netbox_vlan":                     {"role_id": "netbox_ipam_role", "group_id": "netbox_vlan_group"},
	"netbox_available_vlan":           {"role_id": "netbox_ipam_role", "group_id": "netbox_vlan_group"},
	"netbox_site":                     {"group_id": "netbox_site_group"},
	"netbox_tenant":                   {"group_id": "netbox_tenant_group"},
	"netbox_contact":                  {"group_id": "netbox_contact_group"},
	"netbox_wireless_lan":             {"group_id": "netbox_wireless_lan_group"},
	"netbox_fhrp_group_assignment":    {"group_id": "netbox_fhrp_group"},
	"netbox_circuit":                  {"provider_id": "netbox_circuit_provider", "type_id": "netbox_circuit_type"},
	"netbox_circuit_provider_network": {"provider_id": "netbox_circuit_provider"},
	"netbox_device_interface": {
		"lag_device_interface_id":    "netbox_device_interface",
		"parent_device_interface_id": "netbox_device_interface",
	},
	"netbox_device_interface_primary_mac_address":          {"interface_id": "netbox_device_interface"},
	"netbox_virtual_machine_interface_primary_mac_address": {"interface_id": "netbox_interface"},
	"netbox_region":             {"parent_region_id": "netbox_region"},
	"netbox_location":           {"parent_id": "netbox_location"},
	"netbox_site_group":         {"parent_id": "netbox_site_group"},
	"netbox_tenant_group":       {"parent_id": "netbox_tenant_group"},
	"netbox_contact_group":      {"parent_id": "netbox_contact_group"},
	"netbox_wireless_lan_group": {"parent_id": "netbox_wireless_lan_group"},
}

// References adds cross-resource references. A reference is only added when
// the argument exists at the top level of the resource's schema and the
// target resource is generated, so a schema change in a new Terraform
// provider version can't break generation.
func References() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		if r.TerraformResource == nil {
			return
		}
		add := func(field, target string) {
			if _, ok := groupKinds[target]; !ok {
				return
			}
			if _, ok := r.TerraformResource.Schema[field]; !ok {
				return
			}
			if r.References == nil {
				r.References = ujconfig.References{}
			}
			r.References[field] = ujconfig.Reference{TerraformName: target}
		}
		for field, target := range globalRefs {
			add(field, target)
		}
		for field, target := range resourceRefs[r.Name] {
			add(field, target)
		}
	}
}
