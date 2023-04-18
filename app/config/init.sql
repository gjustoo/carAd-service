create table posts (
  id serial not null unique,
  title varchar(64),
  content text,
  primary key(id)
);


insert into posts(title, content)
values
    ('Hello World', 'The obligatory Hello World Post ...'),
    ('Another Post', 'Yet another blog post about something exciting');