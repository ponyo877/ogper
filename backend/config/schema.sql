CREATE DATABASE ogp;
CREATE TABLE sites (
  id           INT           NOT NULL AUTO_INCREMENT,
  hash         VARCHAR(7)    NOT NULL,
  title        VARCHAR(255)  NOT NULL,
  descriptio   VARCHAR(255)  NOT NULL,
  name         VARCHAR(255)  NOT NULL,
  site_url     VARCHAR(8000) NOT NULL,
  image_url    VARCHAR(255)  NOT NULL,
  user_hash    VARCHAR(26)   NOT NULL,
  published_at TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_at   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY hash (hash),
  INDEX (user_hash, published_at)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;