#!/bin/bash

# Mengecek apakah argumen sudah diisi
if [ -z "$1" ] || [ -z "$2" ] || [ -z "$3" ]
  then
    echo "Usage: folder_name facebook_username linkedin_username"
    exit 1
fi



# Membuat folder utama dengan nama folder_name dan tanggal
folder_name="$1 at $(date +"%a %b %-d %T %Z %Y")"
mkdir "$folder_name"

# Membuat folder about_me
about_me="$folder_name/about_me"
mkdir "$about_me"

# Membuat folder personal dan file facebook.txt
personal="$about_me/personal"
mkdir "$personal"
echo  "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/facebook.txt/$2" > "$personal/facebook.txt"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/facebook.txt/" > "$personal/facebook.txt"

# Membuat folder professional dan file linkedin.txt
professional="$about_me/professional"
mkdir "$professional"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/linkedin.txt/" > "$professional/linkedin.txt"

# Membuat folder my_friends dan file list_of_my_friends.txt
my_friends="$folder_name/my_friends"
mkdir "$my_friends"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$my_friends/list_of_my_friends.txt"

# Membuat folder my_system_info
my_systems_info="$folder_name/my_systems_info"
mkdir "$my_systems_info"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/about_this_laptop.txt" > "$my_systems_info/about_this_laptop.txt"


# Membuat file internet_connection.txt
internet_connection="$my_systems_info/internet_connection.txt"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/internet_connection.txt" > "$my_systems_info/internet_connection.txt"

ping -c 3 google.com > my_system_info/internet_connection.txt

echo "Folder $folder_name has been created."
