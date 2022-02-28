 kubectl create ns httpserver

 kubectl create -f httpserver/deployment.yaml -n httpserver

 kubectl get pod -n httpserver

 kubectl describe pod -n httpserver httpserver-8b4856448-xh4j7