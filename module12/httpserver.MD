### 证书配置
kubectl  create -f issuer.yaml -n httpserver

kubectl get issuer -n httpserver

NAME               READY   AGE
letsencrypt-prod   True    25s

kubectl get issuer -n httpserver letsencrypt-prod -oyaml

kubectl replace -f ingress.yaml -n httpserver

kubectl get cert -n httpserver

kubectl describe cert -n httpserver
### 

### Deploy simple

```sh
kubectl create ns httpserver
kubectl create -f simple.yaml -n httpserver
kubectl create -f istio-specs.yaml -n httpserver
```

### Check ingress ip

```sh
kubectl get svc -nistio-system

istio-ingressgateway   LoadBalancer   10.108.31.242
```

### Access the httpserver via ingress

```sh
export INGRESS_IP=10.108.31.242
curl -H "Host: httpserver.cncamp.io" $INGRESS_IP/hello -v
```
