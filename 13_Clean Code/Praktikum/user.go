package main

type user struct {
	id       int
	username int //seharusnya username string
	password int //seharusnya password string
}

//pada line ke-3, username dan password seharusnya bertipe data string

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user { //func (u *userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user { //func (u *userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
//pada line ke-14 dan ke-18, method getallusers dan getuserbyid diganti menggunakan tipe pointer (*userservice) agar data pada userservice bisa diubah pada saat method tersebut dipanggil.

	return user{}
}

//data pada userservice bisa diubah pada saat method tersebut dipanggil menggunakan pointer receiver.
