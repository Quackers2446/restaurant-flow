alter table `Restaurant` modify column `google_restaurant_id` int unique;
alter table `OpeningHours` add unique `unique_opening_hours`(`google_restaurant_id`, `type`);
alter table `OpeningPeriod` add unique `unique_opening_period`(`opening_hours_id`, `open_day`);
alter table `OpeningPeriod` modify column `open_day` tinyint not null check (`open_day` >= 0 and `open_day` <= 6); -- 0 = Sunday, 6 = Saturday
