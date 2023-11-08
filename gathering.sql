CREATE DATABASE gathering; 

USE gathering;

CREATE TABLE `attendees` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `attendee_id` longtext,
  `member_id` bigint(20) DEFAULT NULL,
  `gathering_id` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_members_id` (`member_id`),
  KEY `fk_gatheringss_id` (`gathering_id`),
  CONSTRAINT `fk_gatheringss_id` FOREIGN KEY (`gathering_id`) REFERENCES `gatherings` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_members_id` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- gathering.gathering_types definition

CREATE TABLE `gathering_types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gathering_id` bigint(20) NOT NULL,
  `type` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gathering_id` (`gathering_id`),
  CONSTRAINT `fk_gathering_id` FOREIGN KEY (`gathering_id`) REFERENCES `gatherings` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- gathering.gatherings definition

CREATE TABLE `gatherings` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gathering_id` longtext,
  `creator` longtext NOT NULL,
  `scheduled_at` datetime(3) NOT NULL,
  `name` longtext NOT NULL,
  `location` longtext NOT NULL,
  `status` tinyint(1) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- gathering.invitation_statuses definition

CREATE TABLE `invitation_statuses` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `invitation_id` bigint(20) DEFAULT NULL,
  `status` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_invitation_id` (`invitation_id`),
  CONSTRAINT `fk_invitation_id` FOREIGN KEY (`invitation_id`) REFERENCES `invitations` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- gathering.invitations definition

CREATE TABLE `invitations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `invitation_id` longtext,
  `member_id` bigint(20) DEFAULT NULL,
  `gathering_id` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_member_id` (`member_id`),
  KEY `fk_gatherings_id` (`gathering_id`),
  CONSTRAINT `fk_gatherings_id` FOREIGN KEY (`gathering_id`) REFERENCES `gatherings` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_member_id` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- gathering.members definition

CREATE TABLE `members` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `member_id` longtext,
  `first_name` longtext,
  `last_name` longtext,
  `email` longtext,
  `address` longtext,
  `phone_number` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;