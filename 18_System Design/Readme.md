## Section 18_SYSTEM DESIGN

>>Diagram adalah suatu representasi simbolis informasi dalam bentuk geometri 2 (dua) dimensi sesuai teknik visualisasi.

>>Berikut ini tools dalam desain diagram yaitu :
  1. smartdraw
  2. Lucidchart
  3. Whimsical
  4. draw.io
  5. visio

>>Desain software yang sering digunakan yaitu :
  1. Flowchart : merupakan diagram yang menampilkan langkah-langkah dan keputusan untuk melakukan sebuah proses dari suatu program. Setiap langkah digambarkan dalam bentuk diagram dan dihubungkan dengan garis atau arah panah.
  2. Use Cae : merupakan salah satu jenis diagram UML (Unified Modelling Language) yang menggambarkan hubungan interaksi antara sistem dan aktor. Use Case dapat mendeskripsikan tipe interaksi antara si pengguna sistem dengan sistemnya. 
  3. ERD : (Entity Relationship Diagram) merupakan model atau rancangan untuk membuat database, supaya lebih mudah dalam menggambarkan data yang memiliki hubungan atau relasi dalam bentuk sebuah desain. Dengan adanya ERD, maka sistem database yang terbentuk dapat digambarkan dengan lebih terstruktur dan terlihat rapi. 
  4. HLA (High Level Architecture)

## Horizontal Scalling dan Vertical Scalling

>>Beberapa pertimbangan yang dilakukan saat merancang sebuah sistem yang besar yaitu : 
  1. Potongan arsitektur apa saja yang berbeda yang dapat digunakan?  
  2. Bagaimana potongan-potongan ini bekerja satu sama lain?  
  3. Bagaimana kita bisa memanfaatkan potongan-potongan ini dengan sebaik-baiknya lalu apa pengorbanan yang tepat? 
  Membiasakan konsep-konsep ini akan sangat bermanfaat dalam memahami konsep sistem terdistribusi.

>>Kunci Karakteristik Dalam Sistem Terdistribusi yaitu :
  1. Scalability => merupakan kemampuan hardware atau software dalam mendukung berbagai ukuran data atau jumlah pengguna yang menggunakannya.
  2. Reliability => merupakan kemampuan hardware atau software dalam mendukung berbagai ukuran data atau jumlah pengguna yang menggunakannya. Definisi secara umum, merupakan peluang sebuah komponen, sub-sistem atau sistem yang melakukan fungsinya dengan baik, seperti yang dipersyaratkan, dalam kurun waktu tertentu dan dalam kondisi operasi tertentu.
  3. Availability => merupakan waktu sistem tetap beroperasi untuk melakukan fungsi yang diperlukan dalam jangka waktu tertentu. Ini adalah ukuran sederhana persentase waktu sistem, layanan, atau mesin tetap beroperasi dalam kondisi normal.  Contoh : pesawat yang dapat diterbangkan berjam-jam dalam sebulan tanpa downtime yang banyak dapat dikatakan memiliki availability yang tinggi.
  4. Efficiency => untuk memahami bagaimana mengukur efisiensi sistem terdistribusi, anggaplah kita memiliki operasi yang berjalan secara terdistribusi dan mengirimkan sekumpulan item sebagai hasilnya. Dua ukuran standar efisiensi adalah waktu respon (atau latensi) yang menunjukkan penundaan untuk mendapatkan item pertama dan throughput (atau bandwidth) yang menunjukkan jumlah item yang dikirimkan dalam satuan waktu tertentu (misalnya , satu detik) .
  5. Serviceability or Manageability => merupakan kesederhanaan dan kecepatan dimana sistem dapat diperbaiki atau dipelihara. Jika waktu untuk memperbaiki sistem yang gagal bertambah, maka ketersediaannya akan berkurang. Hal-hal yang perlu dipertimbangkan untuk pengelolaan adalah kemudahan mendiagnosa dan memahami masalah ketika terjadi, kemudahan membuat update atau modifikasi, dan seberapa sederhana sistem untuk beroperasi (yaitu, apakah secara rutin beroperasi tanpa kegagalan atau pengecualian).

## Job/Work Queue
Definisi :
a. Job/Work Queue "Dalam perangkat lunak sistem , sebuah job queue ( terkadang antrian batch ) merupakan struktur data yang dikelola oleh software penjadwal pekerjaan yang berisi pekerjaan untuk dijalankan."  
b. Work Queue adalah kerangka kerja untuk membangun aplikasi master - pekerja besar yang  rentang ribuan mesin diambil dari cluster, awan, dan grid.  

## Load Balancing
>>Merupakan komponen penting lainnya dari sistem terdistribusi. Ini membantu untuk menyebarkan lalu lintas di sekelompok server untuk meningkatkan respon dan ketersediaan aplikasi, website atau database. Load Balancing juga melacak status semua sumber daya saat mendistribusikan permintaan. Jika server tidak tersedia untuk menerima permintaan baru atau tidak merespons atau memiliki tingkat kesalahan yang tinggi, Load Balancing akan menghentikan pengiriman lalu lintas ke server tersebut.
>>Load Balancing untuk memanfaatkan skalabilitas dan redundansi penuh, kita dapat mencoba menyeimbangkan beban pada setiap lapisan sistem. 
Berikut ini 3 (tiga) tempat yang perlu ditambahkan yaitu : 
- Antara pengguna dan server web 
- Antara server web dan lapisan platform internal, seperti server aplikasi atau server cache 
- Antara lapisan platform internal dan basis data.

