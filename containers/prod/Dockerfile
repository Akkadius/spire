#############################################
# go workspace
#############################################
FROM debian:12-slim

ARG PUID=1000
ENV PUID ${PUID}
ARG PGID=1000
ENV PGID ${PGID}

ENV GO_VERSION 1.19.1

RUN apt-get update -yqq && \
    groupadd -g ${PGID} go && \
    useradd -u ${PUID} -g go -m go -G go && \
    usermod -p "*" go

#############################################
# install dependencies
#############################################
RUN apt-get update && apt-get install -y \
    bash \
    curl \
    git \
    inetutils-ping \
    wget \
    sudo \
    make \
    mariadb-client \
    zip \
    unzip \
    && rm -rf /var/lib/apt/lists/*

#############################################
# install go
#############################################
RUN cd /tmp && wget --quiet https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
	&& tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && rm -rf /tmp/*

#############################################
# set go env vars
# https://golang.org/doc/code.html
#############################################
ENV GOPATH=/home/go
ENV GOROOT=/usr/local/go/
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
ENV CGO_ENABLED 0

#############################################
# sudo
#############################################
RUN echo "go ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/user && \
     chmod 0440 /etc/sudoers.d/user

#######################################################################
# default shell
#######################################################################
RUN chsh -s /bin/bash go

USER go

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/gobuffalo/packr/packr@latest

#######################################################################
# node
#######################################################################
RUN sudo apt-get update && sudo apt-get install -y curl sudo gnupg nodejs npm \
    && sudo rm -rf /var/lib/apt/lists/*

# Remove sudo; we needed it during build steps to make per-user installs smoother
USER root
RUN rm /etc/sudoers.d/user
USER go

#############################################
# bash
#############################################
ENV PS1 '\[\e]0;\w\a\]\n\[\e[32m\]\u@\h \[\e[33m\]\w\[\e[0m\]\n\$ '

WORKDIR /home/go/src
