UPDATE Actors
SET Number = $2
WHERE Number = $1;

UPDATE Actors
SET Honorar = $2
WHERE Honorar = $1;

UPDATE Names
SET Family = $2
WHERE Family = $1;

UPDATE Names
SET Name = $2
WHERE Name = $1;