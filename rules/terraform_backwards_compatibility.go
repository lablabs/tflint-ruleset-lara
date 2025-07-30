package rules

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	tffs "github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"

	"github.com/lablabs/tflint-ruleset-lara/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type TerraformBackwardsCompatibilityRule struct {
	tflint.DefaultRule

	ctx context.Context
}

// NewTerraformBackwardsCompatibilityRule returns a new rule
func NewTerraformBackwardsCompatibilityRule(ctx context.Context) *TerraformBackwardsCompatibilityRule {
	return &TerraformBackwardsCompatibilityRule{
		ctx: ctx,
	}
}

type terraformBackwardsCompatibilityRuleConfig struct {
	Version    string   `hclext:"version"`
	Exemptions []string `hclext:"exemptions,optional"`
}

// Name returns the rule name
func (r *TerraformBackwardsCompatibilityRule) Name() string {
	return "terraform_backwards_compatibility"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformBackwardsCompatibilityRule) Enabled() bool {
	return true
}

func (r *TerraformBackwardsCompatibilityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *TerraformBackwardsCompatibilityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *TerraformBackwardsCompatibilityRule) config(runner tflint.Runner) (*terraformBackwardsCompatibilityRuleConfig, error) {
	config := &terraformBackwardsCompatibilityRuleConfig{}

	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return nil, err
	}

	return config, nil
}

// Check checks whether module source is blocklisted
func (r *TerraformBackwardsCompatibilityRule) Check(runner tflint.Runner) error {
	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}

	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	config, err := r.config(runner)
	if err != nil {
		return err
	}

	workingDir := path.String()
	if workingDir == "" {
		workingDir, err = runner.GetOriginalwd()
		if err != nil {
			return err
		}
	}

	execPath, err := r.installTerraform(r.ctx, config.Version)
	if err != nil {
		return err
	}

	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return err
	}

	// initialize Terraform if the lock file does not exist, i.e. when running within module
	if _, err := os.Stat(".terraform.lock.hcl"); errors.Is(err, os.ErrNotExist) {
		err = tf.Init(r.ctx, tfexec.Backend(false))
		if err != nil {
			return err
		}

		// cleanup directory after validation
		defer os.RemoveAll(".terraform.lock.hcl")
	}

	output, err := tf.Validate(r.ctx)
	if err != nil {
		return err
	}

	err = r.checkValidationOutput(runner, output, config.Exemptions)
	if err != nil {
		return err
	}

	return nil
}

func (r *TerraformBackwardsCompatibilityRule) installTerraform(ctx context.Context, terraformVersion string) (string, error) {
	installDir := fmt.Sprintf("%s/%s/%s", os.TempDir(), r.Name(), terraformVersion)
	err := os.MkdirAll(installDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Use existing version if available
	execPath, err := install.NewInstaller().Ensure(ctx, []src.Source{
		&tffs.ExactVersion{
			Product:    product.Terraform,
			Version:    version.Must(version.NewVersion(terraformVersion)),
			ExtraPaths: []string{installDir},
		},
	})
	if err != nil {
		// If the version is not available, install it
		execPath, err = install.NewInstaller().Ensure(ctx, []src.Source{
			&releases.ExactVersion{
				Product:    product.Terraform,
				InstallDir: installDir,
				Version:    version.Must(version.NewVersion(terraformVersion)),
			},
		})
		if err != nil {
			return "", err
		}
	}

	return execPath, nil
}

func (r *TerraformBackwardsCompatibilityRule) checkValidationOutput(runner tflint.Runner, output *tfjson.ValidateOutput, exemptions []string) error {
	if output.Valid {
		return nil
	}

	errs := []error{}

	for _, diagnostic := range output.Diagnostics {
		if !slices.Contains(exemptions, diagnostic.Detail) {
			issueRange := hcl.EmptyBody().MissingItemRange()
			if diagnostic.Range != nil {
				issueRange = hcl.Range{
					Filename: diagnostic.Range.Filename,
					Start: hcl.Pos{
						Line:   diagnostic.Range.Start.Line,
						Column: diagnostic.Range.Start.Column,
						Byte:   diagnostic.Range.Start.Byte,
					},
					End: hcl.Pos{
						Line:   diagnostic.Range.End.Line,
						Column: diagnostic.Range.End.Column,
						Byte:   diagnostic.Range.End.Byte,
					},
				}
			}

			errs = append(errs, runner.EmitIssue(
				r,
				fmt.Sprintf("%s. %s", diagnostic.Summary, diagnostic.Detail),
				issueRange,
			))
		}
	}

	return errors.Join(errs...)
}
