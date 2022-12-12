DELETE FROM Person WHERE id NOT IN (SELECT * FROM (SELECT MIN(id) FROM Person GROUP BY email) AS iid);
DELETE p2 FROM Person p1, Person p2 WHERE p1.email = p2.email AND p1.id < p2.id;
DELETE p2 FROM Person p1 INNER JOIN Person p2 ON p1.email = p2.email WHERE p1.id < p2.id;
DELETE p2 FROM Person p1 INNER JOIN Person p2 ON p1.email = p2.email AND p1.id < p2.id;