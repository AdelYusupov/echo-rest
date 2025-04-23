# echo-rest
docker run --name fitnes -e POSTGRES_PASSWORD=yourpassword -p 5432:5432 -d postgres
```
CREATE DATABASE fitnes;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE measurements (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  weight FLOAT NOT NULL,
  height FLOAT NOT NULL,
  body_fat FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
```
