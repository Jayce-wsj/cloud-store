/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : localhost:3306
 Source Schema         : cloudpan

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 13/03/2022 13:04:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `fid` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `size` int(11) NULL DEFAULT NULL COMMENT '文件大小',
  `hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件哈希',
  `uploadtime` datetime NULL DEFAULT NULL COMMENT '上传时间',
  `updatetime` datetime NULL DEFAULT NULL COMMENT '最近更新',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上传者',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`fid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of files
-- ----------------------------
INSERT INTO `files` VALUES (7, 'xlsx_1647143874308155800.xlsx', 102, NULL, '2022-03-13 11:57:54', NULL, 'admin', 'doc');
INSERT INTO `files` VALUES (8, 'xlsx_1647143946642491300.xlsx', 102, NULL, '2022-03-13 11:59:06', NULL, 'admin', 'doc');
INSERT INTO `files` VALUES (9, 'xlsx_1647144034627704100.xlsx', 102, NULL, '2022-03-13 12:00:34', NULL, 'admin1', 'doc');
INSERT INTO `files` VALUES (10, 'mp4_1647144076707585800.mp4', 1440, NULL, '2022-03-13 12:01:16', NULL, 'admin', 'mov');
INSERT INTO `files` VALUES (11, 'png_1647146908000724500.png', 464, NULL, '2022-03-13 12:48:28', '2022-03-13 12:48:28', 'admin', 'pic');
INSERT INTO `files` VALUES (12, 'txt_1647147019526669700.txt', 0, NULL, '2022-03-13 12:50:19', '2022-03-13 12:50:19', 'admin', 'other');
INSERT INTO `files` VALUES (13, 'xlsx_1647147449592143500.xlsx', 102, NULL, '2022-03-13 12:57:29', '2022-03-13 12:57:29', 'admin', 'doc');
INSERT INTO `files` VALUES (14, 'xlsx_1647147495165048700.xlsx', 102, NULL, '2022-03-13 12:58:15', '2022-03-13 12:58:15', 'admin', 'doc');
INSERT INTO `files` VALUES (15, 'xlsx_1647147495165563500.xlsx', 102, NULL, '2022-03-13 12:58:15', '2022-03-13 12:58:15', 'admin', 'doc');
INSERT INTO `files` VALUES (16, 'xlsx_1647147573463361000.xlsx', 102, NULL, '2022-03-13 12:59:33', '2022-03-13 12:59:33', 'admin', 'doc');
INSERT INTO `files` VALUES (17, 'xlsx_1647147573463872600.xlsx', 102, NULL, '2022-03-13 12:59:33', '2022-03-13 12:59:33', 'admin', 'doc');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `regtime` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', 'admin', '2022-03-12 21:05:37');
INSERT INTO `user` VALUES (2, '123123', '123123', '2022-03-12 21:06:25');

SET FOREIGN_KEY_CHECKS = 1;
