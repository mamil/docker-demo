# docker-demo

## container启动已经完成
```
$ sudo ./docker-demo run -ti /bin/sh
```
即可启动container

### 如果/proc下没有进程信息
运行这个挂载，否则无法启动容器
```
mount -t proc proc /proc
```

## 添加资源限制
```
sudo ./docker-demo run -ti -m 100m stress --vm-bytes 200m --vm-keep -m 1
```