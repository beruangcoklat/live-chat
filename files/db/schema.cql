drop keyspace if exists live_chat;
create keyspace live_chat with replication = {'class' : 'SimpleStrategy', 'replication_factor':1};
use live_chat;

drop table if exists channel;
create table channel (
    id uuid,
    name text,
    created_at bigint,
    primary key(id, created_at)
);

drop table if exists chat;
create table chat (
    channel_id uuid,
    sender text,
    message text,
    created_at bigint,
    primary key(channel_id, created_at)
);

insert into channel (id, name, created_at) values (uuid(), 'channel 1',  unixTimestampOf(now()));
insert into channel (id, name, created_at) values (uuid(), 'channel 2',  unixTimestampOf(now()));
