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

### Cheat sheet

| Action | Command |
| --- | --- |
| Run `//service/ping` (as binary) | `bazel run //services/ping` |
| Run `//service/ping` (as Docker container; requires `docker`) | `bazel run //services/ping:docker_image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64` |
| Print `//service/ping` staging manifest | `bazel run //services/ping:staging` |
| Deploy `//service/ping` to staging | `bazel run //services/ping:staging.apply --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64` |
| Run `//cmd/ping` (as binary, with options) | `bazel run //cmd/ping -- --address=1.2.3.4:50051` |

### 🚧 Work in progress! 🚧

This monorepo is functionally complete – you can clone it and get it to work if you look through how it links together and you're willing to battle gcloud a bit. But it's sorely missing documentation for how to set up the infrastructure and deploy to kubernetes, as well as some inline documentation.

Coming soon – please bear with me!

### Detailed guides

 * Development
     * [Quickstart](./docs/development/quickstart.md): running, building and making changes locally.
     * [Working with Go](./docs/development/working-with-go.md): special considerations.
 * Services
     * [Deploying services](./docs/services/deploying.md) to Kubernetes via Bazel.
     * [The anatomy of a service](./docs/services/anatomy.md)
     * [Service manifests with jsonnet](./docs/services/manifests.md)
 * [Deploying infrastructure](./docs/deploying-infrastructure.md) with Terraform.
 * [Potential](./docs/potential.md): how you could extend this.

### ❤️ ️Gratitude & Credits ❤️ ️

I first set out to create a monorepo template in 2017, when the landscape of technologies was quite different. The project was [makesomecloud](https://github.com/enginoid/makesomecloud), and it used Pants, a painstakingly hand-tuned concourse setup, sprinkled with lots of manual commands to make up for how poorly different technologies naturally glued together at the time. 

I built that project with Pants because I knew how to get the parts to fit from previous experience and it fit my needs at the time, but I was already convinced that Bazel was there to stay and would bring a host of new possibilities through its budding ecosystem, good documentation, disciplined roadmap and strong community. Since then, I've halfheartedly come back to to try to reproduce my monorepo in Bazel occasionally, but run into enough stumbling blocks with various ecosystem components (both Bazel and non-Bazel) that I've halted my efforts for a few more months.

I finally came back to the project over the Easter holiday (2020). Where I had built `makesomecloud` over perhaps 2-3 weeks, I got the new repository deployed to Kubernetes in about two days with minimal hurdles and _no hacks_. And the resulting monorepo is much simpler. I was extremely impressed!

**I'd like to say a big thank you to the people who have kept improving the community tooling to where it is now.** I have watched from the sidelines and seen both how much there is to do and how ungrateful the work can sometimes be, but it's amazing to see the strides that the community has made.

I'm grateful for everyone who has worked on improving the many tools that help new teams and projects get started with a strong and sustainable foundation, but I wanted to mention a few areas that have particularly delighted me:

 * `rules_go` and `gazelle` work wonders today and make it effortless to work with Go code. Who knew dependency management between two different dependency management systems could be this seamless? `update-repos` and the integration with `go.mod` is priceless and has saved me so much pain managing dependencies. And gone are the days of hand-fiddling to get gRPC and protobuf to work right!
 * Bazel's ability to cross-compile Go binaries makes it possible to build a container with a Go on a Mac, popping it into a `go_image` (via `rules_docker`), and then deploying it to a Kubernetes cluster (via `rules_k8s`), using `jsonnet` for easy templating. This is just one possible configuration of many – think of how many happy paths are enabled by these tools!
 * I've always struggled with caching in Bazel. This time, I set up the smallest CI config file in the world, and it runs everything in under two minutes following the initial fifteen, using (via [Cirrus CI's HTTP cache](https://cirrus-ci.org/guide/writing-tasks/#http-cache)). And there have been huge strides in remote execution, which are going to revolutionize bazel builds.
 * There are so many `examples/` directories everywhere to demonstrate how things fit together. And you have demo projects like [gke-bazel-demo](https://github.com/GoogleCloudPlatform/gke-bazel-demo) that show you how to do it and work. I got a lot of direction for this repo from that that project in particular, so thank you.
 
Open source has the unique ability to bring together loosely coupled components that together make a strong whole. But often preceding that period, there's a difficult one where things don't fit together as well as you'd like, as The Year of the Linux Desktop will have us remember.

An ecosystem doesn't necessary "click" in a particular main release, but when the individual parts are cohesive enough to make a sum often much greater than its part. For other monorepo adventurers and Bazel users, that time may have come a long time ago, and it may still take some time for others. When these moments happen, they may not be tied to any big release or a celebration. But Bazel has finally just "clicked" for me, and I'm awestruck by what the community has accomplished, and here I am, celebrating that!

❤️ _Thank you all_ ❤️
