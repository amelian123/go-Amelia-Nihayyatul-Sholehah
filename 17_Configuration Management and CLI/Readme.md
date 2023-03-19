## SECTION 17_Configuration Management and CLI

>>Command Line merupakan antarmuka berbasis teks yang cepat, kuat, yang digunakan pengembang untuk berkomunikasi secara lebih efektif dan efisien dengan komputer untuk menyelesaikan serangkaian tugas yang lebih luas.

>>Beberapa alasan menggunakan comand line yaitu :
  1. Kontrol granular dari OS atau aplikasi
  2. Manajemen lebih cepat dari banyaknya OS
  3. Kemampuan menyimpan script untuk otomatisasi(automate) tugas-tugas rutin(regular task)
  4. Informasi antarmuka digunakan untuk membantu pemecahan masalah, seperti masalah koneksi jaringan

## Normal User dan Root User
   1. Normal User => diizinkan untuk memanipulasi direktori / home / username saja 
   2. Root User => diizinkan untuk memanipulasi / dir (semua direktori)
   3. Normal User + Sudoers => diizinkan untuk bertindak sebagai root sementara

## Commands Pada UNIX Shell 
>>A. DIRECTORY COMMANDS
     1. pwd : untuk cek lokasi
     2. ls : untuk melihat file dan direktori pada sistem. Menjalankannya tanpa parameter akan menampilkan konten direktori kerja saat ini.
        Berikut beberapa opsi yang bisa digunakan dengan perintah dasar Linux ls:
        => Is -R : akan mencantumkan semua file yang ada di subdirektori.
        => Is -a : akan menampilkan file yang tersembunyi bersama file yang terlihat.
     3. mkdir : untuk membuat folder baru
     4. cd :  untuk mengubah direktori kerja saat ini
     5. rm : untuk menghapus file
     6. cp : untuk menyalin satu file dari direktori saat ini ke direktori lain, masukkan cp diikuti dengan nama file dan direktori tujuan.
     7. mv : untuk move atau memindahkan file ke folder tujuan
     8. ln : untuk membuat link. Pada command ini ada dua jenis yaitu softlink dan hardlink. Untuk softlink bisa menggunakan perintah ln -s

>>B. FILES COMMANDS
     1. create : touch => untuk membuat file
     2. view : head, cat, tail, less
        >>cat => perintah ini berfungsi untuk mencantumkan, menggabungkan, dan menulis konten atau isi file dalam output standar. Untuk menjalankan command ini, ketik cat diikuti nama dan ekstensi file. Contohnya cat namafile.txt.
        >>head : digunakan untuk menampilkan data yang disalurkan (piped) ke CLI.
        >>tail : mengecek apakah file memiliki data baru, atau untuk membaca pesan error.
     3. editor : vim, nano => untuk edit teks
     4. permission : chown, chmod
        >>chmod merupakan perintah dasar yang digunakan untuk mengubah izin baca(r), tulis(w), dan eksekusi(x) direktori. Di Linux, setiap file dikaitkan dengan tiga kelas user, yaitu pemilik (owner), anggota grup (group member), dan lainnya (others). Contohnya : agar anggota grup dan orang lain bisa ikut membaca, menulis, dan mengeksekusi file, maka ubahlah izin ke jenis izin -rwxrwxrwx, yang nilai numeriknya adalah 777. Bisa menggunakan script:
        chmod 777 note.txt
        >>chown untuk mengubah atau mentransfer kepemilikan file, direktori, atau link simbolik ke username tertentu. Contohnya ingin menjadikan linuxcoba sebagai pemilik namafile.txt, bisa menggunakan script : 
        chown linuxcoba namafile.txt
     5. different : diff => untuk membandingkan isi dari dua file yang tersedia

>>NETWORK COMMANDS
  1. ping : untuk cek koneksi
  2. ssh : untuk remote server, ditemui pada GitHub
  3. netstat : untuk mengetahui di dalam jaringan tersebut sedang terjadi apa saja, sedang mengakes apa apa saja, yang aktif apa saja.
  4. nmap : untuk analisis port
  5. ip addr (ifconfig) : untuk menampilkan IP / mengubah IP / subnet
  6. wget dan curl : untuk mengunduh file dari url
  7. telnet : hampir sama dengan ssh, tetapi secara keamanan lebih rendah sehingga cukup riskan
  8. netcat : untuk aktivitas sebuah jaringan, contohnya monitoring jaringan

>>UTILITY COMMANDS
  1. man : untuk menampilkan informasi commands yang akan diberikan
  2. env : untuk menampilkan env database
  3. echo : untuk memindahkan beberapa data ke dalam satu file. Misalnya, jika ingin menambahkan teks, “Hello word” ke file yang bernama name.txt, yang perlu diketik adalah echo Hello word >> name.txt.
  4. date : untuk menampilkan waktu
  5. which : untuk menampilkan lokasi
  6. watch : untuk monitoring running dari sebuah program
  7. sudo : untuk akses dengan user biasa
  8. history : untuk menampilkan commands yang pernah dibuat atau dirun sebelumnya
  9. grep : untuk mencari sebuah kata melalui pencarian di seluruh teks dalam file tertentu. Contohnya ingin mencari kata biru dalam file notepad.txt , maka pakai script :
  grep blue notepad.txt
  Output yang diberikan perintah ini akan menampilkan baris yang memuat kata biru
  10. locate : untuk pencarian lokasi file/database

>> Shell Program berfungsi sebagai jembatan antara user dan kernel.  
>> Shell Script merupakan bahasa pemrograman yang disusun berdasarkan perintah shell.
