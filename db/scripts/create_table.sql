DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         SERIAL PRIMARY KEY,
  title      VARCHAR (128) UNIQUE NOT NULL,
  artist     VARCHAR (255) NOT NULL,
  rating      INT NOT NULL
);
