# Semantic version compare

```
semver.go
```
Check the Go test suite next to the implementation if you want to see how it handles malformed edge cases.

You need to parse and compare semantic versions without dragging in extra modules.

It uses only the Go standard library. You do not need to provision external services or manage third-party dependencies just to validate a version string.