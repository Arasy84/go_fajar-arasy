dalam pemrograman, terdapat beberapa jenis proses, diantaranya adalah sequential, parallel, dan concurrent
sequential akan menjalankan proses secara synchronous
parallel dan concurrent akan menjalankan proses secara asynchronous, tetapi parallel menggunakan jumlah cpu core resources, sedangkan concurrent juga bisa seperti parallel, namun cara kerjanya yaitu dengan membagi proses menjadi bagian-bagian tertentu
concurrent digunakan untuk melakukan proses yang besar dan berat, seperti api load, data processing, dsb
di golang, sebuah fungsi dapat dijalankan secara concurrent dengan menambah keyword go di pemanggilan fungsinya, dan fungsi yang dijalankan secara concurrent disebut goroutine
untuk melakukan komunikasi antar goroutine misalnya untuk pertukaran data, bisa menggunakan channel
untuk melakukan pengelolaan terhadap goroutine bisa dengan menggunakan waitgroup
untuk mengatasi antar goroutine bentrok dalam mengakses data yang sama, bisa menggunakan mutex