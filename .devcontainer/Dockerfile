FROM ubuntu:jammy

LABEL maintainer="Vanessasaurus <@vsoch>"

# Match the default user id for a single system so we aren't root
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000
ENV USERNAME=${USERNAME}
ENV USER_UID=${USER_UID}
ENV USER_GID=${USER_GID}
USER root

# Install extra buildrequires for flux-sched:
RUN apt-get update && \
    apt-get -qq install -y --no-install-recommends \
        wget \
        curl \ 
        less \
        git \
        locales \
        autoconf \
        libtool \
        automake \
        build-essential \
        ca-certificates

# Assuming installing to /usr/local
ENV LD_LIBRARY_PATH=/usr/lib:/usr/local/lib
ENV PATH=$PATH:/workspaces/hwloc-go/bin

RUN wget --no-check-certificate https://go.dev/dl/go1.21.7.linux-amd64.tar.gz && tar -xvf go1.21.7.linux-amd64.tar.gz && \
         mv go /usr/local && rm go1.21.7.linux-amd64.tar.gz

# Install hwloc
RUN git clone https://github.com/open-mpi/hwloc /opt/hwloc && \
    cd /opt/hwloc && \
    ./autogen.sh && \
    ./configure --enable-static --disable-shared LDFLAGS="-static" && \
    make LDFLAGS=-all-static && \
    make install && ldconfig /usr/local/lib

ENV PATH=$PATH:/usr/local/go/bin:/home/vscode/go/bin

# Add the group and user that match our ids
RUN groupadd -g ${USER_GID} ${USERNAME} && \
    adduser --disabled-password --uid ${USER_UID} --gid ${USER_GID} --gecos "" ${USERNAME} && \
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers