# Working with Go

 - [ ] Describe recommended setup for working with go locally, including installing protobuf.

## Managing Go dependencies: `gazelle`, `go.mod` and `gazelle update-repos`

When you change anything, I recommend you run these commands in order:
```
go build ./...
bazel run gazelle -- update-repos -from_file=go.mod
bazel run gazelle
```

The rest of this section explains why.

---

Because Bazel does its own dependency management, there are some commands that you need to run when your package structure or dependencies change.

The first of them is the main `gazelle` command, which you need to run after you've changed the Go package structure or added new dependencies. Gazelle will automatically create a `BUILD.bazel` file in each Go package and populate it with default build rules from the [bazelbuild/rules_go](https://github.com/bazelbuild/rules_go) package that reflect the structure and dependencies of your Go code.

```
bazel run gazelle
```

Another tool you'll need is `gazelle update-repos`, which we will use to keep `go.mod` and `WORKSPACE` in sync. This introduces sufficient moving parts and room for future confusion, that I'd like to explain how it works through a perilous tale of adding an import.

`go_rules` and `gazelle` have made Go development in Bazel remarkably easy, but it's not intuitive at first how this dependency management works because it has a few interdependent parts to it. This can yield get frustrating errors when you're doing basic tasks like adding a dependency, so it's worth explaining how the pieces interact so you know how to deal with them.

Let's work through an example of adding an import of `"go.uber.org/zap"` into `//cmd/setup`.

The key thing to know is that when you've added an import to a new package, you need to reflect it in three places to be able to build with your new import:
  1. **The package's `BUILD.bazel`.** The target for the library is going to need a new dependency added to the `deps` list when you run `gazelle`, which will look something like `"@org_uber_go_zap//:go_default_library"`.
  2. **In Go modules.** [go.mod](./go.mod) and [go.sum](./go.sum) are automatically updated to include any new imports when you use the go tooling to build.
  3. **Bazel's [WORKSPACE](./WORKSPACE) file.** This file creates a Bazel target for every dependency (eg. `@org_uber_go_zap`) that specifies what exact version of the package to use, in much the same way as `go.mod` and `go.sum`.  The `WORKSPACE` file can be automatically updated to match the `go.mod` file through a call to `gazelle update-repos`, which we'll cover a bit later on.

The tricky thing is that when a dependency is added, all these sources need to be updated. If you added this import and then immediately built the package with Bazel, you would get an error! Let's see what happens!

```
$ bazel run //cmd/setup
[...]
compilepkg: missing strict dependencies:
	/private/var/tmp/_bazel_fred/b8e1d21679e67122bcee32a2b0e93361/sandbox/darwin-sandbox/1350/execroot/__main__/cmd/setup/main.go: import of "go.uber.org/zap"
No dependencies were provided.
Check that imports in Go sources match importpath attributes in deps.
Target //cmd/setup:setup failed to build
[...]
```

The key to this error is:
> Check that imports in Go sources match importpath attributes in deps.

What this is telling us is that the `BUILD.bazel` is missing `"@org_uber_go_zap//:go_default_library"` from the `deps` list.

But we happen to know this is something that `gazelle` can help with, because `gazelle` inspects all the imports in different pacakges and generates the right `BUILD.bazel` file. Let's run Gazelle and see what happens!

```
$ bazel run gazelle
[...]
INFO: Build completed successfully, 1 total action
```

All good, so let's re-run that command!

```
$ bazel run //cmd/setup
ERROR: /Users/fred/src/github.com/enginoid/monorepo-base/cmd/setup/BUILD.bazel:3:1: no such package '@org_uber_go_zap//': The repository '@org_uber_go_zap' could not be resolved and referenced by '//cmd/setup:go_default_library'
ERROR: Analysis of target '//cmd/setup:setup' failed; build aborted: no such package '@org_uber_go_zap//': The repository '@org_uber_go_zap' could not be resolved
INFO: Elapsed time: 0.114s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (1 packages loaded, 0 targets configured)
FAILED: Build did NOT complete successfully (1 packages loaded, 0 targets configured)
```

That didn't quite cut it. `gazelle` has gone in and updated `BUILD.bazel` to include `@org_uber_go_zap` in deps, but now we're seeing this new error:

```
ERROR: Analysis of target '//cmd/setup:setup' failed; build aborted: no such package '@org_uber_go_zap//': The repository '@org_uber_go_zap' could not be resolved
```

This is telling us that `@org_uber_go_zap` hasn't been defined. It would ordinarily be defined in `WORKSPACE`, and it's indeed the `gazelle update-repos` command that's responsible for reflecting it in there.

Since this repository is set up with go modules, this _wonderful_ command will take the packages in our `go.mod` and `go.sum` files and reflect them as targets in `WORKSPACE`:

```
$ bazel run gazelle -- update-repos -from_file=go.mod 
```

Unfortunately, this command would have no effect if we run it now. That's because the `go.mod` file hasn't been updated after we added our import, so there are no changes to be reflected over to `WORKSPACE`.

As you may recall, `go.mod` is updated when we build a go package and a new dependency is detected, **BUT building through Bazel (`bazel build //cmd/setup`) does not trigger an update to `go.mod`!** The [go_rules/README](https://github.com/bazelbuild/rules_go#does-this-work-with-go-modules) tells us why this is:

> Modules are a dependency management feature in cmd/go, the build system that ships with the Go SDK. Bazel uses the Go compiler and linker in the Go toolchain, but it does not use cmd/go.

So what we need to do is to build our project in Go, in order to update `go.mod`:

```
$ go build ./cmd/setup
go: finding module for package go.uber.org/zap
go: found go.uber.org/zap in go.uber.org/zap v1.14.1
```

Now our `go.mod` and `go.sum` files are updated and we can run this to great effect:

```
$ bazel run gazelle -- update-repos -from_file=go.mod 
```

And now we can finally run this successfully:

```
$ bazel build //cmd/setup
[...]
🎉🎉🎉 INFO: Build completed successfully, 1 total action 🎉🎉🎉
```

Okay, so now we got a feel for how the pieces work together. What does that tell us about how to add dependencies in practice?

I think one think it tells us is that we want to make sure that our Go code is able to build independently of Bazel. In particular, we want to be able to build our go modules with `cmd/go` so we can update `go.mod`. In other words, even though we're using Bazel for building, it'd be desirable to make sure this always works:

```
go build ./...
```

This may not seem like a tall ask, but it might limit some of the use cases or benefits of Bazel just a little bit. One implication is that bazel itself isn't self-sufficient to develop. For this repo, you'll definitely need `go` and `protobuf` installed. Another might be that further down the line, it might be tricky to add more complex dependencies to Go targets while keeping `gazelle`'s auto-generation working.

The other thing it tells us is that when we add dependencies, this is the desired sequence of actions is as stated above, before you toggled into this madness!

