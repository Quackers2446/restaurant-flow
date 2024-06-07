CREATE TABLE `User` (
    `user_id` binary(16) PRIMARY KEY,
    `name` varchar(32) NOT NULL,
    `username` varchar(32) NOT NULL,
    `email` varchar(320) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `Restaurant` (
    `restaurant_id` int PRIMARY KEY AUTO_INCREMENT,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `google_restaurant_id` int
);

CREATE TABLE `GoogleRestaurant` (
    `google_restaurant_id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(64) NOT NULL,
    `description` text,
    `address` varchar(512) NOT NULL,
    `phone` varchar(32) NOT NULL,
    `website` varchar(256),
    `google_url` varchar(256),
    `avg_rating` tinyint,
    `business_status` ENUM ('Operational', 'ClosedTemporarily', 'ClosedPermanently'),
    `price_level` ENUM ('Free', 'Inexpensive', 'Moderate', 'Expensive', 'VeryExpensive'),
    `supports_curbside_pickup` boolean DEFAULT null,
    `supports_delivery` boolean DEFAULT null,
    `supports_dine_in` boolean DEFAULT null,
    `supports_reservations` boolean DEFAULT null,
    `serves_breakfast` boolean,
    `serves_brunch` boolean,
    `serves_dinner` boolean,
    `serves_lunch` boolean,
    `serves_vegetarian_food` boolean,
    `serves_wine` boolean,
    `wheelchair_accessible_entrance` boolean,
    `place_id` varchar(512) UNIQUE NOT NULL,
    `coords` point NOT NULL,
    `northeast` point NOT NULL,
    `southwest` point NOT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `restaurant_id` int
);

CREATE TABLE `Tag` (
    `tag_id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(64),
    `restaurant_id` int
);

CREATE TABLE `Review` (
    `review_id` int PRIMARY KEY AUTO_INCREMENT,
    `rating` int NOT NULL,
    `comments` text,
    `restaurant_id` int,
    `user_id` binary(16),
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `OpeningHours` (
    `opening_hours_id` int PRIMARY KEY AUTO_INCREMENT,
    `type` ENUM ('Main', 'DriveThrough', 'HappyHour', 'Delivery', 'Takeout', 'Kitchen', 'Breakfast', 'Lunch', 'Dinner', 'Brunch', 'Pickup', 'SeniorHours') NOT NULL,
    `google_restaurant_id` int
);

CREATE TABLE `OpeningPeriod` (
    `opening_period_id` int PRIMARY KEY AUTO_INCREMENT,
    `open_day` tinyint NOT NULL,
    `open_time` char(4) NOT NULL,
    `close_day` tinyint,
    `close_time` char(4),
    `opening_hours_id` int
);

CREATE INDEX `GoogleRestaurant_index_0` ON `GoogleRestaurant` (`place_id`);

ALTER TABLE `Restaurant` ADD FOREIGN KEY (`google_restaurant_id`) REFERENCES `GoogleRestaurant` (`google_restaurant_id`);

ALTER TABLE `GoogleRestaurant` ADD FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurant` (`restaurant_id`);

ALTER TABLE `Tag` ADD FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurant` (`restaurant_id`);

ALTER TABLE `Review` ADD FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurant` (`restaurant_id`);

ALTER TABLE `Review` ADD FOREIGN KEY (`user_id`) REFERENCES `User` (`user_id`);

ALTER TABLE `OpeningHours` ADD FOREIGN KEY (`google_restaurant_id`) REFERENCES `GoogleRestaurant` (`google_restaurant_id`);

ALTER TABLE `OpeningPeriod` ADD FOREIGN KEY (`opening_hours_id`) REFERENCES `OpeningHours` (`opening_hours_id`);
