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

// GoogleBigQueryRoutineInvalidDeterminismLevelRule checks the pattern is valid
type GoogleBigQueryRoutineInvalidDeterminismLevelRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleBigQueryRoutineInvalidDeterminismLevelRule returns new rule with default attributes
func NewGoogleBigQueryRoutineInvalidDeterminismLevelRule() *GoogleBigQueryRoutineInvalidDeterminismLevelRule {
	return &GoogleBigQueryRoutineInvalidDeterminismLevelRule{
		resourceType:  "google_big_query_routine",
		attributeName: "determinism_level",
	}
}

// Name returns the rule name
func (r *GoogleBigQueryRoutineInvalidDeterminismLevelRule) Name() string {
	return "google_big_query_routine_invalid_determinism_level"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleBigQueryRoutineInvalidDeterminismLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleBigQueryRoutineInvalidDeterminismLevelRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleBigQueryRoutineInvalidDeterminismLevelRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleBigQueryRoutineInvalidDeterminismLevelRule) Check(runner tflint.Runner) error {
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
			validateFunc := validation.StringInSlice([]string{"DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC", ""}, false)

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