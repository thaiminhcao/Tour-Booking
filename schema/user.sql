CREATE TABLE `users` (
  `user_id` integer,
  `username` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `created_at` timestamp,
  `gender` varchar(255),
  `dob` timestamp,
  PRIMARY KEY(user_id)
);
