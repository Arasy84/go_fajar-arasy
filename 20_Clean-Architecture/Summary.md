Clean Architecture adalah konsep desain perangkat lunak yang mempromosikan pemisahan perhatian dan mengikuti prinsip SOLID untuk pemeliharaan yang lebih mudah, skalabilitas, dan mengurangi ketergantungan pada kerangka kerja atau pustaka.

Prinsip dasar Clean Architecture:

    -Encapsulation: Setiap komponen memiliki tanggung jawabnya sendiri dan tidak boleh bergantung pada komponen lain.
    -Dependency Inversion Principle (DIP): Komponen tinggi tingkat ketergantungan pada komponen rendah tingkat.
    -Interface Segregation Principle (ISP): Setiap komponen harus memiliki satu interface yang digunakan oleh komponen lain.
    -Open/Closed Principle (OCP): Komponen harus terbuka untuk modifikasi, tetapi tertutup untuk ekspansi.
    -Liskov Substitution Principle (LSP): Setiap subclass harus dapat menggantikan superclass-nya.

Software testing merupakan suatu proses menganalisis suatu fitur untuk menemukan perbedaan antara hasil aktual dengan hasil yang diperlukan, dan untuk mengevaluasi fitur itu sendiri
Tujuan dari testing: mencegah regresi, meningkatkan rasa percaya diri dalam melakukan refactoring, meningkatkan desain kode, sebagai dokumentasi kode, dan scaling tim
Level testing: unit (single logical operation) -> integration (module/subsystem) -> ui (end to end)
Testify adalah salah satu framework testing golang
Pola umum: meletakkan semua file testing ke dalam satu folder yang sama, atau meletakkan file test di folder yang sama dengan file aslinya.

Keuntungan Clean Architecture

Clean Architecture memiliki beberapa keuntungan, antara lain:

    -Mudah dipelihara: Pemisahan perhatian membuat aplikasi lebih mudah dipelihara dan dimodifikasi.
    -Dapat diskalakan: Aplikasi dapat diskalakan dengan menambahkan komponen baru ke lapisan yang sesuai.
    -Tidak bergantung pada kerangka kerja atau pustaka: Aplikasi dapat diimplementasikan menggunakan kerangka kerja atau pustaka apa pun.
    
Kesimpulan

Clean Architecture adalah pendekatan desain perangkat lunak yang dapat digunakan untuk membuat aplikasi yang lebih mudah dipelihara, dapat diskalakan, dan tidak terlalu bergantung pada kerangka kerja atau pustaka.