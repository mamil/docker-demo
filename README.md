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

## AUFS测试

```
✗ sudo mount -t aufs -o dirs=./container-layer:./image-layer4:./image-layer3:./image-layer2:./image-layer1 none ./mnt
```

### [done]问题- 发现运行之后可执行文件会消失
之前设置cgroup名字有问题，删除的cgroup的时候会把执行文件删掉

### [done]问题- 运行之后需要重新mount proc
这个应该是在容器里面重新mount proc导致的，如果容器里面不mount proc，退出后宿主机是正常的
```go
    defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
    syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
```

---
在退出container之后需要重新在宿主机mount proc，就可以解决这个问题

### 问题- proc现在只能挂载在容器或者宿主机，二选一
在解决上面问题之后，退出容器虽然能在宿主机正常使用proc，但是容器运行时，宿主机不能使用proc，还需要改进


## volume数据卷

```
# sudo ./docker-demo run -ti -v /root/volume:/containerVolume sh
```

---
### 问题- 容器退出之后，资源没有清除，mnt处于无法删除状态
先恢复proc挂载，`mount -t proc proc /proc`
用这个命令可以让文件恢复正常 `sudo umount /root/mnt -l`
然后就可以正常删除了

应该和挂载/proc有关，不挂载可以删除

### 问题- umount 失败，报错如下
```
umount: /root/mnt2/containerVolume: umount failed: No such file or directory.
```

### 运行commit
```
# sudo ./docker-demo commit 123
```
会在/root 生成
```
# ll
total 1.5M
-rw-r--r--  1 root root   20 Aug 12 11:40 123.tar
```

但现在的代码还有问题，会报错
```
Tar folder /root/mnt error exit status 2
```
生成的tar文件里面没有东西
