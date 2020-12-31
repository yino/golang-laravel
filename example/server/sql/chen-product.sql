/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 192.168.14.208:3306
 Source Schema         : chen-product

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 23/12/2020 16:20:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_corp_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_corp_admin`;
CREATE TABLE `sys_corp_admin`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` int(10) NULL DEFAULT NULL,
  `updated_at` int(10) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '1启用 0禁用',
  `salt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '随机字符串',
  `last_login_time` int(10) NULL DEFAULT NULL,
  `token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `mobile`(`account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_corp_admin
-- ----------------------------
INSERT INTO `sys_corp_admin` VALUES (1, 'sun123', '孙', 0, 1608706804, NULL, 'da3177cbd9f064004b6a0d59a3a484bb', 1, 'abcd', 1608706804, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDg3MTA3MDcsInVpZCI6Ilx1MDAwMSJ9.2FT86aKA2PDlWPDVXhBoQMInQhcLiOEHUHfxv5CYvw8');

-- ----------------------------
-- Table structure for sys_member
-- ----------------------------
DROP TABLE IF EXISTS `sys_member`;
CREATE TABLE `sys_member`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` int(10) NULL DEFAULT 0,
  `updated_at` int(10) NULL DEFAULT 0,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `integral` int(10) NULL DEFAULT 0 COMMENT '积分',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_member
-- ----------------------------
INSERT INTO `sys_member` VALUES (1, '15829090356', '孙1', 1, 1607068288, '2020-12-04 16:10:16', 21);
INSERT INTO `sys_member` VALUES (2, '15829090358', '孙5', 1606968039, 1607935587, NULL, 201);
INSERT INTO `sys_member` VALUES (3, '15829090351', '孙6', 1606980352, 1606980352, NULL, 1);
INSERT INTO `sys_member` VALUES (4, '15829090356', '孙7', 1607072630, 1607072630, NULL, 16);
INSERT INTO `sys_member` VALUES (5, '15829090357', '孙17', 1608023929, 1608023929, NULL, 1);
INSERT INTO `sys_member` VALUES (6, '15829090352', '孙康', 1608024175, 1608024175, NULL, 0);
INSERT INTO `sys_member` VALUES (7, '15829090353', '孙康1', 1608024277, 1608024277, NULL, 0);
INSERT INTO `sys_member` VALUES (8, '15829090354', '孙康5', 1608024492, 1608024492, NULL, 0);
INSERT INTO `sys_member` VALUES (9, '15829090359', '孙康6', 1608024571, 1608024571, '2020-12-18 18:45:10', 0);
INSERT INTO `sys_member` VALUES (10, '15829090317', '孙康7', 1608024616, 1608024616, '2020-12-18 18:41:50', 0);
INSERT INTO `sys_member` VALUES (11, '15829090327', '孙康12', 1608024643, 1608710620, NULL, 120);

-- ----------------------------
-- Table structure for sys_member_integral_record
-- ----------------------------
DROP TABLE IF EXISTS `sys_member_integral_record`;
CREATE TABLE `sys_member_integral_record`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `member_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户电话',
  `member_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户姓名',
  `created_at` int(10) NULL DEFAULT 0,
  `updated_at` int(10) NULL DEFAULT 0,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `integral` int(10) NULL DEFAULT 0 COMMENT '当前获得积分',
  `surplus_integral` int(11) NULL DEFAULT NULL COMMENT '剩余积分 不含本次获得',
  `member_id` int(11) NULL DEFAULT NULL COMMENT '用户id',
  `product_id` int(11) NULL DEFAULT NULL,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品名称',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '操作员id',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作员姓名',
  `money` int(255) NULL DEFAULT NULL COMMENT '消费金额/分',
  `date` date NULL DEFAULT NULL COMMENT '每天的日期',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `member_id`(`member_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_member_integral_record
-- ----------------------------
INSERT INTO `sys_member_integral_record` VALUES (1, '15829090356', '孙', 1607068006, 1607068006, NULL, 1, 16, 1, 1, '你好', 1, NULL, 5000, '2020-12-11');
INSERT INTO `sys_member_integral_record` VALUES (2, '15829090356', '孙', 1607068095, 1607068095, NULL, 1, 17, 1, 1, '你好', 1, NULL, 8000, '2020-12-12');
INSERT INTO `sys_member_integral_record` VALUES (3, '15829090356', '孙', 1607068187, 1607068187, NULL, 1, 18, 1, 1, '你好', 1, NULL, 10000, '2020-12-13');
INSERT INTO `sys_member_integral_record` VALUES (4, '15829090356', '孙', 1607068229, 1607068229, NULL, 1, 19, 1, 1, '你好', 1, '孙', 2000, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (5, '15829090356', '孙', 1607068288, 1607068288, NULL, 1, 20, 1, 1, '你好', 1, '孙', 1000, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (6, '15829090358', '孙', 1607934980, 1607934980, NULL, 20, 1, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (7, '15829090358', '孙', 1607935016, 1607935016, NULL, 20, 21, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (8, '15829090358', '孙', 1607935017, 1607935017, NULL, 20, 41, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (9, '15829090358', '孙', 1607935017, 1607935017, NULL, 20, 61, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (10, '15829090358', '孙', 1607935019, 1607935019, NULL, 20, 81, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (11, '15829090358', '孙', 1607935019, 1607935019, NULL, 20, 101, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (12, '15829090358', '孙', 1607935020, 1607935020, NULL, 20, 121, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (13, '15829090358', '孙', 1607935020, 1607935020, NULL, 20, 141, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (14, '15829090358', '孙', 1607935020, 1607935020, NULL, 20, 161, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (15, '15829090358', '孙', 1607935587, 1607935587, NULL, 20, 181, 2, 2, '你好呀', 1, '孙', 200, '2020-12-14');
INSERT INTO `sys_member_integral_record` VALUES (16, '15829090327', '孙康12', 1608109416, 1608109416, NULL, 20, 1, 11, 2, '你好呀', 1, '孙', 10, '2020-12-16');
INSERT INTO `sys_member_integral_record` VALUES (17, '15829090327', '孙康12', 1608109462, 1608109462, NULL, 10, 21, 11, 2, '你好呀', 1, '孙', NULL, '2020-12-16');
INSERT INTO `sys_member_integral_record` VALUES (18, '15829090327', '孙康12', 1608110058, 1608110058, NULL, 20, 31, 11, 2, '你好呀', 1, '孙', 10, '2020-12-16');
INSERT INTO `sys_member_integral_record` VALUES (19, '15829090327', '孙康12', 1608710332, 1608710332, NULL, 16, 51, 11, 2, '你好呀', 1, '孙', 500, '2020-12-23');
INSERT INTO `sys_member_integral_record` VALUES (20, '15829090327', '孙康12', 1608710362, 1608710362, NULL, 21, 67, 11, 3, 'sun17测试1', 1, '孙', 20, '2020-12-23');
INSERT INTO `sys_member_integral_record` VALUES (21, '15829090327', '孙康12', 1608710373, 1608710373, NULL, 16, 88, 11, 2, '你好呀', 1, '孙', 10, '2020-12-23');
INSERT INTO `sys_member_integral_record` VALUES (22, '15829090327', '孙康12', 1608710620, 1608710620, NULL, 16, 104, 11, 2, '你好呀', 1, '孙', 300, '2020-12-23');

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` int(10) NULL DEFAULT 0,
  `updated_at` int(10) NULL DEFAULT 0,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) NULL DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `start_time` int(10) NULL DEFAULT 0,
  `end_time` int(10) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
INSERT INTO `sys_operation_records` VALUES (10, 1607073406, 1607073406, NULL, '127.0.0.1', 'GET', '/corp/member/index?page=1', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '/corp/member/index?page=1', NULL, 1, 1607073406, 1607073406);
INSERT INTO `sys_operation_records` VALUES (11, 1607073485, 1607073485, NULL, '127.0.0.1', 'GET', '/corp/member/index?page=1', 200, 1, 'PostmanRuntime/7.26.8', NULL, '/corp/member/index?page=1', NULL, 1, 1607073485, 1607073485);
INSERT INTO `sys_operation_records` VALUES (12, 1607073536, 1607073536, NULL, '127.0.0.1', 'GET', '/corp/member/index?page=1', 200, 1, 'PostmanRuntime/7.26.8', NULL, '/corp/member/index?page=1', NULL, 1, 1607073536, 1607073536);
INSERT INTO `sys_operation_records` VALUES (13, 1607073577, 1607073577, NULL, '127.0.0.1', 'GET', '/corp/member/index?page=1', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '/corp/member/index?page=1', NULL, 1, 1607073577, 1607073577);
INSERT INTO `sys_operation_records` VALUES (14, 1607073578, 1607073578, NULL, '127.0.0.1', 'GET', '/corp/member/index?page=1', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '/corp/member/index?page=1', NULL, 1, 1607073578, 1607073578);
INSERT INTO `sys_operation_records` VALUES (15, 1607073804, 1607073804, NULL, '127.0.0.1', 'GET', '/corp/member/index?page=1', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '/corp/member/index?page=1', NULL, 1, 1607073804, 1607073804);
INSERT INTO `sys_operation_records` VALUES (16, 1607073815, 1607073815, NULL, '127.0.0.1', 'POST', '/corp/member/edit', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '{\"id\":[\"1\"],\"integral\":[\"16\"],\"mobile\":[\"15829090356\"],\"name\":[\"孙\"]}', NULL, 1, 1607073815, 1607073815);
INSERT INTO `sys_operation_records` VALUES (17, 1607073955, 1607073955, NULL, '127.0.0.1', 'POST', '/corp/member/edit', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '{\"id\":[\"1\"],\"integral\":[\"16\"],\"mobile\":[\"15829090356\"],\"name\":[\"孙\"]}', '{\"code\":50000,\"data\":null,\"msg\":\"数据已存在\"}', 1, 1607073955, 1607073955);
INSERT INTO `sys_operation_records` VALUES (18, 1607074232, 1607074232, NULL, '127.0.0.1', 'POST', '/corp/member/edit', 200, NULL, 'PostmanRuntime/7.26.8', NULL, '{\"id\":[\"1\"],\"integral\":[\"16\"],\"mobile\":[\"15829090356\"],\"name\":[\"孙\"]}', '{\"code\":50000,\"data\":null,\"msg\":\"数据已存在\"}', 1, 1607074232, 1607074232);

-- ----------------------------
-- Table structure for sys_product
-- ----------------------------
DROP TABLE IF EXISTS `sys_product`;
CREATE TABLE `sys_product`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品名称',
  `created_at` int(10) NULL DEFAULT 0,
  `updated_at` int(10) NULL DEFAULT 0,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `integral` int(10) NULL DEFAULT 0 COMMENT '所得积分',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_product
-- ----------------------------
INSERT INTO `sys_product` VALUES (1, '你好', 12, 0, '2020-12-04 16:10:21', 2, NULL);
INSERT INTO `sys_product` VALUES (2, '你好呀', 1607069167, 1607069390, NULL, 16, '');
INSERT INTO `sys_product` VALUES (3, 'sun17测试1', 1608285682, 1608287822, NULL, 21, '/storage/upload/product/20201218183701微信图片_20200819154512.jpg');
INSERT INTO `sys_product` VALUES (4, 'sun17测试2', 1608285700, 1608285700, '2020-12-18 18:49:30', 10, NULL);

SET FOREIGN_KEY_CHECKS = 1;
