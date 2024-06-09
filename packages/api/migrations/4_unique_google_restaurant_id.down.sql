alter table `Restaurant` modify column `google_restaurant_id` int;
alter table `OpeningHours` drop index `unique_opening_hours`;
alter table `OpeningPeriod` drop index `unique_opening_period`;
alter table `OpeningPeriod` modify column `open_day` tinyint not null;
