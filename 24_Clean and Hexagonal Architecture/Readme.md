## Section 24_Clean and Hexagonal Architecture

// Arsitektur yang baik adalah pemisahan perhatian dengan menggunakan lapisan untuk membangun aplikasi modular, dapat diskalakan, dapat dipelihara dan dapat diuji.

// MVC dan MVVM
   Pada bagian ini ada 3 peran dari View, Controller, dan Model. Skemanya seperti di bawah ini :
   1. User berinteraksi dengan view
   2. View peringatan controller pada kejadian tertentu
   3. Controller memperbarui Model
   4. Model memberi peringatan bahwa view ada perubahan
   5. View mengambil model data dan mengupdate diri sendiri

// Kondisi Sekarang yaitu : 
   Setiap tim memiliki strukturnya sendiri 
   -> Bagaimana mempertahankannya?  
   -> Masalah mobilitas : sulit untuk mentransfer pengetahuan pada domain bisnis dan ditambah dengan struktur kode yang unik dan memiliki waktu serta otak ekstra untuk mempelajarinya dari atas ke bawah
   -> Masalah lain seperti : sulit untuk mengimplementasikan unit test terutama tes yang memiliki koneksi dengan database dan kemudian hampir tidak memilih tes integrasi. 

// Kerangka sebelum mendesain Clean Arsitektur yaitu :
   1. Independent of Framework : Hal ini memungkinkan Anda untuk menggunakan kerangka kerja tersebut sebagai alat, daripada memiliki kendala terbatas 
   2. Testable : dapat diuji aturan bisnis dapat diuji tanpa hal lain
   3. Independent of UI : UI dapat berubah dengan mudah, tanpa mengubah bagian sistem lainnya
   4. Independent of Database : aturan bisnis user tidak terikat pada basis data, user dapat menukar database dengan mudah 
   5. Independent of Any External : sebenarnya aturan bisnis user sama sekali tidak tahu apa-apa tentang kata luar

## Keuntungan
   1. Strukturnya standar, sehingga mudah untuk menemukan jalan dalam proyek
   2. Pengembangan lebih cepat dalam jangka panjang
   3. Ketergantungan mocking menjadi sepele dalam pengujian unit
   4. Mudah beralih dari prototipe ke solusi yang tepat (misalnya mengubah penyimpanan dalam memori ke database SQL)

// CA Layer
   Berikut ini layer architecture antara lain :
   1. Entities Layer : objek bisnis karena mencerminkan konsep yang dikelola aplikasi
   2. Use Case - Domain Layer :  berisi logika bisnis
   3. Controller - Presentation Layer : menampilkan data ke layar dan menangani interaksi pengguna
   4. Drivers - Data Layer : mengelola data aplikasi mis.  mengambil data dari jaringan, mengelola cache data

## DDD (Domain Driven Design)
   Ada 2 sudut pandang yaitu :
   1. software engineer
   2. ahli dalam pemecahan suatu masalah seperti produk owner, project manager, client.

// Clean Architecture adalah software architecture sedangkan Domain Driven Design adalah teknik desain software

## Materi Context Golang
Package context adalah suatu package yang membawa deadline, cancellation signal, atau value lain pada request/permintaan API

// Implementasi Context 
   - Membuat root context dengan fungsi background(). Scriptnya : var ctx = context.Background()
   - Fungsi background() akan membalikkan root context dimana user bisa memakainya untuk operasi lain

// Context with Value
  - Menambahkan suatu nilai/value ke dalam root context, menggunakan context.WithValue
  - Implementasi context with value sering ditemui pada middleware. Dapat mengirim data dari mddleware (contoh user_id dari token) ke handler menggunakan context

// Context With Cancelation
   1. API Tanpa Implementasi Context Cancellation : user dapat cancel request, artinya user cancel request tetapi data masih tetap harus diteruskan ke database
   2. API Dengan Implementasi Context Cancellation : user dapat cancel request, artinya user cancel request dan data tidak diteruskan ke database

## Materi MVC to CLean Code_Migrating From MVC to Clean Code
// Sebelum Migrasi :
   Kali ini akan memakai contoh project MVC yang ada pada repository berikut :
   a. MVC : the goal ketika menggunakan clean code adalah kodenya lebih modular, scallable, dan maintainable. Modular artinya bisa dengan mudah mengganti dipendensi satu ke dipendensi lainnya. Scallable artinya dapat dengan mudah menambahkan fitur baru dan lainnya. Serta maintainable artinya dapat dengan mudah memperbaiki isu jika terdapat isu pada kode yang dibuat

// Options Migrating
   Ada 3 opsi saat melakukan migrasi arsitektur kode dari mvc ke clean code yaitu : 
   1. Pertahankan desain sekarang dengan memisahkan dependensi
   2. Pertahankan desainnya tetapi pindahkan juga kodenya ke dalam suatu layer
   3. Ubah desainnya dan pisahkan dependensi

// Struktur Code yang akan dipakai yaitu :
   1. Controller => berisi kode yang berhubungan langsung ke user (interface layer)
    2. Repository => berisi kode yang berhubungan langsung ke database (interface layer)
    3. Usecase => berisi bisnis logik yang dipakai

// Di dalam folder use case berisi file order_usecase.go. 

// Folder Repository
   >> OrderRepository adalah blueprint dari repository order
   >> db *gorm.DB -> perlu memakai sharing by pointer untuk variabel yang menyimpan koneksi daripada memakai sharing by value
   >> ProductRepository adalah blueprint dari repository product
   >> db *gorm.DB -> perlu memakai sharing by pointer untuk variabel yang menyimpan koneksi daripada memakai sharing by value

// Folder Controller
   >> Ganti field struct userController dari gorm.DB ke interface orderUsecase
   >> NewOrderController adalah fungsi yang digunakan untuk membuat instance dari struct orderController

## Materi Write Unite Test_Write Unit Test
>> Membuat mocking repository dengan library testify
>> ret.Get() akan membalikkan data sesuai index nya karena data bertipe interface{} maka perlu untuk mengacastingnya
>> Find adalah nama method di repository yang akan dipanggil dan di set expected resultnya
>> Lalu panggil fungsi yang akan di test