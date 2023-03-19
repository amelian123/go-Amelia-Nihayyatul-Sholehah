--SECTION 14 DATABASE, DDL, DML

--1.
CREATE DATABASE alta_online_shop;

use alta_online_shop;

--2.a
CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    address VARCHAR(255),
    date_of_birth DATE,
    user_status VARCHAR(255),
    gender VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--2.b.1
CREATE TABLE product (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  stock INT NOT NULL,
  product_type_id INT NOT NULL,
  operator_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (product_type_id) REFERENCES product_type(id),
  FOREIGN KEY (operator_id) REFERENCES operators(id)
);

--2.b.2
CREATE TABLE product_type (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

--2.b.3
CREATE TABLE operators (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

--2.b.4
CREATE TABLE product_description (
    id INT PRIMARY KEY AUTO_INCREMENT,
    description TEXT
);

--2.b.5
CREATE TABLE payment_method (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

--2.c.1
CREATE TABLE transaction (
    id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    transaction_date DATE,
    FOREIGN KEY (customer_id) REFERENCES user(id)
);

--2.c.2
CREATE TABLE transaction_detail (
    id INT PRIMARY KEY AUTO_INCREMENT,
    transaction_id INT,
    product_id INT,
    quantity INT,
    price_per_unit INT,
    FOREIGN KEY (transaction_id) REFERENCES transaction(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);

--3.
CREATE TABLE kurir (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--4.
ALTER TABLE kurir ADD ongkos_dasar INT;

--5.
ALTER TABLE kurir RENAME TO shipping;

--6.
DROP TABLE shipping;

--7.a
CREATE TABLE payment_method_description (
    id INT PRIMARY KEY AUTO_INCREMENT,
    payment_method_id INT,
    description TEXT,
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);

--7.b.1
CREATE TABLE user (
    id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    status_user VARCHAR(20) NOT NULL,
    gender ENUM('M', 'F') NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

--7.b.2
CREATE TABLE alamat (
    id INT PRIMARY KEY,
    user_id INT NOT NULL,
    jalan VARCHAR(100) NOT NULL,
    kota VARCHAR(50) NOT NULL,
    provinsi VARCHAR(50) NOT NULL,
    kode_pos VARCHAR(10) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

--7.c
CREATE TABLE user_payment_method_detail (
    user_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    PRIMARY KEY (user_id, payment_method_id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);


--SECTION 16 JOIN, UNION, AGREGASI, SUBQUERY, FUNCTION(DBMS)
--A.DML

--soal 1.a
INSERT INTO operators VALUES ('1', 'Ahmad'), ('2', 'Kelly'), ('3', 'Fero'), ('4', 'Afgan'), ('5', 'Kenzo');

--soal 1.b
INSERT INTO product_type VALUES ('1', 'Balon'), ('2', 'Pensil'), ('3', 'Sepeda');

--soal 1.c
INSERT INTO product
(name, price, stock, product_type_id, operator_id, created_at, updated_at)
values
('piring', 100, 10, 1, 3, current_timestamp(), current_timestamp()),
('wajan', 20000, 4, 1, 3, current_timestamp(), current_timestamp());

--soal 1.d
INSERT INTO product
(name, price, stock, product_type_id, operator_id, created_at, updated_at)
values
('sabun', 5000, 8, 2, 1, current_timestamp(), current_timestamp()),
('sampo', 9000, 6, 2, 1, current_timestamp(), current_timestamp()),
('sandal', 15000, 4, 2, 1, current_timestamp(), current_timestamp());

--soal 1.e
INSERT INTO product
(name, price, stock, product_type_id, operator_id, created_at, updated_at)
values
('baju', 55000, 12, 3, 4, current_timestamp(), current_timestamp()),
('payung', 19000, 7, 3, 4, current_timestamp(), current_timestamp()),
('topi', 67000, 14, 3, 4, current_timestamp(), current_timestamp());

--soal 1.f
INSERT INTO product_description (id, description)
VALUES 
(1, 'menarik'),
(2, 'garansi dua bulan'),
(3, 'tahan banting'),
(4, 'kualitas bagus'),
(5, 'dalam botol'),
(6, 'jumlah banyak'),
(7, 'tahan wangi seharian'),
(8, 'tidak mudah luntur'),
(9, 'anti karat'),
(10, 'anti ketombe'),
(11, 'desain terbaru'),
(12, 'terbuat dari kanvas'),
(13, 'natural alami'),
(14, 'melindungi dari matahari'),
(15, 'aman dipakai');

--soal 1.g
INSERT INTO payment_method (id, name) 
VALUES 
(1, 'cash'), 
(2, 'kredit'), 
(3, 'cash');

--soal 1.h
INSERT INTO user (id, name, address, date_of_birth, user_status, gender, created_at, updated_at) 
VALUES (1, 'Sela', 'Jl.Mawar', '2002-03-09', 'student', 'F', current_timestamp(), current_timestamp()), 
       (2, 'Atun', 'Jl.Huum', '2012-05-04', 'student', 'F', current_timestamp(), current_timestamp()), 
       (3, 'Raja', 'Jl.Ruum', '2009-03-07', 'student', 'M', current_timestamp(), current_timestamp()),
       (4, 'Faiz', 'Jl.Budak', '2001-08-30', 'college', 'M', current_timestamp(), current_timestamp()),
       (5, 'Farhan', 'Jl.Aura', '2000-09-23', 'student', 'M', current_timestamp(), current_timestamp());
       
--soal 1.i  
INSERT INTO transaction (id, customer_id, transaction_date) 
VALUES (1, 1, '2023-03-17'), 
       (2, 2, '2023-03-16'),
       (3, 3, '2023-03-15'),
       (4, 1, '2023-03-17'), 
       (5, 2, '2023-03-16'),
       (6, 3, '2023-03-15'),
       (7, 1, '2023-03-17'), 
       (8, 2, '2023-03-16'),
       (9, 3, '2023-03-15'),
       (10, 1, '2023-03-17'), 
       (11, 2, '2023-03-16'),
       (12, 3, '2023-03-15'),
       (13, 1, '2023-03-17'), 
       (14, 2, '2023-03-16'),
       (15, 3, '2023-03-15');
       
--soal 1.j
INSERT INTO transaction_detail (id, transaction_id, product_id, quantity, price_per_unit) 
VALUES (1, 1, 1, 3, 24000), 
       (2, 2, 2, 4, 15000),
       (3, 3, 3, 5, 9000),
       (4, 4, 4, 5, 4000),
       (5, 5, 5, 36, 2000),
       (6, 6, 6, 8, 6000),
       (7, 7, 7, 2, 39000);
       
--soal 2.a
SELECT name, gender FROM user WHERE gender = 'M';

--soal 2.b
SELECT * FROM product WHERE id = 3;

--soal 2.c
SELECT * FROM user WHERE name LIKE '%a%' AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY);

--soal_2.d
SELECT COUNT(*) FROM user WHERE gender = 'F';

--2.e
SELECT * FROM user ORDER BY name ASC;

--2.f
SELECT * FROM product LIMIT 5;

SOAL 3
--3.a
UPDATE product SET name = 'product dummy' WHERE id = 1;

--3.b
UPDATE transaction_detail SET quantity = 3 WHERE product_id = 1;

SOAL 4 
--4.a
DELETE FROM product WHERE id IN (1);
--4.b
DELETE FROM product WHERE product_type_id = 1;


--B. Join, Union, Sub query, function
--soal 1
SELECT * FROM transaction WHERE id = 1 UNION SELECT * FROM transaction WHERE id = 2;

--soal 2.
SELECT SUM(price_per_unit) as total_harga FROM transaction_detail WHERE id = 1;

--soal 3.
SELECT COUNT(*) as total_transaksi FROM transaction t 
JOIN transaction_detail td ON t.id = td.transaction_id 
JOIN product p ON td.product_id = p.id 
WHERE p.product_type_id = 2;

--soal 4.
SELECT p.*, pt.name as product_type_name FROM product p 
JOIN product_type pt ON p.product_type_id = pt.id;

--soal 5.
SELECT t.*, p.name as product_name, u.name as user_name FROM transaction t 
JOIN transaction_detail td ON t.id = td.transaction_id 
JOIN product p ON td.product_id = p.id 
JOIN user u ON t.id = u.id;

--soal 6.
DELIMITER $$
CREATE FUNCTION delete_transaction_detail_by_transaction_id(t_id INT)
RETURNS BOOLEAN
BEGIN
    DELETE FROM transaction_detail WHERE transaction_id = t_id;
    RETURN TRUE;
END$$
DELIMITER ;

--soal 7.
DELIMITER $$
CREATE FUNCTION update_total_quantity_by_transaction_detail_id(td_id INT)
RETURNS BOOLEAN
BEGIN
    UPDATE transaction t 
    JOIN (SELECT transaction_id, SUM(quantity) as total_quantity FROM transaction_detail GROUP BY transaction_id) td 
    ON t.id = td.transaction_id 
    SET t.total_quantity = td.total_quantity 
    WHERE td.transaction_detail_id = td_id;   
	RETURN TRUE;
END$$
DELIMITER ;

--soal 8.
SELECT * FROM product WHERE id NOT IN (SELECT product_id FROM transaction_detail);
