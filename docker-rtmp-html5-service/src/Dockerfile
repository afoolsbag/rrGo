# ------------------------------------------------------------------------------
# BUILD TIME
FROM centos:8 AS build-time

RUN dnf upgrade --assumeyes
RUN dnf group install --assumeyes 'Development Tools'
RUN dnf install --assumeyes git openssl-devel pcre-devel wget zlib-devel

WORKDIR /usr/local/src
RUN wget https://nginx.org/download/nginx-1.18.0.tar.gz
RUN tar -xavf nginx-1.18.0.tar.gz
RUN git clone https://github.com/arut/nginx-rtmp-module.git

WORKDIR /usr/local/src/nginx-1.18.0
RUN ./configure --prefix='/opt/nginx' --add-module='../nginx-rtmp-module'
RUN sed --in-place 's/-Werror//g' objs/Makefile
RUN make --jobs
RUN make install

# ------------------------------------------------------------------------------
# RUN TIME
FROM centos:8 AS run-time

LABEL maintainer="afoolsbag@gmail.com"

RUN dnf upgrade --assumeyes
RUN ln --symbolic --force /usr/share/zoneinfo/Asia/Chongqing /etc/localtime
RUN dnf install --assumeyes openssl pcre zlib

COPY --from=build-time /opt/nginx                                /opt/nginx
COPY --from=build-time /usr/local/src/nginx-rtmp-module/stat.xsl /opt
COPY                   nginx.conf                                /opt/nginx/conf
RUN mkdir /var/vod
RUN ln --symbolic /opt/nginx/sbin/nginx /usr/local/sbin/nginx

EXPOSE 80/tcp
EXPOSE 1935/tcp
VOLUME /var/vod

WORKDIR /opt/nginx
CMD ["nginx", "-g", "daemon off;"]
