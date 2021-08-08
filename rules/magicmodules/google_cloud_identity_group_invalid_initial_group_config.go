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

// GoogleCloudIdentityGroupInvalidInitialGroupConfigRule checks the pattern is valid
type GoogleCloudIdentityGroupInvalidInitialGroupConfigRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleCloudIdentityGroupInvalidInitialGroupConfigRule returns new rule with default attributes
func NewGoogleCloudIdentityGroupInvalidInitialGroupConfigRule() *GoogleCloudIdentityGroupInvalidInitialGroupConfigRule {
	return &GoogleCloudIdentityGroupInvalidInitialGroupConfigRule{
		resourceType:  "google_cloud_identity_group",
		attributeName: "initial_group_config",
	}
}

// Name returns the rule name
func (r *GoogleCloudIdentityGroupInvalidInitialGroupConfigRule) Name() string {
	return "google_cloud_identity_group_invalid_initial_group_config"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleCloudIdentityGroupInvalidInitialGroupConfigRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleCloudIdentityGroupInvalidInitialGroupConfigRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleCloudIdentityGroupInvalidInitialGroupConfigRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleCloudIdentityGroupInvalidInitialGroupConfigRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
