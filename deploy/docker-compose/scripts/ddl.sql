CREATE DATABASE go_gin_web CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE go_gin_web;

CREATE TABLE sys_users (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                           created_at DATETIME NOT NULL   COMMENT '创建时间',
                           updated_at DATETIME NOT NULL   COMMENT '更新时间',
                           user_name VARCHAR(20) NOT NULL COMMENT '用户名',
                           password VARCHAR(16) NOT NULL COMMENT '密码',
                           mobile VARCHAR(11) NOT NULL COMMENT '手机号码',
                           email VARCHAR(30) NULL COMMENT '邮箱',
                           status TINYINT(1) NOT NULL DEFAULT 1  COMMENT '用户状态',
                           birthday DATE NULL COMMENT '生日',
                           tenant_id INT UNSIGNED NOT NULL COMMENT '租户ID'
);