-- CREATE FUNCTION inserActor() RETURNS void AS '
-- 	DECLARE 
-- 	vAr INTEGER;
-- 	SET vAr= SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names;
-- 	INSERT INTO Names (Id, Family, Given) 
-- 	VALUES  (vAr, "Bor", "Alex");
-- 	INSERT INTO Actors (Id, Nameid, Nationid, Number, Honorar ) 
-- 	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Actors), vAr, 1, 1, 1)
-- ' LANGUAGE SQL;

	INSERT INTO Names (Id, Family, Given) 
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names), 'Doe', 'John');

	INSERT INTO Actors (Id, Nameid, Nationid, Number, Honorar ) 
	VALUES  (
		(SELECT COALESCE(MAX(Id), 0) + 1 FROM  Actors), 
		(SELECT id FROM  Names WHERE (Family LIKE 'Doe') AND (Given LIKE 'John')), 
		1, 
		10, 
		100
		);
