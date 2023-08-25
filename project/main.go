package project

import "fmt"

// Version is ruleset version
const Version string = "1.0.0"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/lablabs/tflint-ruleset-lara/blob/v%s/docs/rules/%s.md", Version, name)
}