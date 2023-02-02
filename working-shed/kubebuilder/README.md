# # Overview

An image to help work with the creation of [controller/operators with kubebuilder](https://book.kubebuilder.io/)

# Building

```bash
docker build  ./working-shed/kubebuilder -t kubebuilder_v2
```

# Run 

```bash 
docker run --rm -it \
    --name kubebuilder \
    -w /go/src \
    -v operator-sdk:/go/pkg \
    -v $(pwd):/go/src \
    -v /var/run/docker.sock:/var/run/docker.sock  \
    --privileged \
    kubebuilder_v2
```



# Using this image

This watcher-operators was created

```bash 
kubebuilder init --domain kubefirst.io --license apache2 --owner "6za" --repo github.com/k1tests/basic-controller

kubebuilder create api --group k1 --version v1beta1 --kind MyEntity --controller --resource

make manifests

make install
```

```bash 
make run
```