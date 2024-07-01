CREATE TABLE `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, username VARCHAR(25) NOT NULL, password VARCHAR(40) NOT NULL, email VARCHAR(75) NOT NULL UNIQUE);
INSERT INTO users(username, password, email) VALUES ('bryan', 'testMDP', 'bryan@sdoo.com'), ('ismael', 'testimamdp', 'ismae@sdj.com');
INSERT INTO users(username, password, email) VALUES ('bryan', 'testMDP', 'bryan@sdoo.com');



CREATE TABLE `posts` (`id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, title VARCHAR(100), content VARCHAR(10000), categories VARCHAR(30));
POST -> title, content, (user foreign key), likes, dislikes , date,

INSERT INTO posts(title, content, categories) VALUES ('salut', 'tesyo le rap', '["foot", "manga"]');


table foot (title, content)

table manga 

table lecture

table posts
table comms -> post(foreign key)
