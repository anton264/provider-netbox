package config

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// groupKind is the API group (named after the NetBox app) and Kind for a
// Terraform resource. Only resources listed here are generated.
type groupKind struct {
	group string
	kind  string
}

// groupKinds covers every resource of e-breuninger/netbox v5.8.0. When you bump
// the Terraform provider version, add new resources here.
var groupKinds = map[string]groupKind{
	// dcim
	"netbox_device":                               {"dcim", "Device"},
	"netbox_device_interface":                     {"dcim", "DeviceInterface"},
	"netbox_device_interface_primary_mac_address": {"dcim", "DeviceInterfacePrimaryMACAddress"},
	"netbox_device_type":                          {"dcim", "DeviceType"},
	"netbox_manufacturer":                         {"dcim", "Manufacturer"},
	"netbox_interface_template":                   {"dcim", "InterfaceTemplate"},
	"netbox_platform":                             {"dcim", "Platform"},
	"netbox_device_primary_ip":                    {"dcim", "DevicePrimaryIP"},
	"netbox_device_oob_ip":                        {"dcim", "DeviceOOBIP"},
	"netbox_device_tag":                           {"dcim", "DeviceTag"},
	"netbox_device_role":                          {"dcim", "DeviceRole"},
	"netbox_site":                                 {"dcim", "Site"},
	"netbox_region":                               {"dcim", "Region"},
	"netbox_location":                             {"dcim", "Location"},
	"netbox_site_group":                           {"dcim", "SiteGroup"},
	"netbox_rack":                                 {"dcim", "Rack"},
	"netbox_rack_type":                            {"dcim", "RackType"},
	"netbox_rack_role":                            {"dcim", "RackRole"},
	"netbox_rack_reservation":                     {"dcim", "RackReservation"},
	"netbox_cable":                                {"dcim", "Cable"},
	"netbox_device_console_port":                  {"dcim", "DeviceConsolePort"},
	"netbox_device_console_server_port":           {"dcim", "DeviceConsoleServerPort"},
	"netbox_device_power_port":                    {"dcim", "DevicePowerPort"},
	"netbox_device_power_outlet":                  {"dcim", "DevicePowerOutlet"},
	"netbox_device_front_port":                    {"dcim", "DeviceFrontPort"},
	"netbox_device_rear_port":                     {"dcim", "DeviceRearPort"},
	"netbox_rear_port_template":                   {"dcim", "RearPortTemplate"},
	"netbox_device_module_bay":                    {"dcim", "DeviceModuleBay"},
	"netbox_device_bay":                           {"dcim", "DeviceBay"},
	"netbox_device_bay_template":                  {"dcim", "DeviceBayTemplate"},
	"netbox_console_port_template":                {"dcim", "ConsolePortTemplate"},
	"netbox_console_server_port_template":         {"dcim", "ConsoleServerPortTemplate"},
	"netbox_module":                               {"dcim", "Module"},
	"netbox_module_type":                          {"dcim", "ModuleType"},
	"netbox_module_bay_template":                  {"dcim", "ModuleBayTemplate"},
	"netbox_power_feed":                           {"dcim", "PowerFeed"},
	"netbox_power_panel":                          {"dcim", "PowerPanel"},
	"netbox_power_port_template":                  {"dcim", "PowerPortTemplate"},
	"netbox_power_outlet_template":                {"dcim", "PowerOutletTemplate"},
	"netbox_inventory_item_role":                  {"dcim", "InventoryItemRole"},
	"netbox_inventory_item":                       {"dcim", "InventoryItem"},
	"netbox_virtual_chassis":                      {"dcim", "VirtualChassis"},
	"netbox_mac_address":                          {"dcim", "MACAddress"},

	// ipam
	"netbox_available_ip_address":  {"ipam", "AvailableIPAddress"},
	"netbox_vrf":                   {"ipam", "VRF"},
	"netbox_ip_address":            {"ipam", "IPAddress"},
	"netbox_service":               {"ipam", "Service"},
	"netbox_prefix":                {"ipam", "Prefix"},
	"netbox_available_prefix":      {"ipam", "AvailablePrefix"},
	"netbox_vlan":                  {"ipam", "VLAN"},
	"netbox_vlan_group":            {"ipam", "VLANGroup"},
	"netbox_available_vlan":        {"ipam", "AvailableVLAN"},
	"netbox_ipam_role":             {"ipam", "Role"},
	"netbox_fhrp_group":            {"ipam", "FHRPGroup"},
	"netbox_fhrp_group_assignment": {"ipam", "FHRPGroupAssignment"},
	"netbox_ip_range":              {"ipam", "IPRange"},
	"netbox_aggregate":             {"ipam", "Aggregate"},
	"netbox_rir":                   {"ipam", "RIR"},
	"netbox_route_target":          {"ipam", "RouteTarget"},
	"netbox_asn":                   {"ipam", "ASN"},

	// virtualization
	"netbox_virtual_machine":                               {"virtualization", "VirtualMachine"},
	"netbox_virtual_machine_interface_primary_mac_address": {"virtualization", "VirtualMachineInterfacePrimaryMACAddress"},
	"netbox_cluster_type":                                  {"virtualization", "ClusterType"},
	"netbox_cluster":                                       {"virtualization", "Cluster"},
	"netbox_cluster_group":                                 {"virtualization", "ClusterGroup"},
	"netbox_interface":                                     {"virtualization", "VirtualMachineInterface"},
	"netbox_primary_ip":                                    {"virtualization", "VirtualMachinePrimaryIP"},
	"netbox_virtual_disk":                                  {"virtualization", "VirtualDisk"},

	// tenancy
	"netbox_contact":            {"tenancy", "Contact"},
	"netbox_contact_group":      {"tenancy", "ContactGroup"},
	"netbox_contact_assignment": {"tenancy", "ContactAssignment"},
	"netbox_contact_role":       {"tenancy", "ContactRole"},
	"netbox_tenant":             {"tenancy", "Tenant"},
	"netbox_tenant_group":       {"tenancy", "TenantGroup"},

	// circuits
	"netbox_circuit":                  {"circuits", "Circuit"},
	"netbox_circuit_type":             {"circuits", "CircuitType"},
	"netbox_circuit_provider_network": {"circuits", "CircuitProviderNetwork"},
	"netbox_circuit_provider":         {"circuits", "CircuitProvider"},
	"netbox_circuit_termination":      {"circuits", "CircuitTermination"},

	// users
	"netbox_user":       {"users", "User"},
	"netbox_group":      {"users", "Group"},
	"netbox_permission": {"users", "Permission"},
	"netbox_token":      {"users", "Token"},

	// extras
	"netbox_tag":                     {"extras", "Tag"},
	"netbox_custom_field":            {"extras", "CustomField"},
	"netbox_custom_field_choice_set": {"extras", "CustomFieldChoiceSet"},
	"netbox_webhook":                 {"extras", "Webhook"},
	"netbox_config_template":         {"extras", "ConfigTemplate"},
	"netbox_event_rule":              {"extras", "EventRule"},
	"netbox_config_context":          {"extras", "ConfigContext"},

	// vpn
	"netbox_vpn_tunnel_group":       {"vpn", "TunnelGroup"},
	"netbox_vpn_tunnel":             {"vpn", "Tunnel"},
	"netbox_vpn_tunnel_termination": {"vpn", "TunnelTermination"},

	// wireless
	"netbox_wireless_lan_group": {"wireless", "WirelessLANGroup"},
	"netbox_wireless_lan":       {"wireless", "WirelessLAN"},
}

// GroupKindOverrides groups resources by NetBox app (dcim, ipam, ...) instead
// of upjet's default of using the first word of the resource name.
func GroupKindOverrides() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		if gk, ok := groupKinds[r.Name]; ok {
			r.ShortGroup = gk.group
			r.Kind = gk.kind
		}
	}
}
