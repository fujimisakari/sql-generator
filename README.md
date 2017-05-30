# SQL Generator for the Go language

Introduction
------------

SQL test data generation tool


Installation and usage
----------------------

```
$ make setup

$ make create
go run output.go const.go model.go main.go ./schema.yaml create
CREATE TABLE `accounts` (
  `account_id` SERIAL PRIMARY KEY,
  `account_name` VARCHAR(20),
  `first_name` VARCHAR(20),
  `last_name` VARCHAR(20),
  `email` VARCHAR(100),
  `password_hash` CHAR(64),
  `point` INTEGER,
  `created_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

$ make drop
go run output.go const.go model.go main.go ./schema.yaml drop
DROP TABLE IF EXISTS `accounts`;
```

Example
-------

This example will generate the following output:

```
$ make example
go run output.go const.go model.go main.go ./schema.yaml example
 make example
go run output.go const.go model.go main.go ./schema.yaml example
INSERT INTO `accounts` (account_id, account_name, first_name, last_name, email, password_hash, point, created_at) VALUES
(1, '1-account', 'risa', 'morita', '1-account@gmail.com', 'hogehoge', 498, '2017-05-21 20:22:58'),
(2, '2-account', 'ryo', 'gondou', '2-account@gmail.com', 'hogehoge', 1034, '2017-05-21 20:12:16'),
(3, '3-account', 'ryo', 'uehara', '3-account@gmail.com', 'hogehoge', 1084, '2017-05-21 20:08:32'),
(4, '4-account', 'ryo', 'maehara', '4-account@gmail.com', 'hogehoge', 1102, '2017-05-23 01:59:37'),
(5, '5-account', 'takezo', 'morita', '5-account@gmail.com', 'hogehoge', 871, '2017-05-21 05:11:20'),
(6, '6-account', 'takezo', 'fujimoto', '6-account@gmail.com', 'hogehoge', 1134, '2017-05-21 11:37:22'),
(7, '7-account', 'takezo', 'kondou', '7-account@gmail.com', 'hogehoge', 1167, '2017-05-22 22:48:09'),
(8, '8-account', 'ryo', 'gondou', '8-account@gmail.com', 'hogehoge', 793, '2017-05-22 23:39:08'),
(9, '9-account', 'risa', 'maehara', '9-account@gmail.com', 'hogehoge', 1042, '2017-05-21 23:57:13'),
:
:
```
