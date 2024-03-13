# UPGRADE KUBEADM VERSION FROM 1.28.X TO 1.29.X

"""
Upgrade the current version of kubernetes from 1.28.0 to 1.29.0 exactly using the kubeadm utility.
Make sure that the upgrade is carried out one node at a time starting with the controlplane node. 
To minimize downtime, the deployment gold-nginx should be rescheduled on an alternate node before 
upgrading each node.

Upgrade controlplane node first and drain node node01 before upgrading it. Pods for gold-nginx 
should run on the controlplane node subsequently.
"""

1. Configure APT to install version 1.29, by editing the kubernetes.list file

   - vi /etc/apt/sources.list.d/kubernetes.list
   - deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.28/deb/ /

   Change 1.28 to 1.29 in the file

>>> Follow official documentation for how to continue the upgrade [https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/]

2. Determine what version to upgrade to
   
   - sudo apt update
   - sudo apt-cache madison kubeadm
   '''
   kubeadm | 1.29.2-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   kubeadm | 1.29.1-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   kubeadm | 1.29.0-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   '''

## 3. Upgrade the controlplane
   1. Upgrade kubeadm
      - sudo apt-mark unhold kubeadm && \
        sudo apt-get update && sudo apt-get install -y kubeadm='1.29.0-1.1' && \
        sudo apt-mark hold kubeadm

   2. Verify that the download works and has the expected version:
      - kubeadm version

   3. Verify the upgrade plan:
      - sudo kubeadm upgrade plan
     
   4. Choose a version to upgrade to, and run the appropriate command.
      # replace x with the patch version you picked for this upgrade
      - sudo kubeadm upgrade apply v1.29.x
    
## 4. For the other controlplane nodes
   1. Do the same steps as above but instead of 
      - sudo kubeadm upgrade apply
      Do:
      - sudo kubeadm upgrade node  #RUN THIS IN THE NODE
   2. Drain the node
     `kubectl drain <node-to-drain> --ignore-daemonsets` #RUN THIS IN THE CONTROLPLANE
   3. Upgrade kubelet and kubectl
      # replace x in 1.29.x-* with the latest patch version
      - sudo apt-mark unhold kubelet kubectl && \
        sudo apt-get update && sudo apt-get install -y kubelet='1.29.x-*' kubectl='1.29.x-*' && \
        sudo apt-mark hold kubelet kubectl
   4. Restart the kubelet:
      - sudo systemctl daemon-reload
      - sudo systemctl restart kubelet
   5. Uncordon the node
      Bring the node back online by marking it schedulable:
      # replace <node-to-uncordon> with the name of your node
      - kubectl uncordon <node-to-uncordon>

5. Verify the status of the cluster
   - kubectl get nodes
   '''
   NAME           STATUS   ROLES           AGE   VERSION
   controlplane   Ready    control-plane   73m   v1.29.0
   node01         Ready    <none>          72m   v1.29.0
   '''

    