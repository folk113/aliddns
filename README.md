# 阿里云域名动态解析

## 前提
- 拥有阿里云域名
- 拥有动态外网IP地址
- 阿里云创建RAM用户访问控制权限，要达到在[Ram访问控制/授权]中看到以下４栏内容，否则可能出现更新错误的问题
  |RAM角色 | AliyunDNSFullAccess  |系统策略 | 管理云解析（DNS）的权限|
  | -----  | ---- |---- | ---- |
  | RAM角色  |AliyunDomainFullAccess  |系统策略  |管理域名服务的权限|
  | 用户  | AliyunDNSFullAccess  |系统策略 | 管理云解析（DNS）的权限|
  | 用户  |  AliyunDomainFullAccess  |系统策略 | 管理域名服务的权限|
## 编译环境配置
- 下载golang1.13以后版本，配置golang开发环境，在工程根目录执行go get

- 配置编译目标的运行环境
    
    > 以下为windos编译环境，如果在类bsd unix环境下编译，把Set换成export
    
    - linux环境: <br/>Set GOOS=linux
- windows环境: <br/>Set GOOS=windows
    - CPU指令集：<br/>arm: Set GOARCH=arm
                 <br/>x86-64: Set GOARCH=amd64
    
    - 最后执行go build即可
    
## 配置文件
把config目录下的config.sample.json文件改名config.json，并在里面添加相应的字段，具体内容不做解释

## 运行
sudo crontab -e 写入以下内容,意思是每５分钟运行一次程序，并把结果写入到同级目录下的log.log中<br/>
*/5 * * * *  /project-absolute-path/run.sh

## FAQ
有问题，请联系:<br/>
邮箱:8980728@qq.com<br>
微信:<br/>![](./wechat.png)
