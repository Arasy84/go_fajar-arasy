package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
}

Dalam kode di atas, saya telah melakukan beberapa perubahan untuk membuatnya lebih terbaca dan rapi:

1. Menggunakan huruf kapital pada awal kata untuk mengikuti konvensi penamaan yang umum digunakan dalam bahasa pemrograman Go.
2. Menggunakan spasi antara nama tipe data dan kurung kurawal pembuka.
3. Menggunakan spasi sebelum dan setelah operator penugasan (=) untuk meningkatkan keterbacaan.
4. Menggunakan penamaan yang lebih deskriptif untuk fungsi dan variabel.
5. Menggunakan operator penugasan tambah dan sama dengan (+=) untuk mempersingkat penulisan.
6. Menghapus variabel yang tidak digunakan.

Dengan melakukan perubahan ini, kode menjadi lebih mudah dipahami dan mengikuti prinsip clean code.