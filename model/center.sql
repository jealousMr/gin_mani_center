create table rule(
 id bigint(20) auto_increment,
 rule_id varchar(32) not null ,
 user varchar(32) not null,
 rule_type tinyint(4) unsigned not null,
 rule_state tinyint(4) unsigned not null,
 execute_type tinyint(4) unsigned not null,
 desc_text text default null,
 source_name varchar(128) default "",
 source_url varchar(256) default "",
 source_state tinyint(4) unsigned not null,
 create_at timestamp default CURRENT_TIMESTAMP,
 update_at timestamp default CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 UNIQUE (`rule_id`),
 UNIQUE (`rule_id`,`user`,`rule_type`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;