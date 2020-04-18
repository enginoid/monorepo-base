local microservice = import "../../microservice.libsonnet";

microservice.Simple(
    name="buildfarm-server",
    environment="staging",
    port=8980,
    replicas=1,
)
