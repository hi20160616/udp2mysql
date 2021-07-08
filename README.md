# udp2mysql
Listen UDP and save packet to mysql or mariadb

# gRPC

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/udp2mysql/v1/udp2mysql.proto
```

# DB

## MariaDB

Refer: [mariadb docker official images](!https://hub.docker.com/_/mariadb)
```bash
$ docker pull mariadb
docker run -p 3306:3306 --name=udp2mysql -v /my/own/datadir:/var/lib/mysql -e MYSQL_ROOT_PASSWORD='<password>' -d mariadb:latest
$ docker exec -it udp2mysql mysql -uroot -p'<password>'
```

Create Database  
Refer: https://dev.mysql.com/doc/refman/8.0/en/entering-queries.html
```bash
MariaDB [(none)]> CREATE database udp2mysql;
Query OK, 1 row affected (0.000 sec)

MariaDB [(none)]> use udp2mysql;
Database changed
```
Create Tables
```bash
MariaDB [udp2mysql]> DROP TABLE udp_packets;
Query OK, 0 rows affected (0.003 sec)

MariaDB [udp2mysql]> CREATE TABLE udp_packets (id VARCHAR(32) NOT NULL, name VARCHAR(32), title VARCHAR(255), content TEXT(65535), update_time BIGINT(19), UNIQUE KEY (id));
Query OK, 0 rows affected (0.005 sec)

MariaDB [udp2mysql]> DESC udp_packets;
+-------------+--------------+------+-----+---------+-------+
| Field       | Type         | Null | Key | Default | Extra |
+-------------+--------------+------+-----+---------+-------+
| id          | varchar(32)  | NO   | PRI | NULL    |       |
| name        | varchar(32)  | YES  |     | NULL    |       |
| title       | varchar(255) | YES  |     | NULL    |       |
| content     | mediumtext   | YES  |     | NULL    |       |
| update_time | bigint(19)   | YES  |     | NULL    |       |
+-------------+--------------+------+-----+---------+-------+
5 rows in set (0.001 sec)
```
Create User for the database  
Refer: 
- https://dev.mysql.com/doc/refman/8.0/en/create-user.html
- https://dev.mysql.com/doc/refman/8.0/en/grant.html#grant-database-privileges
```bash
MariaDB [udp2mysql]> CREATE USER 'udp2mysql'@'%' IDENTIFIED BY 'yourpassword';
Query OK, 0 rows affected (0.001 sec)

MariaDB [udp2mysql]> GRANT ALL ON udp2mysql.* TO 'udp2mysql'@'%';
Query OK, 0 rows affected (0.001 sec)
```
Creating database dumps
```
$ docker exec udp2mysql sh -c 'exec mysqldump --all-databases -uroot -p"$MARIADB_ROOT_PASSWORD"' > ./all-databases.sql
```
Restoring data from dump files
```
$ docker exec -i udp2mysql sh -c 'exec mysql -uroot -p"$MARIADB_ROOT_PASSWORD"' < ./all-databases.sql
```
