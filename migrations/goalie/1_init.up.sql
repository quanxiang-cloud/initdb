-- ----------------------------
-- Table structure for func
-- ----------------------------
DROP TABLE IF EXISTS `func`;
CREATE TABLE `func` (
  `id` varchar(36) NOT NULL,
  `p_id` varchar(36) DEFAULT NULL,
  `func_tag` varchar(255) NOT NULL,
  `name` varchar(60) NOT NULL,
  `describe` text,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `nk_func_pid` (`p_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` varchar(36) NOT NULL,
  `name` varchar(30) NOT NULL,
  `tag` varchar(24) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for role_func
-- ----------------------------
DROP TABLE IF EXISTS `role_func`;
CREATE TABLE `role_func` (
  `id` varchar(36) NOT NULL,
  `role_id` varchar(36) NOT NULL,
  `func_id` varchar(36) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_func` (`role_id`,`func_id`),
  KEY `nk_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for role_owner
-- ----------------------------
DROP TABLE IF EXISTS `role_owner`;
CREATE TABLE `role_owner` (
  `id` varchar(36) NOT NULL,
  `role_id` varchar(36) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `owner_id` varchar(36) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_owner` (`role_id`,`owner_id`),
  KEY `nk_role` (`role_id`),
  KEY `nk_owner` (`owner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;

