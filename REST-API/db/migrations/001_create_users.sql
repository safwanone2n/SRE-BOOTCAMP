-- Migrating table users;

  CREATE TABLE users (
  id serial primary key,
  first_name varchar not null,
  last_name varchar not null,
  email text not null,
  phone_number varchar
);

---- create above / drop below ----

DROP TABLE users;
