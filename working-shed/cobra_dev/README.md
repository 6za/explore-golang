# Overview 

This image is used help use [kubefirst CLI](https://github.com/kubefirst/kubefirst) in contianer enviorment for tests and running builds of it. 

This image is meant only to test the cluud version of the installer, as the local requires some docker magic to allow it to work properly with k3d. 


# Building

```bash
docker build working-shed/cobra_dev -t cobra_dev_v2
```
# Requirements

```bash
export AWS_CREDS_PATH=$HOME/.aws
export BRANCH=SOMENAME
export SSH_CREDS=$HOME/.ssh

```
- `SSH_CREDS` can be used for git operations  
- `AWS_CREDS_PATH` is used for AWS auth

# Running 

```bash 
docker run \
  -v  go-pkg:/go/pkg \
  -v $SSH_CREDS:/home/developer/.ssh/ \
  -it --name ${BRANCH}  \
  --dns="1.0.0.1" --dns="208.67.222.222" --dns="8.8.8.8" \
  -v $(PWD):/home/developer/app \
  -e "TERM=xterm-256color" \
  -v $AWS_CREDS_PATH:/home/developer/.aws \
   cobra_dev_v2
```
