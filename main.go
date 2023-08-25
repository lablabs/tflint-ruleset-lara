package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/lablabs/tflint-ruleset-lara/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "lara",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformModuleBlocklistedSourcesRule(),
			},
		},
	})
}
