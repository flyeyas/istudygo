### 证书配置
kubectl  create -f issuer.yaml -n httpserver

kubectl get issuer -n httpserver

NAME               READY   AGE
letsencrypt-prod   True    25s

kubectl get issuer -n httpserver letsencrypt-prod -oyaml

kubectl replace -f ingress.yaml -n httpserver

kubectl get cert -n httpserver

kubectl describe cert -n httpserver

kubectl get CertificateRequest -n httpserver

kubectl describe CertificateRequest -n httpserver bbtgo-tls-xlg55

### 部署gateway并且暴露服务

istioctl install --set profile=demo -y

kubectl create ns httpmesh

kubectl label ns httpmesh istio-injection=enabled

kubectl get ns -L istio-injection

NAME               STATUS   AGE     ISTIO-INJECTION
calico-system      Active   12d
cert-manager       Active   141m
cloudnative        Active   5d16h
default            Active   12d
httpmesh           Active   15s     enabled
httpserver         Active   2d16h
ingress            Active   2d15h
istio-system       Active   4m21s
jenkins            Active   12d
kube-node-lease    Active   12d
kube-public        Active   12d
kube-system        Active   12d
prometheus-stack   Active   12d
tigera-operator    Active   12d

kubectl -n httpmesh create deployment my-nginx --image=nginx

kubectl -n httpmesh get pod

NAME                        READY   STATUS    RESTARTS   AGE
my-nginx-6b74b79f57-474hb   2/2     Running   0          8s


### 作业
kubectl get pod -n httpserver

NAME                         READY   STATUS    RESTARTS   AGE
cm-acme-http-solver-2gjjs    1/1     Running   0          77m
httpserver-698d494d4-lfmnt   1/1     Running   0          16m

kubectl delete pod httpserver-698d494d4-lfmnt -n httpserver

kubectl create -f service.yaml
kubectl create -f virtualserver.yaml


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
