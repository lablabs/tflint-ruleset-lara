package rules

import (
	"fmt"
	"regexp"
	"github.com/lablabs/tflint-ruleset-lara/project"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
)

type TerraformModuleBlocklistedSourcesRule struct {
	tflint.DefaultRule
}

// NewTerraformModuleBlocklistedSourcesRule returns a new rule
func NewTerraformModuleBlocklistedSourcesRule() *TerraformModuleBlocklistedSourcesRule {
	return &TerraformModuleBlocklistedSourcesRule{}
}

type TerraformModuleBlocklistedSourcesRuleConfig struct {
	Blocklist []string `hclext:"blocklist,optional"`
}

// Name returns the rule name
func (r *TerraformModuleBlocklistedSourcesRule) Name() string {
	return "terraform_module_blocklisted_sources"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformModuleBlocklistedSourcesRule) Enabled() bool {
	return true
}

func (r *TerraformModuleBlocklistedSourcesRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *TerraformModuleBlocklistedSourcesRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether module source is blocklisted
func (r *TerraformModuleBlocklistedSourcesRule) Check(rr tflint.Runner) error {
	runner := terraform.NewRunner(rr)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	config := TerraformModuleBlocklistedSourcesRuleConfig{}
	config.Blocklist = append(config.Blocklist, "^git::https://github.com/lablabs/")
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
		logger.Info("Found module " + call.Name + " with source " + call.Source)
	}
	return nil
}

func (r *TerraformModuleBlocklistedSourcesRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall, config TerraformModuleBlocklistedSourcesRuleConfig) error {

	for _, blocked := range config.Blocklist {
		if regexp.MustCompile(string(blocked)).MatchString(module.Source) {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("Module source %s is in the block list.", module.Source),
				module.DefRange,
			)
		}
	}

	return nil

}
