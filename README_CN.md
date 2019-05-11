Gmirror
====== 

[English](README.md) | 简体中文

用于在 Git 仓库间进行同步

**警告！大多数功能仍在开发阶段**

# 安装

## 从二进制安装

## 从源码安装

需要的工具

* [Git](https://git-scm.com/)
* [Go (Go 1.11 或之后的版本)](https://golang.org/dl/)

### 1.从Github获取源码

```bash
git clone https://github.com/countstarlight/gmirror.git
```

### 2.编译安装

```bash
cd gmirror
go install
```

# 常见问题

## 1.提示没有指定 `SSH_AUTH_SOCK`

```bash
FATA[14:37:03] error creating SSH agent: "SSH agent requested but SSH_AUTH_SOCK not-specified" 
```
需要确保 `ssh-agent` 正在运行：
```bash
eval `ssh-agent` # 输出: Agent pid xxxx
```
并且已经添加SSH私钥：
```bash
ssh-add
# 输出: Identity added: /home/orawlings/.ssh/id_rsa (/root/.ssh/id_rsa)
```