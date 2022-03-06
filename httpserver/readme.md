 kubectl create ns httpserver

 kubectl create -f httpserver/deployment.yaml -n httpserver

  kubectl create -f httpserver/service.yaml -n httpserver

 kubectl get pod -n httpserver

kubectl describe pod -n httpserver httpserver-7c5d458bb5-vzb86

docker run -d -p 8080:80 maizui216/bbtgo:2.0.0