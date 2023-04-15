CREATE TABLE `payment` (
  `payment_id` integer PRIMARY KEY,
  `payment_method` integer,
  `deposit` integer,
  `amount_owed` integer,
  PRIMARY KEY(payment_id)
);




