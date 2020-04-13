# Potential and inspiration

This section describes where you could take this monorepo, by explaining capabilities (that you can use to create something cool) and possibilities (cool things you could do).

## Possibilities

 - [ ] TODO: Phoenix environments (per-developer, per-build etc.)
 - [ ] TODO: Monorepo tools and shorthands that could be useful.

## Capabilities

### Querying

Bazel includes a powerful query language. One example of what you can do with it is find all go libraries and binaries:

```
$ bazel query 'kind("go_(library|binary) rule", //...)'
//services/ping:ping
//services/ping:docker_image.binary
//services/ping:go_default_library
//cmd/setup:setup
//cmd/setup:go_default_library
//cmd/setup/internal/gcloud:go_default_library
//cmd/ping:ping
//cmd/ping:go_default_library
//services/ping/proto:go_default_library
```

You could use this to build all Go binaries!

```
$ bazel build $(bazel query 'kind("go_(library|binary) rule", //...)')
```

You could use something more specific to build all service binaries, say before running a script that runs integration tests against services locally.

