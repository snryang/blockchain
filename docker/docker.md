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
