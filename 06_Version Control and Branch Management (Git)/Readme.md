VERSION CONTROL AND BRANCH MANAGEMENT (GIT)

>>Versioning : digunakan untuk mengatur versi dari source code program
>>Tools yang dipakai seperti :
  1. VCS (Version Control System)
  2. SCM (Source Code Manager)
  3. RCS (Revision Control System)
 >>GITadalah salah satu version control system populer yang digunakan para developer untuk mengembangkan software secara bersama-bersama
 >>GIT sendiri juga bisa digunakan melalui windows, linux, serta mac installer
 >>Fungsi GitHub yaitu untuk memudahkan kolaborasi pengerjaan proyek, mencegah perubahan kode yang bisa merubah kode asli, dan sebagai portofolio developer.
 >>Berikut ini beberapa istilah dalam GitHub antara lain :
   1. Repository -> Direktori atau folder yang berisi file dan riwayat perubahan kode project 
   2. Commit -> Rekaman riwayat perubahan pada file Anda, meliputi siapa, apa, dan kapan perubahan terjadi.
   3. Clone	-> Salinan repository di komputer agar dapat diedit secara offline di perangkat.
   4. Fork	-> Menyalin repository orang lain ke akun GitHub. Biasanya digunakan untuk bereksperimen pada suatu project yang dianggap menarik.
   5. Remote -> Versi repository yang disimpan di server GitHub. Bisa melakukan sinkronisasi dengan versi clone sehingga perubahan offline tetap tercatat.
   6. Branch	Cabang dari repository utama Anda. Di branch, kode yang Anda utak-atik tak akan berefek ke repository utama. Jadi, Anda bisa bebas bereksperimen atau memperbaiki bug di sini.
   7. Merge	Menggabungkan kode yang sudah diubah pada suatu branch ke repository lainnya. Jadi, setelah Anda bereksperimen dengan kode di branch, Anda bisa langsung memasukkannya pada repository utama dengan merge.
   8. Pull request	Mengusulkan suatu perubahan pada repository ke pemilik/pemimpin project. Lalu, ia berhak menerima atau menolak usulan tersebut.
   9. Issue	Saran, pertanyaan, atau permintaan yang berhubungan dengan repository. Bisa dibuat oleh anggota tim Anda ataupun semua orang (untuk public repository).
>>Perintah Dasar pada GIT yaitu :
  1. git config : salah satu perintah git yang paling banyak digunakan adalah git config, yang bisa digunakan untuk mengatur konfigurasi tertentu sesuai keinginan pengguna, seperti email, algoritma untuk diff, username, format file, dll. Contohnya, perintah berikut bisa digunakan untuk mengatur email
     git config --global user.email bulbul@google.com
  2. git init : perintah ini digunakan untuk membuat repositori baru.
     git init
  3. git add : perintah git add bisa digunakan untuk menambahkan file ke index. Contohnya akan menambahkan file bernama coba.html yang ada di direktori lokal ke index
     git add coba.html
  4. git clone : perintah git clone digunakan untuk checkout repositori. Jia repositori berada di remote server
     git clone bulbul@93.188.160.58:/path/to/repository
  5. git commit : perintah git commit digunakan untuk melakukan commit pada perubahan ke head. Ingat bahwa perubahan apapun yang di-commit tidak akan langsung ke remote repository. Gunakan:
     git commit –m “Isi dengan keterangan untuk commit” . Misalnya git commit –m “first commit”
  6. git status : perintah git status menampilkan daftar file yang berubah bersama dengan file yang ingin di tambahkan atau di-commit. Gunakan:
     git status
  7. git push : perintah push akan mengirimkan perubahan ke master branch dari remote repository yang berhubungan dengan direktori kerja Anda. Misalnya:
     git push origin master
  8. git checkout : perintah git checkout bisa digunakan untuk membuat branch atau untuk berpindah diantaranya. Misalnya, perintah berikut ini akan membuat branch baru dan berpindah ke dalamnya:
     git checkout -b <nama-branch>
     Untuk berpindah dari branch satu ke lainnya, gunakan: git checkout <branch-name>
  9. git branch : perintah git branch bisa digunakan untuk me-list, membuat atau menghapus branch. Untuk menampilkan semua branch yang ada di repository, gunakan:
     git branch
     Untuk menghapus branch: git branch -d <branch-name>
  10. git pull : untuk menggabungkan semua perubahan yang ada di remote repository ke direktori lokal, gunakan perintah pull: git pull
  11. git merge : perintah merge digunakan untuk menggabungkan sebuah branch ke branch aktif. Gunakan : git merge <nama-branch>
  12. git diff : perintah git diff digunakan untuk menampilkan conflicts. Untuk melihat conflicts dengan file dasar, gunakan : git diff --base <nama-file>
      Perintah berikut digunakan untuk menampilkan conflicts diantara branch yang akan dimerge: git diff <source-branch> <target-branch>
