{
   Simple:: function(name, containerSpec, replicas=3) {
      "apiVersion": "apps/v1beta1",
      "kind": "Deployment",
      "metadata": {
         "name": name
      },
      "spec": {
         "replicas": replicas,
         "template": {
            "metadata": {
               "labels": {
                  "app": name
               }
            },
            "spec": {
               "containers": [containerSpec]
            }
         }
      }
   }
}