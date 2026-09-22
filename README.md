# Semantic version compare

```
semver.go
```
When routing SMS payloads or managing OTP API versions, you need to compare semver strings without pulling in a massive dependency tree. Check the Go test files next to the implementation to see how it handles edge cases.

This parses and compares semantic versions with zero external dependencies.

The Go implementation relies entirely on the standard library. You do not need to install any extra services or third-party packages to get it running.