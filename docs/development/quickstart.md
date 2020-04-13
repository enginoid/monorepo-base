
## Running a service and its client binary

1. In one tab, run the server: `bazel run //services/ping`
2. In another tab, run the client: `bazel run //cmd/ping`

The client will print:
```
2020/04/12 22:01:44 dialing localhost:50051...
2020/04/12 22:01:44 ping message: ""
2020/04/12 22:01:44 ping reply: ""
```

The server will have printed:
```
2020/04/12 22:01:41 listening on localhost:50051
2020/04/12 22:01:44 received ping: ""
```

## Running a container

If you have `docker` installed, you can run the same service as a container:

```
bazel run //services/ping:docker_image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```

The `--platforms=@io_bazel_rules_go//go/toolchain:linux_amd64` flag is significant, since it signals to Bazel that the binary needs to be built for Linux. This is important if you're not running Linux, since it ensures that the built binary can be run by the container. Otherwise, the container will crash on startup!

## Generating manifests

The service has been configured to have a manifest for `staging` and `production` environments. You can see what these look like by running these commands:

```
bazel run //services/ping:staging
bazel run //services/ping:production
```

The manifests are almost identical. The only difference I've configured so far is that staging has one replica of the container, but production has three.

### Deploying the service

**To run these commands, you need to update your [config/k8s.bzl](./config/k8s.bzl).** To learn how, see [Deploying services](../services/deploying.md).

This means that these commands will deploy the manifests (as printed in the commands above) to these clusters upon running these commands:

```
bazel run //services/ping:staging.apply --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
bazel run //services/ping:production.apply --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```