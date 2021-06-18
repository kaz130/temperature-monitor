set client_encoding = 'UTF8';

CREATE TABLE chat (
  id          SERIAL NOT NULL,
  temperature float8,
  pressure    float8,
  humidity    float8,
  time        timestamp with time zone,
  PRIMARY KEY (id)
);
