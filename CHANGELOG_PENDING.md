SDK Feature
===
* `aws/session`: Add support for AssumeRoles with MFA (#1088)
  * Adds support for assuming IAM roles with MFA enabled. A TokenProvider func was added to stscreds.AssumeRoleProvider that will be called each time the role's credentials need to be refreshed. A basic token provider that sources the MFA token from stdin as stscreds.StdinTokenProvider.
