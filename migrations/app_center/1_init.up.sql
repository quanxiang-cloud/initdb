create table t_app_center
(
    id          varchar(64)  not null
        primary key,
    app_name    varchar(80)  null,
    access_url  varchar(200) null,
    app_icon    text         null,
    create_by   varchar(64)  null,
    update_by   varchar(64)  null,
    create_time bigint       null,
    update_time bigint       null,
    use_status  bigint       null,
    del_flag   TINYINT       null,
    delete_time  BIGINT      null,
    app_sign varchar(30) null,
    extension text null,
    description text null,
    per_poly BOOLEAN default  false,
    constraint t_app_center_app_name_uindex
        unique (app_name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



create table t_app_user_relation
(
    user_id varchar(64) null,
    app_id  varchar(64) null
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table t_app_scope(
    app_id varchar(64) null ,
    scope_id varchar(64) null
    type varchar(64) null
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- auto-generated definition
create table t_app_template
(
    id           varchar(64)  not null primary key,
    name         varchar(80)  null comment 'template name',
    app_icon     text         null comment 'app icon',
    path         varchar(200) null comment 'file server path',
    source_id    varchar(64)  null comment 'source app id',
    source_name  varchar(80)  null comment 'source app name',
    version      varchar(64)  null comment 'template version',
    group_id     varchar(64)  null comment 'group id',
    created_by   varchar(64)  null,
    created_name varchar(64)  null,
    created_time bigint       null,
    updated_by   varchar(64)  null,
    updated_name varchar(64)  null,
    updated_time bigint       null,
    status       int          null comment 'publish status：0:private 1:public',
    constraint t_app_template_id_uindex
        unique (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment 'app template table';


ALTER TABLE t_app_center ADD COLUMN server INT COMMENT 'initialized modules of app' AFTER use_status;
DROP INDEX t_app_center_app_name_uindex ON t_app_center;