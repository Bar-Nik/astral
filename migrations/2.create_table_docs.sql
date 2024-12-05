-- up
create table docs
(
    id         uuid      not null,
    owner_id   uuid      not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),

    foreign key (owner_id) references users on delete cascade,
    unique (id),
    unique (owner_id, id)
);

alter table users
    add column current_docs_id uuid;

-- down
alter table users
drop
column current_docs_id;

drop table docs;