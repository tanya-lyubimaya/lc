SELECT DISTINCT author_id AS id FROM Views WHERE author_id = viewer_id ORDER BY id;
SELECT author_id AS id FROM Views WHERE author_id = viewer_id GROUP BY id ORDER BY id;
SELECT author_id AS id FROM Views GROUP BY id HAVING SUM(CASE WHEN author_id = viewer_id THEN 1 ELSE 0 END) > 0 ORDER BY id;