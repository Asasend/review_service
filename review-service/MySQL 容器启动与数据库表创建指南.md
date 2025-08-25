# MySQL 容器启动与数据库表创建指南
## 1. 启动 MySQL 容器
### 1.1 清理已存在的容器
如果之前已经创建过同名容器，需要先停止并删除：

```bash
docker stop mysql
docker rm mysql
```
### 1.2 启动 MySQL 容器
使用以下命令启动 MySQL 容器（端口映射为 3307:3306）：

```bash
docker run --name mysql -p 3307:3306 -e 
MYSQL_ROOT_PASSWORD=root1234 -d mysql
```
参数说明：

- --name mysql ：容器名称
- -p 3307:3306 ：端口映射，宿主机 3307 端口映射到容器 3306 端口
- -e MYSQL_ROOT_PASSWORD=root1234 ：设置 root 用户密码
- -d ：后台运行
- mysql ：使用最新版本的 MySQL 镜像
### 1.3 验证容器状态
检查容器是否正常运行：

```bash

docker ps -a
```
## 2. 连接到 MySQL 容器
### 2.1 进入容器内部连接
使用以下命令进入容器并连接到 MySQL：


```bash
docker exec -it mysql mysql -uroot -p
```
输入密码： root1234

## 3. 创建数据库和表
### 3.1 创建数据库

```sql
CREATE DATABASE review_service CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;
USE review_service;
```
### 3.2 创建评价信息表 (review_info)
```sql
CREATE TABLE review_info (
  `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT 
  COMMENT '主键ID',
  `create_by` varchar(48) NOT NULL DEFAULT '' COMMENT 
  '创建者',
  `update_by` varchar(48) NOT NULL DEFAULT '' COMMENT 
  '更新者',
  `create_at` timestamp NOT NULL DEFAULT 
  CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT 
  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
  COMMENT '更新时间',
  `delete_at` timestamp COMMENT '删除时间',
  `version` int(10) unsigned NOT NULL DEFAULT '0' 
  COMMENT '版本号',
  `review_id` bigint(32) NOT NULL DEFAULT '0' COMMENT 
  '评价ID',
  `content` varchar(512) NOT NULL COMMENT '评价内容',
  `score` tinyint(4) NOT NULL DEFAULT '0' COMMENT '评
  分',
  `service_score` tinyint(4) NOT NULL DEFAULT '0' 
  COMMENT '服务评分',
  `express_score` tinyint(4) NOT NULL DEFAULT '0' 
  COMMENT '物流评分',
  `has_media` tinyint(4) NOT NULL DEFAULT '0' COMMENT 
  '是否有媒体文件',
  `order_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '
  订单ID',
  `sku_id` bigint(32) NOT NULL DEFAULT '0' COMMENT 
  'SKU ID',
  `spu_id` bigint(32) NOT NULL DEFAULT '0' COMMENT 
  'SPU ID',
  `store_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '
  店铺ID',
  `user_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '
  用户ID',
  `anonymous` tinyint(4) NOT NULL DEFAULT '0' COMMENT 
  '是否匿名',
  `tags` varchar(1024) NOT NULL DEFAULT '' COMMENT '标
  签JSON',
  `pic_info` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '图片信息',
  `video_info` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '视频信息',
  `status` tinyint(4) NOT NULL DEFAULT '10' COMMENT '
  状态:10待审核;20已通过;30已拒绝;40已删除',
  `is_default` tinyint(4) NOT NULL DEFAULT '0' 
  COMMENT '是否默认评价',
  `has_reply` tinyint(4) NOT NULL DEFAULT '0' COMMENT 
  '是否有回复:0无;1有',
  `op_reason` varchar(512) NOT NULL DEFAULT '' 
  COMMENT '操作原因',
  `op_remarks` varchar(512) NOT NULL DEFAULT '' 
  COMMENT '操作备注',
  `op_user` varchar(64) NOT NULL DEFAULT '' COMMENT '
  操作用户',
  `goods_snapshoot` varchar(2048) NOT NULL DEFAULT '' 
  COMMENT '商品快照',
  `ext_json` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '扩展JSON',
  `ctrl_json` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '控制JSON',
  PRIMARY KEY (`id`),
  KEY `idx_delete_at` (`delete_at`) COMMENT '删除时间索
  引',
  UNIQUE KEY `uk_review_id` (`review_id`) COMMENT '评
  价ID唯一索引',
  KEY `idx_order_id` (`order_id`) COMMENT '订单ID索引',
  KEY `idx_user_id` (`user_id`) COMMENT '用户ID索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价
信息表';
```
### 3.3 创建评价回复表 (review_reply_info)
```sql
CREATE TABLE review_reply_info (
  `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT 
  COMMENT '主键ID',
  `create_by` varchar(48) NOT NULL DEFAULT '' COMMENT 
  '创建者',
  `update_by` varchar(48) NOT NULL DEFAULT '' COMMENT 
  '更新者',
  `create_at` timestamp NOT NULL DEFAULT 
  CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT 
  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
  COMMENT '更新时间',
  `delete_at` timestamp COMMENT '删除时间',
  `version` int(10) unsigned NOT NULL DEFAULT '0' 
  COMMENT '版本号',
  `reply_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '
  回复ID',
  `review_id` bigint(32) NOT NULL DEFAULT '0' COMMENT 
  '评价ID',
  `content` varchar(512) NOT NULL COMMENT '回复内容',
  `pic_info` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '图片信息',
  `video_info` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '视频信息',
  `reply_type` tinyint(4) NOT NULL DEFAULT '0' 
  COMMENT '回复类型',
  `store_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '
  店铺ID',
  `status` tinyint(4) NOT NULL DEFAULT '10' COMMENT '
  状态',
  `ext_json` varchar(1024) NOT NULL DEFAULT '' 
  COMMENT '扩展JSON',
  PRIMARY KEY (`id`),
  KEY `idx_delete_at` (`delete_at`) COMMENT '删除时间索
  引',
  UNIQUE KEY `uk_reply_id` (`reply_id`) COMMENT '回复
  ID唯一索引',
  KEY `idx_review_id` (`review_id`) COMMENT '评价ID索引
  '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价
