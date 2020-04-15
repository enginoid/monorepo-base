# Deploying services

## Deployment commands

To deploy a service to each environment, use:

```
bazel run //services/ping:staging.apply --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
bazel run //services/ping:production.apply --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```

To deploy all services to an environment, use:

```
bazel run $(bazel query 'filter('staging.apply', kind("k8s_object", //services/...))') --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
bazel run $(bazel query 'filter('production.apply', kind("k8s_object", //services/...))') --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```

**Note: Always include the `--platform` flag.** The `--platforms=@io_bazel_rules_go//go/toolchain:linux_amd64` flag is necessary to make sure the Go binary is built for the Docker container that will run the binary. As a part of running the `:*.apply` target, a new Docker image will be built (if necessary). If you are on a different platform (eg. macOS) and omit this flag, the binary will crash on container start. 

## Setting up for deployment

In order to be able to deploy to staging and production, you will need to create two new contexts for your `kubectl` command for each deployment environment:

  * `monorepo-base-staging`
  * `monorepo-base-production`

If you'd like, you can change the name of the required contexts by editing [`/config/k8s.bzl`](/config/k8s.bzl).

### Creating a new context

You can create a new `kubectl` context with the following command:

    kubectl config set-context [context_name] \
        --cluster=[cluster] \
        --user=[cluster] \
        --namespace=[namespace]

### Creating contexts that use your local Docker Desktop cluster

If you have Docker Desktop installed with Kubernetes enabled, you can set your contexts up to deploy to your local Kubernetes cluster.

Docker Desktop should have added the following to your local config:
  * a `docker-desktop` user
  * a `docker-desktop` cluster

You can check that these exist by running `kubectl config view` and checking for the user and cluster definition:

```
$ kubectl config view
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://kubernetes.docker.internal:6443
  name: docker-desktop
[...]
users:
- name: docker-desktop
  user:
    client-certificate-data: REDACTED
    client-key-data: REDACTED
```

If your output matches the above, you can use the user and cluster to create a context for each environment pointing to different namespaces for the cluster.

```
kubectl config set-context monorepo-base-staging 
    --cluster=docker-desktop \ 
    --user=docker-desktop \
    --namespace=monorepo-staging

kubectl config set-context monorepo-base-production 
    --cluster=docker-desktop \ 
    --user=docker-desktop \
    --namespace=monorepo-production
```

### Creating contexts using other clusters

To use a different cluster and user that's been set up in your Kubernetes config, use `kubectl config view` to inspect your existing contexts, and then create your new context from the users and clusters defined therein.

For example, I happen to have a cluster set up on Google Cloud Platform, so I can copy the cluster and user info:

```
$ kubectl config view

apiVersion: v1
[...]
contexts:
[...]
- context:
    cluster: gke_monorepo-base_europe-west2-c_monorepo-base-staging
    user: gke_monorepo-base_europe-west2-c_monorepo-base-staging
  name: gke_monorepo-base_europe-west2-c_monorepo-base-staging
```

From this, I can create my staging configuration as:

```
kubectl config set-context monorepo-base-staging 
    --cluster=gke_monorepo-base_europe-west2-c_monorepo-base-staging \ 
    --user=gke_monorepo-base_europe-west2-c_monorepo-base-staging
```

If you're sharing a staging cluster among multiple developers and even integration test runs, you may want to set up your staging cluster to have different namespaces for different environments, rather than risking conflicts when everyone deploys to the same cluster.

For example:
  * The `default` namespace may be used for the "main" staging environment, and could be updated automatically by CI on merge.
  * Each developer might have a `developer-[name]` namespace for testing their changes.
  * Each integration test run might have a `build-[git_commit_hash]` namespace.
  * Each feature branch might have a `branch-[name]` namespace.
  
If you use this arrangement, you might want to set the namespace for your context to something like `developer-[name]`. 

# TODO

 - [ ] Why `--platform` is necessary.