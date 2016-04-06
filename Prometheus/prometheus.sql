/*
Navicat MySQL Data Transfer

Source Server         : a_troy测试
Source Server Version : 50621
Source Host           : localhost:3306
Source Database       : prometheus

Target Server Type    : MYSQL
Target Server Version : 50621
File Encoding         : 65001

Date: 2016-04-06 17:15:21
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

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
  `father_device_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `uniq_name` (`device_name`) USING BTREE,
  KEY `device_fk_father_device_id` (`father_device_id`),
  KEY `device_fk_device_model_id` (`device_model_id`) USING BTREE,
  CONSTRAINT `device_fk_device_model_id` FOREIGN KEY (`device_model_id`) REFERENCES `devicemodel` (`device_model_id`) ON UPDATE CASCADE,
  CONSTRAINT `device_fk_father_device_id` FOREIGN KEY (`father_device_id`) REFERENCES `device` (`device_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of device
-- ----------------------------
INSERT INTO `device` VALUES ('1', '192.168.33.81', '1', null);

-- ----------------------------
-- Table structure for `devicemodel`
-- ----------------------------
DROP TABLE IF EXISTS `devicemodel`;
CREATE TABLE `devicemodel` (
  `device_model_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_model_name` varchar(255) NOT NULL,
  `device_type` enum('Other','Router','Switch','Server','VM') NOT NULL DEFAULT 'Other',
  PRIMARY KEY (`device_model_id`),
  UNIQUE KEY `uniq_device_model_name` (`device_model_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of devicemodel
-- ----------------------------
INSERT INTO `devicemodel` VALUES ('1', 'IBM', 'Server');
INSERT INTO `devicemodel` VALUES ('2', '华为', 'Router');
INSERT INTO `devicemodel` VALUES ('3', 'cisco', 'Switch');
INSERT INTO `devicemodel` VALUES ('11', '1asdasdasda', 'Server');
INSERT INTO `devicemodel` VALUES ('13', '333', 'Router');

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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of location
-- ----------------------------
INSERT INTO `location` VALUES ('1', '大连IDC', '1.jpg', '2');
INSERT INTO `location` VALUES ('2', '长春IDC', '1.jpg', null);
INSERT INTO `location` VALUES ('3', 'singapore123', 'asas', '1');

-- ----------------------------
-- Table structure for `netport`
-- ----------------------------
DROP TABLE IF EXISTS `netport`;
CREATE TABLE `netport` (
  `mac` varchar(17) DEFAULT NULL,
  `ipv4_int` bigint(64) DEFAULT NULL,
  `device_id` int(11) DEFAULT NULL,
  `type` enum('vip','eth','bond') NOT NULL,
  UNIQUE KEY `un_ipv4_int` (`ipv4_int`),
  KEY `ipv4_fk_device_id` (`device_id`),
  CONSTRAINT `ipv4_fk_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of netport
-- ----------------------------
INSERT INTO `netport` VALUES ('00:CF:E0:31:D5:D5', '3232244049', '1', 'vip');
INSERT INTO `netport` VALUES ('00:CF:E0:31:D5:D5', '3232244050', '1', 'vip');

-- ----------------------------
-- Table structure for `space`
-- ----------------------------
DROP TABLE IF EXISTS `space`;
CREATE TABLE `space` (
  `cabinet_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `u_position` int(11) NOT NULL,
  `position` enum('rear','front','interior') NOT NULL,
  KEY `atom_fk_cabinet_id` (`cabinet_id`),
  KEY `atom_fk_device_id` (`device_id`),
  CONSTRAINT `atom_fk_cabinet_id` FOREIGN KEY (`cabinet_id`) REFERENCES `cabinet` (`cabinet_id`) ON UPDATE CASCADE,
  CONSTRAINT `atom_fk_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of space
-- ----------------------------

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
