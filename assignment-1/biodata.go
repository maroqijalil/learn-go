package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

type Trainee struct {
	name    string
	address string
	job     string
	reason  string
}

func (trainee Trainee) print() {
	fmt.Println("Nama:", trainee.name)
	fmt.Println("Alamat:", trainee.address)
	fmt.Println("Pekerjaan:", trainee.job)
	fmt.Println("Alasan:", trainee.reason)
}

func main() {
	arg := os.Args[1]
	number, _ := strconv.Atoi(arg)

	trainees := []Trainee{}

	for i := 0; i < 50; i++ {
		job := gofakeit.JobTitle()

		trainees = append(trainees, Trainee{
			name:    gofakeit.Name(),
			address: gofakeit.Address().Address,
			job:     job,
			reason:  fmt.Sprintf("Saya ingin memaksimalkan posisi saya sebagai %s di %s.", job, gofakeit.Company()),
		})
	}

	if (number <= 50) && (number >= 1) {
		trainees[number-1].print()
	} else {
		fmt.Println("Peserta pelatihan dengan no.", number, "tidak ditemukan!")
	}
}
