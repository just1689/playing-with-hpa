# Playing with HPA

This small demo is an example of how to use horizontal pod autoscaling for scaling out workloads by CPU usage metrics in Kubernetes.

The applications included:
- batch: for starting new jobs
- worker: for "calculating interest"
- counter: keeping track of how many jobs have reported being done.

All three are built as a single binary and they use the variable `job` to determine what to run.

The queue chosen for this project is NSQ.

The `k8s/` directory contains yaml for deployment to a Kubernetes environment.

## Before you run

### Install metrics-server
Clone https://github.com/kubernetes-sigs/metrics-server.git
```shell script
kubectl apply -f deploy/1.8+
```
### Edit the deployment 

```shell script
kubectl edit deployment/metrics-server -n kube-system
```

add the flags within args:
```yaml
 - --kubelet-insecure-tls
 - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
```

Save and exit


## How to run

### Build

Build the image and push it to a registry that your cluster has access to
`docker build -t MY_IMAGE:TAG .`

Modify the k8s/*.yaml files' image bit to reference your image.


### Deploy to K8s
```shell script
kubectl apply -f k8s/
```

Port forward the batch bit
```shell script
kubectl port-forward service/batch 8080
```

Open http://localhost:8080/ in the browser

Ctrl + C to stop it

Now forward the counter

Port forward the batch bit
```shell script
kubectl port-forward service/counter 8080
```

Open http://localhost:8080/get in the browser


### Observing

```shell script
kubectl top pods
kubectl top nodes
kubectl logs -f <pods/BATCH_POD_ID>
```
