# sipgo - agent guide

fertig.ai's public fork of `emiago/sipgo`.

## Updates and security

- Renovate (`renovate.json`, shared preset) opens update PRs; patch/minor merge automatically once CI passes.
- `security.yml` runs a secret and dependency scan on every PR and blocks a PR that adds a leaked secret.
- Acknowledging a finding: at most 14 days; the owner decides.
