CREATE TABLE recipes (
    RecipeID SERIAL PRIMARY KEY,
    NAME varchar(40) NOT NULL,
    PrepTime varchar(40) NOT NULL,
    Difficulty integer NOT NULL,
    isVeg boolean NOT NULL
);