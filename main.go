package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/lablabs/tflint-ruleset-lara/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(0)
	}()

	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "lara",
			Version: "1.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformModuleBlocklistedSourcesRule(),
				rules.NewTerraformBackwardsCompatibilityRule(ctx),
			},
		},
	})
}
