# Publishing `harness`

This guide explains how to release and publish new versions of the `github.com/px0-ai/harness` Go module.

## Single Source of Truth: `VERSION`

The version is defined in the [`VERSION`](VERSION) file at the root of the repository:

```
0.0.1
```

This file is embedded directly into the Go module and exposed as [`harness.Version`](version.go). Any version updates should be made in this file.

## Versioning Scheme

Go modules follow [Semantic Versioning (SemVer)](https://semver.org/) with a required `v` prefix in git tags:

`vMAJOR.MINOR.PATCH` (e.g., `v0.0.1`, `v0.1.0`, `v1.0.0`)

- **Patch (`v0.0.x`)**: Bug fixes and minor internal improvements with no breaking changes.
- **Minor (`v0.x.0`)**: New features, new harness presets, or enhancements. (While on `v0.x.x`, API changes may occur as the interface stabilizes).
- **Major (`v1.0.0`+)**: Stable production release with backwards compatibility guarantees. Bumping to `v2.0.0` or higher requires changing the module path in `go.mod` (e.g., `github.com/px0-ai/harness/v2`).

## Release Checklist

### 1. Update the `VERSION` file
Set the new version in [`VERSION`](VERSION) without the `v` prefix (e.g., `0.0.2`):

```sh
echo "0.0.2" > VERSION
```

### 2. Run tests and checks
Run all tests and vetting from the repository root:

```sh
go test -v ./...
go vet ./...
```

### 3. Commit version change
Commit the `VERSION` file update and ensure working tree is clean:

```sh
git add VERSION
git commit -m "Bump version to $(tr -d '[:space:]' < VERSION)"
git push origin master
```

### 4. Create and push the Git tag
Read the tag version directly from [`VERSION`](VERSION) and push it to `origin`:

```sh
VERSION=$(tr -d '[:space:]' < VERSION)
git tag -a "v$VERSION" -m "Release v$VERSION"
git push origin "v$VERSION"
```

### 5. Trigger Go Module Proxy indexing
Go fetches modules via the public proxy `proxy.golang.org`. Querying the proxy triggers immediate caching and indexes the version on [pkg.go.dev](https://pkg.go.dev):

```sh
VERSION=$(tr -d '[:space:]' < VERSION)
GOPROXY=https://proxy.golang.org go list -m "github.com/px0-ai/harness@v$VERSION"
```

### 6. Verify availability
Check that consumers can fetch the new release:

```sh
VERSION=$(tr -d '[:space:]' < VERSION)
go get "github.com/px0-ai/harness@v$VERSION"
```

You can view the package documentation at:
- `https://pkg.go.dev/github.com/px0-ai/harness`
