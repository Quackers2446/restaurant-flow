alter table `restaurant` modify column `google_restaurant_id` int unique;
alter table `opening_hours` add unique `unique_opening_hours`(`google_restaurant_id`, `type`);
alter table `opening_period` add unique `unique_opening_period`(`opening_hours_id`, `open_day`);
alter table `opening_period` modify column `open_day` tinyint not null check (`open_day` >= 0 and `open_day` <= 6); -- 0 = Sunday, 6 = Saturday
