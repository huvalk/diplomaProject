-- pg_ctl -D /usr/local/var/postgres start
-- pg_ctl -D /usr/local/var/postgres stop
create table users
(
    id        bigserial primary key,
    firstname varchar(80)        not null,
    lastName  varchar(80)        not null,
    email     varchar(80) unique not null
);

create table event
(
    id          bigserial primary key,
    name        varchar(80) not null,
    description varchar(80) not null,
    founder     integer REFERENCES users (id),
    date_start  timestamp,
    date_end    timestamp,
    state       varchar(80),
    place       varchar(80)
-- participants_count integer
);
-- insert into event values(default,'event1','descr1',1,'2021-02-25 10:23:54+02','2021-02-25 15:23:54+02','place1');

create table event_users
(
    event_id integer REFERENCES event (id),
    user_id  integer REFERENCES users (id),
    CONSTRAINT uniq_pair UNIQUE (event_id, user_id)
);

create table feed
(
    id    bigserial primary key,
    event integer REFERENCES event (id)
);

create table feed_users
(
    feed_id integer REFERENCES feed (id),
    user_id integer REFERENCES users (id),
    CONSTRAINT uniq_pair2 UNIQUE (feed_id, user_id)

);

create table team
(
    id    bigserial primary key,
    name  varchar(80) not null,
    event integer REFERENCES event (id)
);

create table team_users
(
    team_id integer REFERENCES team (id),
    user_id integer REFERENCES users (id),
    CONSTRAINT uniq_pair3 UNIQUE (team_id, user_id)
);

create table notification
(
    type integer not null default 0,
    user_id integer REFERENCES users (id),
    message varchar(320) not null default '',
    created timestamp not null,
    watched bool not null default false
);



