create table `party` (
    `party_id` int primary key auto_increment,
    `description` text,
    `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `restaurant_id` int NOT NULL,
    `max_members` int NOT NULL DEFAULT 4
);

create table `party_members` (
    `party_id` int NOT NULL,
    `user_id` binary(16) NOT NULL,
    `joined_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`party_id`, `user_id`)
);

ALTER TABLE `party` ADD FOREIGN KEY (`restaurant_id`) REFERENCES `restaurant` (`restaurant_id`);
ALTER TABLE `party_members` ADD FOREIGN KEY (`party_id`) REFERENCES `party` (`party_id`);
