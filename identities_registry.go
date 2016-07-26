// Copyright 2016 Aporeto Inc.

package gaia

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(MyNamespaceIdentity)
	elemental.RegisterIdentity(HealthReportIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(ServerIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(MapEdgeIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(MapNodeIdentity)
	elemental.RegisterIdentity(AuthenticatorIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(UserIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(RenderedPoliciesIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
}
