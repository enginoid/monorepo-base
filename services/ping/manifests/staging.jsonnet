local microservice = import "../../microservice.libsonnet";

microservice.Simple(
    name="service-ping",
    environment="staging",
    port=50051,
    replicas=1,
)
