// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	frominstance "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ami/frominstance"
	launchpermission "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ami/launchpermission"
	policy "github.com/multinik/provider-crossplane-k2cloud/internal/controller/autoscaling/policy"
	ami "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/ami"
	eip "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/eip"
	instance "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/instance"
	lb "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/lb"
	route "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/route"
	subnet "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/subnet"
	vpc "github.com/multinik/provider-crossplane-k2cloud/internal/controller/aws/vpc"
	plan "github.com/multinik/provider-crossplane-k2cloud/internal/controller/backup/plan"
	selection "github.com/multinik/provider-crossplane-k2cloud/internal/controller/backup/selection"
	vaultdefault "github.com/multinik/provider-crossplane-k2cloud/internal/controller/backup/vaultdefault"
	metricalarm "github.com/multinik/provider-crossplane-k2cloud/internal/controller/cloudwatch/metricalarm"
	gateway "github.com/multinik/provider-crossplane-k2cloud/internal/controller/customer/gateway"
	networkacl "github.com/multinik/provider-crossplane-k2cloud/internal/controller/default/networkacl"
	routetable "github.com/multinik/provider-crossplane-k2cloud/internal/controller/default/routetable"
	securitygroup "github.com/multinik/provider-crossplane-k2cloud/internal/controller/default/securitygroup"
	vpcdefault "github.com/multinik/provider-crossplane-k2cloud/internal/controller/default/vpc"
	vpcdhcpoptions "github.com/multinik/provider-crossplane-k2cloud/internal/controller/default/vpcdhcpoptions"
	gatewaydx "github.com/multinik/provider-crossplane-k2cloud/internal/controller/dx/gateway"
	gatewayassociation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/dx/gatewayassociation"
	transitvirtualinterface "github.com/multinik/provider-crossplane-k2cloud/internal/controller/dx/transitvirtualinterface"
	snapshot "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ebs/snapshot"
	snapshotimport "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ebs/snapshotimport"
	volume "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ebs/volume"
	tag "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/tag"
	transitgateway "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/transitgateway"
	transitgatewayroute "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/multinik/provider-crossplane-k2cloud/internal/controller/ec2/transitgatewayvpcattachment"
	association "github.com/multinik/provider-crossplane-k2cloud/internal/controller/eip/association"
	cluster "github.com/multinik/provider-crossplane-k2cloud/internal/controller/eks/cluster"
	nodegroup "github.com/multinik/provider-crossplane-k2cloud/internal/controller/eks/nodegroup"
	group "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/group"
	groupmembership "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/grouppolicyattachment"
	policyiam "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/policy"
	project "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/project"
	user "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/user"
	usergroupmembership "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/usergroupmembership"
	userpolicyattachment "github.com/multinik/provider-crossplane-k2cloud/internal/controller/iam/userpolicyattachment"
	gatewayinternet "github.com/multinik/provider-crossplane-k2cloud/internal/controller/internet/gateway"
	gatewayattachment "github.com/multinik/provider-crossplane-k2cloud/internal/controller/internet/gatewayattachment"
	pair "github.com/multinik/provider-crossplane-k2cloud/internal/controller/key/pair"
	template "github.com/multinik/provider-crossplane-k2cloud/internal/controller/launch/template"
	listener "github.com/multinik/provider-crossplane-k2cloud/internal/controller/lb/listener"
	targetgroup "github.com/multinik/provider-crossplane-k2cloud/internal/controller/lb/targetgroup"
	targetgroupattachment "github.com/multinik/provider-crossplane-k2cloud/internal/controller/lb/targetgroupattachment"
	routetableassociation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/main/routetableassociation"
	acl "github.com/multinik/provider-crossplane-k2cloud/internal/controller/network/acl"
	aclassociation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/network/aclassociation"
	aclrule "github.com/multinik/provider-crossplane-k2cloud/internal/controller/network/aclrule"
	backup "github.com/multinik/provider-crossplane-k2cloud/internal/controller/paas/backup"
	service "github.com/multinik/provider-crossplane-k2cloud/internal/controller/paas/service"
	groupplacement "github.com/multinik/provider-crossplane-k2cloud/internal/controller/placement/group"
	providerconfig "github.com/multinik/provider-crossplane-k2cloud/internal/controller/providerconfig"
	table "github.com/multinik/provider-crossplane-k2cloud/internal/controller/route/table"
	tableassociation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/route/tableassociation"
	record "github.com/multinik/provider-crossplane-k2cloud/internal/controller/route53/record"
	zone "github.com/multinik/provider-crossplane-k2cloud/internal/controller/route53/zone"
	bucket "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucket"
	bucketacl "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketacl"
	bucketcorsconfiguration "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketcorsconfiguration"
	bucketlifecycleconfiguration "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketlifecycleconfiguration"
	bucketobject "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketobject"
	bucketpolicy "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketpolicy"
	bucketrequestpaymentconfiguration "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/object"
	objectcopy "github.com/multinik/provider-crossplane-k2cloud/internal/controller/s3/objectcopy"
	groupsecurity "github.com/multinik/provider-crossplane-k2cloud/internal/controller/security/group"
	grouprule "github.com/multinik/provider-crossplane-k2cloud/internal/controller/security/grouprule"
	createvolumepermission "github.com/multinik/provider-crossplane-k2cloud/internal/controller/snapshot/createvolumepermission"
	attachment "github.com/multinik/provider-crossplane-k2cloud/internal/controller/volume/attachment"
	dhcpoptions "github.com/multinik/provider-crossplane-k2cloud/internal/controller/vpc/dhcpoptions"
	dhcpoptionsassociation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/vpc/dhcpoptionsassociation"
	connection "github.com/multinik/provider-crossplane-k2cloud/internal/controller/vpn/connection"
	gatewayroutepropagation "github.com/multinik/provider-crossplane-k2cloud/internal/controller/vpn/gatewayroutepropagation"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		frominstance.Setup,
		launchpermission.Setup,
		policy.Setup,
		ami.Setup,
		eip.Setup,
		instance.Setup,
		lb.Setup,
		route.Setup,
		subnet.Setup,
		vpc.Setup,
		plan.Setup,
		selection.Setup,
		vaultdefault.Setup,
		metricalarm.Setup,
		gateway.Setup,
		networkacl.Setup,
		routetable.Setup,
		securitygroup.Setup,
		vpcdefault.Setup,
		vpcdhcpoptions.Setup,
		gatewaydx.Setup,
		gatewayassociation.Setup,
		transitvirtualinterface.Setup,
		snapshot.Setup,
		snapshotimport.Setup,
		volume.Setup,
		tag.Setup,
		transitgateway.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		association.Setup,
		cluster.Setup,
		nodegroup.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		policyiam.Setup,
		project.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userpolicyattachment.Setup,
		gatewayinternet.Setup,
		gatewayattachment.Setup,
		pair.Setup,
		template.Setup,
		listener.Setup,
		targetgroup.Setup,
		targetgroupattachment.Setup,
		routetableassociation.Setup,
		acl.Setup,
		aclassociation.Setup,
		aclrule.Setup,
		backup.Setup,
		service.Setup,
		groupplacement.Setup,
		providerconfig.Setup,
		table.Setup,
		tableassociation.Setup,
		record.Setup,
		zone.Setup,
		bucket.Setup,
		bucketacl.Setup,
		bucketcorsconfiguration.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		object.Setup,
		objectcopy.Setup,
		groupsecurity.Setup,
		grouprule.Setup,
		createvolumepermission.Setup,
		attachment.Setup,
		dhcpoptions.Setup,
		dhcpoptionsassociation.Setup,
		connection.Setup,
		gatewayroutepropagation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
