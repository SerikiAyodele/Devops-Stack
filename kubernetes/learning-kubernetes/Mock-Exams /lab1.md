## FIRST MOCK EXAM LAB

*1.* 

```controlplane ~ ➜  k run nginx-pod --image=nginx:alpine
pod/nginx-pod created
controlplane ~ ➜  k get pods
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   1/1     Running   0          6s
```

*2.* Deploy a messaging pod using the redis:alpine image with the labels set to tier=msg.
`k run messaging --image=redis:alpine --labels tier=msg`

```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: "2024-03-18T06:29:42Z"
  labels:
    tier: msg
  name: messaging
```

*3.* Create a namespace named apx-x9984574
`k create ns apx-x9984574`

```
controlplane ~ ➜  k get ns
NAME              STATUS   AGE
apx-x9984574      Active   3s
```

*4.* Get the list of nodes in JSON format and store it in a file at /opt/outputs/nodes-z3444kd9.json.
`k get nodes -o json > /opt/outputs/nodes-z3444kd9.json`

*5.* Create a service messaging-service to expose the messaging application within the cluster on port 6379
`k expose pod messaging --name=messaging-service --port=6379 --type=ClusterIP`

*6.* Create a deployment named hr-web-app using the image kodekloud/webapp-color with 2 replicas.
`k create deploy hr-web-app --image kodekloud/webapp-color --replicas 2`

*7.* Create a static pod named static-busybox on the controlplane node that uses the busybox image and the command sleep 1000
`k run static-busybox --image busybox --command -- sleep 1000`

*8.* Create a POD in the finance namespace named temp-bus with the image redis:alpine
`k run temp-bus -n finance --image redis:alpine`

*9.* A new application orange is deployed. There is something wrong with it. Identify and fix the issue.

*10.* Expose the hr-web-app as service hr-web-app-service application on port 30082 on the nodes on the cluster.
The web application listens on port 8080.
`k expose deploy hr-web-app --name hr-web-app-service --port 30082 --target-port 8080 --type NodePort`

*11*.* Use JSON PATH query to retrieve the osImages of all the nodes and store it in a file /opt/outputs/nodes_os_x43kj56.txt.
The osImages are under the nodeInfo section under status of each node
k get nodes -o jsonpath= '{.status.nodeInfo.osImages}' > /opt/outputs/nodes_os_x43kj56.txt

kubectl get nodes -o jsonpath='{range .items[*]}{.status.nodeInfo.osImage}' > /opt/outputs/nodes_os_x43kj56.txt

*13.* Create a Persistent Volume with the given specification: -
```
Volume name: pv-analytics
Storage: 100Mi
Access mode: ReadWriteMany
Host path: /pv/data-analytics

k create pvc pv-analytics

apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-analytics
spec:
  capacity:
    storage: 100Mi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Recycle
  storageClassName: slow
  mountOptions:
    - hard
    - nfsvers=4.1
  nfs:
    path:  /pv/data-analytics
    server: 172.17.0.2
```