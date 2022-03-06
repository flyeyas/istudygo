 kubectl create ns httpserver

 kubectl create -f httpserver/deployment.yaml -n httpserver

 kubectl get pod -n httpserver

kubectl describe pod -n httpserver httpserver-f95c85fd7-8r2jx

docker run -d -p 8080:80 maizui216/bbtgo:2.0.0