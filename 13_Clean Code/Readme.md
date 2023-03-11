SECTION 13_CLEAN CODE

## Pengertian
Clean code adalah istilah untuk kode yang mudah dibaca, dipahami, dan diubah oleh programmer. 

## Penggunaan Clean Code
Berikut ini penggunaan clean code antara lain :
1. Kolaborasi kerja
2. Pengembangan fitur
3. Pengembangan yang lebih cepat

## Karakteristik Clean Code
Berikut ini karakteristik clean code yaitu :
1. Mudah dipahami : artinya penamaan kode mudah dipahami 
2. Mudah dieja dan dicari
3. Singkat namun mendeskripsikan konteks
4. Konsisten : mulai dari deklarasi tulisan (kapital/tidaknya huruf), tanda baca, method fungsi harus konsiten/sama
5. Menghindari penambahan konteks yang tidak perlu
6. Adanya komentar
7. Menggunakan fungsi yang sesuai/baik
8. Menggunakan konvensi 
9. Menggunakan formatting

>>Saran Formatting
  a. Menggunakan prettier atau formatter
  b. Deklarasi variabel berdekatan dengan penggunanya
  c. Memperhatikan indentasi
  d. Baris kode yang berhubungan saling berdekatan
  e. Satu class berisi 300 sampai 500 baris
  f. Dekatkan fungsi dengan pemanggilnya
  g. Lebar baris pada kode yaitu 80 sampai 120 karakter

## Prinsip Clean Code : yaitu menghindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mencetak fungsi C, dan seterusnya.

Berikut ini tips sederhana dalam clean code yaitu :
a. Jangan menggunakan banyak argumen pada fungsi
b. Fungsi atau class harus kecil
c. Harus diperhatikan untuk mencapai kondisi yang seimbang, kecil, dan jumlahnya minimal
d. Fungsi dibuat untuk melakukan satu tugas saja

## Duplikasi kode terjadi karena sering copy paste. Untuk menghindarinya perlu membuat fungsi yang dapat digunakan secara berulang.

## Refactoring 
Refactoring merupakan proses restrukturisasi kode yang dibuat dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal.

>>Teknik Refactoring antara lain :
  1. Memecah kode dengan fungsi atau class
  2. Deteksi kode yang memiliki duplikasi
  3. Membuat sebuah abstraksi
  4. Memperbaiki penamaan dan lokasi kode