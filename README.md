# Playing with HPA

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