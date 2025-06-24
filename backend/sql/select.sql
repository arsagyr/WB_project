SELECT Actors.id, Names.Family AS "Family name", Names.Given AS "Given name", Nations.Name AS "Nation", Number, Honorar  FROM Actors 
JOIN Names ON Actors.Nameid=Names.id
JOIN Nations ON Actors.Nationid=Nations.id