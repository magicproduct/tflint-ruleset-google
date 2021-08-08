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

// GoogleBigqueryRoutineInvalidRoutineTypeRule checks the pattern is valid
type GoogleBigqueryRoutineInvalidRoutineTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleBigqueryRoutineInvalidRoutineTypeRule returns new rule with default attributes
func NewGoogleBigqueryRoutineInvalidRoutineTypeRule() *GoogleBigqueryRoutineInvalidRoutineTypeRule {
	return &GoogleBigqueryRoutineInvalidRoutineTypeRule{
		resourceType:  "google_bigquery_routine",
		attributeName: "routine_type",
	}
}

// Name returns the rule name
func (r *GoogleBigqueryRoutineInvalidRoutineTypeRule) Name() string {
	return "google_bigquery_routine_invalid_routine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleBigqueryRoutineInvalidRoutineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleBigqueryRoutineInvalidRoutineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleBigqueryRoutineInvalidRoutineTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleBigqueryRoutineInvalidRoutineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"SCALAR_FUNCTION", "PROCEDURE", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
