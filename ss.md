Kali2.0 update到最新版本后安装shadowsocks服务报错问题

https://blog.csdn.net/blackfrog_unique/article/details/60320737

```
vim /usr/local/lib/python2.7/dist-packages/shadowsocks/crypto/openssl.py
```

## ss配置
安装
```
apt-get update
apt-get install python-pip python-setuptools m2crypto
pip install setuptools --upgrade
pip install wheel shadowsocks
```
配置文件： 
```
shadowsocks.json
{
    "server":"::",
    "server_port":8388,
    "local_port":1080,
    "password":"digitaloceanss",
    "timeout":300,
    "method":"aes-256-cfb",
    "fast_open": true
}
scp -r shadowsocks.json root@144.202.102.232:/root/shadowsocks.json
```
后台运行
```
ssserver -c shadowsocks.json -d start
ssserver -c shadowsocks.json -d stop
```

## 更换内核步骤
(vultr ubuntu 16.04)

1：查看当前安装的内核
```
dpkg -l|grep linux-image
```

2：安装新内核
```
apt-get install linux-image-4.4.0-47-generic linux-image-extra-4.4.0-47-generic
```

3：卸载不要的内核
```
apt-get purge linux-image-4.4.0-62-generic linux-image-extra-4.4.0-62-generic linux-image-4.4.0-87-generic linux-image-extra-4.4.0-87-generic
apt-get purge linux-image-4.4.0-104-generic linux-image-extra-4.4.0-104-generic
```

4：更新 grub引导
```
update-grub
```

重启系统查看

5：查看内核版本
```
cat /proc/version
```
或
```
uname -a
```

## 装锐速

安装
```
wget --no-check-certificate -O appex.sh https://raw.githubusercontent.com/0oVicero0/serverSpeeder_Install/master/appex.sh && chmod +x appex.sh && bash appex.sh install
```
卸载
```
wget --no-check-certificate -O appex.sh https://raw.githubusercontent.com/0oVicero0/serverSpeeder_Install/master/appex.sh && chmod +x appex.sh && bash appex.sh uninstall
```
更新许可证
```
bash /appex/bin/serverSpeeder.sh renewLic
```
加速模块检索
```
https://github.com/0oVicero0/serverSpeeder_kernel/blob/master/serverSpeeder.txt
```
