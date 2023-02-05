# About

うちのAPとWiFiルータを再起動するもの

ぽちぽちログインして再起動する手動作業が手間だったので、playwright化しました

## Environment
- GW/Router: ZTE/ZXHN F660A
- Wifi-AP: Netgear/Orbi RBK50-100JPS

## needs
- playwright-go
- ~task(taskfile)~

## usage

envファイルに認証情報を設定
```
cp ./env.template ./env
vi ./env
```

実行
```
go run main
```

ログ
````
go run main

2023/02/05 23:52:28
 f660aAuth: {hoge fuga  192.168.1.1}
 orbiAuth: {piyo moge 192.168.2.1}
2023/02/05 23:52:30 Downloading browsers...
2023/02/05 23:52:30 Downloaded browsers successfully
2023/02/05 23:52:31 START to reboot Orbi
2023/02/05 23:52:31 STEP: goto 192.168.2.1
2023/02/05 23:52:32 STEP: Confirm Reboot
2023/02/05 23:52:32 restart!
2023/02/05 23:52:34 START to reboot f660a
2023/02/05 23:52:34 STEP: LOGIN goto http://192.168.1.1/
2023/02/05 23:52:35 STEP: LOGIN with authInfo
2023/02/05 23:52:36 STEP: GOTO KANRI iframe page http://192.168.1.1/template.gch?pid=1002&nextpage=manager_dev_conf_t.gch
2023/02/05 23:52:37 STEP: REBOOT Reboot/Submit
2023/02/05 23:52:37 restart!
``
