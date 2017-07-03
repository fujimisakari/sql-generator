# SQL Generator for the Go language

Introduction
------------

SQL test data generation tool


Installation and usage
----------------------

```
$ make setup
$ make build or make install
$ sql-generator schema > ./schema.yaml

$ sql-generator create ./schema.yaml
CREATE TABLE `accounts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `account_name` VARCHAR(20),
  `first_name` VARCHAR(20),
  `last_name` VARCHAR(20),
  `email` VARCHAR(100),
  `password_hash` VARCHAR(32),
  `point` INT,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

$ sql-generator drop ./schema.yaml
DROP TABLE IF EXISTS `accounts`;
```

Example
-------

This example will generate the following output:

```
$ sql-generator example ./schema.yaml
INSERT INTO `accounts` (id, account_name, first_name, last_name, email, password_hash, point, created_at) VALUES
(1, '1-account', 'takahiro', 'maehara', '1-account@gmail.com', 'hogehoge', 481, '2017-05-23 17:35:52'),
(2, '2-account', 'risa', 'gondou', '2-account@gmail.com', 'hogehoge', 1436, '2017-04-13 15:13:44'),
(3, '3-account', 'risa', 'fujimoto', '3-account@gmail.com', 'hogehoge', 434, '2017-04-23 21:09:14'),
(4, '4-account', 'takahiro', 'fujimoto', '4-account@gmail.com', 'hogehoge', 1370, '2017-04-22 14:36:17'),
(5, '5-account', 'misaka', 'uehara', '5-account@gmail.com', 'hogehoge', 1139, '2017-05-22 22:40:43'),
(6, '6-account', 'takezo', 'kondou', '6-account@gmail.com', 'hogehoge', 532, '2017-04-04 00:06:13'),
(7, '7-account', 'misaka', 'morita', '7-account@gmail.com', 'hogehoge', 504, '2017-05-03 10:04:02'),
(8, '8-account', 'misaka', 'fujimoto', '8-account@gmail.com', 'hogehoge', 1210, '2017-05-02 07:22:02'),
(9, '9-account', 'risa', 'fujimoto', '9-account@gmail.com', 'hogehoge', 312, '2017-05-13 20:59:38'),
(10, '10-account', 'risa', 'morita', '10-account@gmail.com', 'hogehoge', 371, '2017-04-20 07:48:47');

:
:
```

How to set TableSchema
-------

You can define these by setting and changing on yaml file.


```
$ sql-generator schema

sql-generator schema

table-schema:
  name: accounts
  params: ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
  ex-number: 5000

  columns:
    - name: id
      type: INT NOT NULL AUTO_INCREMENT
      ex-type: int-inc

    - name: account_name
      type: VARCHAR(20)
      ex-type: string-unique
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
      ex-type: string-unique
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
    - value: PRIMARY KEY (`id`)
    # - value: INDEX `idx_first_name_last_name` (`first_name`, `last_name`)
    # - value: INDEX `idx_point` (`point`)
    # - value: UNIQUE KEY `email` (`email`)
```
