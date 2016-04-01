/*
Navicat MySQL Data Transfer

Source Server         : a_troy测试
Source Server Version : 50621
Source Host           : localhost:3306
Source Database       : prometheus

Target Server Type    : MYSQL
Target Server Version : 50621
File Encoding         : 65001

Date: 2016-04-01 14:46:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `cabinet`
-- ----------------------------
DROP TABLE IF EXISTS `cabinet`;
CREATE TABLE `cabinet` (
  `cabinet_id` int(11) NOT NULL AUTO_INCREMENT,
  `cabinet_name` varchar(255) NOT NULL,
  `iscloud` enum('true','false') NOT NULL,
  `capacity_total` bigint(20) DEFAULT NULL,
  `capacity_used` bigint(20) DEFAULT NULL,
  `location_id` int(11) NOT NULL,
  PRIMARY KEY (`cabinet_id`),
  UNIQUE KEY `uniq_name_location_id` (`cabinet_name`,`location_id`),
  KEY `catbinet_fk_location_id` (`location_id`),
  CONSTRAINT `catbinet_fk_location_id` FOREIGN KEY (`location_id`) REFERENCES `location` (`location_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of cabinet
-- ----------------------------
INSERT INTO `cabinet` VALUES ('1', 'asdasd', 'false', '1023', '1000', '1');
INSERT INTO `cabinet` VALUES ('7', 'yyyy', 'false', '2047', '0', '1');
INSERT INTO `cabinet` VALUES ('9', '', 'false', '2047', '0', '2');

-- ----------------------------
-- Table structure for `device`
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device` (
  `device_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_name` varchar(255) DEFAULT NULL,
  `device_model_id` int(11) DEFAULT NULL,
  `fater_device_id` int(11) DEFAULT NULL,
  `cabinet_id` int(11) DEFAULT NULL,
  `u_position` bigint(20) DEFAULT '0',
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `uniq_name` (`device_name`) USING BTREE,
  KEY `device_fk_cabinet_id` (`cabinet_id`),
  KEY `device_fk_father_device_id` (`fater_device_id`),
  KEY `device_fk_device_model_id` (`device_model_id`) USING BTREE,
  CONSTRAINT `device_fk_cabinet_id` FOREIGN KEY (`cabinet_id`) REFERENCES `cabinet` (`cabinet_id`) ON UPDATE CASCADE,
  CONSTRAINT `device_fk_device_model_id` FOREIGN KEY (`device_model_id`) REFERENCES `device_model` (`device_model_id`) ON UPDATE CASCADE,
  CONSTRAINT `device_fk_father_device_id` FOREIGN KEY (`fater_device_id`) REFERENCES `device` (`device_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of device
-- ----------------------------

-- ----------------------------
-- Table structure for `device_model`
-- ----------------------------
DROP TABLE IF EXISTS `device_model`;
CREATE TABLE `device_model` (
  `device_model_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_model_name` varchar(255) NOT NULL,
  `device_type` enum('Other','Router','Switch','Server','VM') NOT NULL DEFAULT 'Other',
  PRIMARY KEY (`device_model_id`),
  UNIQUE KEY `uniq_device_model_name` (`device_model_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of device_model
-- ----------------------------
INSERT INTO `device_model` VALUES ('1', 'IBM', 'Server');
INSERT INTO `device_model` VALUES ('2', '华为', 'Router');
INSERT INTO `device_model` VALUES ('3', 'cisco', 'Switch');
INSERT INTO `device_model` VALUES ('11', '1asdasdasda', 'Server');
INSERT INTO `device_model` VALUES ('13', '333', 'Router');

-- ----------------------------
-- Table structure for `ipv4`
-- ----------------------------
DROP TABLE IF EXISTS `ipv4`;
CREATE TABLE `ipv4` (
  `ipv4_int` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  PRIMARY KEY (`ipv4_int`),
  KEY `ipv4_fk_device_id` (`device_id`),
  CONSTRAINT `ipv4_fk_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of ipv4
-- ----------------------------

-- ----------------------------
-- Table structure for `location`
-- ----------------------------
DROP TABLE IF EXISTS `location`;
CREATE TABLE `location` (
  `location_id` int(11) NOT NULL AUTO_INCREMENT,
  `location_name` varchar(255) NOT NULL,
  `picture` varchar(255) NOT NULL,
  `father_location_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`location_id`),
  UNIQUE KEY `uniq_name` (`location_name`) USING BTREE,
  KEY `fk_location_father_location_id` (`father_location_id`),
  CONSTRAINT `fk_location_father_location_id` FOREIGN KEY (`father_location_id`) REFERENCES `location` (`location_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of location
-- ----------------------------
INSERT INTO `location` VALUES ('1', '大连IDC', '1.jpg', '2');
INSERT INTO `location` VALUES ('2', '长春IDC', '1.jpg', null);
INSERT INTO `location` VALUES ('3', 'singapore123', 'asas', '1');

-- ----------------------------
-- Table structure for `tag`
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) NOT NULL,
  PRIMARY KEY (`tag_id`),
  UNIQUE KEY `uniq_name` (`tag_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tag
-- ----------------------------
