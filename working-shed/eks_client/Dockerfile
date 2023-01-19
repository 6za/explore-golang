FROM ubuntu:22.04
RUN \
    apt-get update && apt-get install -y \
    curl jq vim unzip less \
    && rm -rf /var/lib/apt/lists/* 

RUN \
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.21.3/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl

RUN \
    curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp \
    &&  mv /tmp/eksctl /usr/local/bin  

RUN \
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"   \
    && unzip awscliv2.zip \
    && ./aws/install \
    && aws --version
