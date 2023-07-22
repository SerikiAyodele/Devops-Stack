Definition
An open source system for automating deployment, scaling and management of containarized application

Orchestration tools such as Kubernetes offer:
- High availability
- Scalability: Applications have higher performance e.g. load time
- Disaster Recovery: The architecture has to have a way to back-up the data and restore the state of the application at any point in time

3 processes have to be present on all nodes
- container runtime e.g docker
- kubelet
- kube proxy: for communication between the nodes and pods