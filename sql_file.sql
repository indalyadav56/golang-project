CREATE TABLE users (
  id INT PRIMARY KEY, 
  name VARCHAR(50)
);

INSERT INTO users VALUES (1, 'John');
INSERT INTO users VALUES (2, 'Jane');

SELECT * FROM users;