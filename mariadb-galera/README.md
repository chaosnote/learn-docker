# Docker 上運用 mariadb-galera

## 官方文件

[mariadb-galera](https://github.com/bitnami/containers/tree/main/bitnami/mariadb-galera)

## Step 0

step_0.sh

``` shell
#!/bin/bash

## @see https://github.com/bitnami/containers/tree/main/bitnami/mariadb-galera#setting-up-a-multi-master-cluster

sudo docker run -d --rm --name mariadb-galera-0 \
  -e MARIADB_GALERA_CLUSTER_BOOTSTRAP=yes \
  -e MARIADB_GALERA_CLUSTER_NAME=my_galera \
  -e MARIADB_GALERA_MARIABACKUP_USER=my_mariabackup_user \
  -e MARIADB_GALERA_MARIABACKUP_PASSWORD=my_mariabackup_password \
  -e MARIADB_ROOT_PASSWORD=my_root_password \
  -e MARIADB_USER=my_user \
  -e MARIADB_PASSWORD=my_password \
  -e MARIADB_DATABASE=my_database \
  bitnami/mariadb-galera:latest
```

## Step 1

step_1.sh

``` shell
#!/bin/bash

sudo docker run -d --rm --name mariadb-galera-1 --link mariadb-galera-0:mariadb-galera \
  -e MARIADB_GALERA_CLUSTER_NAME=my_galera \
  -e MARIADB_GALERA_CLUSTER_ADDRESS=gcomm://mariadb-galera \
  -e MARIADB_GALERA_MARIABACKUP_USER=my_mariabackup_user \
  -e MARIADB_GALERA_MARIABACKUP_PASSWORD=my_mariabackup_password \
  bitnami/mariadb-galera:latest
```

## Test

test.sh

``` shell
#!/bin/bash
sudo docker exec -it mariadb-galera-0 mariadb -uroot -pmy_root_password -e "SHOW STATUS LIKE 'wsrep_%';"
```

## Stop

stop.sh

``` shell
sudo docker stop mariadb-galera-1
sleep 3s
sudo docker stop mariadb-galera-0
```
