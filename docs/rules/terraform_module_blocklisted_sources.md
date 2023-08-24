# terraform_module_blocklisted_sources

Disallow specific Terraform module sources.


## Configuration

Name | Default | Value
--- | --- | ---
enabled | true | Boolean
blocklist | `["^git::https://github.com/lablabs"]`| list of regexps/strings

> **Note**  
> To exclude specific Terraform module from being checked by this rule, add following comment above the module code in the terraform file.  
> `# tflint-ignore: terraform_module_blocklisted_source`

## Example

```hcl
rule "terraform_module_blocklisted_sources" {
    enabled = true
    blocklist = ["^../../", "git::ssh://github.com/lablabs"]
}
```

Match Terraform module source path starting with `../../` and module source which contains `git::ssh://github.com/lablabs` somewhere in the string.
