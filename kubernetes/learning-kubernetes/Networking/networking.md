# Networking

## Kubernetes Service
- A service is created on top of an existing deployment
- A Kubernetes service is an abstarction that provides a way to to expose an        application running within a set of pods 
- ClusterIP exposes the services only within the cluster, it prevents external communication and makes the cluster more secure
 - ClusterIP is the default and most common type of Service in Kubernetes.
 - It provides a stable, internal IP address (ClusterIP) that is only accessible from within the Kubernetes cluster.
- Nodeport allows access to the service from outside the Kubernetes cluster and within the cluster