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

### 操作Cgroups
```
# mkdir cgroup-test

# 挂载一个hierarchy
# sudo mount -t cgroup -o none,name=cgroup-test cgroup-test ./cgroup-test/

# 挂载之后系统在这个目录下生成了默认文件
# ls ./cgroup-test
cgroup.clone_children  cgroup.procs  cgroup.sane_behavior  notify_on_release  release_agent  tasks
```

#### 扩展子cgroup
```
# cd cgroup-test
# sudo mkdir cgroup-1
# sudo mkdir cgroup-2

# cd cgroup-1
# ls
cgroup.clone_children  cgroup.procs  notify_on_release  tasks
```
只要创建文件夹就会自动创建需要的文件
kernel会把文件夹标记为这个cgroup的子cgroup，会继承父cgroup的属性

#### 在cgroup中移动进程
只需要把进程id写到目标cgroup的tasks文件中即可