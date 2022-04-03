/*
 Navicat Premium Data Transfer

 Source Server         : www.lffwl.com
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : 127.0.0.1:3306
 Source Schema         : www.lffwl.com

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 03/04/2022 18:48:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for www_admin
-- ----------------------------
DROP TABLE IF EXISTS `www_admin`;
CREATE TABLE `www_admin`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '明文密码',
  `hash_password` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'hash密码',
  `created_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  `updated_id` int(11) NOT NULL DEFAULT 0 COMMENT '最近更新人',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '最近更新时间',
  `deleted_id` int(11) NOT NULL DEFAULT 0 COMMENT '软删除人',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `www_admin_user_name`(`user_name`) USING BTREE,
  INDEX `www_admin_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of www_admin
-- ----------------------------
INSERT INTO `www_admin` VALUES (1, 'admin', '123456', '$2a$10$Ro5PR9mLfcCFTQ30sW6qVOCXpIN1gRvKNJXARTWLGMztJ7bhjbrOS', 0, '2021-08-07 18:03:19', 1, '2021-08-27 18:01:16', 0, NULL);

-- ----------------------------
-- Table structure for www_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `www_admin_role`;
CREATE TABLE `www_admin_role`  (
  `admin_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY (`admin_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of www_admin_role
-- ----------------------------
INSERT INTO `www_admin_role` VALUES (1, 1);

-- ----------------------------
-- Table structure for www_api
-- ----------------------------
DROP TABLE IF EXISTS `www_api`;
CREATE TABLE `www_api`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `menu` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单地址',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `key` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标识',
  `router` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求地址',
  `methods` tinyint(4) NULL DEFAULT NULL COMMENT '请求类型',
  `type` tinyint(4) NULL DEFAULT 0 COMMENT '菜单类型',
  `pid` int(11) NULL DEFAULT 0 COMMENT '上级ID',
  `created_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  `updated_id` int(11) NOT NULL DEFAULT 0 COMMENT '最近更新人',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '最近更新时间',
  `deleted_id` int(11) NOT NULL DEFAULT 0 COMMENT '软删除人',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `www_api_router`(`router`) USING BTREE,
  INDEX `www_api_methods`(`methods`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of www_api
-- ----------------------------
INSERT INTO `www_api` VALUES (15, '/page/api/index', '权限列表', 'manage.auth.api', '/api', 1, 0, 16, 1, '2021-08-19 08:24:35', 1, '2021-08-25 14:42:30', 0, NULL);
INSERT INTO `www_api` VALUES (16, '', '授权管理', 'manage.auth', '', 1, 0, 0, 0, NULL, 1, '2021-08-25 15:01:14', 0, NULL);
INSERT INTO `www_api` VALUES (17, NULL, '新增权限', 'manage.auth.api.store', '/api', 2, 1, 15, 0, NULL, 1, '2021-08-25 14:42:36', 0, NULL);
INSERT INTO `www_api` VALUES (18, '', '编辑权限', 'manage.auth.api.update', '/api', 3, 1, 15, 1, '2021-08-25 10:38:42', 1, '2021-08-25 15:04:51', 0, NULL);
INSERT INTO `www_api` VALUES (19, '', '权限删除', 'manage.auth.api.destroy', '/api', 4, 1, 15, 1, '2021-08-25 10:48:17', 1, '2021-08-25 15:04:45', 0, NULL);
INSERT INTO `www_api` VALUES (20, '/page/role/index', '角色管理', 'manage.auth.role', '/role', 1, 0, 16, 1, '2021-08-25 14:43:43', 1, '2021-08-25 14:52:31', 0, NULL);
INSERT INTO `www_api` VALUES (21, '', '新增角色', 'manage.auth.role.store', '/role', 2, 1, 20, 1, '2021-08-25 14:46:51', 1, '2021-08-25 15:04:41', 0, NULL);
INSERT INTO `www_api` VALUES (22, '', '编辑角色', 'manage.auth.role.update', '/role', 3, 1, 20, 1, '2021-08-25 14:54:29', 1, '2021-08-25 15:04:38', 0, NULL);
INSERT INTO `www_api` VALUES (23, '', '删除角色', 'manage.auth.role.destroy', '/role', 4, 1, 20, 1, '2021-08-25 14:54:57', 1, '2021-08-25 15:04:32', 0, NULL);
INSERT INTO `www_api` VALUES (24, '', '设置权限', 'manage.auth.role.set-api', '', 1, 1, 20, 1, '2021-08-25 15:30:21', 0, '2021-08-25 15:30:21', 0, NULL);
INSERT INTO `www_api` VALUES (25, '/page/admin/index', '管理员管理', 'manage.auth.admin', '/admin', 1, 0, 16, 1, '2021-08-25 16:50:43', 0, '2021-08-25 16:50:43', 0, NULL);
INSERT INTO `www_api` VALUES (26, '', '新增管理员', 'manage.auth.admin.store', '/admin', 2, 1, 25, 1, '2021-08-25 17:08:23', 0, '2021-08-25 17:08:23', 0, NULL);
INSERT INTO `www_api` VALUES (27, '', '编辑管理员', 'manage.auth.admin.update', '/admin', 3, 1, 25, 1, '2021-08-25 17:08:50', 0, '2021-08-25 17:08:50', 0, NULL);
INSERT INTO `www_api` VALUES (28, '', '删除管理员', 'manage.auth.admin.destroy', '/admin', 4, 1, 25, 1, '2021-08-25 17:09:22', 0, '2021-08-25 17:09:22', 0, NULL);
INSERT INTO `www_api` VALUES (29, '', '设置角色', 'manage.auth.admin.set-role', '', 2, 1, 25, 1, '2021-08-26 09:55:56', 0, '2021-08-26 09:55:56', 0, NULL);
INSERT INTO `www_api` VALUES (30, '', '内容管理', 'manage.content', '', 1, 0, 0, 1, '2021-08-26 11:24:14', 1, '2021-08-26 11:26:24', 0, NULL);
INSERT INTO `www_api` VALUES (31, '/page/blog-type/index', '博客类型管理', 'manage.content.blog-type', '/blog-type', 1, 0, 30, 1, '2021-08-26 11:34:26', 1, '2021-08-26 14:11:47', 0, NULL);
INSERT INTO `www_api` VALUES (32, '', '新增博客类型', 'manage.content.blog-type.store', '/manage/blog-type', 2, 1, 31, 1, '2021-08-26 11:42:35', 0, '2021-08-26 11:42:35', 0, NULL);
INSERT INTO `www_api` VALUES (33, '', '编辑博客类型', 'manage.content.blog-type.update', '/blog-type', 3, 1, 31, 1, '2021-08-26 11:44:05', 0, '2021-08-26 11:44:05', 0, NULL);
INSERT INTO `www_api` VALUES (34, '', '删除博客类型', 'manage.content.blog-type.destroy', '/blog-type', 4, 1, 31, 1, '2021-08-26 11:44:42', 0, '2021-08-26 11:44:42', 0, NULL);
INSERT INTO `www_api` VALUES (35, '/page/blog/index', '博客管理', 'manage.content.blog', '/blog', 1, 0, 30, 1, '2021-08-26 14:07:37', 0, '2021-08-26 14:07:37', 0, NULL);
INSERT INTO `www_api` VALUES (36, '', '新增博客', 'manage.content.blog.store', '/blog', 2, 1, 35, 1, '2021-08-26 14:08:47', 0, '2021-08-26 14:08:47', 0, NULL);
INSERT INTO `www_api` VALUES (37, '', '编辑博客', 'manage.content.blog.update', '/blog', 3, 1, 35, 1, '2021-08-26 14:09:18', 0, '2021-08-26 14:09:18', 0, NULL);
INSERT INTO `www_api` VALUES (38, '', '删除博客', 'manage.content.blog.destroy', '/blog', 4, 1, 35, 1, '2021-08-26 14:09:52', 0, '2021-08-26 14:09:52', 0, NULL);

-- ----------------------------
-- Table structure for www_blog
-- ----------------------------
DROP TABLE IF EXISTS `www_blog`;
CREATE TABLE `www_blog`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type_id` int(11) NOT NULL COMMENT '博客类型ID',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `seo_keyword` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'seo关键字',
  `seo_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'seo描叙',
  `created_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  `updated_id` int(11) NOT NULL DEFAULT 0 COMMENT '最近更新人',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '最近更新时间',
  `deleted_id` int(11) NOT NULL DEFAULT 0 COMMENT '软删除人',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `www_blog_type_id`(`type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '博客内容表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for www_blog_type
-- ----------------------------
DROP TABLE IF EXISTS `www_blog_type`;
CREATE TABLE `www_blog_type`  (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `seo_keyword` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'seo关键字',
  `seo_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'seo描叙',
  `created_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  `updated_id` int(11) NOT NULL DEFAULT 0 COMMENT '最近更新人',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '最近更新时间',
  `deleted_id` int(11) NOT NULL DEFAULT 0 COMMENT '软删除人',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '博客类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of www_blog_type
-- ----------------------------
INSERT INTO `www_blog_type` VALUES (1, 'GO', 'Go,GoLang,Go笔记', 'Go语言的开发过程中的一些问题，Go开发环境配置等', 0, '2021-08-13 16:01:26', 3, '2021-08-14 14:55:27', 0, NULL);
INSERT INTO `www_blog_type` VALUES (2, 'PHP', '', '', 3, '2021-08-14 14:54:51', 1, '2021-08-26 14:05:41', 0, NULL);
INSERT INTO `www_blog_type` VALUES (3, 'JavaScript', '', '', 1, '2021-08-26 13:58:10', 0, '2021-08-26 13:58:10', 1, '2021-08-28 14:44:21');
INSERT INTO `www_blog_type` VALUES (4, 'Nginx', '', '', 1, '2021-08-28 14:27:51', 0, '2021-08-28 14:27:51', 0, NULL);
INSERT INTO `www_blog_type` VALUES (5, 'Other', '', '', 1, '2021-08-28 16:54:59', 1, '2021-08-29 12:15:05', 0, NULL);
INSERT INTO `www_blog_type` VALUES (6, 'JavaScript', '', '', 1, '2021-08-29 12:14:38', 0, '2021-08-29 12:14:38', 1, '2021-08-29 12:14:54');
INSERT INTO `www_blog_type` VALUES (7, 'Linux', '', '', 1, '2022-03-01 16:07:48', 0, '2022-03-01 16:07:48', 0, NULL);

-- ----------------------------
-- Table structure for www_role
-- ----------------------------
DROP TABLE IF EXISTS `www_role`;
CREATE TABLE `www_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `created_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  `updated_id` int(11) NOT NULL DEFAULT 0 COMMENT '最近更新人',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '最近更新时间',
  `deleted_id` int(11) NOT NULL DEFAULT 0 COMMENT '软删除人',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `www_role_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of www_role
-- ----------------------------
INSERT INTO `www_role` VALUES (1, '超级管理员', 0, '2021-08-12 14:06:54', 1, '2021-08-25 15:12:53', 0, NULL);

-- ----------------------------
-- Table structure for www_role_api
-- ----------------------------
DROP TABLE IF EXISTS `www_role_api`;
CREATE TABLE `www_role_api`  (
  `api_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY (`api_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色接口关联表' ROW_FORMAT = Dynamic;


SET FOREIGN_KEY_CHECKS = 1;
