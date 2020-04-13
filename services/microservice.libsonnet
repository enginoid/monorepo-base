local container = import "./container.libsonnet";
local deployment = import "./deployment.libsonnet";
local service = import "./service.libsonnet";

{
   Simple:: function(name, environment, port, replicas=3) [
      service.Simple(name, port),
      deployment.Simple(name,
          container.Simple(
              name,
              "eu.gcr.io/monorepo-base/%s:%s" % [name, environment],
              port),
          replicas=replicas)
  ]
}