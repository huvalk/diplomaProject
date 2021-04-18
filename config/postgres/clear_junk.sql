delete from notification
where created < NOW() - INTERVAL '3 DAY';

delete from invite
where (
    team_id is null
    and user_id is null
)
or (
    guest_team_id is null
    and guest_user_id is null
)