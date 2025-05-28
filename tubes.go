package main

import "fmt"

const NMAX int = 1000

var i int

type team struct {
	namatim, bestplayer string
	menang, selisihSkor int
}

type Team [NMAX]team

func main() {
	var data Team
	var jumlahdata int
	menu(&data, &jumlahdata)
}

func menu(data *Team, jumlahdata *int) {
	var pilih, selisihSkor, menang int
	var namatim, bestplayer string
	var hasil *team

	for {
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Println("        Selamat Datang ke Aplikasi Pengelolaan Data E-Sports Tournament\n      ")
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Println("1. Input data <>")
		fmt.Println("2. Update data <>")
		fmt.Println("3. Delete data <>")
		fmt.Println("4. Tampilan Klasemen Menurut Kemenangan (Selection Sort) <>")
		fmt.Println("5. Tampilan Klasemen Menurut Skor (Insertion Sort) <>")
		fmt.Println("6. Search Nama Team (Sequential Search) <>")
		fmt.Println("7. Search Nama Team (Binary Search) <>")
		fmt.Println("8. Keluar Aplikasi !! <>")
		fmt.Println("Silahkan Pilih Menu yang ingin anda akses :)")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Input Nama data :")
			fmt.Scan(&namatim)
			InputTeam(data, jumlahdata, namatim)

		case 2:
			fmt.Print("Nama data yang ingin diupdate :")
			fmt.Scan(&namatim)
			fmt.Print("Input Kemenangan :")
			fmt.Scan(&menang)
			fmt.Print("Input Best Player :")
			fmt.Scan(&bestplayer)
			UpdateTeam(data, *jumlahdata, namatim, menang, selisihSkor, bestplayer)

		case 3:
			fmt.Print("Nama data yang ingin di delete :")
			fmt.Scan(&namatim)
			DeleteTeam(data, jumlahdata, namatim)

		case 4:
			SelectionSortBerdasarkanMenang(data, *jumlahdata)
			MenampilkanData(data, *jumlahdata)

		case 5:
			InsertionSortBerdasarkanSkor(data, *jumlahdata)
			MenampilkanData(data, *jumlahdata)

		case 6:
			fmt.Print("Input Nama data : ")
			fmt.Scan(&namatim)
			hasil = SequentialSearch(data, *jumlahdata, namatim)
			if hasil != nil {
				fmt.Printf(" data : %s \n Menang :  %d \n Selisih Skor : %d \n Best Player : %s\n", hasil.namatim, hasil.menang, hasil.selisihSkor, hasil.bestplayer)
			} else {
				fmt.Println("data tidak ada")
			}

		case 7:
			fmt.Print("Input Nama data : ")
			fmt.Scan(&namatim)
			hasil = BinarySearch(data, *jumlahdata, namatim)
			if hasil != nil {
				fmt.Printf(" data : %s \n Menang :  %d \n Selisih Skor : %d \n Best Player : %s\n", hasil.namatim, hasil.menang, hasil.selisihSkor, hasil.bestplayer)
			} else {
				fmt.Println("data tidak ada")
			}

		case 8:
			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi Pengelolaan Data E-Sports Tournament")
			return
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}
}

func InputTeam(data *Team, jumlah *int, nama string) {
	data[*jumlah].namatim = nama
	*jumlah++
}

func UpdateTeam(data *Team, jumlah int, nama string, menang int, skor int, player string) {
	for i = 0; i < jumlah; i++ {
		if data[i].namatim == nama {
			data[i].menang = menang
			data[i].selisihSkor = menang * 3
			data[i].bestplayer = player
			return
		}
	}
}

func DeleteTeam(data *Team, jumlah *int, nama string) {
	var j int
	for i = 0; i < *jumlah; i++ {
		if data[i].namatim == nama {
			for j = i; j < *jumlah-1; j++ {
				data[j] = data[j+1]
			}
			*jumlah--
			break
		}
	}
}

func SequentialSearch(data *Team, jumlah int, nama string) *team {
	var found bool = false
	var i int
	i = 0
	for i < jumlah && !found {
		if data[i].namatim == nama {
			found = true
		} else {
			i++
		}
	}
	if found {
		return &data[i]
	} else {
		return nil
	}
}

func BinarySearch(data *Team, jumlah int, nama string) *team {
	var j, kr, kn, med int
	var found bool = false
	var temp team

	kr = 0
	kn = jumlah - 1

	for i = 0; i < jumlah-1; i++ {
		for j = 0; j < jumlah-i-1; j++ {
			if data[j].namatim > data[j+1].namatim {
				temp = data[j]
				data[j] = data[j+1]
				data[j+1] = temp

			}
		}
	}

	for kr <= kn && !found {
		med = (kr + kn) / 2
		if data[med].namatim == nama {
			return &data[med]
		} else if data[med].namatim < nama {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	return nil
}

func SelectionSortBerdasarkanMenang(data *Team, jumlah int) {
	var j, idx_Max int
	var temp team
	for i = 0; i < jumlah-1; i++ {
		idx_Max = i
		for j = i + 1; j < jumlah; j++ {
			if data[j].menang > data[idx_Max].menang {
				idx_Max = j
			}
		}
		temp = data[i]
		data[i] = data[idx_Max]
		data[idx_Max] = temp
	}
}

func InsertionSortBerdasarkanSkor(data *Team, jumlah int) {
	var j int
	var temp team
	for i = 1; i < jumlah; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].selisihSkor < temp.selisihSkor {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func MenampilkanData(data *Team, jumlah int) {
	fmt.Println("\nKlasemen:")
	fmt.Printf("%-3s %-20s %-10s %-15s %-20s\n", "No", "Nama data", "Menang", "Selisih Skor", "Best Player")
	fmt.Println("-------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-3d %-20s %-10d %-15d %-20s\n",
			i+1, data[i].namatim, data[i].menang, data[i].selisihSkor, data[i].bestplayer)
	}
}
