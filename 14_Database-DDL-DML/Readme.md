## SECTION 14_Database – DDL – DML

>>Database merupakan kumpulan data yang disimpan dengan sistem tertentu, dan saling berhubungan, sehingga dapat dikelola dengan mudah. 

>>Foreign Key dan Primary Key:
  1. Harus unik
  2. Tabel hanya boleh memiliki satu primary key, namun dalam beberapa kasus boleh lebih dari 1 primary key (composite key)
  3. Tabel boleh memiliki lebih dari satu foreign key
  4. Foreign key digunakan untuk membuat relasi antar tabel

>>Beberapa jenis relasi database antara lain :
  a. One to One : merupakan relasi yang mana setiap satu baris data pada tabel pertama hanya berhubungan dengan satu baris pada tabel kedua.
  b. One to Many :  merupakan relasi yang mana setiap satu baris data pada tabel pertama berhubungan dengan lebih dari satu baris pada tabel kedua. 
  c. Many to Many : merupakan relasi yang mana setiap lebih dari satu baris data dari tabel pertama berhubungan dengan lebih dari satu baris data pada tabel kedua. Artinya, kedua tabel masing-masing dapat mengakses banyak data dari tabel yang direlasikan. Dalam hal ini, relasi Many to Many akan menghasilkan tabel ketiga sebagai perantara tabel kesatu dan tabel kedua sebagai tempat untuk menyimpan foreign key dari masing-masing tabel.

## RDBMS (Relational Database Management Systems)
>>Contoh software yang menggunakan Relational Databse Model sebagai dasarnya yaitu MySQL
>>Berikut ini jenis perintah pada SQL antara lain :
  1. DDL (Data Definition Language)
  2. DML(Data Manipulation Language)
  3. DCL(Data Control Language)

>>DDL Statement antara lain :
  a. create database database_name;
  b. use database_name;
  c. create table....
  d. drop table....
  e. rename table...
>>Berikut ini tipe data pada MySQL yaitu :
  1. Num
  2. Huruf
  3. Date

>>DML : digunakan untuk memanipulasi data dalam tabel dari suatu database.
>>Statement Operation dari DML yaitu INSERT, SELECT, UPDATE, DELETE
>>DML Statement seperti LIKE/BETWEEN,AND/OR, ORDERED BY, LIMIT