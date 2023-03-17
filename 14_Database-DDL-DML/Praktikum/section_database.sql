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