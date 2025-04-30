# 使用 Goalng 官方映像檔

後端編譯

## 下載映像檔

    sudo docker pull golang:1.24-bullseye

## 首次編譯

額外設置 -ti，進入後運行 sh compile.sh

    sudo docker run -v 專案位置:/source -v 輸出位置:/dist -w /source --rm -ti --name compile_go golang:1.24-bullseye bash

產生映像檔，節省下次 module 下載的時間

    sudo docker commit -m "after update golang module" -a "Ur name" Container_ID TAG
    參數說明
        -m 修改記錄
        -a (修改者/作者)名稱

## 之後編譯代碼

    sudo docker run -v 專案位置:/source -v 輸出位置:/dist -w /source -ti --rm --name compile_go TAG sh compile.sh
