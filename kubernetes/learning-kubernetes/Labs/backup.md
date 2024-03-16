Take the backup of ETCD at the location /opt/etcd-backup.db on the controlplane node.

BUILT IN SNAPSHOP
ETCDCTL_API=3 etcdctl --endpoints $ENDPOINT snapshot save /opt/etcd-backup.db 

verify the snapshot
ETCDCTL_API=3 etcdctl --write-out=table snapshot status /opt/etcd-backup.db 