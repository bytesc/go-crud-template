/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : crud-list

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 10/11/2023 10:53:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crud_list
-- ----------------------------
DROP TABLE IF EXISTS `crud_list`;
CREATE TABLE `crud_list`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `birthday` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_crud_list_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of crud_list
-- ----------------------------
INSERT INTO `crud_list` VALUES (1, '2023-11-07 21:27:35.394', '2023-11-07 21:27:35.394', '2023-11-07 22:45:32.469', '张三', '', '', '', '', '');
INSERT INTO `crud_list` VALUES (2, '2023-11-07 21:31:27.696', '2023-11-08 20:01:59.105', '2023-11-09 22:56:03.954', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (3, '2023-11-07 21:31:32.675', '2023-11-07 21:31:32.675', '2023-11-09 23:11:37.320', '张三', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (4, '2023-11-07 21:37:03.731', '2023-11-07 21:37:03.731', '2023-11-07 22:46:06.996', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (5, '2023-11-07 21:37:17.170', '2023-11-07 21:37:17.170', '2023-11-09 22:57:01.751', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (6, '2023-11-07 21:39:26.496', '2023-11-07 21:39:26.496', '2023-11-09 22:57:03.488', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (7, '2023-11-08 20:46:11.399', '2023-11-08 20:46:11.399', '2023-11-09 22:57:04.104', '王五', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (8, '2023-11-09 13:58:56.743', '2023-11-09 13:58:56.743', '2023-11-09 23:12:32.704', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (9, '2023-11-09 13:58:57.457', '2023-11-09 13:58:57.457', '2023-11-09 23:12:34.160', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (10, '2023-11-09 13:58:58.690', '2023-11-09 13:58:58.690', '2023-11-09 23:08:07.592', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (11, '2023-11-09 13:58:59.310', '2023-11-09 13:58:59.310', '2023-11-09 23:08:09.054', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (12, '2023-11-09 13:59:01.872', '2023-11-09 13:59:01.872', '2023-11-09 23:08:09.915', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (13, '2023-11-09 13:59:13.883', '2023-11-09 13:59:13.883', '2023-11-09 23:08:10.590', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (14, '2023-11-09 13:59:14.622', '2023-11-09 13:59:14.622', '2023-11-09 23:12:35.036', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (15, '2023-11-09 13:59:15.253', '2023-11-09 13:59:15.253', '2023-11-09 23:12:56.670', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (16, '2023-11-09 13:59:15.799', '2023-11-09 13:59:15.799', '2023-11-09 23:12:57.249', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (17, '2023-11-09 13:59:16.305', '2023-11-09 13:59:16.305', '2023-11-09 23:14:06.045', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (18, '2023-11-09 13:59:16.818', '2023-11-09 13:59:16.818', '2023-11-09 23:13:35.029', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (19, '2023-11-09 13:59:17.215', '2023-11-09 13:59:17.215', '2023-11-09 23:15:25.639', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (20, '2023-11-09 13:59:17.663', '2023-11-09 13:59:17.663', '2023-11-09 23:15:27.337', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (21, '2023-11-09 13:59:18.049', '2023-11-09 13:59:18.049', '2023-11-09 23:15:28.320', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (22, '2023-11-09 13:59:18.401', '2023-11-09 13:59:18.401', '2023-11-09 23:15:51.203', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (23, '2023-11-09 13:59:28.167', '2023-11-09 13:59:28.167', '2023-11-09 23:24:46.475', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (24, '2023-11-09 13:59:28.781', '2023-11-09 13:59:28.781', '2023-11-09 23:16:48.551', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (25, '2023-11-09 22:23:41.724', '2023-11-09 22:23:41.724', '2023-11-09 23:27:28.649', '1wq', '', '', '', '', '');
INSERT INTO `crud_list` VALUES (26, '2023-11-09 22:34:25.917', '2023-11-09 22:34:25.917', '2023-11-09 23:25:06.288', '章', '', '', '', '2023-11-25', '');
INSERT INTO `crud_list` VALUES (27, '2023-11-09 22:38:03.718', '2023-11-09 22:38:03.718', '2023-11-09 23:24:48.936', 'hello', 'sss', '13213213', '1214324', '', '');
INSERT INTO `crud_list` VALUES (28, '2023-11-09 22:40:25.083', '2023-11-09 22:40:25.083', '2023-11-09 23:26:28.712', 'hello', 'sss', '13213213', '1214324', '', '');
INSERT INTO `crud_list` VALUES (29, '2023-11-09 22:41:28.863', '2023-11-09 22:41:28.863', '2023-11-09 23:25:51.951', '吴', '', '', '', '', '');
INSERT INTO `crud_list` VALUES (30, '2023-11-09 22:58:42.714', '2023-11-09 22:58:42.714', '2023-11-09 23:25:39.839', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (31, '2023-11-09 22:58:43.480', '2023-11-10 08:35:26.555', NULL, '狗蛋', '55', '2222@22.mail', '13988889999', '', '');
INSERT INTO `crud_list` VALUES (32, '2023-11-09 22:58:44.106', '2023-11-09 22:58:44.106', '2023-11-09 23:25:15.943', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (33, '2023-11-09 22:58:45.548', '2023-11-10 08:42:57.166', NULL, '李四', '55', '2222@22.mail', '18989898989', '', '');
INSERT INTO `crud_list` VALUES (34, '2023-11-09 22:58:46.004', '2023-11-09 22:58:46.004', '2023-11-09 23:26:44.827', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (35, '2023-11-09 22:58:46.393', '2023-11-09 22:58:46.393', '2023-11-10 08:44:04.621', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (36, '2023-11-09 22:58:46.716', '2023-11-09 22:58:46.716', '2023-11-09 22:59:18.211', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (37, '2023-11-09 23:09:17.051', '2023-11-09 23:09:17.051', '2023-11-10 08:35:00.519', '12', '', '', '', '', '');
INSERT INTO `crud_list` VALUES (38, '2023-11-09 23:28:47.790', '2023-11-09 23:28:47.790', '2023-11-10 08:44:04.621', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (39, '2023-11-09 23:28:48.486', '2023-11-09 23:28:48.486', NULL, '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (40, '2023-11-09 23:28:55.466', '2023-11-09 23:28:55.466', '2023-11-10 10:23:40.838', '王五', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (41, '2023-11-09 23:28:59.515', '2023-11-10 10:24:06.850', NULL, '张飞', 'ss', '2ff@22.mail', '1998787322', '2023-11-22', '');
INSERT INTO `crud_list` VALUES (42, '2023-11-09 23:29:03.675', '2023-11-09 23:29:03.675', '2023-11-10 08:44:13.917', '王五', '55', '55522@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (43, '2023-11-09 23:29:05.722', '2023-11-09 23:29:05.722', NULL, '王五', '55', '55522@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (44, '2023-11-09 23:29:20.496', '2023-11-10 10:22:10.980', NULL, '孙新', '595', '5ssx@e2.mail', '', '', '');
INSERT INTO `crud_list` VALUES (45, '2023-11-09 23:29:21.926', '2023-11-10 10:24:50.176', NULL, '李思', '09', 'ls@22.mail', '18990908787', '2023-11-07', '');
INSERT INTO `crud_list` VALUES (46, '2023-11-09 23:29:22.967', '2023-11-09 23:29:22.967', '2023-11-10 10:22:20.003', '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (47, '2023-11-09 23:29:23.373', '2023-11-09 23:29:23.373', '2023-11-10 10:22:25.201', '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (48, '2023-11-09 23:29:24.977', '2023-11-10 10:48:26.070', NULL, '张c', '595', '5ss@22.mail', '16767980909', '', '');
INSERT INTO `crud_list` VALUES (49, '2023-11-09 23:29:25.385', '2023-11-09 23:29:25.385', '2023-11-10 10:22:25.203', '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (50, '2023-11-09 23:29:38.552', '2023-11-09 23:29:38.552', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (51, '2023-11-09 23:29:39.270', '2023-11-09 23:29:39.270', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (52, '2023-11-09 23:29:39.978', '2023-11-09 23:29:39.978', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (53, '2023-11-09 23:29:40.638', '2023-11-09 23:29:40.638', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (54, '2023-11-09 23:29:45.245', '2023-11-09 23:29:45.245', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (55, '2023-11-09 23:29:47.022', '2023-11-09 23:29:47.022', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (56, '2023-11-09 23:29:47.677', '2023-11-09 23:29:47.677', NULL, '孙', '595', '5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (57, '2023-11-09 23:30:09.748', '2023-11-09 23:30:09.748', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (58, '2023-11-09 23:30:10.439', '2023-11-09 23:30:10.439', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (59, '2023-11-09 23:30:11.082', '2023-11-09 23:30:11.082', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (60, '2023-11-09 23:30:11.630', '2023-11-09 23:30:11.630', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (61, '2023-11-09 23:30:12.191', '2023-11-09 23:30:12.191', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (62, '2023-11-09 23:30:12.673', '2023-11-09 23:30:12.673', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (63, '2023-11-09 23:30:13.188', '2023-11-09 23:30:13.188', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (64, '2023-11-09 23:30:13.680', '2023-11-09 23:30:13.680', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (65, '2023-11-09 23:30:14.202', '2023-11-09 23:30:14.202', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (66, '2023-11-09 23:30:14.773', '2023-11-09 23:30:14.773', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (67, '2023-11-09 23:30:15.208', '2023-11-09 23:30:15.208', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (68, '2023-11-09 23:30:15.683', '2023-11-09 23:30:15.683', NULL, '唐', '95', 't5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (69, '2023-11-09 23:30:31.471', '2023-11-09 23:30:31.471', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (70, '2023-11-09 23:30:31.935', '2023-11-09 23:30:31.935', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (71, '2023-11-09 23:30:32.343', '2023-11-09 23:30:32.343', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (72, '2023-11-09 23:30:32.710', '2023-11-09 23:30:32.710', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (73, '2023-11-09 23:30:33.100', '2023-11-09 23:30:33.100', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (74, '2023-11-09 23:30:33.544', '2023-11-09 23:30:33.544', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (75, '2023-11-09 23:30:33.982', '2023-11-09 23:30:33.982', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (76, '2023-11-09 23:30:34.417', '2023-11-09 23:30:34.417', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (77, '2023-11-09 23:30:34.862', '2023-11-09 23:30:34.862', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (78, '2023-11-09 23:30:35.245', '2023-11-09 23:30:35.245', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (79, '2023-11-09 23:30:35.612', '2023-11-10 08:45:06.194', NULL, '吴口', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (80, '2023-11-09 23:30:35.970', '2023-11-09 23:30:35.970', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (81, '2023-11-09 23:30:36.381', '2023-11-09 23:30:36.381', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (82, '2023-11-09 23:30:36.769', '2023-11-09 23:30:36.769', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (83, '2023-11-09 23:30:37.146', '2023-11-09 23:30:37.146', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (84, '2023-11-09 23:30:37.530', '2023-11-09 23:30:37.530', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (85, '2023-11-09 23:30:37.909', '2023-11-09 23:30:37.909', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (86, '2023-11-09 23:30:38.261', '2023-11-09 23:30:38.261', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (87, '2023-11-09 23:30:38.641', '2023-11-09 23:30:38.641', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (88, '2023-11-09 23:30:38.965', '2023-11-09 23:30:38.965', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (89, '2023-11-09 23:30:39.301', '2023-11-09 23:30:39.301', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (90, '2023-11-09 23:30:39.685', '2023-11-09 23:30:39.685', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (91, '2023-11-09 23:30:40.014', '2023-11-09 23:30:40.014', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (92, '2023-11-09 23:30:40.360', '2023-11-09 23:30:40.360', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (93, '2023-11-09 23:30:40.693', '2023-11-09 23:30:40.693', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (94, '2023-11-09 23:30:41.030', '2023-11-09 23:30:41.030', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (95, '2023-11-09 23:30:41.387', '2023-11-09 23:30:41.387', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (96, '2023-11-09 23:30:41.734', '2023-11-09 23:30:41.734', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (97, '2023-11-09 23:30:42.045', '2023-11-09 23:30:42.045', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (98, '2023-11-09 23:30:43.386', '2023-11-09 23:30:43.386', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (99, '2023-11-09 23:30:43.744', '2023-11-09 23:30:43.744', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (100, '2023-11-09 23:30:44.150', '2023-11-09 23:30:44.150', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (101, '2023-11-09 23:30:44.495', '2023-11-09 23:30:44.495', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (102, '2023-11-09 23:30:44.829', '2023-11-09 23:30:44.829', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (103, '2023-11-09 23:30:45.197', '2023-11-09 23:30:45.197', NULL, '吴', '5695', 'wwt5ss@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (104, '2023-11-09 23:31:28.582', '2023-11-09 23:31:28.582', NULL, '张三', 'g', '123213', '12313', '2023-11-28', '');

SET FOREIGN_KEY_CHECKS = 1;
