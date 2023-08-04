# Networking

## Kubernetes Service
- A service is created on top of an existing deployment
- A Kubernetes service is an abstarction that provides a way to to expose an        application running within a set of pods 
- ClusterIP exposes the services only within the cluster, it prevents external communication and makes the cluster more secure
 - ClusterIP is the default and most common type of Service in Kubernetes.
 - It provides a stable, internal IP address (ClusterIP) that is only accessible from within the Kubernetes cluster.
- Nodeport allows access to the service from outside the Kubernetes cluster and within the cluster

## Kubernetes Ingress
- Ingress is responsible for managing external access to our clusters

## Service Mesh
- A service mesh is an infrastructure layer in your application that facilitate  communication between services
- If one service fails, traffic has to be routed to another service, this is the essence of a service mesh 