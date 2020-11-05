/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : b2b_dev_common

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 05/11/2020 17:44:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_corp_admin
-- ----------------------------
DROP TABLE IF EXISTS `t_corp_admin`;
CREATE TABLE `t_corp_admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账户',
  `password` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '盐',
  `last_login_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后一次登录时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态（1:正常，2:禁用）',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '平台管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_corp_admin
-- ----------------------------
INSERT INTO `t_corp_admin` VALUES (1, '测试人员', 'admin', 'a2a6c163513ecd4fa8588d27ffb13de5', '1292eb', 1604569425, 1, 1565582315, 1603941682, 0);

SET FOREIGN_KEY_CHECKS = 1;
