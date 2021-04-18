delete from notification
where created < NOW() - INTERVAL '3 DAY';

delete from invite
    using event
where (
        invite.team_id is null
        and invite.user_id is null
    )
   or (
        invite.guest_team_id is null
        and invite.guest_user_id is null
    )
   or (
            event.id = invite.event_id
        and event.state != 'Open');