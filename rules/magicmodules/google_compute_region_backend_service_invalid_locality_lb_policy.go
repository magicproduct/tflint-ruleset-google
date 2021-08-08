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
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule checks the pattern is valid
type GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule returns new rule with default attributes
func NewGoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule() *GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule {
	return &GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule{
		resourceType:  "google_compute_region_backend_service",
		attributeName: "locality_lb_policy",
	}
}

// Name returns the rule name
func (r *GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule) Name() string {
	return "google_compute_region_backend_service_invalid_locality_lb_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
