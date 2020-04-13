[![Build Status](https://api.cirrus-ci.com/github/enginoid/monorepo-base.svg)](https://cirrus-ci.com/github/enginoid/monorepo-base)

### What you're looking at

A Bazel monorepo you can clone, adapt, and use...

  - [x] ...with an example service (`//services/ping`)
    - [x] ...that uses Go + gRPC + Proto
      - [x] (...although you don't need these system dependencies to build!)
    - [x] ...that can build a really slim Docker image
      - [x] (...also without a system dependency on docker!)
    - [x] ...that can generate manifests for staging and production
      - [x] ...with different manifests per environment
      - [x] ...using jsonnet to abstract away the detail for simple cases
        - [x] ...without limiting your options
    - [x] ...that lets you build and deploy the service in one Bazel command
  - [x] ...and some code to help you set up a Kubernetes cluster to use in GCP
    - [x] ...through a script that sets up your project (`//cmd/setup`)
    - [x] ...and some Terraform definitions to set up a cluster (`./terraform`)
  - [x] ...that tests everything in Cirrus CI
    - [x] ...with a [really simple config](./.cirrus.yml), even as your project gets more complex
    - [x] ...and caching that makes your builds quick, even as your project grows

### Prerequisites

If you just want to build and test:
   * `bazel` ([instructions](https://docs.bazel.build/versions/master/install.html))

If also you want to run docker containers locally...
   * `docker` if you're planning to run containers ([instructions](https://docs.docker.com/get-docker/))

If you want to deploy to a kubernetes cluster...
   * `kubectl` (`brew install kubectl` or [instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl/))
 
If you want to set up a kubernetes cluster in GCP with Terraform...
   * `terraform` (`brew install terraform` or [instructions](https://learn.hashicorp.com/terraform/getting-started/install.html#installing-terraform))

## Cheat sheet

| Action | Command |
| --- | --- |
| Run `//service/ping` (as binary) | `bazel run //services/ping` |
| Run `//service/ping` (as Docker container; requires `docker`) | `bazel run //services/ping:docker_image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64` |
| Print `//service/ping` staging manifest | `bazel run //services/ping:staging` |
| Deploy `//service/ping` to staging | `bazel run //services/ping:staging.apply --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64` |
| Run `//cmd/ping` (as binary, with options) | `bazel run //cmd/ping -- --address=1.2.3.4:50051` |

# Detailed guides

 * Development
     * [Quickstart](./docs/development/quickstart.md): running, building and making changes locally.
     * [Working with Go](./docs/development/working-with-go.md): special considerations.
 * Services
     * [Deploying services](./docs/services/deploying.md) to Kubernetes via Bazel.
     * [The anatomy of a service](./docs/services/anatomy.md)
     * [Service manifests with jsonnet](./docs/services/manifests.md)
 * [Deploying infrastructure](./docs/deploying-infrastructure.md) with Terraform.
 * [Potential](./docs/potential.md): how you could extend this.



