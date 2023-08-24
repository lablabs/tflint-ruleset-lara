package rules

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"regexp"

	"github.com/hashicorp/go-getter"
	// "github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/project"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
)

type TerraformModuleBlocklistedSourcesRule struct {
	tflint.DefaultRule
	attributeName string
}

// NewTerraformModuleBlocklistedSourcesRule returns a new rule
func NewTerraformModuleBlocklistedSourcesRule() *TerraformModuleBlocklistedSourcesRule {
	return &TerraformModuleBlocklistedSourcesRule{
		attributeName: "source",
	}
}

type TerraformModuleBlocklistedSourcesRuleConfig struct {
	Blocklist []string `hclext:"blocklist,optional"`
}

// Name returns the rule name
func (r *TerraformModuleBlocklistedSourcesRule) Name() string {
	return "terraform_module_blocklisted_source"
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
	// return ""err
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
	config.Blocklist = append(config.Blocklist, "git::https://github.com/lablabs/")
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}
	// if err := strings.Contains(call.Source, config.Blocklist); err != nil {
	// 	return err
	// }

	calls, diags := runner.GetModuleCalls()
	if diags.HasErrors() {
		return diags
	}

	for _, call := range calls {
		if err := r.checkModule(runner, call, config); err != nil {
			return err
		}
		// logger.Info("YIIII", config.Blocklist )
	}
	return nil
}

func (r *TerraformModuleBlocklistedSourcesRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall, config TerraformModuleBlocklistedSourcesRuleConfig) error {
	// logger.Info(fmt.Sprintf("MODULE SOURCE %s", module.Source))
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

	u, err := url.Parse(source)
	if err != nil {
		return err
	}

	if u.Opaque != "" {
		// for git:: or hg:: pseudo-URLs, Opaque is :https, but query will still be parsed
		query := u.RawQuery
		u, err = url.Parse(strings.TrimPrefix(u.Opaque, ":"))
		if err != nil {
			return err
		}

		u.RawQuery = query
	}
	for _, blocked := range config.Blocklist {
		if regexp.MustCompile(blocked).MatchString(module.Source) {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("Module source %s is in the block list.", module.Source),
				module.DefRange,
			)
		}
	}

	return nil

}

// # tflint-ignore: terraform_module_blocklisted_source