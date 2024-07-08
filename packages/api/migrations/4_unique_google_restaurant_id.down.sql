alter table `restaurant` modify column `google_restaurant_id` int;
alter table `opening_hours` drop index `unique_opening_hours`;
alter table `opening_period` drop index `unique_opening_period`;
alter table `opening_period` modify column `open_day` tinyint not null;
