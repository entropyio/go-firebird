/* symbol_info */

drop table if exists symbol_info;
create table symbol_info
(
  id           INTEGER     not null primary key autoincrement,
  gmt_create   DATETIME    not null,
  gmt_modified DATETIME    not null,
  symbol_name  VARCHAR(20) not null,
  symbol_desc  VARCHAR(32) not null,
  status       int         not null,
  symbol_icon  VARCHAR(10),
  symbol_group VARCHAR(10)
);

drop index if exists symbol_info__symbol;
create unique index symbol_info__symbol
  on symbol_info (symbol_name);

drop index if exists symbol_info__time;
create index symbol_info__time
  on symbol_info (gmt_create);

INSERT INTO symbol_info (id, gmt_create, gmt_modified, symbol_name, symbol_desc, status, symbol_icon, symbol_group)
VALUES (1, '1576581827473', '1576581830743', 'eosusdt', 'EOS/USDT', 1, 'eos', 'EOS');
INSERT INTO symbol_info (id, gmt_create, gmt_modified, symbol_name, symbol_desc, status, symbol_icon, symbol_group)
VALUES (2, '1576582464918', '1576582466902', 'btcusdt', 'BTC/USDT', 1, 'btc', 'BTC');


/* user_account */

drop table if exists user_account;
create table user_account
(
  id            INTEGER  not null primary key autoincrement,
  gmt_create    DATETIME not null,
  gmt_modified  DATETIME not null,
  user_id       INTEGER  not null,
  symbol_id     INTEGER  not null,
  hold_price    double                        default 0,
  hold_amount   double                        default 0,
  yest_benefit  double                        default 0,
  total_benefit double                        default 0,
  price         double                        default 0,
  amount        double                        default 0,
  total         double                        default 0,
  benefit       double                        default 0,
  rate          double                        default 0,
  status        int      not null
);

drop index if exists user_account__time;
create index user_account__time
  on user_account (gmt_create);

drop index if exists user_account__user;
create unique index user_account__user
  on user_account (user_id, symbol_id);

INSERT INTO user_account (id, gmt_create, gmt_modified, user_id, symbol_id, hold_price, hold_amount, yest_benefit, total_benefit, price, amount, total, benefit, rate, status)
VALUES (1, '1578028057644', '1578028057644', 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1);
INSERT INTO user_account (id, gmt_create, gmt_modified, user_id, symbol_id, hold_price, hold_amount, yest_benefit, total_benefit, price, amount, total, benefit, rate, status)
VALUES (2, '1578034351472', '1578034351472', 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1);


/* user_info */

drop table if exists user_info;
create table user_info
(
  id           INTEGER     not null primary key autoincrement,
  gmt_create   DATETIME    not null,
  gmt_modified DATETIME    not null,
  user_name    VARCHAR(20) not null,
  user_desc    VARCHAR(32) not null,
  status       int         not null
);

drop index if exists user_info__name;
create unique index user_info__name
  on user_info (user_name);

drop index if exists user_info__time;
create index user_info__time
  on user_info (gmt_create);

INSERT INTO user_info (id, gmt_create, gmt_modified, user_name, user_desc, status)
VALUES (1, '1576582464918', '1576582464918', 'star', 'star', 1);
INSERT INTO user_info (id, gmt_create, gmt_modified, user_name, user_desc, status)
VALUES (2, '1576582464918', '1576582464918', 'eric', 'eric', 1);


/* user_trade */
drop table if exists user_trade;
create table user_trade
(
  id           INTEGER  not null primary key autoincrement,
  gmt_create   DATETIME not null,
  gmt_modified DATETIME not null,
  user_id      INTEGER  not null,
  symbol_id    INTEGER  not null,
  price        double   not null,
  amount       double   not null,
  status       int      not null,
  hold_price   double                        default 0 not null,
  hold_amount  double                        default 0 not null,
  type         int                           default 0 not null,
  schedule_id  INTEGER                       default 0 not null,
  reason       VARCHAR(100) default ""
);

drop index if exists user_trade__symbol;
create index user_trade__symbol
  on user_trade (symbol_id);

drop index if exists user_trade__time;
create index user_trade__time
  on user_trade (gmt_create);

drop index if exists user_trade__user;
create index user_trade__user
  on user_trade (user_id);

drop table if exists user_schedule;
create table user_schedule
(
  id           INTEGER     not null primary key autoincrement,
  gmt_create   DATETIME    not null,
  gmt_modified DATETIME    not null,
  user_id      INTEGER     not null,
  symbol_id    INTEGER     not null,
  name         VARCHAR(50) not null,
  type         int    default 0 not null,
  amount       double default 0,
  success      int                              default 0,
  failed       int                              default 0,
  status       int         not null
);

drop index if exists user_schedule__symbol;
create index user_schedule__symbol
  on user_schedule (symbol_id);

drop index if exists user_schedule__user;
create index user_schedule__user
  on user_schedule (user_id);

-- auto-generated definition
drop table if exists rule_item;
create table rule_item
(
  id           INTEGER      not null primary key autoincrement,
  gmt_create   DATETIME     not null,
  gmt_modified DATETIME     not null,
  user_id      INTEGER      not null,
  symbol_id    INTEGER      not null,
  schedule_id  INTEGER      not null,
  rule_type    int          not null,
  join_type    int          not null,
  op_type      int          not null,
  value        VARCHAR(255) NOT NULL,
  status       int          not null
);

drop index if exists rule_item__symbol;
create index rule_item__symbol
  on rule_item (symbol_id);

drop index if exists rule_item__user;
create index rule_item__user
  on rule_item (user_id);

drop index if exists rule_item__schedule;
create index rule_item__schedule
  on rule_item (schedule_id);

-- auto-generated definition
drop table if exists user_data;
create table user_data
(
  id           INTEGER  not null primary key autoincrement,
  gmt_create   DATETIME not null,
  gmt_modified DATETIME not null,
  user_id      INTEGER  not null,
  symbol_id    INTEGER  not null,
  open_price   double default 0,
  close_price  double default 0,
  high_price   double default 0,
  low_price    double default 0,
  hold_price   double default 0,
  hold_amount  double default 0,
  hold_rate    double default 0,
  hold_benefit double default 0,
  status       int      default 0 not null
);

drop index if exists user_data__time;
create unique index user_data__time
  on user_data (gmt_create, user_id, symbol_id);

drop index if exists user_data__symbol;
create index user_data__symbol
  on user_data (symbol_id);

drop index if exists user_data__user;
create index user_data__user
  on user_data (user_id);
