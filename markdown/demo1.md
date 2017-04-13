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

```shell
sudo setsebool -P virt_use_nfs 1
sudo yum -y install nfs-utils libnfsidmap
sudo systemctl enable rpcbind nfs-server
sudo systemctl start rpcbind nfs-server rpc-statd nfs-idmapd
sudo mkdir /nfsfileshare
sudo chmod 777 /nfsfileshare/
sudo vi /etc/exports
sudo exportfs -r
```

```shell
/nfsfileshare 192.168.122.9(rw,sync)
```

```shell
cd $HOME
git clone https://github.com/CrunchyData/crunchy-containers.git
cd crunchy-containers/examples/kube/statefulset[;w
```

```shell
docker pull crunchydata/crunchy-postgres:centos7-9.5-1.2.6
```

```shell
export BUILDBASE=$HOME/crunchy-containers
export CCP_IMAGE_TAG=centos7-9.5-1.2.6
```

```shell
./run.sh
```

```shell
$ kubectl get pod
NAME      READY     STATUS    RESTARTS   AGE
pgset-0   1/1       Running   0          2m
pgset-1   1/1       Running   1          2m
```

```shell
sudo yum -y install postgresql
```

```shell
ql -h pgset-master -U postgres postgres -c 'table pg_stat_replication'
```

```shell
psql -h pgset-replica -U postgres postgres  -c 'create table foo (id int)' 
```

```shell
kubectl scale statefulset pgset --replicas=3
```

```shell
$ ls -l /nfsfileshare/
total 12
drwx------ 20   26   26 4096 Jan 17 16:35 pgset-0
drwx------ 20   26   26 4096 Jan 17 16:35 pgset-1
drwx------ 20   26   26 4096 Jan 17 16:48 pgset-2
```
