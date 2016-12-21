-- MySQL dump 10.13  Distrib 5.1.73, for redhat-linux-gnu (x86_64)
--
-- Host: 192.168.33.13    Database: promethues_golang
-- ------------------------------------------------------
-- Server version	5.6.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `VM`
--

DROP TABLE IF EXISTS `VM`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `VM` (
  `device_id` int(11) NOT NULL COMMENT '设备id',
  `hostname` varchar(255) NOT NULL COMMENT '主机名',
  `memsize` int(11) NOT NULL COMMENT '内存大小（MB）',
  `processor` int(11) NOT NULL DEFAULT '0',
  `os` enum('CentOS','Unknow') NOT NULL DEFAULT 'Unknow' COMMENT '操作系统',
  `release` float NOT NULL COMMENT '操作系统版本',
  `last_change_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后一次探测时间',
  `checksum` varchar(255) NOT NULL DEFAULT 'Never' COMMENT '校验系统各方面是否改变',
  PRIMARY KEY (`device_id`),
  CONSTRAINT `fk_vm_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `audit`
--

DROP TABLE IF EXISTS `audit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `audit` (
  `ctime` bigint(20) NOT NULL COMMENT '创建时间',
  `operate_type` enum('get','update','delete','add') NOT NULL COMMENT '操作类型',
  `operate_object_type` enum('device','ipv4') NOT NULL COMMENT '操作对象的类别',
  `user_name` varchar(255) NOT NULL COMMENT '操作人',
  `operate_object` text NOT NULL COMMENT '操作对象明细',
  `object_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `blade_cabinet`
--

DROP TABLE IF EXISTS `blade_cabinet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blade_cabinet` (
  `device_id` int(11) NOT NULL COMMENT '刀箱的设备ID',
  `capacity` int(10) NOT NULL COMMENT '刀箱容量',
  PRIMARY KEY (`device_id`),
  CONSTRAINT `device_id_fk_4_blade_cabinet` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `blade_server`
--

DROP TABLE IF EXISTS `blade_server`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blade_server` (
  `device_id` int(11) NOT NULL COMMENT '设备ID',
  `hostname` varchar(255) NOT NULL DEFAULT 'Unknow' COMMENT '主机名',
  `memsize` int(11) NOT NULL DEFAULT '0' COMMENT '内存大小（MB）',
  `processor` int(11) NOT NULL DEFAULT '0',
  `os` enum('Unknow','Fedora','Solaris','Windows','Unix','CentOS','RedHat') NOT NULL DEFAULT 'Unknow' COMMENT '操作系统型号',
  `release` float NOT NULL DEFAULT '0' COMMENT '操作系统版本',
  `last_change_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后一次探测时间',
  `checksum` varchar(255) NOT NULL DEFAULT 'Never' COMMENT '校验是否更新',
  `relative_position` int(11) DEFAULT NULL COMMENT '相对在刀箱里的位置',
  `blade_cabinet_id` int(11) DEFAULT NULL COMMENT '刀箱ID（与device表的father_device_id一致，冗余字段为了跟刀箱相对位置做uniq key）',
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `uniq_relative_position` (`relative_position`),
  KEY `blade_cabinet_id_fk` (`blade_cabinet_id`),
  CONSTRAINT `blade_cabinet_id_fk` FOREIGN KEY (`blade_cabinet_id`) REFERENCES `blade_cabinet` (`device_id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `device_id_fk_4_blade_server` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cabinet`
--

DROP TABLE IF EXISTS `cabinet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cabinet` (
  `cabinet_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '机柜ID',
  `cabinet_name` varchar(255) NOT NULL COMMENT '机柜名',
  `capacity_total` int(20) NOT NULL COMMENT '机柜总容量U数',
  `capacity_used` int(20) DEFAULT NULL COMMENT '机柜容量的使用情况（缓存字段）',
  `idc_id` int(11) NOT NULL COMMENT '机柜属于的机房ID',
  `ctime` bigint(20) NOT NULL COMMENT '机柜创建时间',
  PRIMARY KEY (`cabinet_id`),
  UNIQUE KEY `uniq_name_location_id` (`cabinet_name`,`idc_id`),
  KEY `catbinet_fk_idc_id` (`idc_id`) USING BTREE,
  CONSTRAINT `catbinet_fk_idc_id` FOREIGN KEY (`idc_id`) REFERENCES `idc` (`idc_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cloud`
--

DROP TABLE IF EXISTS `cloud`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cloud` (
  `device_model_id` int(11) NOT NULL COMMENT '云的设备型号ID（云也是一种设备型号）',
  `cloud_type` enum('huawei','other','openstack') NOT NULL DEFAULT 'other' COMMENT '云的类型',
  PRIMARY KEY (`device_model_id`),
  CONSTRAINT `cloud_id_2_device_model_id` FOREIGN KEY (`device_model_id`) REFERENCES `deviceModel` (`device_model_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `device`
--

DROP TABLE IF EXISTS `device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `device` (
  `device_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '设备ID',
  `device_name` varchar(255) NOT NULL COMMENT '设备名',
  `device_model_id` int(11) NOT NULL COMMENT '设备型号',
  `father_device_id` int(11) DEFAULT NULL COMMENT '父设备ID（vm或者刀片会有）',
  `ctime` int(11) NOT NULL COMMENT '设备创建时间',
  `group_id` int(11) DEFAULT NULL COMMENT '设备组ID',
  `env` int(11) DEFAULT NULL COMMENT '设备所属环境 0:offline 1:dev 2:qa 3:prod',
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `uniq_name` (`device_name`) USING BTREE,
  KEY `device_fk_father_device_id` (`father_device_id`),
  KEY `device_model_id_fk` (`device_model_id`),
  CONSTRAINT `device_fk_father_device_id` FOREIGN KEY (`father_device_id`) REFERENCES `device` (`device_id`) ON UPDATE CASCADE,
  CONSTRAINT `device_model_id_fk` FOREIGN KEY (`device_model_id`) REFERENCES `deviceModel` (`device_model_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1211 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `deviceModel`
--

DROP TABLE IF EXISTS `deviceModel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deviceModel` (
  `device_model_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '设备型号ID',
  `device_model_name` varchar(255) NOT NULL COMMENT '设备型号名',
  `device_type` enum('Other','Router','Switch','Server','BladeCabinet','BladeServer','Cloud','VM') NOT NULL DEFAULT 'Other' COMMENT '设备型号类型',
  `u` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备型号所占的物理U位数',
  `half_full` enum('full','half') NOT NULL DEFAULT 'full' COMMENT '设备型号所占用的U位面积',
  PRIMARY KEY (`device_model_id`),
  UNIQUE KEY `uniq_device_model_name` (`device_model_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=188 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `deviceTag`
--

DROP TABLE IF EXISTS `deviceTag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deviceTag` (
  `device_id` int(11) NOT NULL COMMENT '设备ID',
  `tag` varchar(11) NOT NULL COMMENT '设备标签',
  UNIQUE KEY `u_relation` (`device_id`,`tag`),
  CONSTRAINT `fk_deviceTag_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `group_id` int(11) NOT NULL COMMENT '设备组ID',
  `group_name` varchar(255) NOT NULL COMMENT '设备组名字',
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `idc`
--

DROP TABLE IF EXISTS `idc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `idc` (
  `idc_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '机房的id',
  `idc_name` varchar(255) NOT NULL COMMENT '机房的id名',
  `location_id` int(11) DEFAULT NULL COMMENT '机房的地理位置ID',
  PRIMARY KEY (`idc_id`),
  UNIQUE KEY `u_idc_name` (`idc_name`) USING BTREE,
  KEY `fk_location_id` (`location_id`),
  CONSTRAINT `fk_location_id` FOREIGN KEY (`location_id`) REFERENCES `location` (`location_id`)
) ENGINE=InnoDB AUTO_INCREMENT=177 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location` (
  `location_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '地理位置ID',
  `location_name` varchar(255) NOT NULL COMMENT '地理位置名',
  `father_location_id` int(11) DEFAULT NULL COMMENT '上级地理位置ID',
  `location_type` enum('云机房','自建机房','租用机房') DEFAULT '租用机房' COMMENT '地理位置如果是机房需要有机房类型',
  `location_interface` varchar(255) DEFAULT NULL COMMENT '地理位置如果是机房的接入方式',
  `location_contact` varchar(255) DEFAULT NULL COMMENT '地理位置如果是机房 联络人',
  `location_contact_num` varchar(255) DEFAULT NULL COMMENT '地理位置如果是机房 联系电话',
  `location_address` varchar(255) DEFAULT NULL COMMENT '地理位置如果是机房 地址',
  PRIMARY KEY (`location_id`),
  UNIQUE KEY `uniq_name` (`location_name`) USING BTREE,
  KEY `fk_location_father_location_id` (`father_location_id`),
  CONSTRAINT `fk_location_father_location_id` FOREIGN KEY (`father_location_id`) REFERENCES `location` (`location_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=211 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `netPort`
--

DROP TABLE IF EXISTS `netPort`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `netPort` (
  `mac` varchar(17) DEFAULT NULL COMMENT 'mac地址',
  `ipv4_int` bigint(64) NOT NULL COMMENT 'ipv4整型',
  `device_id` int(11) NOT NULL COMMENT '设备ID',
  `netPort_type` enum('vip','ether','Unknow','bond') NOT NULL DEFAULT 'Unknow' COMMENT '网口类型',
  `function_type` enum('business','manage') NOT NULL DEFAULT 'manage' COMMENT '网络接口功能',
  `ctime` bigint(20) NOT NULL COMMENT '创建时间或分配时间',
  UNIQUE KEY `un_ipv4_int` (`ipv4_int`),
  KEY `ipv4_fk_device_id` (`device_id`),
  CONSTRAINT `fk_netPort_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `router`
--

DROP TABLE IF EXISTS `router`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `router` (
  `device_id` int(11) NOT NULL COMMENT '设备ID',
  `serial` varchar(255) DEFAULT NULL COMMENT '序列号',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述字段',
  PRIMARY KEY (`device_id`),
  CONSTRAINT `fk_device_id_4_router` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `server`
--

DROP TABLE IF EXISTS `server`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `server` (
  `serial` varchar(255) DEFAULT 'Unknow' COMMENT '序列号',
  `device_id` int(11) NOT NULL COMMENT '设备id',
  `hostname` varchar(255) NOT NULL DEFAULT 'Unknow' COMMENT '主机名',
  `memsize` int(11) NOT NULL DEFAULT '0' COMMENT '内存大小（MB）',
  `processor` int(11) NOT NULL DEFAULT '0',
  `os` enum('Unknow','Fedora','Solaris','Windows','Unix','CentOS','RedHat') NOT NULL DEFAULT 'Unknow' COMMENT '操作系统',
  `_release` float NOT NULL DEFAULT '0' COMMENT '操作系统版本',
  `last_change_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后一次探测时间',
  `checksum` varchar(255) NOT NULL DEFAULT 'Never' COMMENT '校验系统各方面是否改变',
  PRIMARY KEY (`device_id`),
  CONSTRAINT `fk_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `space`
--

DROP TABLE IF EXISTS `space`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `space` (
  `cabinet_id` int(11) NOT NULL COMMENT '机柜ID',
  `device_id` int(11) NOT NULL COMMENT '设备ID',
  `u_position` int(11) NOT NULL COMMENT '具体U位点',
  `position` enum('rear','front','interior') NOT NULL COMMENT 'U位坐标点',
  UNIQUE KEY `atom_fk_position` (`cabinet_id`,`u_position`,`position`),
  KEY `atom_fk_cabinet_id` (`cabinet_id`),
  KEY `atom_fk_device_id` (`device_id`),
  CONSTRAINT `atom_fk_cabinet_id` FOREIGN KEY (`cabinet_id`) REFERENCES `cabinet` (`cabinet_id`) ON UPDATE CASCADE,
  CONSTRAINT `atom_fk_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `switch`
--

DROP TABLE IF EXISTS `switch`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `switch` (
  `serial` varchar(255) DEFAULT NULL COMMENT '序列号',
  `device_id` int(11) NOT NULL COMMENT '交换机设备ID',
  `portals` int(11) NOT NULL DEFAULT '0' COMMENT '交换机口数量',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述字段',
  PRIMARY KEY (`device_id`),
  KEY `device_id` (`device_id`),
  CONSTRAINT `device_id_fk` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `vlan`
--

DROP TABLE IF EXISTS `vlan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vlan` (
  `vlan_id` int(11) NOT NULL COMMENT '交换机vlan的id号',
  `vlan_name` varchar(255) NOT NULL COMMENT 'vlan自定义名字',
  `net` bigint(11) NOT NULL COMMENT 'vlan网段',
  `mask` int(11) NOT NULL COMMENT 'vlan掩码',
  `domain` varchar(255) NOT NULL COMMENT 'vlan域（自定义说明字段）',
  `ctime` bigint(20) NOT NULL COMMENT 'vlan创建时间',
  PRIMARY KEY (`vlan_name`,`net`),
  UNIQUE KEY `network_segment` (`net`,`mask`),
  KEY `net` (`net`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `vlan_switch`
--

DROP TABLE IF EXISTS `vlan_switch`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vlan_switch` (
  `vlan_net` bigint(20) NOT NULL COMMENT 'vlan网段',
  `device_id` int(11) NOT NULL COMMENT '交换机设备ID',
  KEY `device_id` (`device_id`),
  KEY `vlan_net_fk` (`vlan_net`),
  CONSTRAINT `vlan_net_fk` FOREIGN KEY (`vlan_net`) REFERENCES `vlan` (`net`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `vlan_switch_ibfk_1` FOREIGN KEY (`device_id`) REFERENCES `switch` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-12-21 17:16:11
