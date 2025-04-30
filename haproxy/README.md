# 測試 HAProxy

## cmd

    預設工作資料夾 /go 、可參考官方映像檔
    此為測試用，並未多做區分
    之後在做新映像檔再加入
    working_dir: /app

``` shell
sudo docker run -v /home/chris/git/learn-docker/haproxy/server:/go --rm -ti --name compile_go golang:1.24-bullseye bash
```
