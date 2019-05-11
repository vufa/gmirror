Gmirror
====== 

English | [简体中文](README_CN.md)

Mirror between git repositories

**WARNING! Most of the functions are still in development.**

# Install

## From release(Binary Install)

## From source

Prerequisite Tools

* [Git](https://git-scm.com/)
* [Go (at least Go 1.11)](https://golang.org/dl/)

### 1.Fetch from GitHub

```bash
git clone https://github.com/countstarlight/gmirror.git
```

### 2.Build and install

```bash
cd gmirror
go install
```

# Troubleshooting

## 1.Need to specifie `SSH_AUTH_SOCK`:

```bash
FATA[14:37:03] error creating SSH agent: "SSH agent requested but SSH_AUTH_SOCK not-specified" 
```
need to ensure that `ssh-agent` running：
```bash
eval `ssh-agent` # Output: Agent pid xxxx
```
and private ssh keys added:
```bash
ssh-add
# Output: Identity added: /home/orawlings/.ssh/id_rsa (/root/.ssh/id_rsa)
```