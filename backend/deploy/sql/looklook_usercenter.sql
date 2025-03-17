/*
 Navicat MySQL Data Transfer

 Source Server         : looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : looklook_usercenter

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/03/2022 17:14:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `user_id` bigint NOT NULL DEFAULT '0',
  `auth_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一id',
  `auth_type` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_key` (`auth_type`,`auth_key`) USING BTREE,
  UNIQUE KEY `idx_userId_key` (`user_id`,`auth_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户授权表';

SET FOREIGN_KEY_CHECKS = 1;

--
-- Table structure for table `user_address`
--
DROP TABLE IF EXISTS `user_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_address` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `contact_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人姓名',
  `contact_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人手机号码',
  `district` json NOT NULL COMMENT '地区信息',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `postcode` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为默认地址 1是 0否',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户收货地址表';
/*!40101 SET character_set_client = @saved_cs_client */;