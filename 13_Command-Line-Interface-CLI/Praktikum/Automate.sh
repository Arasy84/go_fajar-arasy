#!/bin/bash

# Mendapatkan waktu saat ini
current_time=$(date +"%Y-%m-%d %H:%M:%S")

# Membuat folder utama dengan nama yang sesuai dengan waktu saat ini
main_folder="fajar arasy at $current_time"
mkdir "$main_folder"

# Pindah ke dalam folder utama
cd "$main_folder"

# Membuat struktur folder yang diperlukan
mkdir -p "about_me/personal"
mkdir -p "about_me/professional"
mkdir -p "my_friends"
mkdir -p "my_system_info"

# Membuat file facebook.txt dan linkedin.txt
echo "https://www.facebook.com/$2" > "about_me/personal/facebook.txt"
echo "www.linkedin.com/in/$3" > "about_me/professional/linkedin.txt"

# Membuat file list_of_my_friends.txt dengan isi dari teman-teman
curl -o "my_friends/list_of_my_friends.txt" https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt

# Membuat file about_this_laptop.txt dengan informasi uname -a
echo "My username: $1" > "my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "my_system_info/about_this_laptop.txt"

# Membuat file internet_connection.txt dengan hasil ping ke google.com
echo "Connection to google:" > "my_system_info/internet_connection.txt"
ping -c 3 google.com >> "my_system_info/internet_connection.txt"

# Menampilkan struktur folder setelah semua pekerjaan selesai
tree

# Kembali ke direktori awal
cd ..
