SELECT player_id, MIN(event_date) AS first_login FROM Activity GROUP BY player_id;
/* TODO: Add window functions solution */