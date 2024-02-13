// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule checks the pattern is valid
type GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule returns new rule with default attributes
func NewGoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule() *GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule {
	return &GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule{
		resourceType:  "google_vmwareengine_network_peering",
		attributeName: "peer_network_type",
	}
}

// Name returns the rule name
func (r *GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule) Name() string {
	return "google_vmwareengine_network_peering_invalid_peer_network_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validation.StringInSlice([]string{"STANDARD", "VMWARE_ENGINE_NETWORK", "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE"}, false)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
