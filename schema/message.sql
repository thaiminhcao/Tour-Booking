CREATE TABLE `message_information` (
  `message_id` integer PRIMARY KEY,
  `title` varchar(255),
  `user_id` integer,
  `tour_id` integer,
  `payment_id` integer,
  `created_at` timestamp,
  PRIMARY KEY(message_id)
);
