FROM golang:1.15.2 as build
LABEL maintainer="SunXiaoHu <m15829090357@163.com>"
# 环境变量
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
# 将目录加载到镜像里----不需要要这么做 或者每次进去需要git pull 一下 不然代码永远是最开始加载的代码
#ADD . /go/src/yurendao.api

# 工作目录
WORKDIR /go/src/yurendao.api

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone