
CREATE TABLE `stream_logstash` (
  `id` int AUTO_INCREMENT,
  `users_id` int(25),
  `name` varchar(50) NOT NULL,
  `address` text,
  `phone_number` varchar(15),
  PRIMARY KEY (`id`, `users_id`)
);