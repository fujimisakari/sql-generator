package main

import (
	"fmt"
)

var template = `
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
    - value: PRIMARY KEY (%s)
    # - value: INDEX %s (%s, %s)
    # - value: INDEX %s (%s)
    # - value: UNIQUE KEY %s (%s)
`

var sampleSchema = fmt.Sprintf(template, "`id`",
	"`idx_first_name_last_name`", "`first_name`", "`last_name`",
	"`idx_point`", "`point`",
	"`email`", "`email`")
