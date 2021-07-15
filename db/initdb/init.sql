CREATE TABLE cats (
    ID serial PRIMARY KEY,
    Name VARCHAR(50),
    PhotoURL VARCHAR(100),
    Rating FLOAT
);
INSERT INTO cats (ID, Name, PhotoURL, Rating) VALUES
    (1, 'Dr Nice', 'assets/images/1.jpg', 0),
    (2, 'Narco', 'assets/images/2.jpg', 0),
    (3, 'Bombasto', 'assets/images/3.jpg', 0),
    (4, 'Celeritas', 'assets/images/4.jpg', 0),
    (5, 'Magneta', 'assets/images/5.jpg', 0),
    (6, 'RubberMan', 'assets/images/6.jpg', 0);