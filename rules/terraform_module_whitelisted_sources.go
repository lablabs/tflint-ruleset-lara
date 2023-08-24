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

func (r *TerraformModuleWhitelistedSourcesRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall, config TerraformModuleWhitelistedSourcesRuleConfig) error {
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
