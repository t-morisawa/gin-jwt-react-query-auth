USE database;
DROP DATABASE IF EXISTS users;
CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(45) NULL,
  last_name VARCHAR(45) NULL,
  email VARCHAR(45) NOT NULL,
    data_created DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    status TINYINT(1) NOT NULL,
  PRIMARY KEY (id),
    UNIQUE INDEX email_UNIQUE (email ASC));
