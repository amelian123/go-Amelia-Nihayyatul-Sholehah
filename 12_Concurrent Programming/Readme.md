SECTION 12_CONCURRENT PROGRAMMING

>>Materi yang dipelajari antara lain "big search" website, perbedaan sequential, paralel, dan concurrent, goroutine, channels dan selsect, race condition, fixing data (seperti channels blocking, waitgroups, dan mutex).

>>Sequential Program => dalam program ini, sebelum memulai tugas yang baru baiknya jika tugas sebelumnya harus diselesaikan terlebih dahulu.

>>Parallel Program => dalam program paralel, beberapa tugas dapat dijalankan secara bersamaan. Paralel berarti simultan.

>>Concurrent Program => dalam program ini, banyak tugas dapat dijalankan secara independen dan mungkin muncul bersamaan. Concurrent berarti independen atau mandiri.

## GOROUTINES
Goroutine adalah fungsi atau metode yang dijalankan secara bersamaan (independen) dengan fungsi atau metode lain.  Biaya pembuatan Goroutine sangat kecil. Thread adalah proses yang ringan atau dengan kata lain thread adalah unit yang menjalankan kode di bawah program.

>> Fitur dalam Go yaitu :
   1. Eksekusi secara bersamaan disebut Goroutines
   2.Sinkronisasi dan pesan disebut Channels
   3. Kontrol serentak banyak arah/multi-way disebut Select

>>GOMAXPROCS : Gomaxprocs digunakan untuk mengontrol jumlah utas sistem operasi yang tersedia untuk eksekusi Goroutine ke program go tertentu.Contoh perintahnya yaitu : 
    time GOMAXPROCS=1 go run main.go

## CHANNELS AND SELECT
>>Channels : adalah objek komunikasi yang digunakan oleh goroutine untuk saling berkomunikasi.
>>Select : memudahkan untuk mengontrol komunikasi data melalui satu atau banyak channels.

## RACE CONDITION
Race Condition adalah kondisi dimana dua thread mengakses memori secara bersamaan, salah satunya yaitu menulis. Kondisi balapan (race condition) terjadi karena akses yang tidak sinkron ke memori bersama.

## FIXING DATA RACE
A. WaitGroups
   >>Memblokir dengan waitgroups : cara paling mudah untuk menyelesaikan data race adalah dengan memblok akses read sampai operasi write selesai.
B. Channles Blocking
   >>Hal ini memungkinkan goroutine untuk melakukan sinkronisasi satu sama lain untuk sementara.
C. Mutex /Mutual exclusion
   >>adalah sebuah mekanisme yang di gunakan untuk mengunci dan membuka kunci, yang di maksud mengunci dan membuka kunci ialah sebuah data.