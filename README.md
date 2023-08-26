# TFLint Ruleset LARA
[![Build Status](https://github.com/lablabs/tflint-ruleset-lara/workflows/build/badge.svg?branch=main)](https://github.com/terraform-linters/tflint-ruleset-lara/actions)

This is a repository with LARA tflint ruleset.

## Requirements

- TFLint v0.42+
- Go v1.20

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "lara" {
  enabled = true

  version = "1.0.0"
  source  = "github.com/lablabs/tflint-ruleset-lara"
```

## Rules

|Name|Description|Severity|Enabled|Docs|
| --- | --- | --- | --- | --- |
|terraform_module_blocklisted_sources|Block specific TF Module Sources (regexp/path/url) |ERROR|✔|[yes](docs/rules/terraform_module_blocklisted_sources.md)


## Building the plugin local

Clone the repository locally and run the following command:

```sh
$ make
```

You can easily install the built plugin with the following:

```sh
$ make install
```

You can run the built plugin like the following:

```sh
$ cat << EOS > .tflint.hcl
plugin "lara" {
  enabled = true
}
EOS
$ tflint
```