传统虚拟机技术是虚拟出一套硬件后，在其上运行一个完整操作系统，在该系统上再运行所需应用进程；而容器内的应用进程直接运行于宿主的内核，容器内没有自己的内核，而且也没有进行硬件虚拟。因此容器要比传统虚拟机更为轻便。

### Docker核心概念： 镜像Image 容器Container 仓库Repository
-镜像类似与类，基础环境定义
-容器相当于类的实例
-仓库相当于nuget maven npm pip 存放定义好的镜像

-ubuntu docker安装脚本
```
$ curl -fsSL get.docker.com -o get-docker.sh
$ sudo sh get-docker.sh --mirror Aliyun
```

下载ubuntu16.04版本Docker镜像，和git类似
docker pull ubuntu:16.04  
运行镜像
docker run -it --rm \
ubuntu:16.04 \
bash

--rm 参数：容器退出后将其删除，节省空间，默认情况下并不会删除，可以执行命令 docker rm
列出已经下载的镜像 ls 和linux列出目录内容一样
docker image ls   
清除虚悬镜像，官方镜像重命名，放弃了旧的名称
docker image ls -f dangling=true  或者 docker image prune
查看中间层镜像
docker image ls -a
对容器内容修改之后执行以下命令将容器保存为镜像
docker commit \
--author "Tao Wang <twang2218@gmail.com>" \
--message "修改了默认网页" \
webserver \
nginx:v2

- Union FS 最多127层 写Dockerfile时需要注意
- 执行Docker命令和普通的linux命令不一样，docker是c/s设计，执行run指令时，经常会需要将一些本地文件传到服务器，然后在服务器端进行构建
- docker save 保存镜像  docker load 加载保存的镜像
- docker export 导出容器 docker import 导入容器

### 数据卷
- 数量卷 可以在容器之间共享和重用
- 对 数据卷 的修改会立马生效
- 对 数据卷 的更新，不会影响镜像
- 数据卷 默认会一直存在，即使窗口被删除
#### 命令
- docker volume create my-vol 创建一个数据卷
- docker volume ls 查看所有的数据卷
- docker volume inspect my-vol 查看指定的数据卷
- 创建一个名为 web 的窗口，加载一个 数据卷 到容器的 /webapp 目录
``` bash
docker run -d -p \
--name web \
--mount source=my-vol,target=/webapp \
training/webapp \
python app.py
```
- docker volume rum my-vol 删除数据卷
- docker volume prune 清除无主的数据卷
#### 挂载主机目录
- 挂载一个本地主机的目录到容器中
``` bash
docker run -d -p \
--name web \
--mount type=bind,source=/src/webapp,target=/opt/webapp \
training/webapp \
python app.py
```
### 网络功能
- docker container ls 查看容器列表，可以看到本地主机和容器的端口映射
- docker network create -d bridge my-net 创建一个名为my-net的docker网络
- docker run -it --rm --name busybox1 --network my-net busybox sh 运行一个容器并链接到新建的my-net网络 busybox1
- docker run -it --rm --name busybox2 --network my-net busybox sh 运行一个容器并链接到新建的my-net网络 busybox2
- busybox1中输入命令：ping busybox2  可进行测试
- 多个容器之间互联推荐使用Docker Compose.
