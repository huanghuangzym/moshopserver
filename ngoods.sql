SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0; 


DROP TABLE IF EXISTS `nideshop_goods`;
CREATE TABLE `nideshop_goods` (
  `id` int(11) unsigned NOT NULL,
  `name` varchar(120) NOT NULL DEFAULT '',
  `goods_desc` text,
  `primary_pic_url` varchar(255) NOT NULL COMMENT '商品主图',
  `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价格',
  `ku_cun` int(11) unsigned NOT NULL DEFAULT '10' COMMENT '库存',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `nideshop_goods` VALUES (1009027,  '陀螺', '奥迪的陀螺', 'http://yanxuan.nosdn.127.net/35ad21679dbd30a23a8308287ffd4673.jpg', 79.00, 109);
INSERT INTO `nideshop_goods` VALUES (1009023,  '赛车', '奥迪的赛车', 'http://yanxuan.nosdn.127.net/35ad21679dbd30a23a8308287ffd4673.jpg', 81.00, 209);



COMMIT;
