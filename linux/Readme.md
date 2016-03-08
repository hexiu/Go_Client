# Linux 客户端使用要求

- 在/mnt下创建目录client(`mkdir /mnt/client ; cd /mnt/client`)
- 将[Linux 客户端](https://github.com/hexiu/Go_Client/linux/main "Linux 客户端")，[Linux 启动脚本](https://github.com/hexiu/Go_Client/linux/startClient.sh "Linux 启动脚本") 下载至本机，放在 /mnt/client 目录下
  
  ```
  wget https://github.com/hexiu/Go_Client/linux/main
  wget  https://github.com/hexiu/Go_Client/linux/startClient.sh
  ```
- 下载 startClient.sh 脚本放在/etc/profile.d 目录下
  - Ubuntu 
    
    `sudo mv /mnt/client/startClient.sh /etc/profile.d/`
  - Redhat/CentOS:  
      
      `su - root`
      
      `mv /mnt/client/startClient.sh /etc/profile.d/`