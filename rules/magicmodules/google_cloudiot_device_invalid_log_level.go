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

// GoogleCloudiotDeviceInvalidLogLevelRule checks the pattern is valid
type GoogleCloudiotDeviceInvalidLogLevelRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleCloudiotDeviceInvalidLogLevelRule returns new rule with default attributes
func NewGoogleCloudiotDeviceInvalidLogLevelRule() *GoogleCloudiotDeviceInvalidLogLevelRule {
	return &GoogleCloudiotDeviceInvalidLogLevelRule{
		resourceType:  "google_cloudiot_device",
		attributeName: "log_level",
	}
}

// Name returns the rule name
func (r *GoogleCloudiotDeviceInvalidLogLevelRule) Name() string {
	return "google_cloudiot_device_invalid_log_level"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleCloudiotDeviceInvalidLogLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleCloudiotDeviceInvalidLogLevelRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleCloudiotDeviceInvalidLogLevelRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleCloudiotDeviceInvalidLogLevelRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"NONE", "ERROR", "INFO", "DEBUG", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
