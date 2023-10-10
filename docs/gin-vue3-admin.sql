/*
 Navicat Premium Data Transfer

 Source Server         : n9e
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : 10.77.64.33:3306
 Source Schema         : gin-vue3-admin

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 10/10/2023 14:00:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典id',
  `dict_name` varchar(100) DEFAULT NULL COMMENT '字典名称',
  `dict_code` varchar(100) NOT NULL COMMENT '字典代码',
  `dict_type` tinyint(4) DEFAULT NULL COMMENT '字典类型（0int，1string）',
  `status` tinyint(4) NOT NULL COMMENT '状态（0启用 1停用）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict` VALUES (1, '配置状态', 'config_status', 0, 0, '2023-04-06 16:07:38', '2023-04-06 16:07:38', NULL, '');
INSERT INTO `sys_dict` VALUES (2, '字典状态', 'dict_status', 0, 0, '2023-04-06 16:13:27', '2023-04-06 16:13:27', NULL, '');
INSERT INTO `sys_dict` VALUES (3, '字典类型', 'dict_type', 0, 0, '2023-04-06 16:13:50', '2023-04-06 16:13:50', NULL, '');
INSERT INTO `sys_dict` VALUES (4, '字典选项状态', 'dict_item_status', 0, 0, '2023-04-06 16:14:41', '2023-04-06 16:14:41', NULL, '');
INSERT INTO `sys_dict` VALUES (5, '菜单状态', 'menu_status', 0, 0, '2023-04-06 16:22:15', '2023-04-06 16:22:15', NULL, '');
INSERT INTO `sys_dict` VALUES (6, '菜单类型', 'menu_type', 0, 0, '2023-04-06 16:22:31', '2023-04-06 16:22:31', NULL, '');
INSERT INTO `sys_dict` VALUES (7, '角色状态', 'role_status', 0, 0, '2023-04-06 16:23:00', '2023-04-06 16:23:00', NULL, '');
INSERT INTO `sys_dict` VALUES (8, '用户状态', 'user_status', 0, 0, '2023-04-06 16:23:16', '2023-04-06 16:23:16', NULL, '');
INSERT INTO `sys_dict` VALUES (9, '用户类型', 'user_type', 0, 0, '2023-04-06 16:23:44', '2023-04-06 16:23:44', NULL, '');
INSERT INTO `sys_dict` VALUES (10, '用户性别', 'user_gender', 0, 0, '2023-04-06 16:23:53', '2023-04-06 16:23:53', NULL, '');
INSERT INTO `sys_dict` VALUES (12, '测试', 'test', 0, 0, '2023-10-09 11:09:07', '2023-10-09 11:09:07', NULL, '测试');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父菜单ID',
  `menu_name` varchar(50) NOT NULL COMMENT '菜单名称',
  `menu_type` tinyint(4) NOT NULL COMMENT '菜单类型（0菜单，1按钮）',
  `path` varchar(200) DEFAULT '' COMMENT '路由地址',
  `redirect` varchar(200) DEFAULT '' COMMENT '重定向路由地址',
  `link` varchar(255) DEFAULT NULL COMMENT '外部地址',
  `is_iframe` tinyint(4) DEFAULT NULL COMMENT '是否内嵌',
  `is_hide` tinyint(4) DEFAULT NULL COMMENT '是否隐藏',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(100) DEFAULT NULL COMMENT '菜单图标',
  `permission` varchar(255) DEFAULT NULL COMMENT '按钮权限（权限即接口）',
  `request_method` varchar(255) DEFAULT NULL COMMENT '请求方法',
  `status` tinyint(4) NOT NULL COMMENT '菜单状态（0启用 1停用）',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10016 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='菜单信息表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES (1, 0, '主页', 1, 'home', NULL, NULL, 0, 0, 'home/index', 'iconfont icon-shouye', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-02 22:01:01', NULL, NULL);
INSERT INTO `sys_menu` VALUES (10, 0, '系统管理', 1, 'system', '/system/user', NULL, 0, 0, NULL, 'iconfont icon-xitongshezhi', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-07 22:59:47', NULL, NULL);
INSERT INTO `sys_menu` VALUES (80, 0, '外链', 1, 'link1', NULL, 'https://www.json.cn/', 0, 0, '', 'ele-ChromeFilled', '', '', 0, 0, '2022-12-05 21:21:34', '2023-04-07 23:05:25', NULL, '');
INSERT INTO `sys_menu` VALUES (81, 0, '内链', 1, 'link2', NULL, 'https://www.json.cn/', 1, 0, NULL, 'ele-ChromeFilled', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-07 22:59:29', NULL, NULL);
INSERT INTO `sys_menu` VALUES (90, 0, '公共接口', 1, 'common', NULL, '', 0, 1, '', 'iconfont icon-neiqianshujuchucun', '', '', 0, 0, '2023-03-27 15:14:23', '2023-04-07 23:32:05', NULL, '');
INSERT INTO `sys_menu` VALUES (100, 10, '用户管理', 1, 'system/user', NULL, NULL, 0, 0, 'system/user/index', 'iconfont icon-icon-', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-02 21:59:50', NULL, NULL);
INSERT INTO `sys_menu` VALUES (110, 10, '菜单管理', 1, 'system/menu', NULL, NULL, 0, 0, 'system/menu/index', 'iconfont icon-caidan', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-02 22:00:14', NULL, NULL);
INSERT INTO `sys_menu` VALUES (120, 10, '角色管理', 1, 'system/role', NULL, NULL, 0, 0, 'system/role/index', 'ele-Avatar', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-07 22:59:41', NULL, NULL);
INSERT INTO `sys_menu` VALUES (901, 90, '用户信息', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_auth_userInfo', 'GET', 0, 0, '2023-03-27 15:21:03', '2023-03-27 15:21:03', NULL, NULL);
INSERT INTO `sys_menu` VALUES (902, 90, '用户菜单', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_auth_menu', 'GET', 0, 0, '2023-03-27 15:22:17', '2023-03-27 15:22:17', NULL, NULL);
INSERT INTO `sys_menu` VALUES (903, 90, '修改密码', 2, NULL, '', '', 0, 0, '', '', 'system_auth_updatePwd', 'POST', 0, 0, '2023-04-07 23:30:43', '2023-04-07 23:31:50', NULL, '');
INSERT INTO `sys_menu` VALUES (1000, 100, '用户查询', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_query', 'GET', 0, 0, '2022-12-05 21:21:34', '2023-04-02 21:32:55', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1001, 100, '用户新增', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_add', 'POST', 0, 0, '2022-12-05 21:21:34', '2023-04-07 23:00:01', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1002, 100, '用户更新', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_update', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1003, 100, '用户删除', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_delete', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1004, 100, '用户详情', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_detail', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1005, 100, '用户列表', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_list', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1006, 100, '重置用户密码', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_resetPwd', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1007, 100, '设置用户状态', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_setStatus', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1100, 110, '菜单查询', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_query', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1101, 110, '菜单新增', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_add', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1102, 110, '菜单更新', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_update', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1103, 110, '菜单删除', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_delete', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1104, 110, '菜单详情', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_detail', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1105, 110, '菜单列表', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_list', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1106, 110, '菜单树形', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_tree', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1200, 120, '角色查询', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_role_query', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1201, 120, '角色新增', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_role_add', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1202, 120, '角色更新', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_role_update', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1203, 120, '角色删除', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_role_delete', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1204, 120, '角色详情', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_role_detail', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1205, 120, '角色列表', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_role_list', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (10012, 0, 'test', 1, 'test', '', '', 0, 0, '/test', 'iconfont icon-siweidaotu', '', '', 0, 0, '2023-08-03 10:29:42', '2023-08-03 10:29:42', NULL, '');
INSERT INTO `sys_menu` VALUES (10014, 10, '字典管理', 1, 'system/dict', '', '', 0, 0, 'system/dict/index', 'ele-EditPen', '', '', 0, 0, '2023-10-09 17:46:41', '2023-10-09 17:46:41', NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(100) NOT NULL COMMENT '角色名称',
  `nickname` varchar(30) NOT NULL COMMENT '显示名称',
  `note` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, 'Admin', '管理员', '');
INSERT INTO `sys_role` VALUES (10, 'Test', '测试', '测试功能');
INSERT INTO `sys_role` VALUES (12, 'DEV', '开发', '开发人员');
INSERT INTO `sys_role` VALUES (52, 'Tftest', '测试2', '测试功能更新');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色和菜单关联id',
  `role_name` varchar(128) NOT NULL COMMENT '角色名称',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  `operation` varchar(128) NOT NULL COMMENT '操作路由',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_role_menu` (`role_name`,`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1390 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` VALUES (52, 'Admin', 1, '/home');
INSERT INTO `sys_role_menu` VALUES (53, 'Admin', 10, '/system');
INSERT INTO `sys_role_menu` VALUES (54, 'Admin', 100, '/system/user');
INSERT INTO `sys_role_menu` VALUES (55, 'Admin', 1000, '/system_user_query');
INSERT INTO `sys_role_menu` VALUES (56, 'Admin', 1001, '/system_user_add');
INSERT INTO `sys_role_menu` VALUES (57, 'Admin', 1002, '/system_user_update');
INSERT INTO `sys_role_menu` VALUES (58, 'Admin', 1003, '/system_user_delete');
INSERT INTO `sys_role_menu` VALUES (59, 'Admin', 1004, '/system_user_detail');
INSERT INTO `sys_role_menu` VALUES (60, 'Admin', 1005, '/system_user_list');
INSERT INTO `sys_role_menu` VALUES (61, 'Admin', 1006, '/system_user_resetPwd');
INSERT INTO `sys_role_menu` VALUES (63, 'Admin', 110, '/system/menu');
INSERT INTO `sys_role_menu` VALUES (64, 'Admin', 1100, '/system_menu_query');
INSERT INTO `sys_role_menu` VALUES (65, 'Admin', 1101, '/system_menu_add');
INSERT INTO `sys_role_menu` VALUES (66, 'Admin', 1102, '/system_menu_update');
INSERT INTO `sys_role_menu` VALUES (67, 'Admin', 1103, '/system_menu_delete');
INSERT INTO `sys_role_menu` VALUES (68, 'Admin', 1104, '/system_menu_detail');
INSERT INTO `sys_role_menu` VALUES (69, 'Admin', 1105, '/system_menu_list');
INSERT INTO `sys_role_menu` VALUES (70, 'Admin', 1106, '/system_menu_tree');
INSERT INTO `sys_role_menu` VALUES (71, 'Admin', 120, '/system/role');
INSERT INTO `sys_role_menu` VALUES (72, 'Admin', 1200, '/system_role_query');
INSERT INTO `sys_role_menu` VALUES (73, 'Admin', 1201, '/system_role_add');
INSERT INTO `sys_role_menu` VALUES (74, 'Admin', 1202, '/system_role_update');
INSERT INTO `sys_role_menu` VALUES (75, 'Admin', 1203, '/system_role_delete');
INSERT INTO `sys_role_menu` VALUES (76, 'Admin', 1204, '/system_role_detail');
INSERT INTO `sys_role_menu` VALUES (77, 'Admin', 1205, '/system_role_list');
INSERT INTO `sys_role_menu` VALUES (98, 'Admin', 80, '/link1');
INSERT INTO `sys_role_menu` VALUES (99, 'Admin', 81, '/link2');
INSERT INTO `sys_role_menu` VALUES (100, 'Admin', 90, '/common');
INSERT INTO `sys_role_menu` VALUES (101, 'Admin', 901, '/system_auth_userInfo');
INSERT INTO `sys_role_menu` VALUES (102, 'Admin', 902, '/system_auth_menu');
INSERT INTO `sys_role_menu` VALUES (103, 'Admin', 903, '/system_auth_updatePwd');
INSERT INTO `sys_role_menu` VALUES (1106, 'Test', 1, '/home');
INSERT INTO `sys_role_menu` VALUES (1108, 'Test', 10, '/system');
INSERT INTO `sys_role_menu` VALUES (1110, 'Test', 100, '/system/user');
INSERT INTO `sys_role_menu` VALUES (1112, 'Test', 1002, '/system_user_update');
INSERT INTO `sys_role_menu` VALUES (1114, 'Test', 1003, '/system_user_delete');
INSERT INTO `sys_role_menu` VALUES (1116, 'Test', 1004, '/system_user_detail');
INSERT INTO `sys_role_menu` VALUES (1118, 'Test', 1005, '/system_user_list');
INSERT INTO `sys_role_menu` VALUES (1120, 'Test', 1006, '/system_user_resetPwd');
INSERT INTO `sys_role_menu` VALUES (1124, 'Test', 1000, '/system_user_query');
INSERT INTO `sys_role_menu` VALUES (1126, 'Test', 1001, '/system_user_add');
INSERT INTO `sys_role_menu` VALUES (1128, 'Test', 110, '/system/menu');
INSERT INTO `sys_role_menu` VALUES (1130, 'Test', 1100, '/system_menu_query');
INSERT INTO `sys_role_menu` VALUES (1132, 'Test', 1101, '/system_menu_add');
INSERT INTO `sys_role_menu` VALUES (1134, 'Test', 1102, '/system_menu_update');
INSERT INTO `sys_role_menu` VALUES (1136, 'Test', 1103, '/system_menu_delete');
INSERT INTO `sys_role_menu` VALUES (1138, 'Test', 1104, '/system_menu_detail');
INSERT INTO `sys_role_menu` VALUES (1140, 'Test', 1105, '/system_menu_list');
INSERT INTO `sys_role_menu` VALUES (1142, 'Test', 1106, '/system_menu_tree');
INSERT INTO `sys_role_menu` VALUES (1144, 'Test', 120, '/system/role');
INSERT INTO `sys_role_menu` VALUES (1146, 'Test', 1200, '/system_role_query');
INSERT INTO `sys_role_menu` VALUES (1148, 'Test', 1201, '/system_role_add');
INSERT INTO `sys_role_menu` VALUES (1150, 'Test', 1202, '/system_role_update');
INSERT INTO `sys_role_menu` VALUES (1152, 'Test', 1203, '/system_role_delete');
INSERT INTO `sys_role_menu` VALUES (1154, 'Test', 1204, '/system_role_detail');
INSERT INTO `sys_role_menu` VALUES (1156, 'Test', 1205, '/system_role_list');
INSERT INTO `sys_role_menu` VALUES (1158, 'Test', 90, '/common');
INSERT INTO `sys_role_menu` VALUES (1160, 'Test', 901, '/system_auth_userInfo');
INSERT INTO `sys_role_menu` VALUES (1162, 'Test', 902, '/system_auth_menu');
INSERT INTO `sys_role_menu` VALUES (1164, 'Test', 903, '/system_auth_updatePwd');
INSERT INTO `sys_role_menu` VALUES (1166, 'Test', 10012, '/Test');
INSERT INTO `sys_role_menu` VALUES (1334, 'DEV', 1, '');
INSERT INTO `sys_role_menu` VALUES (1336, 'DEV', 10, '');
INSERT INTO `sys_role_menu` VALUES (1338, 'DEV', 100, '');
INSERT INTO `sys_role_menu` VALUES (1340, 'DEV', 1002, '');
INSERT INTO `sys_role_menu` VALUES (1342, 'DEV', 1003, '');
INSERT INTO `sys_role_menu` VALUES (1344, 'DEV', 1004, '');
INSERT INTO `sys_role_menu` VALUES (1346, 'DEV', 1005, '');
INSERT INTO `sys_role_menu` VALUES (1348, 'DEV', 1006, '');
INSERT INTO `sys_role_menu` VALUES (1350, 'DEV', 1007, '');
INSERT INTO `sys_role_menu` VALUES (1352, 'DEV', 1000, '');
INSERT INTO `sys_role_menu` VALUES (1354, 'DEV', 1001, '');
INSERT INTO `sys_role_menu` VALUES (1356, 'DEV', 110, '');
INSERT INTO `sys_role_menu` VALUES (1358, 'DEV', 1100, '');
INSERT INTO `sys_role_menu` VALUES (1360, 'DEV', 1101, '');
INSERT INTO `sys_role_menu` VALUES (1362, 'DEV', 1102, '');
INSERT INTO `sys_role_menu` VALUES (1364, 'DEV', 1103, '');
INSERT INTO `sys_role_menu` VALUES (1366, 'DEV', 1104, '');
INSERT INTO `sys_role_menu` VALUES (1368, 'DEV', 1105, '');
INSERT INTO `sys_role_menu` VALUES (1370, 'DEV', 1106, '');
INSERT INTO `sys_role_menu` VALUES (1372, 'DEV', 120, '');
INSERT INTO `sys_role_menu` VALUES (1374, 'DEV', 1200, '');
INSERT INTO `sys_role_menu` VALUES (1376, 'DEV', 1201, '');
INSERT INTO `sys_role_menu` VALUES (1378, 'DEV', 1202, '');
INSERT INTO `sys_role_menu` VALUES (1380, 'DEV', 1203, '');
INSERT INTO `sys_role_menu` VALUES (1382, 'DEV', 1204, '');
INSERT INTO `sys_role_menu` VALUES (1384, 'DEV', 1205, '');
INSERT INTO `sys_role_menu` VALUES (1386, 'DEV', 80, '');
INSERT INTO `sys_role_menu` VALUES (1388, 'DEV', 81, '');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(30) NOT NULL COMMENT '用户账号',
  `nickname` varchar(30) NOT NULL COMMENT '用户昵称',
  `email` varchar(50) DEFAULT NULL COMMENT '用户邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号码',
  `roles` varchar(20) DEFAULT NULL COMMENT '角色',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `portrait` varchar(255) DEFAULT NULL COMMENT '头像地址',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(64) NOT NULL DEFAULT '',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `update_by` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'superadmin', '超级管理员', '999@qq.com', '19999999999', 'Admin', 'be7bd798ee215d01f98092cc88d03fdd', 'https://img1.baidu.com/it/u=948325104,3657174403&fm=253&fmt=auto&app=138&f=JPEG?w=388&h=514', '2022-11-08 14:27:47', 'system', '2022-11-08 14:27:47', 'system');
INSERT INTO `users` VALUES (2, 'admin', '管理员', '999@qq.com', '19999999999', 'Admin', 'be7bd798ee215d01f98092cc88d03fdd', NULL, '2022-11-08 14:27:47', 'system', '2022-11-08 14:27:47', 'system');
INSERT INTO `users` VALUES (3, 'test', '测试用户', '999@qq.com', '19999999999', 'Test', 'be7bd798ee215d01f98092cc88d03fdd', NULL, '2022-11-08 14:27:47', 'system', '2022-11-08 14:27:47', 'system');
INSERT INTO `users` VALUES (78, 'ceshi2', '测试用户', '123456@qq.com', '13566777888', 'Test', 'be7bd798ee215d01f98092cc88d03fdd', '', '2023-09-01 09:04:32', 'system', '2023-09-01 15:22:53', 'system');
INSERT INTO `users` VALUES (80, 'ceshi', '测试新增', '123456@qq.com', '13566666666', '', 'be7bd798ee215d01f98092cc88d03fdd', '', '2023-09-14 10:37:39', 'superadmin', '2023-09-14 10:37:39', 'superadmin');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
