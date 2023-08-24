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

    signing_key = <<-KEY
    -----BEGIN PGP PUBLIC KEY BLOCK-----

    mQINBGTnr3YBEACixM6iqRbM7vU4LwbQDi5l87vHZM5T2iWzhLaE6Jg8DCiCzdBv
    vn8msWvmVovXNNTLEfpdieuiPmqP51mantZY6JBzbG1J5i4w0Q/wPRdex8SJGsi5
    fwYouZA9D0FgSy/ZoHuYME3lXiro2odWmrbXV/4QflilzRML76byImq3zLZpSfUr
    oQNwVEieANCvCZWm/720kZqc88Mlaxc3O3WVMDubBa8TdU+kInGmZPQQVzYt5CM1
    1fppKUTAsMS1yvqIxt4KQwlgTDibgRMqDeUwppn/Jk8JZ0+rnhxcr0D671S3q9b0
    jroqXNxGJcSGeCc6com0eKAjOEEM0zFIwalVrpFkNYfRB/wRTx66tXOh6vyo24H0
    eKBvLxwQ/ZEZG5hCUNyeuYx6vDv4hcmkLSIEnCZ+QRMzaeKocYujs+jlMpUGVkFa
    1loxEwaphxSXybSi0N23If3bJ4QZRkK7ka8oPz49RCzh4IzI/hH/ZcqacZ/U1SSe
    yuHW4IXqLoHm5eWKKBl68/JuGwAmgzD/dfShnZX4EyCHZ9Jhslqo/Wx0nK+7LuG6
    Cr5scEYVDTB1QtwW35okU3o5nst0TUVTmgfL6phzlKy1ebjP7UuzCNVzU2ojDvDU
    O7N40gI4b4nVLVM5XTUYAjf/2y3MqBSwQaQpD8od5dCGX/N7ZkLBiEuOQQARAQAB
    tCBMYWJ5cmludGggTGFicyA8aW5mb0BsYWJsYWJzLmlvPokCVwQTAQgAQRYhBNY0
    SHVATkvcLh5aIOq1Zo5VtwrcBQJk5692AhsDBQkDwmcABQsJCAcCAiICBhUKCQgL
    AgQWAgMBAh4HAheAAAoJEOq1Zo5VtwrcKqYP/Rwge/zG/Fm9MWxzcbnGWGE+Yo3U
    qKWJUfdBj8EXu5CX/Us6TghuNpGZntj5Z4rr+RBOel2IZN80mczN/ao3+pQuZ7HQ
    gYnlSNlGWccwWVYsHTMktN2MSDN6WABBTiBIGbOeUCqiglT9kyWmc++/UYTEiYpn
    nHsB2zpDXvf+AI3h1VfiXldnafrOp7SM4/5O7WeJQpMQabgN/FN7Y0QaPmU8WTZ9
    v1MJXTTKyWd3gPukuf6OFBMcCbsf9LtodUC/ywkLJ2ccQhbOKHbguAWw5AQMPD0c
    toglak3iJ+MyW23r0U3C9oxLl8pRuCXK6FPSrzqNNK9n+wLdFryNaxU4uWEVtGYt
    KysCVLW1P3BcoqWH560Sv5MIhRHQaXUzFvFdMO1f9hXAv8luzMBWT/Otid/MAnGO
    u46ZCiARIBmiehJJGmMO3Q7Cm1IUOz0llqeRy2IeN2vv7KMAsgRstU5VqHXMCrUR
    g0GvZBIIDkcMPi1AwnTOH3ZUjeyF2dJRmOcf9C4FYnOLkjXNPDhWygge4wIOU3X8
    FsTKflVY/i/jCK9qpWGXVuObSLHKzPlAwC4vWmBi8SM33RZohNXm7N+QrxCz5Mvv
    31o90QomfblXONZktgVoE9d+2Cm9tO2x2z2LxTEdsELuORK6UB1RXpMzNXkRYCnD
    DGiH0rPJO+FIhmy/
    =qm71
    -----END PGP PUBLIC KEY BLOCK-----
    KEY
}
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