回复表';
```
### 3.4 创建评价回复表 (review_appeal_info)
```sql
CREATE TABLE review_appeal_info (
  `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_by` varchar(48) NOT NULL DEFAULT '' COMMENT '创建方标识',
  `update_by` varchar(48) NOT NULL DEFAULT '' COMMENT '更新方标识',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` timestamp COMMENT '逻辑删除标记',
  `version` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁标记',
  `appeal_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '回复id',
  `review_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '评价id',
  `store_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '店铺id',
  `status` tinyint(4) NOT NULL DEFAULT '10' COMMENT '状态:10待审核; 20申诉通过; 30申诉驳回',
  `reason` varchar(255) NOT NULL COMMENT '申诉原因类别',
  `content` varchar(255) NOT NULL COMMENT '申诉内容描述',
  `pic_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息: 图片',
  `video_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息: 视频',
  `op_remarks` varchar(512) NOT NULL DEFAULT '' COMMENT '运营备注',
  `op_user` varchar(64) NOT NULL DEFAULT '' COMMENT '运营者标识',
  `ext_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '信息扩展',
  `ctrl_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '控制扩展',
  PRIMARY KEY (`id`),
  KEY `idx_delete_at` (`delete_at`) COMMENT '逻辑删除索引',
  KEY `idx_appeal_id` (`appeal_id`) COMMENT '申诉id索引',
  UNIQUE KEY `uk_review_id` (`review_id`) COMMENT '评价id索引',
  KEY `idx_store_id` (`store_id`) COMMENT '店铺id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价商家申诉表';
```
## 4. 验证表创建
### 4.1 查看数据库
```sql
SHOW DATABASES;
```
### 4.2 查看表结构
```sql
USE review_service;
SHOW TABLES;
DESC review_info;
DESC review_reply_info;
```
## 5. 退出 MySQL
```sql
exit
```
## 6. 注意事项
1. 端口配置 ：MySQL 容器使用端口 3307，请确保应用程序配置文件中的数据库端口设置为 3307
2. 字符集 ：数据库使用 utf8mb4 字符集，支持完整的 Unicode 字符
3. 密码安全 ：生产环境中请使用更安全的密码
4. 数据持久化 ：如需数据持久化，建议添加数据卷映射参数 -v /path/to/data:/var/lib/mysql
## 7. 常用管理命令
# 查看容器状态
```bash
docker ps -a
```
# 停止容器
```bash
docker stop mysql
```
# 启动容器
```bash
docker start mysql
```
# 重启容器
```bash
docker restart mysql
```
# 查看容器日志
```bash
docker logs mysql
```