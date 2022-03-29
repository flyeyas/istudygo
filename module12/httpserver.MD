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

#### 证书 dns01

kubectl create -f cert-manager-secret.yaml
kubectl.exe get secret -n cert-manager

NAME                                  TYPE                                  DATA   AGE
cloudflare-api-token-secret           Opaque                                1      28s

#### 创建 ClusterIssuer
kubectl create -f issuer-dns.yaml -n httpserver
kubectl get  ClusterIssuer -n httpserver

NAME                READY   AGE
letsencrypt-dns01   True    43s

#### 创建 Certificate
kubectl create -f bbtgo-51-cafe-certificate-dns.yaml -n httpserver
kubectl get certificate -n httpserver

kubectl describe CertificateRequest -n httpserver bbtgo-yunyuansheng-xyz

kubectl get secret bbtgo-yunyuansheng-xyz-tls -n default

查看日志
kubectl logs -f $(kubectl get pods -n cert-manager | grep cert-manager | grep -v 'cainjector\|webhook' | awk '{print $1}') -n cert-manager

E0329 00:13:59.666405       1 controller.go:163] cert-manager/challenges "msg"="re-queuing item due to error processing" "error"="while attempting to find Zones for domain _acme-challenge.bbtgo.yunyuansheng.xyz.\nwhile querying the Cloudflare API for GET \"/zones?name=_acme-challenge.bbtgo.yunyuansheng.xyz\" \n\t Error: 9109: Max auth failures reached, please check your Authorization header." "key"="httpserver/bbtgo-yunyuansheng-xyz-c578l-2832687487-2344801065"




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
