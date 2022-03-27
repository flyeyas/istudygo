### Deploy simple

```sh
kubectl create ns httpserver
kubectl create -f simple.yaml -n httpserver
kubectl create -f istio-specs.yaml -n httpserver
```

### Check ingress ip

```sh
k get svc -nistio-system

istio-ingressgateway   LoadBalancer   10.108.31.242
```

### Access the httpserver via ingress

```sh
export INGRESS_IP=10.108.31.242
curl -H "Host: httpserver.cncamp.io" $INGRESS_IP/hello -v
```
