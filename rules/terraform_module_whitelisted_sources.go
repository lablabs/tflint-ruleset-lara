package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
)

type TerraformModuleWhitelistedSourcesRule struct {
	tflint.DefaultRule
	attributeName string
}

// NewTerraformModuleWhitelistedSourcesRule returns a new rule
func NewTerraformModuleWhitelistedSourcesRule() *TerraformModuleWhitelistedSourcesRule {
	return &TerraformModuleWhitelistedSourcesRule{
		attributeName: "source",
	}
}

type TerraformModuleWhitelistedSourcesRuleConfig struct {
	Whitelist []string `hclext:"whitelist,optional"`
}

// Name returns the rule name
func (r *TerraformModuleWhitelistedSourcesRule) Name() string {
	return "terraform_module_whitelisted_source"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformModuleWhitelistedSourcesRule) Enabled() bool {
	return true
}

func (r *TerraformModuleWhitelistedSourcesRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *TerraformModuleWhitelistedSourcesRule) Link() string {
	// return project.ReferenceLink(r.Name())
	return ""
}

// Check checks whether module source is whitelisted
func (r *TerraformModuleWhitelistedSourcesRule) Check(rr tflint.Runner) error {
	// resources, err := runner.GetResourceContent("aws_route53_record", &hclext.BodySchema{
	// 	Attributes: []hclext.AttributeSchema{
	// 		{Name: "source"},
	// 	},
	// }, nil)
	// if err != nil {
	// 	return err
	// }

	// // Put a log that can be output with `TFLINT_LOG=debug`
	// logger.Error(fmt.Sprintf("Get %d instances", len(resources.Blocks)))

	// for _, resource := range resources.Blocks {
	// 	attribute, exists := resource.Body.Attributes["source"]
	// 	if !exists {
	// 		logger.Info("Attribute does not exist")
	// 		continue
	// 	}

	// 	// err := runner.EvaluateExpr(attribute.Expr, func(instanceType string) error {
	// 	// 	return runner.EmitIssue(
	// 	// 		r,
	// 	// 		fmt.Sprintf("instance type is %s", instanceType),
	// 	// 		attribute.Expr.Range(),
	// 	// 	)
	// 	// }, nil)
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	return runner.EmitIssue(
	// 		r,
	// 		fmt.Sprintf("instance type is %s", "test"),
	// 		attribute.Expr.Range(),
	// 	)
	// }
	// return nil

	// modules, err := runner.GetModuleContent(&hclext.BodySchema{
	// 	Attributes: []hclext.AttributeSchema{
	// 		{Name: "enabled"},
	// 	}}, nil)
	// if err != nil {
	// 	return err
	// }

	// if !path.IsRoot() {
	// 	// This rule does not evaluate child modules.
	// 	return nil
	// }

	// config := TerraformModuleWhitelistedSourcesRuleConfig{Whitelist: []string{"lablabs.io"}}
	// if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
	// 	return err
	// }

	// logger.Info(fmt.Sprint("PAATH", len(modules.Blocks)))

	// calls, diags := runner.GetModuleCalls()
	// if diags.HasErrors() {
	// 	return diags
	// }

	// for _, call := range calls {
	// 	if err := r.checkModule(runner, call, config); err != nil {
	// 		return err
	// 	}
	// }

	runner := terraform.NewRunner(rr)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	config := TerraformModuleWhitelistedSourcesRuleConfig{Whitelist: []string{"lablabs.io"}}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	calls, diags := runner.GetModuleCalls()
	if diags.HasErrors() {
		return diags
	}

	for _, call := range calls {
		if err := r.checkModule(runner, call, config); err != nil {
			return err
		}
		logger.Info("YIIII", call)
	}
	return nil
}

// func (r *TerraformModuleWhitelistedSourcesRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall, config TerraformModuleWhitelistedSourcesRuleConfig) error {

// 	logger.Info(fmt.Sprintf("SOURCE: %v", module.Source))
// 	return nil
// 	return runner.EmitIssue(
// 		r,
// 		fmt.Sprintf(`Module source "%s" is not pinned`, module.Source),
// 		module.SourceAttr.Expr.Range(),
// 	)
// }

func (r *TerraformModuleWhitelistedSourcesRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall, config TerraformModuleWhitelistedSourcesRuleConfig) error {
	// source, err := getter.Detect(module.Source, filepath.Dir(module.DefRange.Filename), []getter.Detector{
	// 	// https://github.com/hashicorp/terraform/blob/51b0aee36cc2145f45f5b04051a01eb6eb7be8bf/internal/getmodules/getter.go#L30-L52
	// 	new(getter.GitHubDetector),
	// 	new(getter.GitDetector),
	// 	new(getter.BitBucketDetector),
	// 	new(getter.GCSDetector),
	// 	new(getter.S3Detector),
	// 	new(getter.FileDetector),
	// })
	// if err != nil {
	// 	return err
	// }

	logger.Info(fmt.Sprintf("MODULE SOURCE %s", module.Source))

	// if !slices.Contains(config.Whitelist, source) {
	// 	return runner.EmitIssue(
	// 		r,
	// 		fmt.Sprintf("Module source %s is not whitelisted", source),
	// 		module.DefRange,
	// 	)
	// }

	return nil

}
