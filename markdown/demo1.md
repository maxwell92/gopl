```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-pd
spec:
  containers:
  - image:gcr.io/google_containers/test-webserver
    name:test-container
    volumeMounts:
    - mountPath:/cache
      name:cache-volume
  volumes:
  - name:cache-volume
emptyDir: {}
```

```yaml
apiVersion:           v1
kind: Pod
metadata:
name:          test-gcespec:
containers:
-         image:         nginx:latest
           ports:
-         containerPort:      80
           name:          test-gce
volumeMounts:
 -       mountPath:           /usr/share/nginx/html
           name:          gce-pd
volumes:
-         name:          gce-pd
           gcePersistentDisk:
           pdName:    mysite-volume-1
           fsType:       ext4

```
