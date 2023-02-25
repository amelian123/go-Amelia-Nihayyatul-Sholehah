DATA STRUCTURE

>>Data structure adalah sebuah cara untuk menyimpan dan mengorganisir data dalam program komputer. 

>>Berikut ini jenis data structure diantaranya yaitu :
  1. Array: adalah struktur data yang digunakan untuk menyimpan kumpulan data dengan tipe data yang sama. Dalam golang tipe data ini bersifat statis, artinya ukuran array harus didefinisikan saat deklarasi dan ukuran array tidak dapat diubah setelahnya.
  2. Slice: adalah struktur data yang mirip dengan array, namun bersifat dinamis dan ukuran slice dapat diubah setelahnya. Slice pada Golang dapat digunakan untuk melakukan manipulasi data yang lebih fleksibel dibandingkan dengan array.
  3. Map: adalah struktur data yang digunakan untuk menyimpan data dalam bentuk pasangan key-value. Map pada Golang bersifat dinamis dan ukuran map dapat diubah setelahnya.
  4. Struct: adalah struktur data yang digunakan untuk menyimpan kumpulan data dengan tipe data yang berbeda. Struct pada Golang mirip dengan konsep record pada bahasa pemrograman lainnya.
  5. Queue: adalah struktur data yang digunakan untuk menyimpan data dalam bentuk antrian (first in, first out). Queue pada Golang dapat diimplementasikan dengan menggunakan slice.
  6. Stack: adalah struktur data yang digunakan untuk menyimpan data dalam bentuk tumpukan (last in, first out). Stack pada Golang dapat diimplementasikan dengan menggunakan slice.

>>Untuk mendeklarasikan array, perlu menentukan jumlah elemen yang disimpannya dalam tanda kurung siku [], diikuti dengan jenis elemen yang disimpan array.
    var <variable_name> [<size_of_array>]<tipe_variable>

>>Fungsi adalah bagian dari kode yang dipanggil dengan nama. Fungsi adalah cara mudah untuk membagi kode menjadi blok-blok yang berguna. 
  Deklarasi dari function yaitu :
        func <name_function> () { <statements>}
            func main() {}
        func <name_function> () <type_return> { <statements> }
            func phi() float64 { return 3.14 }
        func <name_function> (<parameter>) { <statements> }
            func square(value int) int {
                return value * value
        }
