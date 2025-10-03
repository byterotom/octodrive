# Using ubuntu base image after testing with a lot of different images
FROM ubuntu:22.04

WORKDIR /root/app

# Installing all the gui dependencies
RUN apt update && \
    apt install -y curl build-essential libgtk-3-dev libwebkit2gtk-4.0-dev pkg-config

# Installing and setting up go
RUN curl -LO https://dl.google.com/go/go1.25.1.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf go1.25.1.linux-amd64.tar.gz \
    && rm go1.25.1.linux-amd64.tar.gz
ENV PATH=/usr/local/go/bin:/root/go/bin:$PATH
ENV GOPATH=/root/go

# Installing and setting up npm and nodejs through nvm 
ENV NVM_DIR=/root/.nvm
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash \
    && . "$NVM_DIR/nvm.sh" \
    && nvm install 22.18.0 \
    && nvm alias default 22.18.0

# Installing wails
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest