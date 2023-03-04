SECTION 9 (STRING, ADVANCE FUNCTION , POINTER , METHOD, STRUCT AND INTERFACE)

>> A. String

>> B. Advance Function

>> C. Pointer : Pointer adalah variabel yang menyimpan alamat memori dari variabel lain. Pointer memiliki kekuatan untuk mengubah data yang ditunjuk.

Deklarasi pointer : var <variable_name> *<variable_type>
    var nameAddress *string 

Use : var <variable_name> *<variable_type>
	var name = “John” 
    	var nameAddress *string
   	nameAddress = &name 

Operator dalam Pointer ada dua yaitu :
1. tanda (&) : And atau Address of Operator merupakan operator yang digunakan untuk mendapatkan suatu alamat dari sebuah variable.
2. tanda (*) : Bintang atau Dereference Operator merupakan operator yang digunakan untuk mengambil nilai asli dari sebuah alamat memori.

>> D. Package : adalah kumpulan fungsi dan data.

>> E. Error Handling : terdiri dari panic dan recover
1. Panic : Saat runtime Go mendeteksi kesalahan ini, ia akan panik.
2. Recover : Untuk menambahkan kemampuan pemulihan dari kesalahan panik, tambahkan fungsi anonim atau tentukan fungsi khusus dan panggil dengan kata kunci 'tunda' dari dalam metode, di mana kepanikan mungkin terjadi dari panggilan internal lainnya.

>> F. Struct : adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map , hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

>> G. Interface : adalah kumpulan tanda tangan metode yang dapat diimplementasikan objek. Oleh karena itu, antarmuka mendefinisikan perilaku objek.
