ALTER TABLE flow  ADD  app_status varchar(20) NOT NULL DEFAULT 'ACTIVE' COMMENT 'app状态' after app_id;
ALTER TABLE flow_instance  ADD  app_status varchar(20) NOT NULL DEFAULT 'ACTIVE' COMMENT 'app状态' after app_id;

alter table flow add cron varchar(20) null;

alter table flow_variables
    modify code varchar(200) default '' not null comment '变量标识';

alter table flow_instance_variables
    modify code varchar(200) default '' not null comment '变量标识';

alter table flow_instance add request_id varchar(200) null;