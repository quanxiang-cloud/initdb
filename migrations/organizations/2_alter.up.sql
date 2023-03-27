--  超管数据
INSERT INTO org_user
(id, name, phone, email, self_email, id_card, address, use_status, tenant_id, position, avatar, job_number, gender, source, password_status, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by)
VALUES
    ('1', 'SuperAdmin', '13888886666', 'admin@quanxiang.dev', 'admin@quanxiang.dev', '', '', 1, '', '', '', '', 1, '', 1, 1649429260908, 1649433071, 0, '', '', '');


INSERT INTO org_department
(id, name, use_status, attr, pid, super_pid, grade, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by, tenant_id)
VALUES
    ('1', 'QCC', 1, 1, '', '1', 1, 1641966964, 1649433066, 0, '', '', '', '');


-- 密码 654321a..
INSERT INTO org_user_account
(id, account, user_id, password, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by, tenant_id)
VALUES
    ('1', 'admin@quanxiang.dev', '1', '24d04ec3c9f0e285791035a47ba3e61a', 1635761030, 1649433067, 0, '', '', '', '');




INSERT INTO org_user_department_relation (id, user_id, dep_id, attr)
VALUES ('1', '1', '1', '直属领导');

INSERT INTO message.template (id, name, title, content, created_at, updated_at, create_by, status) VALUES ('org_new_code', '新用户密码', '新用户密码', '您好：</br>欢迎您使用全象云系统，当前密码是 {{.code}}', null, null, null, null);