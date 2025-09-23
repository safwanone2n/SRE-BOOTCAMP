-- This is a sample migration.

create table users (
  id serial primary key,
  first_name varchar not null,
  last_name varchar not null,
  email text not null,
  phone_number varchar 12
);

---- create above / drop below ----

drop table user;
