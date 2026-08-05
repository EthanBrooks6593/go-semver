# Semantic version compare

```
semver.go
```

See the test next to the source for usage.

Parse and compare semantic versions — no dependencies.

This is a small utility for parsing and comparing semantic version strings. It uses only the standard library, so there's nothing to install, no external service to call, and no hidden state to worry about.

The core function takes two version strings, splits them into their numeric components, and compares them component by component. Pre-release identifiers and build metadata are handled per the SemVer spec, so you get predictable behavior even for versions like `1.0.0-alpha.1` or `2.0.0+build.5`.

Usage is straightforward:

```python
from semver_compare import compare

result = compare("1.2.3", "1.2.4")
# result is -1, 0, or 1
```

The test file next to the source shows more examples, including edge cases like versions with different lengths and pre-release suffixes. If you're building release tooling, dependency checks, or anything that needs to order versions reliably, this should cover it without pulling in a package manager or a network call.