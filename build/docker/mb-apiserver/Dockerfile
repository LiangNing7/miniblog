# 定义生产环境的基础镜像为 Debian Trixie
ARG prod_image=debian:trixie

# 使用指定的基础镜像
FROM ${prod_image}
# 设置维护者信息
LABEL maintainer="<xan277@foxmail.com>"

# 设置工作目录为 /opt/miniblog
WORKDIR /opt/miniblog

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone  # 设置时区为上海

RUN mkdir -p /opt/miniblog/log # 创建日志目录

# 将 mb-apiserver 可执行文件复制到工作目录下的 bin 目录
COPY mb-apiserver /opt/miniblog/bin/

# 设置容器启动时执行的命令
ENTRYPOINT ["/opt/miniblog/bin/mb-apiserver"]
