# Write your MySQL query statement below
/*
SELECT a.name AS Employee
FROM Employee AS a, Employee AS b
WHERE a.managerId = b.id AND a.salary > b.salary
*/

SELECT
    a.name AS Employee
FROM Employee AS a JOIN Employee as b 
    ON a.managerId = b.id 
    AND a.salary > b.salary
