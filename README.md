# CNCAMP Module 10 Homework
## 需求
- 为 HTTPServer 添加 0-2 秒的随机延时
- 为 HTTPServer 项目添加延时 Metric
- 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
- 从 Prometheus 界面中查询延时指标数据
- （可选）创建一个 Grafana Dashboard 展现延时分配情况

## 安装Prometheus
```
# create namespace monitoring
make pre-apply

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# 配置grafana ingress grafana.zzzzzsy.xyz
# ./kube-prometheus-stack/values.yaml
# line 701-767

# install prometheus with helm
make install-prom
```

## 执行代码
```
#本地测试
make test

#生成docker镜像
make docker-build

#上传镜像
make docker-push

#本地启动docker镜像
make docker-run

#根据`module08/manifests/`里到文件 发布到k8s集群
#发布前请确保您当前的k8s集群是目标集群
make apply

#清理
#删除cncamp namespace下所有资源
make cleanup
```

## 代码实现
- hello接口添加随机等待函数（0-2秒）
- 增加Prometheus指标httpclient_response_time_seconds来监测访问/hello接口的响应时间
- 新增一条路由/metrics用来显示Prometheus指标
- Patch Deployment文件，新增Prometheus annotations配置
- 新增./manifests/prometheus.yaml配置grafana.zzzzzsy.xyz自签证书

## 测试
- Prometheus
![Alt text](./img/prometheus.jpg?raw=true "Prometheus")

- Grafana
![Alt text](./img/grafana.jpg?raw=true "Grafana")