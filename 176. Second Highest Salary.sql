SELECT IFNULL((SELECT DISTINCT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 1), NULL) AS SecondHighestSalary;
SELECT MAX(salary) AS SecondHighestSalary FROM Employee WHERE salary NOT IN (SELECT MAX(salary) FROM Employee);