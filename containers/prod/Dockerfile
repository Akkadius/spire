#############################################
# go workspace
#############################################
FROM debian:10-slim

ARG PUID=1000
ENV PUID ${PUID}
ARG PGID=1000
ENV PGID ${PGID}

ENV GO_VERSION 1.17.3

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
    make \
    mariadb-client \
    zip \
    unzip \
 && rm -rf /var/lib/apt/lists/*

RUN curl -sL https://deb.nodesource.com/setup_12.x | bash && \
    apt-get update && apt-get install -y nodejs npm node-gyp && rm -rf /var/lib/apt/lists/*

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

#######################################################################
# default shell
#######################################################################
RUN chsh -s /bin/bash go

USER go

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/gobuffalo/packr/packr

#############################################
# bash
#############################################
ENV PS1 '\[\e]0;\w\a\]\n\[\e[32m\]\u@\h \[\e[33m\]\w\[\e[0m\]\n\$ '

WORKDIR /home/go/src
