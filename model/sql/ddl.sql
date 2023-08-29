-- definition
CREATE TABLE `user`
(
    `id`           BIGINT(11) AUTO_INCREMENT COMMENT 'ID',
    `uid`          BIGINT(11) NOT NULL DEFAULT 0 COMMENT 'uid',
    `name`         VARCHAR(32)  NOT NULL DEFAULT '' COMMENT '收款人姓名',
    `password`     VARCHAR(255)  NOT NULL DEFAULT '' COMMENT '密码',
    `phone`        VARCHAR(32) NOT NULL DEFAULT '' COMMENT '电话',
    `referrer_uid` BIGINT(11) NOT NULL DEFAULT 0 COMMENT '推荐人uid',
    `type`         INT(8) NOT NULL DEFAULT 0 COMMENT '用户类型 0-普通｜1-销售｜2-销售经理｜99-超级管理',
    `coin`         BIGINT(11) NOT NULL DEFAULT 0 COMMENT '积分',
    `vip_type`     INT(8) NOT NULL DEFAULT 0 COMMENT 'vip类型 0-非会员｜1-月会员｜2-年会员｜99-终身会员',
    `vip_end_at`   DATETIME     NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT 'vip结束时间',
    `created_at`   TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
    `updated_at`   TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) COMMENT '更新时间',
    `deleted_at`   DATETIME NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid` (`uid`),
    UNIQUE KEY `uniq_phone` (`phone`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

CREATE TABLE `order`
(
    `id`         BIGINT(11) AUTO_INCREMENT COMMENT 'ID',
    `order_no`   VARCHAR(11)  NOT NULL DEFAULT '0' COMMENT '订单id',
    `uid`        BIGINT(11) NOT NULL DEFAULT 0 COMMENT 'uid',
    `sku`        INT(8) NOT NULL DEFAULT 0 COMMENT 'sku',
    `amount`     BIGINT(20) NOT NULL DEFAULT 0 COMMENT '金额',
    `status`     INT(8) NOT NULL DEFAULT 0 COMMENT '订单状态 0-初始化｜100-待支付｜200-付款成功｜400-订单关闭',
    `pay_url`    VARCHAR(255) NOT NULL DEFAULT '' COMMENT '支付链接',
    `pay_time`   TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '付款时间',
    `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
    `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) COMMENT '更新时间',
    `deleted_at` DATETIME NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_order_no` (`order_no`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='订单表';

