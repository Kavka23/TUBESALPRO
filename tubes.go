package main

import "fmt"

type team struct {
	namatim, bestplayer string
	menang, selisihsk   int
}

var i, it int
var tim [1000]team
var jumlaht int

func main() {
	menu()
}

func menu() {
	var pilih, selisihsk, menang int
	var namatim, bestplayer string

	for {
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Println("        Selamat Datang ke Aplikasi Pengelolaan Data E-Sports Tournament\n      ")
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Println("1. Input Tim <>")
		fmt.Println("2. Update Tim <>")
		fmt.Println("3. Delete Tim <>")
		fmt.Println("4. Tampilan Klasemen Menurut Kemenangan (Selection Sort) <>")
		fmt.Println("5. Tampilan Klasemen Menurut Skor (Insertion Sort) <>")
		fmt.Println("6. Search Nama Tim (Sequential Search) <>")
		fmt.Println("7. Search Nama Tim (Binary Search) <>")
		fmt.Println("8. Keluar Aplikasi !! <>")
		fmt.Println("Silahkan Pilih Menu yang ingin anda akses :)")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Input Nama Tim :")
			fmt.Scan(&namatim)
			input(namatim)

		case 2:
			fmt.Print("Nama tim yang ingin diupdate :")
			fmt.Scan(&namatim)
			fmt.Print("Input Kemenangan :")
			fmt.Scan(&menang)
			fmt.Print("Input Best Player :")
			fmt.Scan(&bestplayer)
			fmt.Print("Input Selisih Skor :")
			fmt.Scan(&selisihsk)
			update(namatim, menang, selisihsk, bestplayer)

		case 3:
			fmt.Print("Nama tim yang ingin di delete :")
			fmt.Scan(&namatim)
			deletetim(namatim)

		case 4:
			sortmenang()
			tampilkan()

		case 5:
			sortskor()
			tampilkan()

		case 6:
			fmt.Print("Input Nama tim : ")
			fmt.Scan(&namatim)
			ti := seqsearch(namatim)
			if ti != nil {
				fmt.Printf(" TIM : %s \n Menang :  %d \n Selisih Skor : %d \n Best Player : %s\n", ti.namatim, ti.menang, ti.selisihsk, ti.bestplayer)
			} else {
				fmt.Println("Tim tidak ada")
			}

		case 7:
			fmt.Print("Input Nama tim : ")
			fmt.Scan(&namatim)
			ti := binsearch(namatim)
			if ti != nil {
				fmt.Printf(" TIM : %s \n Menang :  %d \n Selisih Skor : %d \n Best Player : %s\n", ti.namatim, ti.menang, ti.selisihsk, ti.bestplayer)
			} else {
				fmt.Println("Tim tidak ada")
			}

		case 8:
			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi Pengelolaan Data E-Sports Tournament")
			return
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}
}

func input(namatim string) {
	tim[jumlaht] = team{namatim: namatim}
	jumlaht = jumlaht + 1
}

func update(namatim string, menang, selisihsk int, bestplayer string) {
	for i = 0; i < jumlaht; i++ {
		if tim[i].namatim == namatim {
			tim[i].menang = menang
			tim[i].selisihsk = selisihsk
			tim[i].bestplayer = bestplayer
			return
		}
	}
}

func deletetim(namatim string) {
	for i = 0; i < jumlaht; i++ {
		if tim[i].namatim == namatim {
			for it = i; it < jumlaht-1; it++ {
				tim[it] = tim[it+1]
			}
			jumlaht = jumlaht - 1
			return
		}
	}
}

func seqsearch(namatim string) *team {
	for i = 0; i < jumlaht; i++ {
		if tim[i].namatim == namatim {
			return &tim[i]
		}
	}
	return nil
}

func binsearch(namatim string) *team {
	var l, m, h int
	for i = 0; i < jumlaht-1; i++ {
		for it = 0; it < jumlaht-i-1; it++ {
			if tim[it].namatim > tim[it+1].namatim {
				temp := tim[it]
				tim[it] = tim[it+1]
				tim[it+1] = temp

			}
		}
	}
	l = 0
	h = jumlaht - 1

	for l <= h {
		m = (l + h) / 2
		if tim[m].namatim == namatim {
			return &tim[m]
		} else if tim[m].namatim < namatim {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return nil
}

func sortmenang() {
	var maxIdx, j int
	for i = 0; i < jumlaht-1; i++ {
		maxIdx = i
		for j = i + 1; j < jumlaht; j++ {
			if tim[j].menang > tim[maxIdx].menang {
				maxIdx = j
			}
		}
		tim[i], tim[maxIdx] = tim[maxIdx], tim[i]
	}
}

func sortskor() {
	var j int
	for i = 1; i < jumlaht; i++ {
		key := tim[i]
		j = i - 1
		for j >= 0 && tim[j].selisihsk < key.selisihsk {
			tim[j+1] = tim[j]
			j = j - 1
		}
		tim[j+1] = key
	}
}

func tampilkan() {
	fmt.Println("\nKlasemen:")
	fmt.Printf("%-3s %-20s %-10s %-15s %-20s\n", "No", "Nama Tim", "Menang", "Selisih Skor", "Best Player")
	fmt.Println("-------------------------------------------------------------------------------")
	for i := 0; i < jumlaht; i++ {
		fmt.Printf("%-3d %-20s %-10d %-15d %-20s\n",
			i+1, tim[i].namatim, tim[i].menang, tim[i].selisihsk, tim[i].bestplayer)
	}
}
