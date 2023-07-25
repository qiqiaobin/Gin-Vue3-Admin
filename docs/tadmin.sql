/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : go_fast_admin

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 08/04/2023 11:39:28
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `menu_type` tinyint NOT NULL COMMENT '菜单类型（0菜单，1按钮）',
  `path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '路由地址',
  `redirect` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '重定向路由地址',
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '外部地址',
  `is_iframe` tinyint NULL DEFAULT NULL COMMENT '是否内嵌',
  `is_hide` tinyint NULL DEFAULT NULL COMMENT '是否隐藏',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '按钮权限（权限即接口）',
  `request_method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `status` tinyint NOT NULL COMMENT '菜单状态（0启用 1停用）',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10012 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 0, '主页', 1, 'home', NULL, NULL, 0, 0, 'home/index', 'iconfont icon-shouye', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-02 22:01:01', NULL, NULL);
INSERT INTO `sys_menu` VALUES (10, 0, '系统管理', 1, 'system',  '/system/user',NULL, 0, 0, NULL, 'iconfont icon-xitongshezhi', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-07 22:59:47', NULL, NULL);
INSERT INTO `sys_menu` VALUES (80, 0, '外链', 1, 'link1',  NULL,'https://www.json.cn/', 0, 0, '', 'ele-ChromeFilled', '', '', 0, 0, '2022-12-05 21:21:34', '2023-04-07 23:05:25', NULL, '');
INSERT INTO `sys_menu` VALUES (81, 0, '内链', 1, 'link2', NULL, 'https://www.json.cn/', 1, 0, NULL, 'ele-ChromeFilled', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-07 22:59:29', NULL, NULL);
INSERT INTO `sys_menu` VALUES (90, 0, '公共接口', 1, 'common',  NULL,'', 0, 1, '', 'iconfont icon-neiqianshujuchucun', '', '', 0, 0, '2023-03-27 15:14:23', '2023-04-07 23:32:05', NULL, '');
INSERT INTO `sys_menu` VALUES (100, 10, '用户管理', 1, 'system/user', NULL,NULL, 0, 0, 'system/user/index', 'iconfont icon-icon-', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-02 21:59:50', NULL, NULL);
INSERT INTO `sys_menu` VALUES (110, 10, '菜单管理', 1, 'system/menu', NULL, NULL, 0, 0, 'system/menu/index', 'iconfont icon-caidan', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-02 22:00:14', NULL, NULL);
INSERT INTO `sys_menu` VALUES (120, 10, '角色管理', 1, 'system/role',  NULL,NULL, 0, 0, 'system/role/index', 'ele-Avatar', NULL, NULL, 0, 0, '2022-12-05 21:21:34', '2023-04-07 22:59:41', NULL, NULL);
INSERT INTO `sys_menu` VALUES (901, 90, '用户信息', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_auth_userInfo', 'GET', 0, 0, '2023-03-27 15:21:03', '2023-03-27 15:21:03', NULL, NULL);
INSERT INTO `sys_menu` VALUES (902, 90, '用户菜单', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_auth_menu', 'GET', 0, 0, '2023-03-27 15:22:17', '2023-03-27 15:22:17', NULL, NULL);
INSERT INTO `sys_menu` VALUES (903, 90, '修改密码', 2,  NULL,'', '', 0, 0, '', '', 'system_auth_updatePwd', 'POST', 0, 0, '2023-04-07 23:30:43', '2023-04-07 23:31:50', NULL, '');
INSERT INTO `sys_menu` VALUES (1000, 100, '用户查询', 2,  NULL,NULL, NULL, 0, 0,  NULL, NULL, 'system_user_query', 'GET', 0, 0, '2022-12-05 21:21:34', '2023-04-02 21:32:55', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1001, 100, '用户新增', 2,  NULL,NULL, NULL, 0, 0,  NULL, NULL, 'system_user_add', 'POST', 0, 0, '2022-12-05 21:21:34', '2023-04-07 23:00:01', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1002, 100, '用户更新', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_update', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1003, 100, '用户删除', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_user_delete', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1004, 100, '用户详情', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_detail', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1005, 100, '用户列表', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_user_list', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1006, 100, '重置用户密码', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_user_resetPwd', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1007, 100, '设置用户状态', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_user_setStatus', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1100, 110, '菜单查询', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_menu_query', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1101, 110, '菜单新增', 2, NULL, NULL, NULL, 0, 0, NULL, NULL, 'system_menu_add', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1102, 110, '菜单更新', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_menu_update', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1103, 110, '菜单删除', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_menu_delete', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1104, 110, '菜单详情', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_menu_detail', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1105, 110, '菜单列表', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_menu_list', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1106, 110, '菜单树形', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_menu_tree', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1200, 120, '角色查询', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_role_query', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1201, 120, '角色新增', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_role_add', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1202, 120, '角色更新', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_role_update', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1203, 120, '角色删除', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_role_delete', 'POST', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1204, 120, '角色详情', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_role_detail', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1205, 120, '角色列表', 2,  NULL,NULL, NULL, 0, 0, NULL, NULL, 'system_role_list', 'GET', 0, NULL, '2022-12-05 21:21:34', '2022-12-05 21:21:34', NULL, NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `role_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色代码',
  `status` tinyint NOT NULL COMMENT '角色状态（0启用 1停用）',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '管理员', 'admin', 0, 0, '2022-12-15 21:50:10', '2023-04-07 23:32:20', NULL, '');

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色和菜单关联id',
  `role_id` bigint NOT NULL COMMENT '角色id',
  `menu_id` bigint NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_role_menu`(`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1043 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (52, 1, 1);
INSERT INTO `sys_role_menu` VALUES (53, 1, 10);
INSERT INTO `sys_role_menu` VALUES (98, 1, 80);
INSERT INTO `sys_role_menu` VALUES (99, 1, 81);
INSERT INTO `sys_role_menu` VALUES (100, 1, 90);
INSERT INTO `sys_role_menu` VALUES (54, 1, 100);
INSERT INTO `sys_role_menu` VALUES (63, 1, 110);
INSERT INTO `sys_role_menu` VALUES (71, 1, 120);
INSERT INTO `sys_role_menu` VALUES (78, 1, 130);
INSERT INTO `sys_role_menu` VALUES (91, 1, 140);
INSERT INTO `sys_role_menu` VALUES (101, 1, 901);
INSERT INTO `sys_role_menu` VALUES (102, 1, 902);
INSERT INTO `sys_role_menu` VALUES (103, 1, 903);
INSERT INTO `sys_role_menu` VALUES (55, 1, 1000);
INSERT INTO `sys_role_menu` VALUES (56, 1, 1001);
INSERT INTO `sys_role_menu` VALUES (57, 1, 1002);
INSERT INTO `sys_role_menu` VALUES (58, 1, 1003);
INSERT INTO `sys_role_menu` VALUES (59, 1, 1004);
INSERT INTO `sys_role_menu` VALUES (60, 1, 1005);
INSERT INTO `sys_role_menu` VALUES (61, 1, 1006);
INSERT INTO `sys_role_menu` VALUES (62, 1, 1007);
INSERT INTO `sys_role_menu` VALUES (64, 1, 1100);
INSERT INTO `sys_role_menu` VALUES (65, 1, 1101);
INSERT INTO `sys_role_menu` VALUES (66, 1, 1102);
INSERT INTO `sys_role_menu` VALUES (67, 1, 1103);
INSERT INTO `sys_role_menu` VALUES (68, 1, 1104);
INSERT INTO `sys_role_menu` VALUES (69, 1, 1105);
INSERT INTO `sys_role_menu` VALUES (70, 1, 1106);
INSERT INTO `sys_role_menu` VALUES (72, 1, 1200);
INSERT INTO `sys_role_menu` VALUES (73, 1, 1201);
INSERT INTO `sys_role_menu` VALUES (74, 1, 1202);
INSERT INTO `sys_role_menu` VALUES (75, 1, 1203);
INSERT INTO `sys_role_menu` VALUES (76, 1, 1204);
INSERT INTO `sys_role_menu` VALUES (77, 1, 1205);
INSERT INTO `sys_role_menu` VALUES (79, 1, 1300);
INSERT INTO `sys_role_menu` VALUES (80, 1, 1301);
INSERT INTO `sys_role_menu` VALUES (81, 1, 1302);
INSERT INTO `sys_role_menu` VALUES (82, 1, 1303);
INSERT INTO `sys_role_menu` VALUES (83, 1, 1304);
INSERT INTO `sys_role_menu` VALUES (84, 1, 1305);
INSERT INTO `sys_role_menu` VALUES (85, 1, 1306);
INSERT INTO `sys_role_menu` VALUES (86, 1, 1307);
INSERT INTO `sys_role_menu` VALUES (87, 1, 1308);
INSERT INTO `sys_role_menu` VALUES (88, 1, 1309);
INSERT INTO `sys_role_menu` VALUES (89, 1, 1310);
INSERT INTO `sys_role_menu` VALUES (90, 1, 1311);
INSERT INTO `sys_role_menu` VALUES (92, 1, 1400);
INSERT INTO `sys_role_menu` VALUES (93, 1, 1401);
INSERT INTO `sys_role_menu` VALUES (94, 1, 1402);
INSERT INTO `sys_role_menu` VALUES (95, 1, 1403);
INSERT INTO `sys_role_menu` VALUES (96, 1, 1404);
INSERT INTO `sys_role_menu` VALUES (97, 1, 1405);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `user_type` tinyint NOT NULL COMMENT '用户类型（0普通账号，1超级管理员）',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码盐',
  `gender` tinyint NULL DEFAULT NULL COMMENT '用户性别（0未知，1男，2女）',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `status` tinyint NOT NULL COMMENT '帐号状态（0启用 1停用）',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 74 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'superAdmin', '超级管理员', 1, '999@qq.com', '19999999999', 'ec71d37f4340c223896afd45aaf3cf06', '41b278387b85404', 1, 'https://img1.baidu.com/it/u=948325104,3657174403&fm=253&fmt=auto&app=138&f=JPEG?w=388&h=514', 0, '2022-11-08 14:27:47', '2022-11-08 14:27:51', NULL, NULL);
INSERT INTO `sys_user` VALUES (2, 'admin', '管理员', 0, '999@qq.com', '19999999999', 'ec71d37f4340c223896afd45aaf3cf06', '41b278387b85404', 1, NULL, 0, '2022-11-08 14:27:47', '2023-04-08 11:35:08', NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户和角色关联id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `role_id` bigint NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_user_role`(`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (10, 2, 1);

SET FOREIGN_KEY_CHECKS = 1;
