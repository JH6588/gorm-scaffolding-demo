/*
 Navicat Premium Data Transfer

 Source Server         : test loacal
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 07/03/2023 20:43:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for people
-- ----------------------------
DROP TABLE IF EXISTS `people`;
CREATE TABLE `people`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `age` tinyint NULL DEFAULT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of people
-- ----------------------------
INSERT INTO `people` VALUES (4, 'xxx', 19, 'xxx@.com');
INSERT INTO `people` VALUES (5, 'xxx', 19, 'xxx@.com');
INSERT INTO `people` VALUES (6, 'ss', 27, 'ww@gmail.com');
INSERT INTO `people` VALUES (7, 'sss', 66, 'ww@gmail.com');
INSERT INTO `people` VALUES (8, 'java', 92, 'ww@gmail.com');

SET FOREIGN_KEY_CHECKS = 1;
