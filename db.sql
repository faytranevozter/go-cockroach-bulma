CREATE DATABASE tugas_cockroach;

CREATE TABLE tugas_cockroach.user (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(50),
	email VARCHAR(50)
);

INSERT INTO tugas_cockroach.user(name, email) VALUES
('Panjul', 'panjoel@gmail.com'),
('Fahrur', 'fahrur@gmail.com'),
('Indra', 'indra@gmail.com'),
('Oky', 'oky@gmail.com'),
('Laily', 'laily@gmail.com'),
('Rismi', 'rismi@gmail.com');