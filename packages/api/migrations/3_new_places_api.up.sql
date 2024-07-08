alter table `google_restaurant` add `supports_takeout` boolean default null;

-- Not necessary, just a style thing
alter table `google_restaurant` modify column `serves_breakfast` boolean default null;
alter table `google_restaurant` modify column `serves_brunch` boolean default null;
alter table `google_restaurant` modify column `serves_dinner` boolean default null;
alter table `google_restaurant` modify column `serves_lunch` boolean default null;
alter table `google_restaurant` modify column `serves_vegetarian_food` boolean default null;
alter table `google_restaurant` modify column `serves_wine` boolean default null;
alter table `google_restaurant` add column `serves_beer` boolean default null;
alter table `google_restaurant` add column `serves_cocktails` boolean default null;
alter table `google_restaurant` add column `serves_coffee` boolean default null;
alter table `google_restaurant` add column `serves_dessert` boolean default null;

alter table `google_restaurant` add column `good_for_groups` boolean default null;
alter table `google_restaurant` add column `good_for_watching_sports` boolean default null;
alter table `google_restaurant` add column `has_outdoor_seating` boolean default null;
alter table `google_restaurant` add column `has_restroom` boolean default null;

alter table `google_restaurant` add column `accepts_credit_cards` boolean default null;
alter table `google_restaurant` add column `accepts_debit_cards` boolean default null;
alter table `google_restaurant` add column `accepts_cash_only` boolean default null;
alter table `google_restaurant` add column `accepts_nfc` boolean default null;

alter table `google_restaurant` modify column `wheelchair_accessible_entrance` boolean default null;
alter table `google_restaurant` add column `wheelchair_accessible_seating` boolean default null;
