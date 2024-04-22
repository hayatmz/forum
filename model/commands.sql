table users => id PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, username  VARCHAR(25) NOT NULL , password  VARCHAR(40) NOT NULL, email VARCHAR(320) NOT NULL UNIQUE

table posts => id PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, title VARCHAR(100) , content VARCHAR(10000), user (foreign key), likes int, dislikes int, date datetime

table categories => id PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, categorie VARCHAR(30)

table comments => id PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL , user (foreign key), content VARCHAR(5000), post (foreign key), date datetime

table post_categories => id post(foreign key), id categorie (foreign key)


CREATE TABLE `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, username VARCHAR(25) NOT NULL, password VARCHAR(40) NOT NULL, 
email VARCHAR(320) NOT NULL UNIQUE);

CREATE TABLE `posts` (`id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, title VARCHAR(100), 
content VARCHAR(10000), user_id INTEGER, likes int, dislikes int, date datetime DEFAULT current_timestamp, FOREIGN KEY(user_id) REFERENCES users(id));

CREATE TABLE `categories` (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, category VARCHAR(30));

CREATE TABLE `comments` (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, user_id INTEGER, content VARCHAR(5000), post_id INTEGER, 
date datetime DEFAULT current_timestamp, FOREIGN KEY(user_id) REFERENCES users(id), FOREIGN KEY(post_id) REFERENCES posts(id));

CREATE TABLE `post_categories` (post_id INTEGER, category_id, FOREIGN KEY(post_id) REFERENCES posts(id), FOREIGN KEY(category_id) REFERENCES categories(id), 
PRIMARY KEY(post_id, category_id));