## Monolithic dan Microservices
    A. Monolithic
      >>Aplikasi Monolitik memiliki basis kode tunggal dengan banyak modul. Modul dibagi menjadi fitur bisnis atau fitur teknis. Memiliki sistem build tunggal yang membangun seluruh aplikasi dan/atau dependensi. Ini juga memiliki biner tunggal yang dapat dieksekusi atau digunakan.
   B. Microservices
      >>merupakan layanan yang dapat diterapkan secara independen yang dimodelkan di seputar domain bisnis. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur menawarkan banyak pilihan untuk memecahkan masalah yang mungkin Anda hadapi. Oleh karena itu, arsitektur layanan mikro didasarkan pada beberapa layanan mikro yang berkolaborasi.

## SQL dan NoSQL
SQL dan NoSQL (atau database relasional dan database non-relasional). 
>> Database relasional terstruktur dan memiliki skema standar seperti buku telepon yang menyimpan nomor telepon dan alamat.  
>> Basis data non-relasional tidak terstruktur, dan memiliki skema dinamis seperti folder file yang menampung segala sesuatu mulai dari alamat dan nomor telepon seseorang hingga 'suka' Facebook mereka dan preferensi belanja online. 

>>Keuntungan Relasional Database yaitu :
  1. Dirancang untuk segala keperluan
  2. SQL memiliki standar yang jelas
  3. Memiliki banyak tool seperti administrasi, reporting, dan framework.

>> Istilah-istilah (sifat ACID): 
   - Atomicity => merupakan transaksi yang terjadi semua atau tidak terjadi sama sekali
   - Consistency => data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu
   - Isolation => pada saat terjadi request yang bersamaan/cocurrent, memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial
   - Durability => merupakan jaminan bahwa transaksi yang telah tersimpan tetap tersimpan.

>>Tidak hanya SQL, DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan dengan model RDBMS (sifat ACID). 
Berikut ini cara untuk menghindarinya yaitu :
1. Effort pada sifat transaksi ACID
2. Kompleksitas SQL
3. Design schema di depan
4. Transaksi ditangani oleh aplikasi

>>Kelebihannya yaitu kurangnya skemanya, perkembangan cepat
>>Kapan digunakannya yaitu membutuhkan skema fleksibel, ACID tidak diperlukan, data logging (json), data sementara (chace)
>>Kapan tidak direkomendasikan yaitu untuk data finansial, data transaksi, business critical, ACID-compliant

## Schema-less
A. Tradisonal RDBMS
   1. Tidak bisa menambah data yang tidak sesuai skema
   2. Perlu menambahkan data NULL pada item yang tidak memiliki data 
   3. Memiliki tipe data yang strict
   4. Tidak dapat menambah beberapa item data pada sebuah field
B. NoSQL DBMS
   1. Tidak memiliki skema ketika menambahkan data
   2. Aplikasi menangani proses validasi tipe data
   3. Mendukung proses agregasi dokumen pada item

>>Tipe/Kategori
  - Key/value
  - Column-family
  - Graph
  - Document-based
  - Other geospatial file-system object

Dalam teknologi basis data, tidak ada solusi yang cocok untuk semua. Itu sebabnya banyak bisnis bergantung pada database relasional dan non-relasional untuk kebutuhan yang berbeda. Bahkan ketika database NoSQL semakin populer karena kecepatan dan skalabilitasnya, masih ada situasi dimana database SQL yang sangat terstruktur dapat bekerja lebih baik,  memilih teknologi yang tepat bergantung pada use case.

## Caching
 Cache yang digunakan dalam data yang baru saja diminta kemungkinan besar akan diminta lagi. Cache digunakan hampir setiap lapisan komputasi seperti hardware, sistem operasi, web browser, aplikasi web, dan lainnya. Cache seperti memori jangka pendek maksudnya yaitu memiliki ruang terbatas, tetapi biasanya lebih cepat daripada sumber data asli dan berisi item yang paling baru diakses. Cache dapat ada di semua level dalam arsitektur, tetapi sering ditemukan di level yang paling dekat dengan front end di mana cache diimplementasikan untuk mengembalikan data dengan cepat tanpa membebani level downstream.

## Database Replication
>> Redundancy dan Replication
   Redundansi adalah duplikasi komponen atau fungsi penting dari suatu sistem dengan maksud untuk meningkatkan kehandalan sistem, biasanya dalam bentuk backup atau fail-safe, atau untuk meningkatkan kinerja sistem yang sebenarnya. Misalnya, jika hanya ada satu salinan file yang disimpan di satu server, maka kehilangan server tersebut berarti kehilangan file tersebut. Karena kehilangan data jarang merupakan hal yang baik, maka dapat membuat duplikat atau salinan berlebihan dari file tersebut untuk mengatasi masalah ini.

## Database Indexing
Pada database, index merupakan sebuah struktur data yang berisi kumpulan keys beserta referensinya ke actual data di table. Tujuannya untuk mempercepat proses penentuan lokasi data tanpa melakukan pencarian secara penuh ke seluruh data (full scan).

>>Keuntungan dari indeks database yaitu :
  1. Sorted and Less -> jumlah data pada index jauh lebih kecil daripada data aslinya (tergantung kondisi datanya). Dan data-data ini juga tersimpan secara berurutan dalam bentuk struktur apapun.
  2. (Mostly) Stored in Memory (RAM) -> Dengan ukuran dataset yang jauh lebih kecil, index besar kemungkinan dapat disimpan di RAM, dimana cost untuk baca/tulis pada RAM jauh lebih kecil daripada pada Harddisk. Sehingga proses pencarian keys pada index sangatlah cepat.