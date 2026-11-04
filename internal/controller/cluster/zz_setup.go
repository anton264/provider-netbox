// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	circuit "github.com/anton264/provider-netbox/internal/controller/cluster/circuits/circuit"
	circuitprovider "github.com/anton264/provider-netbox/internal/controller/cluster/circuits/circuitprovider"
	circuitprovidernetwork "github.com/anton264/provider-netbox/internal/controller/cluster/circuits/circuitprovidernetwork"
	circuittermination "github.com/anton264/provider-netbox/internal/controller/cluster/circuits/circuittermination"
	circuittype "github.com/anton264/provider-netbox/internal/controller/cluster/circuits/circuittype"
	cable "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/cable"
	consoleporttemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/consoleporttemplate"
	consoleserverporttemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/consoleserverporttemplate"
	device "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/device"
	devicebay "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicebay"
	devicebaytemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicebaytemplate"
	deviceconsoleport "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/deviceconsoleport"
	deviceconsoleserverport "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/deviceconsoleserverport"
	devicefrontport "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicefrontport"
	deviceinterface "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/deviceinterface"
	deviceinterfaceprimarymacaddress "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/deviceinterfaceprimarymacaddress"
	devicemodulebay "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicemodulebay"
	deviceoobip "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/deviceoobip"
	devicepoweroutlet "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicepoweroutlet"
	devicepowerport "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicepowerport"
	deviceprimaryip "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/deviceprimaryip"
	devicerearport "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicerearport"
	devicerole "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicerole"
	devicetag "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicetag"
	devicetype "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/devicetype"
	interfacetemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/interfacetemplate"
	inventoryitem "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/inventoryitem"
	inventoryitemrole "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/inventoryitemrole"
	location "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/location"
	macaddress "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/macaddress"
	manufacturer "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/manufacturer"
	module "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/module"
	modulebaytemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/modulebaytemplate"
	moduletype "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/moduletype"
	platform "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/platform"
	powerfeed "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/powerfeed"
	poweroutlettemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/poweroutlettemplate"
	powerpanel "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/powerpanel"
	powerporttemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/powerporttemplate"
	rack "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/rack"
	rackreservation "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/rackreservation"
	rackrole "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/rackrole"
	racktype "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/racktype"
	rearporttemplate "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/rearporttemplate"
	region "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/region"
	site "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/site"
	sitegroup "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/sitegroup"
	virtualchassis "github.com/anton264/provider-netbox/internal/controller/cluster/dcim/virtualchassis"
	configcontext "github.com/anton264/provider-netbox/internal/controller/cluster/extras/configcontext"
	configtemplate "github.com/anton264/provider-netbox/internal/controller/cluster/extras/configtemplate"
	customfield "github.com/anton264/provider-netbox/internal/controller/cluster/extras/customfield"
	customfieldchoiceset "github.com/anton264/provider-netbox/internal/controller/cluster/extras/customfieldchoiceset"
	eventrule "github.com/anton264/provider-netbox/internal/controller/cluster/extras/eventrule"
	tag "github.com/anton264/provider-netbox/internal/controller/cluster/extras/tag"
	webhook "github.com/anton264/provider-netbox/internal/controller/cluster/extras/webhook"
	aggregate "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/aggregate"
	asn "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/asn"
	availableipaddress "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/availableipaddress"
	availableprefix "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/availableprefix"
	availablevlan "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/availablevlan"
	fhrpgroup "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/fhrpgroup"
	fhrpgroupassignment "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/fhrpgroupassignment"
	ipaddress "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/ipaddress"
	iprange "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/iprange"
	prefix "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/prefix"
	rir "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/rir"
	role "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/role"
	routetarget "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/routetarget"
	service "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/service"
	vlan "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/vlan"
	vlangroup "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/vlangroup"
	vrf "github.com/anton264/provider-netbox/internal/controller/cluster/ipam/vrf"
	providerconfig "github.com/anton264/provider-netbox/internal/controller/cluster/providerconfig"
	contact "github.com/anton264/provider-netbox/internal/controller/cluster/tenancy/contact"
	contactassignment "github.com/anton264/provider-netbox/internal/controller/cluster/tenancy/contactassignment"
	contactgroup "github.com/anton264/provider-netbox/internal/controller/cluster/tenancy/contactgroup"
	contactrole "github.com/anton264/provider-netbox/internal/controller/cluster/tenancy/contactrole"
	tenant "github.com/anton264/provider-netbox/internal/controller/cluster/tenancy/tenant"
	tenantgroup "github.com/anton264/provider-netbox/internal/controller/cluster/tenancy/tenantgroup"
	group "github.com/anton264/provider-netbox/internal/controller/cluster/users/group"
	permission "github.com/anton264/provider-netbox/internal/controller/cluster/users/permission"
	token "github.com/anton264/provider-netbox/internal/controller/cluster/users/token"
	user "github.com/anton264/provider-netbox/internal/controller/cluster/users/user"
	cluster "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/cluster"
	clustergroup "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/clustergroup"
	clustertype "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/clustertype"
	virtualdisk "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/virtualdisk"
	virtualmachine "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/virtualmachine"
	virtualmachineinterface "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/virtualmachineinterface"
	virtualmachineinterfaceprimarymacaddress "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/virtualmachineinterfaceprimarymacaddress"
	virtualmachineprimaryip "github.com/anton264/provider-netbox/internal/controller/cluster/virtualization/virtualmachineprimaryip"
	tunnel "github.com/anton264/provider-netbox/internal/controller/cluster/vpn/tunnel"
	tunnelgroup "github.com/anton264/provider-netbox/internal/controller/cluster/vpn/tunnelgroup"
	tunneltermination "github.com/anton264/provider-netbox/internal/controller/cluster/vpn/tunneltermination"
	wirelesslan "github.com/anton264/provider-netbox/internal/controller/cluster/wireless/wirelesslan"
	wirelesslangroup "github.com/anton264/provider-netbox/internal/controller/cluster/wireless/wirelesslangroup"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		circuit.Setup,
		circuitprovider.Setup,
		circuitprovidernetwork.Setup,
		circuittermination.Setup,
		circuittype.Setup,
		cable.Setup,
		consoleporttemplate.Setup,
		consoleserverporttemplate.Setup,
		device.Setup,
		devicebay.Setup,
		devicebaytemplate.Setup,
		deviceconsoleport.Setup,
		deviceconsoleserverport.Setup,
		devicefrontport.Setup,
		deviceinterface.Setup,
		deviceinterfaceprimarymacaddress.Setup,
		devicemodulebay.Setup,
		deviceoobip.Setup,
		devicepoweroutlet.Setup,
		devicepowerport.Setup,
		deviceprimaryip.Setup,
		devicerearport.Setup,
		devicerole.Setup,
		devicetag.Setup,
		devicetype.Setup,
		interfacetemplate.Setup,
		inventoryitem.Setup,
		inventoryitemrole.Setup,
		location.Setup,
		macaddress.Setup,
		manufacturer.Setup,
		module.Setup,
		modulebaytemplate.Setup,
		moduletype.Setup,
		platform.Setup,
		powerfeed.Setup,
		poweroutlettemplate.Setup,
		powerpanel.Setup,
		powerporttemplate.Setup,
		rack.Setup,
		rackreservation.Setup,
		rackrole.Setup,
		racktype.Setup,
		rearporttemplate.Setup,
		region.Setup,
		site.Setup,
		sitegroup.Setup,
		virtualchassis.Setup,
		configcontext.Setup,
		configtemplate.Setup,
		customfield.Setup,
		customfieldchoiceset.Setup,
		eventrule.Setup,
		tag.Setup,
		webhook.Setup,
		aggregate.Setup,
		asn.Setup,
		availableipaddress.Setup,
		availableprefix.Setup,
		availablevlan.Setup,
		fhrpgroup.Setup,
		fhrpgroupassignment.Setup,
		ipaddress.Setup,
		iprange.Setup,
		prefix.Setup,
		rir.Setup,
		role.Setup,
		routetarget.Setup,
		service.Setup,
		vlan.Setup,
		vlangroup.Setup,
		vrf.Setup,
		providerconfig.Setup,
		contact.Setup,
		contactassignment.Setup,
		contactgroup.Setup,
		contactrole.Setup,
		tenant.Setup,
		tenantgroup.Setup,
		group.Setup,
		permission.Setup,
		token.Setup,
		user.Setup,
		cluster.Setup,
		clustergroup.Setup,
		clustertype.Setup,
		virtualdisk.Setup,
		virtualmachine.Setup,
		virtualmachineinterface.Setup,
		virtualmachineinterfaceprimarymacaddress.Setup,
		virtualmachineprimaryip.Setup,
		tunnel.Setup,
		tunnelgroup.Setup,
		tunneltermination.Setup,
		wirelesslan.Setup,
		wirelesslangroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		circuit.SetupGated,
		circuitprovider.SetupGated,
		circuitprovidernetwork.SetupGated,
		circuittermination.SetupGated,
		circuittype.SetupGated,
		cable.SetupGated,
		consoleporttemplate.SetupGated,
		consoleserverporttemplate.SetupGated,
		device.SetupGated,
		devicebay.SetupGated,
		devicebaytemplate.SetupGated,
		deviceconsoleport.SetupGated,
		deviceconsoleserverport.SetupGated,
		devicefrontport.SetupGated,
		deviceinterface.SetupGated,
		deviceinterfaceprimarymacaddress.SetupGated,
		devicemodulebay.SetupGated,
		deviceoobip.SetupGated,
		devicepoweroutlet.SetupGated,
		devicepowerport.SetupGated,
		deviceprimaryip.SetupGated,
		devicerearport.SetupGated,
		devicerole.SetupGated,
		devicetag.SetupGated,
		devicetype.SetupGated,
		interfacetemplate.SetupGated,
		inventoryitem.SetupGated,
		inventoryitemrole.SetupGated,
		location.SetupGated,
		macaddress.SetupGated,
		manufacturer.SetupGated,
		module.SetupGated,
		modulebaytemplate.SetupGated,
		moduletype.SetupGated,
		platform.SetupGated,
		powerfeed.SetupGated,
		poweroutlettemplate.SetupGated,
		powerpanel.SetupGated,
		powerporttemplate.SetupGated,
		rack.SetupGated,
		rackreservation.SetupGated,
		rackrole.SetupGated,
		racktype.SetupGated,
		rearporttemplate.SetupGated,
		region.SetupGated,
		site.SetupGated,
		sitegroup.SetupGated,
		virtualchassis.SetupGated,
		configcontext.SetupGated,
		configtemplate.SetupGated,
		customfield.SetupGated,
		customfieldchoiceset.SetupGated,
		eventrule.SetupGated,
		tag.SetupGated,
		webhook.SetupGated,
		aggregate.SetupGated,
		asn.SetupGated,
		availableipaddress.SetupGated,
		availableprefix.SetupGated,
		availablevlan.SetupGated,
		fhrpgroup.SetupGated,
		fhrpgroupassignment.SetupGated,
		ipaddress.SetupGated,
		iprange.SetupGated,
		prefix.SetupGated,
		rir.SetupGated,
		role.SetupGated,
		routetarget.SetupGated,
		service.SetupGated,
		vlan.SetupGated,
		vlangroup.SetupGated,
		vrf.SetupGated,
		providerconfig.SetupGated,
		contact.SetupGated,
		contactassignment.SetupGated,
		contactgroup.SetupGated,
		contactrole.SetupGated,
		tenant.SetupGated,
		tenantgroup.SetupGated,
		group.SetupGated,
		permission.SetupGated,
		token.SetupGated,
		user.SetupGated,
		cluster.SetupGated,
		clustergroup.SetupGated,
		clustertype.SetupGated,
		virtualdisk.SetupGated,
		virtualmachine.SetupGated,
		virtualmachineinterface.SetupGated,
		virtualmachineinterfaceprimarymacaddress.SetupGated,
		virtualmachineprimaryip.SetupGated,
		tunnel.SetupGated,
		tunnelgroup.SetupGated,
		tunneltermination.SetupGated,
		wirelesslan.SetupGated,
		wirelesslangroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		circuit.SetupWebhookWithManager,
		circuitprovider.SetupWebhookWithManager,
		circuitprovidernetwork.SetupWebhookWithManager,
		circuittermination.SetupWebhookWithManager,
		circuittype.SetupWebhookWithManager,
		cable.SetupWebhookWithManager,
		consoleporttemplate.SetupWebhookWithManager,
		consoleserverporttemplate.SetupWebhookWithManager,
		device.SetupWebhookWithManager,
		devicebay.SetupWebhookWithManager,
		devicebaytemplate.SetupWebhookWithManager,
		deviceconsoleport.SetupWebhookWithManager,
		deviceconsoleserverport.SetupWebhookWithManager,
		devicefrontport.SetupWebhookWithManager,
		deviceinterface.SetupWebhookWithManager,
		deviceinterfaceprimarymacaddress.SetupWebhookWithManager,
		devicemodulebay.SetupWebhookWithManager,
		deviceoobip.SetupWebhookWithManager,
		devicepoweroutlet.SetupWebhookWithManager,
		devicepowerport.SetupWebhookWithManager,
		deviceprimaryip.SetupWebhookWithManager,
		devicerearport.SetupWebhookWithManager,
		devicerole.SetupWebhookWithManager,
		devicetag.SetupWebhookWithManager,
		devicetype.SetupWebhookWithManager,
		interfacetemplate.SetupWebhookWithManager,
		inventoryitem.SetupWebhookWithManager,
		inventoryitemrole.SetupWebhookWithManager,
		location.SetupWebhookWithManager,
		macaddress.SetupWebhookWithManager,
		manufacturer.SetupWebhookWithManager,
		module.SetupWebhookWithManager,
		modulebaytemplate.SetupWebhookWithManager,
		moduletype.SetupWebhookWithManager,
		platform.SetupWebhookWithManager,
		powerfeed.SetupWebhookWithManager,
		poweroutlettemplate.SetupWebhookWithManager,
		powerpanel.SetupWebhookWithManager,
		powerporttemplate.SetupWebhookWithManager,
		rack.SetupWebhookWithManager,
		rackreservation.SetupWebhookWithManager,
		rackrole.SetupWebhookWithManager,
		racktype.SetupWebhookWithManager,
		rearporttemplate.SetupWebhookWithManager,
		region.SetupWebhookWithManager,
		site.SetupWebhookWithManager,
		sitegroup.SetupWebhookWithManager,
		virtualchassis.SetupWebhookWithManager,
		configcontext.SetupWebhookWithManager,
		configtemplate.SetupWebhookWithManager,
		customfield.SetupWebhookWithManager,
		customfieldchoiceset.SetupWebhookWithManager,
		eventrule.SetupWebhookWithManager,
		tag.SetupWebhookWithManager,
		webhook.SetupWebhookWithManager,
		aggregate.SetupWebhookWithManager,
		asn.SetupWebhookWithManager,
		availableipaddress.SetupWebhookWithManager,
		availableprefix.SetupWebhookWithManager,
		availablevlan.SetupWebhookWithManager,
		fhrpgroup.SetupWebhookWithManager,
		fhrpgroupassignment.SetupWebhookWithManager,
		ipaddress.SetupWebhookWithManager,
		iprange.SetupWebhookWithManager,
		prefix.SetupWebhookWithManager,
		rir.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		routetarget.SetupWebhookWithManager,
		service.SetupWebhookWithManager,
		vlan.SetupWebhookWithManager,
		vlangroup.SetupWebhookWithManager,
		vrf.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		contact.SetupWebhookWithManager,
		contactassignment.SetupWebhookWithManager,
		contactgroup.SetupWebhookWithManager,
		contactrole.SetupWebhookWithManager,
		tenant.SetupWebhookWithManager,
		tenantgroup.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		permission.SetupWebhookWithManager,
		token.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		clustergroup.SetupWebhookWithManager,
		clustertype.SetupWebhookWithManager,
		virtualdisk.SetupWebhookWithManager,
		virtualmachine.SetupWebhookWithManager,
		virtualmachineinterface.SetupWebhookWithManager,
		virtualmachineinterfaceprimarymacaddress.SetupWebhookWithManager,
		virtualmachineprimaryip.SetupWebhookWithManager,
		tunnel.SetupWebhookWithManager,
		tunnelgroup.SetupWebhookWithManager,
		tunneltermination.SetupWebhookWithManager,
		wirelesslan.SetupWebhookWithManager,
		wirelesslangroup.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
