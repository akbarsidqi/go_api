/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 100138
 Source Host           : localhost:3306
 Source Schema         : db_animal

 Target Server Type    : MySQL
 Target Server Version : 100138
 File Encoding         : 65001

 Date: 28/02/2022 11:43:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for animals
-- ----------------------------
DROP TABLE IF EXISTS `animals`;
CREATE TABLE `animals`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `class` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `legs` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of animals
-- ----------------------------
INSERT INTO `animals` VALUES (1, 'lion', 'mammal', 4);
INSERT INTO `animals` VALUES (2, 'lions', 'mammals', 5);
INSERT INTO `animals` VALUES (4, 'lionsis', 'mammalsis', 8);
INSERT INTO `animals` VALUES (5, 'lionsisi', 'mammalsisi', 9);

SET FOREIGN_KEY_CHECKS = 1;
