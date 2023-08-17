# Kubernetes Storage
- We use volumes in kubernetes to persist data
- There are 3 components of kubernetes storage
    - persistent volume
    - persistent volume claim
    - storage class
- Kubernetes does not provide data persistence 
- We need:
    - storage that does not depend on the pod lifecycle and will stll be there if the pod gets recreated
    - The storage to be available on all nodes
    - Storage needs to survive even if cluster crashes

## Persistent Volume
- Is a cluster resource that is used to store data
- PVs are not namespaced and are accessible to the whole cluster 
- PVs are resources that need to be in the cluster before the pods that depend on   it is created

## Persistent Volume claim(PVC)
- Our application has to claim the persistent volume using PVC
- So the pods access storage by using the claim as a volume, the claim
will then try to find a volume oin the cluster that satisfies the claim.
The volume will have the storage backend that it will create the storage resource from, the pod will then be able to use the actual storage backend 

## Storage Class
- SCs provision PVs dynamically when a PVC claims it
- StorageBackend is defined in the SC component via the "provisioner" attribute
- each StorageBackend has it's own provisioner

#### So...
- A pod claims storgae via PVC
- PVC will request storage from SC
- SC creates PV that meets the needs of the claim 

## General Knowledge
- Some local volumes we have are configmaps and secrets
- We can think of storage in kubernetes as an external plugin to your cluster
- Volume is just some directory with data
- A pod can use multiple volumes of differnet types simultaneously