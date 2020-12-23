CREATE TABLE `stream_sequence_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `keyword` varchar(100) NOT NULL,
  `media` varchar(20) NOT NULL,
  `type` varchar(25) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `user_account_id` varchar(120) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_account_id` (`user_account_id`),
  CONSTRAINT `stream_sequence_account_FK` FOREIGN KEY (`user_account_id`) REFERENCES `user_account` (`unique_account`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1