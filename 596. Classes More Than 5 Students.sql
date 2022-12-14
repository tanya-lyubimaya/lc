SELECT class FROM Courses GROUP BY class HAVING COUNT(student) >=5;
SELECT class FROM (SELECT class, COUNT(student) AS counter FROM Courses GROUP BY class) AS temp WHERE counter >= 5;