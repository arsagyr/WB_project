CREATE TABLE IF NOT EXISTS Actors (
    id    SERIAL PRIMARY KEY, 
    Nameid  integer NOT NULL,
    Nationid integer NOT NULL,
	Number     integer NOT NULL,
	Honorar float NOT NULL
);

CREATE TABLE IF NOT EXISTS Nations (
    id    SERIAL PRIMARY KEY, 
    Name   varchar(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS Names (
    id    SERIAL PRIMARY KEY, 
    Family  varchar(30) NOT NULL,
    Given   varchar(30) NOT NULL
);

ALTER TABLE "actors" ADD CONSTRAINT "actors_fk1" FOREIGN KEY ("nameid") REFERENCES "names"("id");
ALTER TABLE "actors" ADD CONSTRAINT "actors_fk2" FOREIGN KEY ("nationid") REFERENCES "nations"("id");


