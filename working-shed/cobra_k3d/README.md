# Overview 

This one is a tricky image needed to allow testing [kubefirst local](https://github.com/kubefirst/kubefirst) from a contianer to validate changes of kubefirst `main`. 


# Building

As we need to "trick" `k3d` in the contianer that it is running on the host, the way that happens is by changing the "defualt" HOME path. So, that is why `FAKE_HOME` enters in play for this image and when usung the container. 

```bash 
export FAKE_HOME=$HOME/_tmp_k3d_user

docker build -t cobra_k3d_v2 working-shed/cobra_k3d --build-arg HOME_PATH=$FAKE_HOME --no-cache

```

# Requirements

```bash
export AWS_CREDS_PATH=$HOME/.aws
export BRANCH=SOMENAME
export SSH_CREDS=$HOME/.ssh
export FAKE_HOME=$HOME/_tmp_k3d_user
export GITHUB_TOKEN='ghp_l....'

```


# Running 

```bash 
docker run \
  -v  go-pkg:/go/pkg \
  -it --name $BRANCH  \
  --dns="1.0.0.1" --dns="208.67.222.222" --dns="8.8.8.8" \
  -v $PWD:/app \
  --network="host" \
  -v /var/run/docker.sock:/var/run/docker.sock  \
  -e "TERM=xterm-256color" \
  -v $FAKE_HOME/.k3d:$FAKE_HOME/.k3d  \
  -e KUBEFIRST_GITHUB_AUTH_TOKEN=$GITHUB_TOKEN \
  -e HOME=$FAKE_HOME \
   cobra_k3d
```
