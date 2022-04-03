# GoFrame Project

lffwl个人博客

GoFrame v1.16.5

sql 文件在 databese 文件夹

部分Nginx 配置，文件和端口转发


`

    location ^~ /file {
        access_log off;
        expires    1d;
        root       /to/path;
    }
    
    location / {
      # 将客户端的 Host 和 IP 信息一并转发到对应节点  
      proxy_http_version 1.1; 
      proxy_set_header Host $http_host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_pass  http://127.0.0.1:8199; # 转发地址
    }

`