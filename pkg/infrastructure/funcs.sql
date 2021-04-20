set schema 'public';

create or replace function inc_event_participants() returns trigger as
$inc_event_participants$
begin
    update event
    set participants_count = participants_count + 1
    where new.event_id = event.id;
    return null;
end;
$inc_event_participants$ language plpgsql;

create or replace function dec_event_participants() returns trigger as
$dec_event_participants$
begin
    update event
    set participants_count = participants_count - 1
    where old.event_id = event.id
      and event.participants_count > 0;
    return null;
end;
$dec_event_participants$ language plpgsql;

create or replace function find_users_team(integer, integer) returns integer
as
'select team_id
 from team_users
          inner join team t on team_users.team_id = t.id
 where team_users.user_id = $1
   and t.event = $2
 union
 select null
 order by team_id
 limit 1;'
    language sql
    immutable
    returns null on null input;

create or replace function find_users_lead_team(integer, integer) returns bigint
as
'select id
 from team t
 where t.lead_id = $1
   and t.event = $2
 union
 select null
 order by id
 limit 1;'
    language sql
    immutable
    returns null on null input;

create trigger added_event_user
    after insert
    on event_users
    for each row
execute procedure inc_event_participants();

create trigger deleted_event_user
    after delete
    on event_users
    for each row
execute procedure dec_event_participants();