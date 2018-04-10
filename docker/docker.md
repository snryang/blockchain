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