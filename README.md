# Cloud Flare Rss

## 1. 创建配置文件的secret
```
k create secret generic cloudflare-conf --from-file=config.json
```

## 2. 创建拉取私有仓库镜像的secret
```
kubectl create secret docker-registry registrykey-u-p \
--docker-server="registry.cn-hongkong.aliyuncs.com/xxx/cloudflare_rss" \
--docker-username="xxx@xxxx" \
--docker-password="xxxxx" \
--docker-email="xxxx@xxx.com"
```

## 3. 部署Deployment
```
kaf cloudflare_rss.yaml
```

## 4. 检查服务是否运行
```
kgp
```

## 5. 进入pod中的container查看
```
keti cloudflarerss-549775c4db-qpztg -c cloudflarerss sh
```



