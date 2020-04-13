{
   Simple:: function(name, port) {
      "apiVersion": "v1",
      "kind": "Service",
      "metadata": {
         "labels": {
           "app": name
         },
         "name": name
      },
      "spec": {
         "ports": [
            {
                "port": port,
                "protocol": "TCP",
                "targetPort": port
            }
         ],
         "selector": {
            "app": name
         },
         "type": "LoadBalancer"
      }
   }
}