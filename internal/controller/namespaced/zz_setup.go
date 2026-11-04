// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	circuit "github.com/anton264/provider-netbox/internal/controller/namespaced/circuits/circuit"
	circuitprovider "github.com/anton264/provider-netbox/internal/controller/namespaced/circuits/circuitprovider"
	circuitprovidernetwork "github.com/anton264/provider-netbox/internal/controller/namespaced/circuits/circuitprovidernetwork"
	circuittermination "github.com/anton264/provider-netbox/internal/controller/namespaced/circuits/circuittermination"
	circuittype "github.com/anton264/provider-netbox/internal/controller/namespaced/circuits/circuittype"
	cable "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/cable"
	consoleporttemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/consoleporttemplate"
	consoleserverporttemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/consoleserverporttemplate"
	device "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/device"
	devicebay "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicebay"
	devicebaytemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicebaytemplate"
	deviceconsoleport "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/deviceconsoleport"
	deviceconsoleserverport "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/deviceconsoleserverport"
	devicefrontport "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicefrontport"
	deviceinterface "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/deviceinterface"
	deviceinterfaceprimarymacaddress "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/deviceinterfaceprimarymacaddress"
	devicemodulebay "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicemodulebay"
	deviceoobip "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/deviceoobip"
	devicepoweroutlet "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicepoweroutlet"
	devicepowerport "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicepowerport"
	deviceprimaryip "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/deviceprimaryip"
	devicerearport "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicerearport"
	devicerole "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicerole"
	devicetag "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicetag"
	devicetype "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/devicetype"
	interfacetemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/interfacetemplate"
	inventoryitem "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/inventoryitem"
	inventoryitemrole "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/inventoryitemrole"
	location "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/location"
	macaddress "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/macaddress"
	manufacturer "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/manufacturer"
	module "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/module"
	modulebaytemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/modulebaytemplate"
	moduletype "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/moduletype"
	platform "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/platform"
	powerfeed "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/powerfeed"
	poweroutlettemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/poweroutlettemplate"
	powerpanel "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/powerpanel"
	powerporttemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/powerporttemplate"
	rack "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/rack"
	rackreservation "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/rackreservation"
	rackrole "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/rackrole"
	racktype "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/racktype"
	rearporttemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/rearporttemplate"
	region "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/region"
	site "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/site"
	sitegroup "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/sitegroup"
	virtualchassis "github.com/anton264/provider-netbox/internal/controller/namespaced/dcim/virtualchassis"
	configcontext "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/configcontext"
	configtemplate "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/configtemplate"
	customfield "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/customfield"
	customfieldchoiceset "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/customfieldchoiceset"
	eventrule "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/eventrule"
	tag "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/tag"
	webhook "github.com/anton264/provider-netbox/internal/controller/namespaced/extras/webhook"
	aggregate "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/aggregate"
	asn "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/asn"
	availableipaddress "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/availableipaddress"
	availableprefix "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/availableprefix"
	availablevlan "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/availablevlan"
	fhrpgroup "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/fhrpgroup"
	fhrpgroupassignment "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/fhrpgroupassignment"
	ipaddress "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/ipaddress"
	iprange "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/iprange"
	prefix "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/prefix"
	rir "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/rir"
	role "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/role"
	routetarget "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/routetarget"
	service "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/service"
	vlan "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/vlan"
	vlangroup "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/vlangroup"
	vrf "github.com/anton264/provider-netbox/internal/controller/namespaced/ipam/vrf"
	providerconfig "github.com/anton264/provider-netbox/internal/controller/namespaced/providerconfig"
	contact "github.com/anton264/provider-netbox/internal/controller/namespaced/tenancy/contact"
	contactassignment "github.com/anton264/provider-netbox/internal/controller/namespaced/tenancy/contactassignment"
	contactgroup "github.com/anton264/provider-netbox/internal/controller/namespaced/tenancy/contactgroup"
	contactrole "github.com/anton264/provider-netbox/internal/controller/namespaced/tenancy/contactrole"
	tenant "github.com/anton264/provider-netbox/internal/controller/namespaced/tenancy/tenant"
	tenantgroup "github.com/anton264/provider-netbox/internal/controller/namespaced/tenancy/tenantgroup"
	group "github.com/anton264/provider-netbox/internal/controller/namespaced/users/group"
	permission "github.com/anton264/provider-netbox/internal/controller/namespaced/users/permission"
	token "github.com/anton264/provider-netbox/internal/controller/namespaced/users/token"
	user "github.com/anton264/provider-netbox/internal/controller/namespaced/users/user"
	cluster "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/cluster"
	clustergroup "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/clustergroup"
	clustertype "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/clustertype"
	virtualdisk "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/virtualdisk"
	virtualmachine "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/virtualmachine"
	virtualmachineinterface "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/virtualmachineinterface"
	virtualmachineinterfaceprimarymacaddress "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/virtualmachineinterfaceprimarymacaddress"
	virtualmachineprimaryip "github.com/anton264/provider-netbox/internal/controller/namespaced/virtualization/virtualmachineprimaryip"
	tunnel "github.com/anton264/provider-netbox/internal/controller/namespaced/vpn/tunnel"
	tunnelgroup "github.com/anton264/provider-netbox/internal/controller/namespaced/vpn/tunnelgroup"
	tunneltermination "github.com/anton264/provider-netbox/internal/controller/namespaced/vpn/tunneltermination"
	wirelesslan "github.com/anton264/provider-netbox/internal/controller/namespaced/wireless/wirelesslan"
	wirelesslangroup "github.com/anton264/provider-netbox/internal/controller/namespaced/wireless/wirelesslangroup"
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
