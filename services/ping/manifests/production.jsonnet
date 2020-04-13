local microservice = import "../../microservice.libsonnet";

microservice.Simple(
    name="service-ping",
    environment="production",
    port=50051,
    replicas=3,
)
