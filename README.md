# SQL Generator for the Go language

Introduction
------------

SQL test data generation tool


Installation and usage
----------------------

```
$ make setup
$ make build
$ ./sql-generator schema > ./schema.yaml

$ ./sql-generator create ./schema.yaml
CREATE TABLE `accounts` (
  `account_id` INT NOT NULL AUTO_INCREMENT,
  `account_name` VARCHAR(20),
  `first_name` VARCHAR(20),
  `last_name` VARCHAR(20),
  `email` VARCHAR(100),
  `password_hash` VARCHAR(32),
  `point` INT,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

$ ./sql-generator drop ./schema.yaml
DROP TABLE IF EXISTS `accounts`;
```

Example
-------

This example will generate the following output:

```
$ ./sql-generator example ./schema.yaml
INSERT INTO `accounts` (account_id, account_name, first_name, last_name, email, password_hash, point, created_at) VALUES
(1, '1-account', 'misaka', 'maehara', '1-account@gmail.com', 'hogehoge', 1058, '2017-04-10 06:53:45'),
(2, '2-account', 'takahiro', 'kondou', '2-account@gmail.com', 'hogehoge', 1222, '2017-04-21 18:03:58'),
(3, '3-account', 'risa', 'fujimoto', '3-account@gmail.com', 'hogehoge', 733, '2017-05-18 14:50:01'),
(4, '4-account', 'misaka', 'uehara', '4-account@gmail.com', 'hogehoge', 805, '2017-05-06 06:31:50'),
(5, '5-account', 'risa', 'morita', '5-account@gmail.com', 'hogehoge', 1000, '2017-05-02 00:32:09'),
(6, '6-account', 'misaka', 'maehara', '6-account@gmail.com', 'hogehoge', 1258, '2017-04-10 03:23:49'),
(7, '7-account', 'ryo', 'maehara', '7-account@gmail.com', 'hogehoge', 500, '2017-05-22 23:16:12'),
(8, '8-account', 'ryo', 'gondou', '8-account@gmail.com', 'hogehoge', 359, '2017-05-18 11:24:20'),
(9, '9-account', 'jun', 'fujimoto', '9-account@gmail.com', 'hogehoge', 1184, '2017-04-05 17:39:11'),
(10, '10-account', 'jun', 'uehara', '10-account@gmail.com', 'hogehoge', 925, '2017-05-20 14:25:51'),
:
:
```

How to set TableSchema
-------

You can define these by setting and changing on yaml file.


```
$ ./sql-generator schema

table-schema:
  name: accounts
  params: ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
  ex-number: 5000

  columns:
    - name: account_id
      type: INT NOT NULL AUTO_INCREMENT
      ex-type: int-inc

    - name: account_name
      type: VARCHAR(20)
      ex-type: string-inc
      ex-text: "account"

    - name: first_name
      type: VARCHAR(20)
      ex-type: string-range
      ex-range:
        - value: "ryo"
        - value: "takezo"
        - value: "risa"
        - value: "misaka"
        - value: "jun"
        - value: "takahiro"

    - name: last_name
      type: VARCHAR(20)
      ex-type: string-range
      ex-range:
        - value: "fujimoto"
        - value: "gondou"
        - value: "uehara"
        - value: "kondou"
        - value: "maehara"
        - value: "morita"

    - name: email
      type: VARCHAR(100)
      ex-type: string-inc
      ex-text: "account@gmail.com"

    - name: password_hash
      type: VARCHAR(32)
      ex-type: string
      ex-text: "hogehoge"

    - name: point
      type: INT
      ex-type: int-range
      ex-range:
        - min: 300
        - max: 1500

    - name: created_at
      type: datetime DEFAULT NULL
      ex-type: date-range
      ex-range:
        - start: "2017-04-01 12:31:24"
        - end: "2017-05-23 23:01:55"

  meta-list:
    - value: PRIMARY KEY (`account_id`)
    # - value: INDEX `idx_first_name_last_name` (`first_name`, `last_name`)
    # - value: INDEX `idx_point` (`point`)
    # - value: UNIQUE KEY `email` (`email`)
```
