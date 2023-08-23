package rules

import (
	"fmt"
	"path/filepath"
	"golang.org/x/exp/slices"

	"github.com/hashicorp/go-getter"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
	"github.com/terraform-linters/tflint-ruleset-terraform/project"
)

type TerraformModuleWhitelistedSourcesRule struct {
	tflint.DefaultRule
	attributeName string
}

type TerraformModuleWhitelistedSourcesRuleConfig struct {
	Whitelist []string `hclext:"whitelist,optional"`
}

// NewTerraformModuleWhitelistedSourcesRule returns a new rule
func NewTerraformModuleWhitelistedSourcesRule() *TerraformModuleWhitelistedSourcesRule {
	return &TerraformModuleWhitelistedSourcesRule{
		attributeName: "source",
	}
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
	return tflint.WARNING
}

func (r *TerraformModuleWhitelistedSourcesRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether module source is whitelisted
func (r *TerraformModuleWhitelistedSourcesRule) Check(rr tflint.Runner) error {
	runner := rr.(*terraform.Runner)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}
	
	config := TerraformModuleWhitelistedSourcesRuleConfig{}
	config.Whitelist = append(config.Whitelist, "git::https://github.com/lablabs/")
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
	}

	return nil
}

func (r *TerraformModuleWhitelistedSourcesRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall, config TerraformModuleWhitelistedSourcesRuleConfig) error {
	source, err := getter.Detect(module.Source, filepath.Dir(module.DefRange.Filename), []getter.Detector{
		// https://github.com/hashicorp/terraform/blob/51b0aee36cc2145f45f5b04051a01eb6eb7be8bf/internal/getmodules/getter.go#L30-L52
		new(getter.GitHubDetector),
		new(getter.GitDetector),
		new(getter.BitBucketDetector),
		new(getter.GCSDetector),
		new(getter.S3Detector),
		new(getter.FileDetector),
	})
	if err != nil {
		return err
	}


	if !slices.Contains(config.Whitelist, source) {
		return runner.EmitIssue(
			r,
			fmt.Sprintf("Module source %s is not whitelisted", source),
			module.DefRange,
		)
	}

	return nil

}