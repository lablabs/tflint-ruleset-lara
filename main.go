package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/lablabs/tflint-ruleset-whitelisted-module-sources/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "whitelisted-module-sources",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformModuleWhitelistedSourcesRule(),
			},
		},
	})
}
