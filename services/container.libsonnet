{
   Simple:: function(name, uri, port) {
      "name": name,
      "image": uri,
      "imagePullPolicy": "Always",
      "ports": [
         {
            "containerPort": port
         }
      ]
   }
}