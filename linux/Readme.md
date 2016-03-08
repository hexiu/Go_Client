# Linux 客户端使用要求

- 切换目录至(`cd /mnt/`)
- 将[Linux 客户端压缩包](http://7xpugm.com1.z0.glb.clouddn.com/client.zip "Linux 客户端") 下载至mnt目录下
- 解压缩`unzip client`
  ```
  也可以：
  wget https://github.com/hexiu/Go_Client/linux/client.zip
  ```
- 将 startClient.sh 脚本放在/etc/profile.d 目录下
  - Ubuntu 
    
    `sudo mv /mnt/client/startClient.sh /etc/profile.d/`
  - Redhat/CentOS:  
      
      `su - root`
      
      `mv /mnt/client/startClient.sh /etc/profile.d/`