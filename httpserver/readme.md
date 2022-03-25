 kubectl create ns httpserver

kubectl create -f secret.yaml -n secret.yaml

 kubectl  get secret -n httpserver

 kubectl create -f httpserver/deployment.yaml -n httpserver

kubectl create -f httpserver/service.yaml -n httpserver

 kubectl get pod -n httpserver


 kubectl get svc -n httpserver


kubectl describe pod -n httpserver httpserver-7c5d458bb5-vzb86

kubectl describe pod -n harbor harbor-registry-86b5499497-4n8vm

kubectl describe svc -n httpserver httpserver-svc

docker run -d -p 8080:80 maizui216/bbtgo:2.0.0

10.110.88.130 core.harbor.domain

 docker push  core.harbor.domain/admin/bbtgo:2.0.0

helm install ingress-nginx nginx-stable/nginx-ingress --create-namespace --namespace ingress --type NodePort --debug

helm list -n ingress

# 地址dns
nslookup bbtgo.51.cafe


# tekton

sudo crictl pull gcr.io/tekton-releases/github.com/tektoncd/triggers/cmd/interceptors:v0.15.1@sha256:726c308bcab1d746b0cb3b90c5f6abbea64da34459681cba2143c34c337cdc94

sudo crictl pull registry.aliyuncs.com/google_containers/interceptors:v0.15.1@sha256:726c308bcab1d746b0cb3b90c5f6abbea64da34459681cba2143c34c337cdc94

kubectl create -f tekton-release.yaml
kubectl create -f interceptors.yaml
kubectl create -f  tekton-dashboard-release.yaml
kubectl create -f trigger-release.yaml

kubectl delete -f tekton-release.yaml
kubectl delete -f interceptors.yaml
kubectl delete -f tekton-dashboard-release.yaml
kubectl delete -f trigger-release.yaml

kubectl -n tekton-pipelines get po

 kubectl -n tekton-pipelines logs -f tekton-dashboard-8588f97764-rrc4b

kubectl apply -f task-hello.yaml

kubectl create -f taskrun-hello.yaml

kubectl delete -f task-hello.yaml

kubectl delete -f taskrun-hello.yaml


kubectl describe taskrun hello-run-kjbf8


sudo vim /etc/containerd/config.toml