# Write your MySQL query statement below

SELECT a.id FROM Weather AS a 
JOIN Weather AS b ON datediff(a.recordDate, b.recordDate) = 1
AND a.temperature > b.temperature