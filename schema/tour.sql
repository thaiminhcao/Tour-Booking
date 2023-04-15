CREATE TABLE `tour` (
  `tour_id` integer,
  `tour_name` varchar(255),
  `desciption` varchar(255),
  `price` integer,
  `start_date` timestamp,
  `end_date` timestamp,
  `location` varchar(255),
  PRIMARY KEY(tour_id)
);