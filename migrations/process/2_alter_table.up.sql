ALTER TABLE proc_instance  ADD  app_status varchar(20) NOT NULL DEFAULT 'ACTIVE' COMMENT 'app状态' after STATUS;