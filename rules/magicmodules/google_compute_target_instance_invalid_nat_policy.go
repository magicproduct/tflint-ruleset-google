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

// GoogleComputeTargetInstanceInvalidNatPolicyRule checks the pattern is valid
type GoogleComputeTargetInstanceInvalidNatPolicyRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeTargetInstanceInvalidNatPolicyRule returns new rule with default attributes
func NewGoogleComputeTargetInstanceInvalidNatPolicyRule() *GoogleComputeTargetInstanceInvalidNatPolicyRule {
	return &GoogleComputeTargetInstanceInvalidNatPolicyRule{
		resourceType:  "google_compute_target_instance",
		attributeName: "nat_policy",
	}
}

// Name returns the rule name
func (r *GoogleComputeTargetInstanceInvalidNatPolicyRule) Name() string {
	return "google_compute_target_instance_invalid_nat_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeTargetInstanceInvalidNatPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeTargetInstanceInvalidNatPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeTargetInstanceInvalidNatPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeTargetInstanceInvalidNatPolicyRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"NO_NAT", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
