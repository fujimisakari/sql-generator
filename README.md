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
  `point` INTEGER
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
INSERT INTO `accounts` (account_id, account_name, first_name, last_name, email, password_hash, point) VALUES
(1, '1-account', 'risa', 'maehara', '1-account@gmail.com', 'hogehoge', 951),
(2, '2-account', 'misaka', 'fujimoto', '2-account@gmail.com', 'hogehoge', 710),
(3, '3-account', 'misaka', 'kondou', '3-account@gmail.com', 'hogehoge', 1232),
(4, '4-account', 'takahiro', 'uehara', '4-account@gmail.com', 'hogehoge', 1471),
(5, '5-account', 'takezo', 'gondou', '5-account@gmail.com', 'hogehoge', 1212),
(6, '6-account', 'misaka', 'morita', '6-account@gmail.com', 'hogehoge', 385),
(7, '7-account', 'risa', 'kondou', '7-account@gmail.com', 'hogehoge', 638),
(8, '8-account', 'jun', 'morita', '8-account@gmail.com', 'hogehoge', 885),
(9, '9-account', 'misaka', 'gondou', '9-account@gmail.com', 'hogehoge', 762),
:
:
```
