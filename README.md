# 阿里云域名动态解析

## 前提
- 拥有阿里云域名
- 拥有动态外网IP地址
- 阿里云创建RAM用户访问控制权限，要达到在Ram访问控制/授权中看到以下４栏内容，否则可能出现更新错误的问题
      RAM角色  AliyunDNSFullAccess  系统策略  管理云解析（DNS）的权限   
      RAM角色  AliyunDomainFullAccess  系统策略  管理域名服务的权限
      AliyunDNSFullAccess  系统策略  管理云解析（DNS）的权限
      AliyunDomainFullAccess  系统策略  管理域名服务的权限
## 编译环境配置
- 下载golang1.13以后版本，配置golang开发环境，执行go get
- 配置编译目标的运行环境
    - linux环境: Set GOOS=linux
    - windows环境: Set GOOS=windows
    - CPU指令集：arm: Set GOARCH=arm
                 x86-64: Set GOARCH=amd64

    - 最后执行go build即可
    
## 配置文件
把config目录下的config.sample.json文件改名config.json，并在里面添加相应的字段，具体内容不做解释

## 运行
sudo crontab -e 写入以下内容,意思是每５分钟运行一次程序
/5 * curl -4 /path-to/aliddns