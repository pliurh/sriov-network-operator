# Quickstart Guide

## Prerequisites

1. Kubernetes of Openshift cluster running on bare metal nodes.
2. Multus-cni is deployed as default CNI plugin, and there is a default CNI plugin (flannel, openshift-sdn etc.) available for Multus-cni.

## Installation

Firstly, clone this GitHub repository.

```bash
go get github.com/pliurh/sriov-network-operator
```

Install the Operator-SDK. The following commands will put operator-sdk to your $GOPATH/bin, please make sure that path is included in your \$PATH.

```bash
cd $GOPATH/github.com/pliurh/sriov-network-operator
make deploy-setup
```

Deploy the operator. 

```bash
make deploy-setup
```

By default, the operator will be deployed in namespace 'sriov-network-operator', you can check if the deployment is finished successfully.

```bash
$ kubectl get -n sriov-network-operator all
NAME                                          READY   STATUS    RESTARTS   AGE
pod/sriov-network-config-daemon-bf9nt         1/1     Running   0          8s
pod/sriov-network-operator-54d7545f65-296gb   1/1     Running   0          10s

NAME                             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/sriov-network-operator   ClusterIP   10.102.53.223   <none>        8383/TCP   9s

NAME                                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                                                 AGE
daemonset.apps/sriov-network-config-daemon   1         1         1       1            1           beta.kubernetes.io/os=linux,node-role.kubernetes.io/worker=   8s

NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/sriov-network-operator   1/1     1            1           10s

NAME                                                DESIRED   CURRENT   READY   AGE
replicaset.apps/sriov-network-operator-54d7545f65   1         1         1       10s
```

## Configuration

After the operator gets installed, you can configure it with creating the custom resource of SriovNetwork and SriovNetworkNodePolicy. But before that, you may want to check the status of SriovNetworkNodeState CRs to find out all the SRIOV capable devices in you cluster.

Here comes an example. As you can see, there are 2 SR-IOV NICs from Intel.

```bash
$ oc get sriovnetworknodestates.sriovnetwork.openshift.io -n sriov-network-operator node-1 -o yaml

apiVersion: sriovnetwork.openshift.io/v1
kind: SriovNetworkNodeState
spec: ...
status:
  interfaces:
  - deviceID: "1572"
    driver: i40e
    mtu: 1500
    pciAddress: "0000:18:00.0"
    totalvfs: 64
    vendor: "8086"
  - deviceID: "1572"
    driver: i40e
    mtu: 1500
    pciAddress: "0000:18:00.1"
    totalvfs: 64
    vendor: "8086"
```

You can choose the NIC you want when creating SriovNetworkNodePolicy CR, by specifying the 'nicSelector'.

```yaml
apiVersion: sriovnetwork.openshift.io/v1
kind: SriovNetworkNodePolicy
metadata:
  name: policy-1
  namespace: sriov-network-operator
spec:
  resourceName: intel-nics
  priority: 99
  mtu: 9000
  numVfs: 2
  nicSelector:
      deviceID: "1572"
      rootDevices:
      - 0000:18:00.1
      vendor: "8086"
  deviceType: netdevice
```

After applying your SriovNetworkNodePolicy CR, check the status of SriovNetworkNodeState again, you should be able to see the NIC has been configured as instructed.

```bash
$ oc get sriovnetworknodestates.sriovnetwork.openshift.io -n sriov-network-operator node-1 -o yaml

...
- Vfs:
    - deviceID: 1572
      driver: iavf
      pciAddress: 0000:18:02.0
      vendor: "8086"
    - deviceID: 1572
      driver: iavf
      pciAddress: 0000:18:02.1
      vendor: "8086"
    - deviceID: 1572
      driver: iavf
      pciAddress: 0000:18:02.2
      vendor: "8086"
    deviceID: "1583"
    driver: i40e
    mtu: 1500
    numVfs: 3
    pciAddress: 0000:18:00.0
    totalvfs: 64
    vendor: "8086"
...
```

At the same time, the SRIOV device plugin and CNI plugin has been provisioned to the worker node. You may check if resource name 'intel-nics' is reported  by device plugin correctly.

```bash
$ kubectl get no -o json | jq -r '[.items[] | {name:.metadata.name, allocable:.status.allocatable}]'
[
  {
    "name": "minikube",
    "allocable": {
      "cpu": "72",
      "ephemeral-storage": "965895780801",
      "hugepages-1Gi": "0",
      "hugepages-2Mi": "0",
      "intel.com/intel-nics": "3",
      "memory": "196706684Ki",
      "openshift.io/sriov": "0",
      "pods": "110"
    }
  }
]
```

To remove the operator related resources.

```bash
make undeploy
```

## Hack

To run the operator locally.

````bash
make run
````

To run the e2e test.

```bash
make test-e2e
```

To build the binary.

```bash
make build
```