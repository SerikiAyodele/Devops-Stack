# GETTING STARTED

### Definition
An open source system for automating deployment, scaling and management of containarized application

### Orchestration tools such as Kubernetes offer:
- High availability
- Scalability: Applications have higher performance e.g. load time
- Disaster Recovery: The architecture has to have a way to back-up the data and restore the state of the application at any point in time

3 processes have to be present on all nodes
- container runtime e.g docker
- kubelet
- kube proxy: for communication between the nodes and pods

### Concepts
- **Pods:** are an abstraction layer over containers and they are the smallest
  unit that a kubernetes user can configure, if a pod dies we can configure
  it to restart, but it will have a new IP
- **Service:** used as an alternative to pod ip, it sits in front of the pod,
  so the pod's lifecycle is not tied to the pod's IP, it serves as an IP
  provider to the pods and a loadbalancer


## Kubernetes Architecture
Responsibilities in a kubernetes cluster are divided between the main node and the worker node
- Operators:They empower kubernetes to automate the entire lifecycle of complex application, making them self-healing and self-managing.
- Controllers:Compares the actual state of cluster to the desired state, and attempts to make the actual state match the desired state.
- ReplicaSet:They are designed to maintain a level of scalability and availability, for anapplication or microservice by creating and 
  managing multiple identical instances of the same pod.
- DaemonSet:Will make sure a single pod is deployed in every node.
- StatefulSet:Used to deploy pods in a particular order, a next pod is deployed only if the previousis reports a ready status.


- kube-apiserver:Responsible for exposing the Kubernetes API, through which users and components can interact with the cluster.
- kube-scheduler:Determines what node a pod should be deployed to.
- etcd Database:Used for managing and storing the cluster's configurationand data.

Each node in a cluster runs two processes:
- Kublet:Handles requests to the containers, manages resources and looks after the local nodes
- Kube-proxy:creates and manages networking rules and it is responsible for managing the network connectivity to containers.

### Additional Security Measures In A Cluster
- Namespaces
- Contexts
- Resource Limits
- Pod Security Policies
- Network Policies

## Kubernetes Pods
- Pods represent a process running in a cluster, it is the smallest unit in a kubernetes cluster. We can have more than one 
  container in a pod, and each container will share the ip address, storage and namespace, each container will normally have
  their own role inside the pod.
- Because containers are ran in parallel, it's hard to tell which started before which, hence we have **initContainers**
  which are used to make sure some containers are ready before others in a pod.
- An example of running multiple containers withing a pod is an app server, were we can have three containers, one in the 
  app server itself, another is a logging adapter and the other a monitoring adapter, in this case the three containers 
  conbined will be providing the same service.
- Replicated pods are created and managed by a controller, such as a deployment.
- All pods in a container are connceted.
- A limit is the maximum amount of CPU or memory that kubernetes grants to a pod.

Each pod has:
- A unique IP address through which they communicate with each other
- Persistent storage volumes (PSV)
- Configuration informan=tion for how a container should run

## Basic KubeCTL Commands
- View Cluster Information:
`kubectl cluster-info`
- View Resources:
`kubectl get <resource>`: List resources of a specific type, e.g., kubectl get pods, kubectl get nodes, kubectl get deployments.
`kubectl describe <resource> <name>`: Show detailed information about a specific resource, e.g., kubectl describe pod my-pod.
- Create and Apply Resources:
`kubectl create`: Create a resource from a file or by specifying details directly, e.g., kubectl create -f pod.yaml.
`kubectl apply`: Create or update a resource from a file or by specifying details directly, e.g., kubectl apply -f deployment.yaml.
- Delete Resources:
`kubectl delete <resource> <name>`: Delete a specific resource, e.g., kubectl delete pod my-pod.
`kubectl delete -f <file>`: Delete resources defined in a file, e.g., kubectl delete -f deployment.yaml.
- Get Cluster Components:
`kubectl get componentstatuses`: List the health status of the cluster's components (e.g., scheduler, controller-manager).
`kubectl get nodes`: List all nodes in the cluster.
- Pod Operations:
`kubectl logs <pod>`: Print the logs of a specific pod, e.g., kubectl logs my-pod.
`kubectl exec -it <pod> -- <command>`: Run a command inside a pod interactively, e.g., kubectl exec -it my-pod -- /bin/bash.
- Scale Deployments:
`kubectl scale deployment <deployment> --replicas=<count>`: Scale the number of replicas in a deployment, e.g., kubectl scale deployment my-deployment --replicas=3.
- Port Forwarding:
`kubectl port-forward <pod> <local-port>:<remote-port>`: Forward a local port to a port on a pod, e.g., kubectl port-forward my-pod 8080:80.
- Get API Resources:
`kubectl api-resources`: List all the supported API resources.