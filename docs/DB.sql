--! 所有id（uid,vid,cid）均使用雪花算法生成全局唯一id

USE test;

--? 性能优化的目标：至少要达到 range 级别，要求是 ref 级别，如果可以是 const

--^ 用户表

--- 用户基本信息表

CREATE TABLE
    user_base_info(
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `user_id` BIGINT UNSIGNED NOT NULL UNIQUE,
        `username` CHAR(30) NOT NULL,
        `sex` TINYINT NOT NULL,
        `birth` DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `area` INT NOT NULL,
        `user_profile` CHAR(255),
        `user_profile_photo_url` CHAR(255),
        `follow_count` INT NOT NULL DEFAULT 0,
        `fan_count` INT NOT NULL DEFAULT 0,
        -- 时间 
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束 
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_user_id` (`user_id`),
        UNIQUE INDEX `uk_username` (`username`)
    );

--- 用户密码表

CREATE TABLE
    user_password(
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `user_id` BIGINT UNSIGNED UNIQUE,
        `salt` CHAR(64) NOT NULL DEFAULT CONCAT(
            SUBSTRING(MD5(RAND()), 1, 32),
            SUBSTRING(MD5(RAND()), 1, 32)
        ),
        `pwd` CHAR(64) NOT NULL DEFAULT CONCAT(
            SUBSTRING(MD5(RAND()), 1, 32),
            SUBSTRING(MD5(RAND()), 1, 32)
        ),
        -- 时间 
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束 
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_user_id` (`user_id`)
    );

--- 用户关注信息 user_id 用户关注 follow_id 用户

CREATE TABLE
    user_follow_info (
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `user_id` BIGINT UNSIGNED NOT NULL,
        `follow_id` BIGINT UNSIGNED NOT NULL,
        `status` TINYINT NOT NULL DEFAULT 1,
        -- 时间 
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束 
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_follows_id` (`user_id`, `follow_id`),
        UNIQUE INDEX `uk_fans_id` (`follow_id`, `user_id`)
    );

--^ 视频表

--- 视频元信息

CREATE TABLE
    video_metadata (
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `video_id` BIGINT NOT NULL,
        `user_id` BIGINT NOT NULL,
        `cover_url` CHAR(255) NOT NULL,
        `video_url` CHAR(255) NOT NULL,
        `video_intro` CHAR(255),
        `video_type` BIGINT NOT NULL,
        `publish_address` INT NOT NULL DEFAULT 0,
        -- 时间
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束 
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_video_id` (`video_id`),
        UNIQUE INDEX `uk_uid` (`user_id`, `video_id`)
    );

--- 视频点赞

CREATE TABLE
    video_like(
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `user_id` BIGINT NOT NULL,
        `video_id` BIGINT NOT NULL,
        `status` TINYINT NOT NULL DEFAULT 1,
        -- 时间
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束 
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_user_id` (`user_id`, `video_id`),
        UNIQUE INDEX `uk_video_id` (`video_id`, `user_id`)
    );

--- 用户收藏 --^ 本质和点赞表一样

CREATE TABLE
    video_collection(
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `user_id` BIGINT NOT NULL,
        `video_id` BIGINT NOT NULL,
        `status` TINYINT NOT NULL DEFAULT 1,
        -- 时间
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束 
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_user_id` (`user_id`, `video_id`),
        UNIQUE INDEX `uk_video_id` (`video_id`, `user_id`)
    );

--^ 互动表

--- 视频评论

CREATE TABLE
    video_comment (
        `id` BIGINT UNSIGNED NOT NULL UNIQUE PRIMARY KEY,
        -- 表项
        `comment_id` BIGINT NOT NULL,
        `pcomment_id` BIGINT NOT NULL,
        `video_id` BIGINT NOT NULL,
        `user_id` BIGINT NOT NULL,
        `comment_text` TEXT,
        -- 时间
        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        -- 约束
        -- CONSTRAINT `pk_id` PRIMARY KEY(id),
        UNIQUE INDEX `uk_comment_id` (`comment_id`),
        UNIQUE INDEX `uk_video_comment` (`video_id`, `comment_id`)
    );

-- ent new UserBaseInfo

-- ent new UserFollowInfo

-- ent new UserPassword

-- ent new VideoComment

-- ent new VideoCollection

-- ent new VideoLike

-- ent new VideoMetadata