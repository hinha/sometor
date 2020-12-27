CREATE TABLE `most_engaged_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL,
  `total_engagement` int(50) NOT NULL,
  `media` varchar(20) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `stream_sequence_account_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `stream_sequence_account_id` (`stream_sequence_account_id`),
  CONSTRAINT `most_engaged_user_FK` FOREIGN KEY (`stream_sequence_account_id`) REFERENCES `stream_sequence_account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `most_mention_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL,
  `count` int(50) NOT NULL,
  `media` varchar(20) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `stream_sequence_account_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `stream_sequence_account_id` (`stream_sequence_account_id`),
  CONSTRAINT `most_engaged_user_1_FK` FOREIGN KEY (`stream_sequence_account_id`) REFERENCES `stream_sequence_account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `most_active_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL,
  `reply_count` int(11) NOT NULL,
  `retweet_count` int(11) NOT NULL,
  `like_count` int(11) NOT NULL,
  `quote_count` int(11) NOT NULL,
  `total_engagement` int(50) NOT NULL,
  `media` varchar(20) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `stream_sequence_account_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `stream_sequence_account_id` (`stream_sequence_account_id`),
  CONSTRAINT `most_engaged_user_FK` FOREIGN KEY (`stream_sequence_account_id`) REFERENCES `stream_sequence_account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1