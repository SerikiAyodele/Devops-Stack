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
- Operators:They empower kubernetes to automate the entire lifecycle of complex application,
            making them self-healing and self-managing.
- Controllers:Compares the actual state of cluster to the desired state, and attempts to make
              the actual state match the desired state.
- ReplicaSet:They are designed to maintain a level of scalability and availability, for an
             application or microservice by creating and managing multiple identical instances
             of the same pod.
- DaemonSet:Will make sure a single pod is deployed in every node.
- StatefulSet:Used to deploy pods in a particular order, a next pod is deployed only if the previous
              is reports a ready status.