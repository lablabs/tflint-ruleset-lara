# terraform_backwards_compatibility

Checks if the code is backwards compatible to a specified version. The check is based on Terraform native `validate` command.

## Configuration

Name | Default | Value
--- | --- | ---
enabled | true | Boolean. Whether the rule is enabled.
version | "" | String. Terraform version to check compatibility against.
exemptions | [] | []String. List of validation exemptions that will be skipped if violated.

## Example

```hcl
rule "terraform_backwards_compatibility" {
  enabled = true

  version = "1.5.7"
  exemptions = [
    "The given configuration is not valid for backend \"s3\": unexpected attribute \"use_lockfile\"."
  ]
}
```
