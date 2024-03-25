-- Active: 1711334461218@@127.0.0.1@5432
CREATE TABLE IF NOT EXISTS links (
			id INTEGER PRIMARY KEY,
			url TEXT NOT NULL,
			downloadFlag BOOLEAN DEFAULT true,
			gid TEXT DEFAULT NULL
		);
CREATE INDEX if not EXISTS `url_index` on `links`(`url`);