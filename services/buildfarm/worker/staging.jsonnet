local deployment = import "../../deployment.libsonnet";

[
  deployment.Simple("buildfarm-worker",
     {
        "name": "buildfarm-worker",
        "image": "eu.gcr.io/monorepo-base/buildfarm-worker:staging",
        "imagePullPolicy": "Always",
     },
    replicas=1)
]
