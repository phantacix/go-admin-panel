/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 10.1.31-MariaDB : Database - go_admin_panel
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`go_admin_panel` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `go_admin_panel`;

/*Table structure for table `sys_account` */

DROP TABLE IF EXISTS `sys_account`;

CREATE TABLE `sys_account` (
  `account_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '帐号id',
  `account_name` varchar(255) DEFAULT NULL COMMENT '帐号名',
  `name` varchar(255) DEFAULT NULL COMMENT '姓名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门id',
  `role_list` varchar(255) DEFAULT NULL COMMENT '帐号所属角色id.逗号分隔多个id',
  `status` int(11) DEFAULT NULL COMMENT '帐号状态',
  PRIMARY KEY (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `sys_dept` */

DROP TABLE IF EXISTS `sys_dept`;

CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '部门父id',
  `name` varchar(255) DEFAULT NULL COMMENT '部门名称',
  `order_num` int(11) DEFAULT NULL COMMENT '排序id',
  `del_flag` tinyint(4) DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `sys_log` */

DROP TABLE IF EXISTS `sys_log`;

CREATE TABLE `sys_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志自增id',
  `account_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `method` varchar(255) DEFAULT NULL COMMENT '执行方法',
  `path` varchar(255) DEFAULT NULL COMMENT '执行路径',
  `params` varchar(255) DEFAULT NULL COMMENT '请求参数',
  `ip` varchar(255) DEFAULT NULL COMMENT '访问ip',
  `create_time` bigint(20) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `sys_menu` */

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '菜单父id',
  `name` varchar(255) DEFAULT NULL COMMENT '菜单名称',
  `url` varchar(255) DEFAULT NULL COMMENT '菜单url',
  `auth_flag` varchar(255) DEFAULT NULL COMMENT '授权标识',
  `type` int(11) DEFAULT NULL COMMENT '菜单类型(0:目录,1:菜单,2:按钮)',
  `icon` varchar(255) DEFAULT NULL COMMENT '菜单图标',
  `order_num` int(11) DEFAULT NULL COMMENT '排序id',
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `role_name` varchar(255) DEFAULT NULL COMMENT '角色名称',
  `menu_list` varchar(255) DEFAULT NULL COMMENT '角色菜单id.逗号分隔menu_id',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
