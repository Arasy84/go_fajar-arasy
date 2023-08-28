 string merupakan tipe data yang dapat menampung karakter-karakter
di golang terdapat package strings yang memiliki fungsi-fungsi untuk melakukan manipulasi string, misalnya: Contains, Replace, dsb.
di golang terdapat fitur anonymous function (fungsi tanpa nama), variadic function (parameter yang bisa menampung banyak argumen), closure (fungsi yang memiliki lexical scope yang mana dapat mengakses data di sekitar fungsi), defer function (fungsi yang dipanggil sebelum fungsi dimana dia dipanggil berakhir)
di golang terdapat pointer yang merupakan tipe data yang menyimpan referensi dari sebuah nilai dengan tipe data tertentu
di golang terdapat struct untuk mendefinisikan struktur data dengan kumpulan data yang bermacam-macam
di golang terdapat method, yang mana mirip dengan function, tapi hanya dapat diakses oleh struct yang telah didefinisikan sebagai receiver dari function terkait
di golang juga terdapat interface, yang mana merupakan kumpulan fungsi abstrak yang perlu diimplementasikan menjadi method yang memiliki definisi fungsi untuk tipe struct mengimplementasikannya
package merupakan sekumpulan kode yang berisi kumpulan fungsi dan data
untuk mendefinisikan access modifier di golang, ditentukan dengan huruf pertama dari penamaan variabel atau function, untuk mendefinisikan variabel atau function yang exported (public) kita perlu menuliskan huruf pertama dengan huruf kapital, sedangkan untuk variabel atau function yang unexported (private) kita perlu menuliskan huruf pertama dengan huruf non kapital
di golang, terdapat panic (exception) dan untuk menanganinya bisa dengan menggunakan recover

Materi tersebut mencakup sejumlah konsep dasar dalam pemrograman Go (Golang). Ini melibatkan pengolahan string, penggunaan fungsi lanjutan, manipulasi data menggunakan struct dan pointer, implementasi metode pada tipe data, definisi dan penggunaan interface untuk polimorfisme, pembuatan dan penggunaan package, serta penanganan kesalahan (error handling) untuk mengelola situasi yang tidak terduga dalam program. Kesemua konsep ini adalah bagian integral dalam pengembangan aplikasi Go yang efisien dan andal.
