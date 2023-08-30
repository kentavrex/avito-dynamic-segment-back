CREATE TABLE users
(
    id SERIAL PRIMARY KEY
);

CREATE TABLE segments
(
    id SERIAL PRIMARY KEY,
    slug varchar(255)
);

CREATE TABLE users_segments
(
    id SERIAL PRIMARY KEY,
    user_id int references users (id) on delete cascade not null,
    segment_id int references segments on delete cascade not null
);