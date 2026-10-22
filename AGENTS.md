# sipgo - agent guide

fertig.ai's public fork of `emiago/sipgo`; sip-proxy consumes it through a `replace` pinned to a tag on the `fertigai` branch.

## Updates and security

- Renovate (`renovate.json`, preset in `fertigai/ci`) opens update PRs; patch and minor updates auto-merge once CI is green, majors wait for the owner.
- `security.yml` (inline, as a public repo cannot call the private `fertigai/ci` workflow) blocks PRs that add a leaked secret or a fixable HIGH/CRITICAL finding in `go.mod`.
- Acknowledging a finding follows the `fertigai/ci` README: `.trivyignore.yaml`, at most 14 days, the owner decides.
- The nightly security report goes to the security Discord channel.
