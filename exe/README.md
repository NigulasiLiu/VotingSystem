# 运行方法

### 环境配置

- node.js: 18.15.0
- yarn: 1.22.19
- MySQL: 8.0.35


后端已打包为可执行文件无需进行环境配置。

### 可执行文件


1. 确保你的MySQL数据库中有`votingsystem`库和`server`库

2. 修改`config`文件夹中的三个配置文件，确定MySQL数据库的账号密码无误

3. 分别运行三个可执行文件

```bash
.\ea
.\server0
.\server1
```


### 前端


在`frontend`文件夹下运行以下指令：

1. 下载依赖

```bash
yarn install
```

2. 运行前端

```bash
yarn serve
```

