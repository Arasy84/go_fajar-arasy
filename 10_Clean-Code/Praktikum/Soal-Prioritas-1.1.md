- Berapa banyak kekurangan dalam penulisan kode tersebut?
  Kode tersebut memiliki 4 kekurangan berikut bagian dan alasanya:
    Kekurangan: Penamaan variabel yang tidak deskriptif.
        Bagian: Pada struct user, username dan password seharusnya memiliki tipe data string bukan int.
        Alasan: Penamaan variabel yang tidak deskriptif dapat menyebabkan kebingungan dan sulit dipahami oleh pengembang lain yang membaca kode tersebut. Selain itu, tipe data yang tidak sesuai juga dapat menyebabkan kesalahan dalam penggunaan variabel.

    Kekurangan: Penamaan struct userservice tidak mengikuti konvensi penamaan yang umum.
        Bagian: Pada struct userservice.
        Alasan: Mengikuti konvensi penamaan yang umum membantu meningkatkan keterbacaan kode dan memudahkan pengembang lain untuk memahami struktur kode tersebut.

    Kekurangan: Metode getallusers dan getuserbyid tidak menggunakan pointer receiver.
        Bagian: Pada metode getallusers dan getuserbyid.
        Alasan: Menggunakan pointer receiver memungkinkan perubahan pada objek yang dipanggil dan menghindari penggunaan memori yang berlebihan saat menyalin objek.

    Kekurangan: Variabel t pada struct userservice tidak diexport.
        Bagian: Pada variabel t pada struct userservice.
        Alasan: Variabel yang tidak diexport tidak dapat diakses dari luar package. Jika variabel ini perlu diakses oleh package lain, maka perlu diexport dengan menggunakan huruf kapital pada awal nama variabel.

Berikut adalah kode yg telah di revisi menggunakan prinsip clean code:

package main

type User struct {
    ID       int
    Username string
    Password string
}

type UserService struct {
    sers []User
}

func (u *UserService) GetAllUsers() []User {
    eturn u.Users
}

func (u *UserService) GetUserByID(id int) User {
    for _, r := range u.Users {
    if id == r.ID {
        return r
        }
    }

    return User{}
}
