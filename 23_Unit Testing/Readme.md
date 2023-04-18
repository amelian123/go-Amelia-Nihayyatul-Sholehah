## Section 23_Unit Testing

// Software Testing
Suatu proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang dibutuhkan (yaitu, cacat) dan untuk mengevaluasi fitur item perangkat lunak. (Menurut standar ANSI / IEEE 1059)

// Tujuan Testing
   Berikut ini tujuan dilakukannya testing yaitu :
   1. Mencegah Regresi 
   2. Tingkat Keyakinan dalam Refactoring 
   3. Meningkatkan Code Design 
   4. Code Documentation 
   5. Scaling antar team

// Level Testing
   a. UI : (End To End) Menguji interaksi antara Keseluruhan melalui Antarmuka pengguna 
   b. Integrasi : Menguji modul atau sub-sistem tertentu melalui api 
   c. Unit : Menguji bagian terkecil yang dapat diuji (operasi logis tunggal) dari aplikasi melalui metode

## Framework
Dasar dari bahasa pemrograman yaitu pemilihan framework untuk unit testing. Framework menyediakan alat, dan struktur yang diperlukan untuk menjalankan testing/pengujian secara efisien.

## Structure
   Pola biasa pada structure yaitu : 
   1. Pusatkan file pengujian Anda di dalam folder tes.  
   2. Simpan file pengujian bersama dengan file produksi.

// Tree Pada Structure
   1. Test File => mengumpulkan tes percobaan/test suites
   2. Test Suites => mengumpulkan test cases
   3. Test Fixtures dan Test Cases => pada test fixtures yang dilakukan yaitu setuo dan teardown, sedangkan pada test cases yang dilakukan yaitu input, proccess, dan output. Kedua testing ini dilakukan bersamaan.

## Runner
   a. Alat yang menjalankan file pengujian 
   b. Gunakan mode tontonan/watch mode (jika ada perubahan basis kode, jalankan pengujian secara otomatis lagi, buat TDD lebih efisien) 
   c. Pilih runner yang dapat berjalan paling cepat

## Mocking
Test cases haruslah independen

// Yang tidak bisa dilakukan testing yaitu :
   1. 3 (tiga) party modules
   2. Database
   3. 3 (tiga) party API atau eksternal sistem
   4. Kelas objek yang akan di tes berada di tempat lain

// Membuat MOCK OBJECT (n) objek palsu yang mensimulasikan perilaku objek nyata. Dalam kondisi seperti fakes, spyes, stubs, mocks

## Coverage
Pembuat kode harus memastikan bahwa kode yang dibuat cukup untuk di tes

// Alat coverage untuk menganalisis kode aplikasi saat pengujian sedang berjalan.

// Coverage Report antara lain :
   1. Function
   2. Statement
   3. Branch
   4. Lines

// Report Format antara lain :
   1. CLI
   2. XML
   3. HTML
   4. LCOV

//Menggunakan tool format seperti sonarqube (static code analyzer) agar terbaca

// Langkah Sederhana Untuk Testing yaitu:
   1. Membuat new test file di go
      a. Akhiran nama library_test.go (misalnya => user_test.go)
      b. File Lokasi : 
         - folder yang sama, package yang sama
         - folder yang sama, package yang berbeda
    2. Menulis test function
       a. Nama : Test dan Kata menggunakan huruf kapital 
       b. Harus memiliki tanda tangan func TestXxx (t*testing.T)

// Yang harus dipahami saat membuat testing yaitu :
   1. Lebih dari 1 tes function
   2. Ikuti ketentuan testing pada go

// Cara run testing yaitu go test ./lib/calculate -cover

// Run dengan Report Coverage => 
   go test ./lib/calculate 
   -coverprofile=cover.out && go tool cover
   -html=cover.out

## Unit Test Rest API Function
   1. Menulis Test Function untuk Echo
      a. Setup Echo (Metode, Params, dll)
      b. Panggil fungsi rest api di package controller 
      c. Buat fungsi uji/test function

// Run dengan Report Coverage All Package =>
   go test -v -coverpkg=./...
   -coverprofile=profile.cov ./...

   go tool cover -func profile.cov