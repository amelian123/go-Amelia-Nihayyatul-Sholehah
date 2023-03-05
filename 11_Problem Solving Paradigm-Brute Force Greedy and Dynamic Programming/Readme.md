SECTION 11_Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming

>>Problem Solving Paradigm : adalah pendekatan yang biasa digunakan untuk memecahkan masalah seperti Complete Search (brute force), Divide and Conquer, Pendekatan Greedy, dan Dynamic Programming. Setiap masalah perlu diselesaikan dengan pendekatan yang sesuai.

>>Complete Search : biasa disebut Bruteforce merupakan metode untuk memecahkan masalah dengan melintasi seluruh ruang pencarian untuk mendapatkan solusi yang dibutuhkan. Bruteforce terjadi ketika tidak ada algoritma lain yang tersedia.  Biasanya mudah ditulis karena lugas.  Secara teoritis semua masalah dapat diselesaikan dengan menggunakan pendekatan Brute Force terutama ketika Anda memiliki waktu yang tidak terbatas.

>>Divide and Conquer : adalah paradigma pemecahan masalah di mana suatu masalah disederhanakan dengan 'membagi' menjadi bagian-bagian yang lebih kecil dan kemudian menaklukkan setiap bagian. 
Langkahnya : 
1. Divide : membagi masalah yang besar menjadi masalah yang lebih kecil.
2. Conquer : ketika masalah sudah cukup kecil untuk diselesaikan, langsung selesaikan
3. Combine : jika dibutuhkan maka perlu menggabungkan solusi dari masalah - masalah yang lebih kecil menjadi solusi untuk masalah yang besar.

>> Binary Search : adalah algoritma pencarian untuk data yang terurut. Pencarian dilakukan dengan cara menebak apakah data yang dicari berada ditengah-tengah data, kemudian membandingkan data yang dicari dengan data yang ada ditengah. Bila data yang ditengah sama dengan data yang dicari, berarti data ditemukan. Namun, bila data yang ditengah lebih besar dari data yang dicari, maka dapat dipastikan bahwa data yang dicari kemungkinan berada disebelah kiri dari data tengah dan data disebelah kanan data tengah dapat diabai. Upper bound dari bagian data kiri yang baru adalah indeks dari data tengah itu sendiri. Sebaliknya, bila data yang ditengah lebih kecil dari data yang dicari, maka dapat dipastikan bahwa data yang dicari kemungkinan besar berada disebelah kanan dari data tengah. Lower bound dari data disebelah kanan dari data tengah adalah indeks dari data tengah itu sendiri ditambah 1, dan seterusnya.

>>Greedy : algoritma dikatakan greedy jika membuat pilihan optimal lokal pada setiap langkah dengan harapan pada akhirnya mencapai solusi optimal global. Dalam beberapa kasus, greedy bekerja - solusinya singkat dan berjalan efisien.

>>Dynamic Programming : adalah teknik algoritma untuk memecahkan masalah optimisasi dengan memecahnya menjadi submasalah yang lebih sederhana dan menggunakan fakta bahwa solusi optimal untuk keseluruhan masalah bergantung pada solusi optimal untuk sub-masalahnya .

>>Fibonacci => persamaan fibbonacci ke-n yaitu : Fib(n) = Fib(n-1) + Fib(n-2), for n>1

>>Metode Dynamic Programming antara lain :
1. Top-down dengan Memoisasi => memecahkan masalah yang lebih besar dengan mencari solusi secara rekursif untuk sub-masalah yang lebih kecil. Setiap kali akan memecahkan sub-masalah, programmer menyimpan hasilnya sehingga tidak perlu menyelesaikannya berulang kali jika itu dipanggil berkali-kali. Sebagai gantinya hanya dapat mengembalikan hasil yang disimpan.  Teknik menyimpan hasil sub problem yang sudah dipecahkan ini disebut Memoization .
2. Bottom-up dengan Tabulasi => tabulasi adalah kebalikan dari pendekatan top-down dan menghindari rekursi.  Dalam pendekatan ini, pemecahan masalah "dari bawah ke atas" (yaitu dengan menyelesaikan semua sub-masalah yang terkait terlebih dahulu). Biasanya dilakukan dengan mengisi tabel n-dimensi. Berdasarkan hasil pada tabel, solusi untuk masalah teratas/awal kemudian dihitung.  Tabulasi adalah kebalikan dari Memoisasi, seperti dalam Memoisasi penyelesaian masalah dan mempertahankan peta dari sub-masalah yang sudah diselesaikan. Dengan kata lain, dalam memoisasi, melakukannya dari atas ke bawah. Artinya peyelesaian masalah teratas terlebih dahulu (yang biasanya berulang ke bawah untuk menyelesaikan sub-masalah).