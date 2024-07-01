# Write your MySQL query statement below
SELECT 
    a.player_id, MIN(a.event_date) AS first_login 
FROM 
    activity as a 
GROUP BY
    a.player_id
