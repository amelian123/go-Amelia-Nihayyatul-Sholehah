## SECTION 16_JOIN, Agregasi, UNION, SUB QUERY, FUNCTION

## JOIN
>>Join adalah sebuah klausa untuk mengkombinasikan record dari dua data atau lebih pada tabel.

>>Join Standar SQL ANSI ada 3 yaitu :
  1. Inner : merupakan penggabungan yang cukup umum digunakan, dimana pada join ini SQL hanya akan mengambil data yang beririsan saja untuk masing-masing tabel. Sementara untuk data yang tidak sama untuk kedua tabel akan diabaikan.
    Scriptnya :
    SELECT t.message FROM users u
    INNER JOIN tweets t
    ON u.id = t.user_id;
  2. Left : merupakan  penggabungan dua tabel atau lebih. Hanya saja pada left join, SQL akan menampilkan semua isi dari tabel pertama kemudian untuk data di tabel kedua akan menyesuaikan dengan kolom yang ada di tabel kedua.
    Scriptnya :
    SELECT u.username, t.message
    FROM users u
    LEFT JOIN tweets t
    ON u.id = t.user_id;
  3. Right : pada right join sebenarnya hampir mirip dengan konsep yang digunakan pada left join. Jika pada left join SQL akan menggabungkan data dengan mengikuti tabel pertama yang dianggap berada di kiri, maka pada right join data akan digabungkan sesuai dengan kolom yang ada pada tabel kedua.
    Scriptnya : 
    SELECT u.username, t.message
    FROM users u
    RIGHT JOIN tweets t
    ON u.id = t.user_id;

## UNION
>>berfungsi untuk menggabungkan dua tabel dalam bentuk baris baru ke bawah, dimana field yang di SELECT antar tabel harus sama. Atau sederhananya yaitu untuk menempatkan baris dari query satu sama lain dan nilainya distinct/unik.
    Scriptnya : 
    SELECT username, fullname
    FROM users WHERE id=1
    UNION
    SELECT username, fullname
    FROM users WHERE id=2

## Agregasi
>>fungsi agregasi digunakan untuk melakukan perhitungan terhadap nilai-nilai hasil suatu query menggunakan SQL. Setiap database mempunyai banyak fungsi agregasi yang spesifik untuk database tersebut.

>>Berikut ini fungsi agregasi SQL yaitu :
  a. MIN : digunakan mengelompokkan baris untuk membentuk nilai ringkasan tunggal.
        -Contoh menampilkan id terkecil dari tabel users :
        SELECT MIN(id) AS id FROM users
  b. MAX : digunakan untuk mendapatkan nilai maksimum atau terbesar dari sebuah data record di tabel.
        -Contoh menampilkan id terbesar dari tabel users :
        SELECT MAX(id) AS id FROM users
  c. SUM : digunakan untuk mendapatkan jumlah total nilai dari sebuah data di tabel.
        -Contoh menampilkan jumlah total favorit_count dari tabel tweets dengan user_id 1 :
        SELECT SUM(favorit_count) 
        FROM tweets WHERE user_id=1
  d. AVG : digunakan untuk mencari nilai rata-rata dari sebuah data di tabel.
        -Contoh menampilkan nilai rata-rata favorit_count dari tabel tweets dengan user_id 1
        SELECT SUM(favorit_count) 
        FROM tweets WHERE user_id=1
  e. COUNT : digunakan untuk mencari jumlah dari sebuah data di tabel. 
        -Contoh menampilkan jumlah data dari tabel tweets dengan user_id 1
        SELECT COUNT(1) FROM tweets 
        WHERE user_id=1
  f. HAVING : digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi agregasi
        -Contoh menampilkan data dari tabel tweets dengan jumlah total favorit_count per user lebih dari 2.
        SELECT user_id FROM tweets 
        GROUP BY user_id
        HAVING SUM(favorit_count)>2

## SUBQUERY
>>Subquery atau Inner query atau nested query merupakan query di dalam query SQL lain.
>>Sebuah subquery digunakan untuk mengembalikkan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil.
>>Statements dalam subquery yaitu SELECT, INSERT, UPDATE, dan DELETE bersama dengan operator seperti IN, BETWEEN, dan sebagainya.

## FUNCTION
>>Function merupakan kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya.