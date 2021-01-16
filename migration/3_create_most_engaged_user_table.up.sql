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


CREATE TABLE `oauth_twitter` (
 `user_id` varchar(50) NOT NULL,
 `name` varchar(50) NOT NULL,
 `username` varchar(50) NOT NULL,
 `profile_image_url` text NOT NULL,
 `access_token` text NOT NULL,
 `access_token_secret` text NOT NULL,
 `user_account_id` varchar(120) NOT NULL,
 `created_at` datetime DEFAULT NULL,
 PRIMARY KEY (`user_id`),
 KEY `user_account_id` (`user_account_id`),
 CONSTRAINT `oauth_twitter_FK` FOREIGN KEY (`user_account_id`) REFERENCES `user_account` (`unique_account`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `twitter_statuses` (
    `statuses_id` varchar(30) NOT NULL,
    `text` text,
    `lang` varchar(20) DEFAULT NULL,
    `permalink` text,
    `created_at` datetime DEFAULT NULL,
    `user_id` varchar(50) NOT NULL,
    PRIMARY KEY (`statuses_id`),
    KEY `user_id` (`user_id`),
    CONSTRAINT `twitter_statuses_FK` FOREIGN KEY (`user_id`) REFERENCES `oauth_twitter` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1