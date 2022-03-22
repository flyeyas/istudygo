1. 创建自动发现规则
   kubectl create ns cloudnative
   kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml -n  prometheus-stack

   kubectl get secret -n prometheus-stack additional-configs -oyaml
2. 注入Prometheus
   1. kubectl get crd | grep prometheus
   2. kubectl get prometheuses.monitoring.coreos.com -n prometheus-stack
   3. kubectl edit prometheuses.monitoring.coreos.com -n prometheus-stack kube-prometheus-stack-prometheus
      1. 在spec:添加additionalScrapeConfigs:
      key: prometheus-additional.yaml
      name: additional-configs
      1. 查看是否生效
         1. kubectl get pod -n prometheus-stack 
         2. kubectl logs -n prometheus-stack  kube-prometheus-stack-operator-7654f9dbc7-78kbc
   4. 查看svc 
      1. kubectl get svc -n prometheus-stack 