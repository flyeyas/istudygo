jenkins.yaml： https://github.com/cncamp/101/blob/master/module10/jenkins/jenkins.yaml

创建 Jenkins Master
kubectl apply –f Jenkins.yaml
kubectl apply –f sa.yaml

等待 Jenkins-0 pod 运行，查看日志查找 root 密码
kubectl logs –f Jenkins-0

查看 Jenkins Service 的 NodePort，登录 Jenkins console
 http://192.168.34.2:<nodePort>

输入 root 密码并创建管理员用户，登录


为了确保管理员安全地安装 Jenkins，密码已写入到日志中（不知道在哪里？）该文件在服务器上：

/var/jenkins_home/secrets/initialAdminPassword