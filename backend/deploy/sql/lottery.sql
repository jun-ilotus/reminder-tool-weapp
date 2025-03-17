/*
 Navicat Premium Data Transfer

 Source Server         : gozero-looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : lottery

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 15/01/2024 18:07:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lottery
-- ----------------------------
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `user_id` int(0) NOT NULL DEFAULT 0 COMMENT '发起抽奖用户ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等奖名称',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等经配图',
  `publish_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '开奖设置：1按时间开奖 2按人数开奖 3即抽即中',
  `publish_time` datetime(0) NOT NULL COMMENT '开奖时间',
  `join_number` int(0) NOT NULL DEFAULT 0 COMMENT '自动开奖人数',
  `introduce` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '抽奖说明',
  `award_deadline` datetime(0) NOT NULL COMMENT '领奖截止时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  `is_selected` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否精选 1是 0否',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '抽奖表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lottery
-- ----------------------------
INSERT INTO `lottery` VALUES (1, 1, '自然堂 气垫BB 自然色', 'https://img10.360buyimg.com/n1/jfs/t1/86934/2/28516/101853/627119fbE03141b18/3a94e4953b95e90d.jpg', 1, '2024-01-14 17:01:35', 0, '官方抽奖，欢迎参与', '2024-01-14 17:02:06', '2024-01-14 17:02:09', '2024-01-15 15:22:12', 1);
INSERT INTO `lottery` VALUES (2, 1, '农夫山泉', 'https://img10.360buyimg.com/n1/jfs/t1/86934/2/28516/101853/627119fbE03141b18/3a94e4953b95e90d.jpg', 1, '2024-01-15 15:21:40', 0, '这是介绍', '2024-01-27 15:21:43', '2024-01-15 15:21:55', '2024-01-15 15:22:05', 1);

-- ----------------------------
-- Table structure for prize
-- ----------------------------
DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `lottery_id` int(0) NOT NULL DEFAULT 0 COMMENT '抽奖ID',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包',
  `name` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品名称',
  `level` int(0) NOT NULL DEFAULT 1 COMMENT '几等奖 默认1',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品图',
  `count` int(0) NOT NULL DEFAULT 0 COMMENT '奖品份数',
  `grant_type` tinyint(1) NOT NULL COMMENT '奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '奖品表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
