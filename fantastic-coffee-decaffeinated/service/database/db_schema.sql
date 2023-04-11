PRAGMA foreign_keys = ON;

CREATE TABLE User
(
Id_user TEXT, 
Nickname TEXT CHECK(length(Nickname) >= 3 AND length(Nickname) <= 13) PRIMARY KEY

);

CREATE TABLE Photo
(
Id_photo TEXT PRIMARY KEY,
User TEXT NOT NULL,
Date TEXT NOT NULL,
Time TEXT NOT NULL,
LocalPath TEXT NOT NULL,
FOREIGN KEY (User)
	REFERENCES User (Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE
);

CREATE TABLE Like
(
User TEXT,
Photo INTEGER,
FOREIGN KEY (User)
	REFERENCES User(Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
FOREIGN KEY (Photo)
	REFERENCES Photo(Id_photo)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
PRIMARY KEY(User, Photo)
);

CREATE TABLE Comment
(
Id_comment INTEGER PRIMARY KEY AUTOINCREMENT,
User TEXT NOT NULL,
Photo TEXT NOT NULL,
Content TEXT NOT NULL CHECK(length(Content) <= 100),
FOREIGN KEY (User)
	REFERENCES User(Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
FOREIGN KEY (Photo)
	REFERENCES Photo(Id_photo)
	ON UPDATE CASCADE
	ON DELETE CASCADE
);

CREATE TABLE Ban
(
Banner TEXT,
Banned TEXT,
FOREIGN KEY (Banner)
	REFERENCES User(Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
FOREIGN KEY (Banned)
	REFERENCES User(Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
PRIMARY KEY (Banner, Banned)
);

CREATE TABLE Follow
(
Follower TEXT,
Followed TEXT,
FOREIGN KEY (Follower)
	REFERENCES User(Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
FOREIGN KEY (Followed)
	REFERENCES User(Nickname)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
PRIMARY KEY (Follower, Followed)
);



