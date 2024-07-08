alter table `google_restaurant` rename column `viewport_high` to `northeast`;
alter table `google_restaurant` rename column `viewport_low` to `southwest`;

alter table `google_restaurant` drop column `supports_takeout`;

alter table `google_restaurant` modify column `serves_breakfast` boolean;
alter table `google_restaurant` modify column `serves_brunch` boolean;
alter table `google_restaurant` modify column `serves_dinner` boolean;
alter table `google_restaurant` modify column `serves_lunch` boolean;
alter table `google_restaurant` modify column `serves_vegetarian_food` boolean;
alter table `google_restaurant` modify column `serves_wine` boolean;
alter table `google_restaurant` drop column `serves_beer`;
alter table `google_restaurant` drop column `serves_cocktails`;
alter table `google_restaurant` drop column `serves_coffee`;
alter table `google_restaurant` drop column `serves_dessert`;

alter table `google_restaurant` drop column `good_for_groups`;
alter table `google_restaurant` drop column `good_for_watching_sports`;
alter table `google_restaurant` drop column `has_outdoor_seating`;
alter table `google_restaurant` drop column `has_restroom`;

alter table `google_restaurant` drop column `accepts_credit_cards`;
alter table `google_restaurant` drop column `accepts_debit_cards`;
alter table `google_restaurant` drop column `accepts_cash_only`;
alter table `google_restaurant` drop column `accepts_nfc`;

alter table `google_restaurant` modify column `wheelchair_accessible_entrance` boolean;
alter table `google_restaurant` drop column `wheelchair_accessible_seating`;
