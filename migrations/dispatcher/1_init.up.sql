DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` varchar(36) NOT NULL,
  `title` varchar(120) DEFAULT NULL,
  `describe` text,
  `type` tinyint(1) NOT NULL,
  `time_bar` varchar(24) DEFAULT NULL,
  `state` tinyint(1) NOT NULL,
  `code` varchar(55) NOT NULL,
  `retry` int(2) DEFAULT NULL,
  `retry_delay` int(6) DEFAULT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;