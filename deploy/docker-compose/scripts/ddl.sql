CREATE DATABASE go_gin_web CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


USE go_gin_web;


CREATE TABLE sys_tenants (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',

                           tenant_name VARCHAR(20) NOT NULL COMMENT '租户名称'

) COMMENT '租户表';


CREATE TABLE sys_orgs (
                             id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                             created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                             updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',
                             tenant_id INT UNSIGNED NOT NULL COMMENT '租户ID',

                             org_name VARCHAR(20) NOT NULL COMMENT '机构名称'
) COMMENT '机构表';


CREATE TABLE sys_users (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',

                           user_name VARCHAR(20) NOT NULL COMMENT '用户名',
                           login_name VARCHAR(20) NOT NULL COMMENT '用户名',
                           password VARCHAR(16) NOT NULL COMMENT '密码',
                           mobile VARCHAR(11) NULL COMMENT '手机号码',
                           email VARCHAR(30) NULL COMMENT '邮箱',
                           enabled TINYINT(1) NOT NULL DEFAULT 1  COMMENT '用户状态 0-禁用，1-启用'

) COMMENT '用户表';


CREATE TABLE sys_user_tenants (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',

                           user_id INT UNSIGNED NOT NULL COMMENT '机构ID',
                           tenant_id INT UNSIGNED NOT NULL COMMENT '租户Id'
) COMMENT '用户租户表';


CREATE TABLE sys_user_orgs (
                                  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                                  created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                                  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',

                                  user_id INT UNSIGNED NOT NULL COMMENT '机构ID',
                                  org_id INT UNSIGNED NOT NULL COMMENT '机构Id'
) COMMENT '用户机构表';


CREATE TABLE sys_roles (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',,
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',
                           tenant_id INT UNSIGNED NOT NULL COMMENT '租户ID',
                           org_id INT UNSIGNED NOT NULL COMMENT '机构ID',

                           role_name VARCHAR(20) NOT NULL COMMENT '角色名称'
) COMMENT '角色表';


CREATE TABLE sys_user_roles (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',,
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',
                           tenant_id INT UNSIGNED NOT NULL COMMENT '租户ID',
                           org_id INT UNSIGNED NOT NULL COMMENT '机构ID',

                           user_id  INT UNSIGNED NOT NULL COMMENT '用户ID',
                           role_id  INT UNSIGNED NOT NULL COMMENT '角色ID'

) COMMENT '用户角色表';


CREATE TABLE sys_authorities  (
                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',,
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',
                           tenant_id INT UNSIGNED NOT NULL COMMENT '租户ID',
                           org_id INT UNSIGNED NOT NULL COMMENT '机构ID',

                           authority VARCHAR(20) NOT NULL COMMENT '权限名称'
) COMMENT '权限表';


CREATE TABLE sys_role_authorities  (
                                  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',,
                                  created_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '创建时间',
                                  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIME()  COMMENT '更新时间',
                                  tenant_id INT UNSIGNED NOT NULL COMMENT '租户ID',
                                  org_id INT UNSIGNED NOT NULL COMMENT '机构ID',

                                  authority_id  INT UNSIGNED NOT NULL COMMENT '权限ID',
                                  role_id  INT UNSIGNED NOT NULL COMMENT '角色ID'
) COMMENT '角色权限表';