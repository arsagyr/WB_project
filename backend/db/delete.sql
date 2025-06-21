DELETE FROM Actors
WHERE id = $1;

DELETE FROM Name
WHERE Family = $1;

UPDATE Names
SET Family = $2
WHERE Family = $1;

UPDATE Names
SET Name = $2
WHERE Name = $1;