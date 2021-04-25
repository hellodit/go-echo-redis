# Go+Echo+Redis
Caching is a method that is often used to improve application performance, cache is used as secondary data storage to store data that is frequently accessed by clients, the cache uses computer memory as a medium for storing data so that it has a very fast data reading speed.

### Context 
This repository made for medium article here (comming soon)

### Prerequisite
* GO 1.15
* Gorm
* Echo v4
* Logrus
* Redis

### Database schema
Database migration also availabe in `db > migration` folder
``` sql 
CREATE TABLE `authors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `birthdate` date NOT NULL,
  `added` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

//Tabel Post
CREATE TABLE `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(500) COLLATE utf8_unicode_ci NOT NULL,
  `content` text COLLATE utf8_unicode_ci NOT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11001 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
```