create table `t_article`(
    `id`: BIGINT NOT NULL,
    `group_id`: TINYINT UNSIGNED NOT NULL,
    `author_id`: BIGINT NOT NULL,
    `title`: VARCHAR(20) NOT NULL,
    `abstract`: VARCHAR(100) NOT NULL,
    `content`: TEXT NOT NULL,
    `create_time`: BIGINT NOT NULL,
    `last_update_time`: BIGINT NOT NULL,
    PRIMARY KEY(`id`),
    KEY index_group_id(`group_id`),
    KEY index_author_id(`author_id`),
)

create table `t_group`(
    `id`: BIGINT NOT NULL,
    `name`: VARCHAR(20) NOT NULL,
    `priority`: INT NOT NULL,
    PRIMARY KEY(`id`),
)

create table `t_user`(
    `id`: BIGINT NOT NULL,
    `username`: VARCHAR(100) NOT NULL,
    `icon_url`: VARCHAR(100) NOT NULL,
    `password`: VARCHAR(100) NOT NULL,
    `is_admin`: TINYINT(1) NOT NULL,
    `create_time`: BIGINT NOT NULL,
    PRIMARY KEY(`id`),
)