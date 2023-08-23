package main

import (
	"github.com/lablabs/tflint-ruleset-whitelisted-module-sources/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "whitelisted-module-sources",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformModuleWhitelistedSourcesRule(),
				// rules.NewTerraformModulePinnedSourceRule(),
			},
		},
	})
}
