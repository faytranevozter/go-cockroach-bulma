-- create table user
CREATE TABLE tugas_cockroach.user (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(50),
	email VARCHAR(50)
);

-- dummy data table user
INSERT INTO tugas_cockroach.user(name, email) VALUES
('Panjul', 'panjoel@gmail.com'),
('Fahrur', 'fahrur@gmail.com'),
('Indra', 'indra@gmail.com'),
('Oky', 'oky@gmail.com'),
('Laily', 'laily@gmail.com'),
('Rismi', 'rismi@gmail.com');


-- create table buku
CREATE TABLE tugas_cockroach.buku (
	id SERIAL PRIMARY KEY, 
	judul VARCHAR(50),
	pengarang VARCHAR(50),
	tahun VARCHAR(4)
);

-- dummy data table buku
INSERT INTO tugas_cockroach.buku(judul, pengarang, tahun) VALUES
('Laskar Biru', 'Oky Riyanto', '2017'),
('Mumet PHP', 'Fahrur Rifai', '2016'),
('Agroculture', 'Indra Purnama', '2019'),
('Antologi Puisi', 'Laily Rahma', '2018'),
('Agro Cabe', 'Rismi', '2017');