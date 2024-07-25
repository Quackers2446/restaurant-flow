create table `user` (
    `user_id` binary(16) primary key,
    `email` varchar(320) not null unique,
    `user_id_text` varchar(36) generated always as (
        insert(
            insert(
                insert(
                    insert(hex(user_id),9,0,'-'),
                    14,0,'-'
                ),
                19,0,'-'
            ),
            24,0,'-'
        )
    ) virtual
);

create table `user_auth` (
    `user_id` binary(16) primary key,
    `password_hash` char(60) character set latin1 collate latin1_bin not null
);

alter table `user_auth` add foreign key (`user_id`) references `user`(`user_id`);

create table `session` (
    `session_id` int primary key auto_increment,
    `user_id` binary(16) not null,
    `ip_addr` char(45) not null,
    `user_agent` text not null,
    `created_at` timestamp not null default current_timestamp,
    `last_used_at` timestamp not null default current_timestamp on update current_timestamp,
    `expires_at` timestamp not null,
    `refresh_key` char(64) not null unique,

    index `session_user_id` (`user_id`)
);

alter table `session` add foreign key (`user_id`) references `user`(`user_id`);
