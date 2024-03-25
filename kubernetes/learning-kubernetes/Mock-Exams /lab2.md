## SECOND MOCK EXAM LAB

**1.** Take a backup of the etcd cluster and save it to /opt/etcd-backup.db
     ```
     ETCDCTL_API=3 etcdctl --endpoints $ENDPOINT snapshot save /opt/etcd-backup.db
     Verify the snapshot:

     ETCDCTL_API=3 etcdctl --write-out=table snapshot status /opt/etcd-backup.db

     TARGET_STORAGE=etcd3
     ETCD_IMAGE=3.3.1
     TARGET_VERSION=3.4.22
     STORAGE_MEDIA_TYPE=application/json
     ```

**2.** Create a Pod called redis-storage with image: redis:alpine with a Volume of
     type emptyDir that lasts for the life of the Pod.

    - Pod named 'redis-storage' created
    - Pod 'redis-storage' uses Volume type of emptyDir
    - Pod 'redis-storage' uses volumeMount with mountPath = /data/redis

    ```
    apiVersion: v1
    kind: Pod
    metadata:
      name: redis-storage
    spec:
      containers:
      - image: redis:alpine
        name: redis-storage
        volumeMounts:
        - mountPath: /data/redis
          name: redis-volume
      volumes:
      - name: redis-volume
        emptyDir:
          sizeLimit: 500Mi
    ```

**3.** Create a new pod called super-user-pod with image busybox:1.28. Allow the pod to be able to set    system_time.

    - The container should sleep for 4800 seconds.
    - Pod: super-user-pod
    - Container Image: busybox:1.28
    - Is SYS_TIME capability set for the container?

    ```
    apiVersion: v1
    kind: Pod
    metadata:
      name: super-user-pod
    spec:
      volumes:
      - name: super-user-pod-vol
        emptyDir: {}
      containers:
      - name: super-user-pod-demo
        image: busybox:1.28
        command: [ "sh", "-c", "sleep 4800" ]
        volumeMounts:
        - name: super-user-pod-vol
          mountPath: /data/demo
        securityContext:
            capabilities:
              add: ["NET_ADMIN", "SYS_TIME"]
    ```

**4.** A pod definition file is created at /root/CKA/use-pv.yaml. Make use of this manifest file and mount the persistent volume called pv-1. Ensure the pod is running and the PV is bound.

    - mountPath: /data
    - persistentVolumeClaim Name: my-pvc
    - persistentVolume Claim configured correctly
    - pod using the correct mountPath
    - pod using the persistent volume claim?

**5.** Create a new deployment called nginx-deploy, with image nginx:1.16 and 1 replica. Next upgrade the deployment to version 1.17 using rolling update.

    - Deployment : nginx-deploy. Image: nginx:1.16
    - Image: nginx:1.16
    - Task: Upgrade the version of the deployment to 1:17
    - Task: Record the changes for the image upgrade
    
    ```
    # initial deployment 
    apiVersion: apps/v1
    kind: Deployment
    metadata:
    name: nginx-deploy
    labels:
      app: nginx
    spec:
    replicas: 1
    selector:
      matchLabels:
        app: nginx
    template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx
          image: nginx:1.16
          ports:
          - containerPort: 80
    strategy:
      type: RollingUpdate
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 1
    ```

    ```
    # Upgrading the deployent from 1.16 to 1.17 will look like
    spec.containers.image: nginx:1.17
    # and applying the change
    ```

**6.** Create a new user called john. Grant him access to the cluster. John should have permission to     create, list, get, update and delete pods in the development namespace . The private key exists in the location:/root/CKA/john.key and csr at /root/CKA/john.csr.

Important Note: As of kubernetes 1.19, the CertificateSigningRequest object expects a signerName.
Please refer the documentation to see an example. The documentation tab is available at the top right of terminal.
    - CSR: john-developer Status:Approved
    - Role Name: developer, namespace: development, Resource: Pods
    - Access: User 'john' has appropriate permissions

**7.** Create a nginx pod called nginx-resolver using image nginx, expose it internally with a service called nginx-resolver-service.Test that you are able to look up the service and pod names from within the cluster. Use the image: busybox:1.28 for dns lookup. Record results in /root/CKA/nginx.svc and /root/CKA/nginx.pod

    - Pod: nginx-resolver created
    - Service DNS Resolution recorded correctly
    - Pod DNS resolution recorded correctly

**8.** Create a static pod on node01 called nginx-critical with image nginx and make sure that it is recreated/restarted automatically in case of a failure.

Use /etc/kubernetes/manifests as the Static Pod path for example.
static pod configured under /etc/kubernetes/manifests ?
Pod nginx-critical-node01 is up and running