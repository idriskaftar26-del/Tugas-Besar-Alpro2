package main

import "fmt"

type Transaksi struct {
	whichWeek int
	jenis     string
	berat     float64
}

type warga struct {
	name       string
	id         int
	setoran    [1000]Transaksi
	jumlahLog  int
	totalBerat float64
}

var (
	dataWarga   [1000]warga
	jumlahWarga int
)

func printFeature() {
	fmt.Printf("=========== WASTE-TRACK ===========\n")
	fmt.Printf("|1. Manajemen Data Warga          |\n")
	fmt.Printf("|2. Catat Setoran Sampah Mingguan |\n")
	fmt.Printf("|3. Cari Data Warga               |\n")
	fmt.Printf("|4. Urutkan Data Warga            |\n")
	fmt.Printf("|5. Tampilkan Statistik Mingguan  |\n")
	fmt.Printf("|9. Exit                          |\n")
	fmt.Printf("===================================\n")
}

func evalChoice(choice string, run *bool) { // main menu
	switch choice {
	case "1":
		menuCRUDWarga()
	case "2":
		catatSetoran()
	case "3":
		menuCariWarga()
	case "4":
		menuUrutWarga()
	case "5":
		tampilkanStatistik()
	case "9":
		*run = false
	}
}

func main() {
	run := true
	for run {
		var choice string
		printFeature()
		fmt.Printf("masukkan pilihan : ")
		fmt.Scan(&choice)

		evalChoice(choice, &run)
	}
	fmt.Printf("Terimakasih\n")
}

func CRUDWarga() { // sub menu
	fmt.Printf("\n=== MANAJEMEN DATA WARGA ===\n")
	fmt.Printf("1. Tampilkan Data Warga\n")
	fmt.Printf("2. Tambah Warga\n")
	fmt.Printf("3. Edit Data Warga\n")
	fmt.Printf("4. Hapus Data Warga\n")
	fmt.Printf("7. Exit\n")
}

func menuCRUDWarga() { // sub menu choice
	subRun := true
	for subRun {
		CRUDWarga()
		fmt.Printf("Masukkan pilihan menu warga: ")
		var subChoice string
		fmt.Scan(&subChoice)

		switch subChoice {
		case "1":
			showDataWarga()
		case "2":
			tambahWarga()
		case "3":
			editWarga()
		case "4":
			hapusWarga()
		case "7":
			subRun = false
		}
	}
}

func showDataWarga() { // printing data warga
	if jumlahWarga == 0 {
		fmt.Printf("Data warga kosong\n")
		return
	}
	for i := 0; i < jumlahWarga; i++ {
		fmt.Printf(
			"Nama : %s\nID : %d\nJumlah Log : %d\nTotal Berat : %.2f\n\n",
			dataWarga[i].name, dataWarga[i].id, dataWarga[i].jumlahLog, dataWarga[i].totalBerat,
		)
	}
}

func findIndexByID(id int) int { // cari index dari warga berdasarkan id
	for i := 0; i < jumlahWarga; i++ {
		if dataWarga[i].id == id {
			return i
		}
	}
	return -1
}

func tambahWarga() {
	if jumlahWarga >= 100 {
		fmt.Printf("Kapasitas penuh\n")
		return
	}
	var newWarga warga
	fmt.Printf("Masukkan ID Warga (Angka): ")
	fmt.Scan(&newWarga.id)

	if findIndexByID(newWarga.id) != -1 {
		fmt.Printf("ID sudah terdaftar\n")
		return
	}

	fmt.Printf("Masukkan Nama Warga: ")
	fmt.Scan(&newWarga.name)
	newWarga.jumlahLog = 0
	newWarga.totalBerat = 0.0

	dataWarga[jumlahWarga] = newWarga
	jumlahWarga++
	fmt.Printf("Warga berhasil ditambahkan\n")
}

func editWarga() {
	fmt.Printf("Masukkan ID Warga yang akan diedit: ")
	var targetID int
	fmt.Scan(&targetID)

	idx := findIndexByID(targetID)
	if idx == -1 {
		fmt.Printf("Warga tidak ditemukan\n")
		return
	}

	fmt.Printf("Masukkan Nama Baru: ")
	fmt.Scan(&dataWarga[idx].name)
	fmt.Printf("Data warga berhasil diperbarui\n")
}

func hapusWarga() {
	fmt.Printf("Masukkan ID Warga yang akan dihapus: ")
	var targetID int
	fmt.Scan(&targetID)

	idx := findIndexByID(targetID)
	if idx == -1 {
		fmt.Printf("Warga tidak ditemukan\n")
		return
	}

	for i := idx; i < jumlahWarga-1; i++ {
		dataWarga[i] = dataWarga[i+1]
	}
	jumlahWarga--
	fmt.Printf("Data warga berhasil dihapus.\n")
}

func catatSetoran() {
	fmt.Printf("Masukkan ID Warga: ")
	var targetID int
	fmt.Scan(&targetID)

	idx := findIndexByID(targetID)
	if idx == -1 {
		fmt.Printf("Warga tidak ditemukan.\n")
		return
	}

	fmt.Printf("Pilih Jenis Sampah (organik, anorganik, b3):\n")
	var pilJenis string
	fmt.Scan(&pilJenis)

	fmt.Printf("Masukkan Berat Sampah (kg): ")
	var berat float64
	fmt.Scan(&berat)

	fmt.Printf("Masukkan Minggu Ke (1-4): ")
	var mgg int
	fmt.Scan(&mgg)

	posLog := dataWarga[idx].jumlahLog
	if posLog >= 100 {
		fmt.Printf("Log transaksi penuh untuk warga ini.\n")
		return
	}

	dataWarga[idx].setoran[posLog] = Transaksi{
		whichWeek: mgg,
		jenis:     pilJenis,
		berat:     berat,
	}
	dataWarga[idx].jumlahLog++
	dataWarga[idx].totalBerat += berat
	fmt.Printf("Setoran sampah berhasil dicatat.\n")
}

func menuCariWarga() {
	fmt.Printf("\n=== PENCARIAN DATA WARGA ===\n")
	fmt.Printf("1. Cari Berdasarkan Nama\n")
	fmt.Printf("2. Cari Berdasarkan ID\n")
	fmt.Printf("Pilih pencarian (1-2): ")
	var sub string
	fmt.Scan(&sub)

	if sub == "1" {
		fmt.Printf("Masukkan Nama Warga yang dicari: ")
		var query string
		fmt.Scan(&query)
		sequentialSearch(query)
	} else if sub == "2" {
		fmt.Printf("Masukkan ID Warga yang dicari: ")
		var queryID int
		fmt.Scan(&queryID)
		binarySearch(queryID)
	}
}

func sequentialSearch(query string) {
	found := false
	for i := 0; i < jumlahWarga; i++ {
		if dataWarga[i].name == query {
			fmt.Printf("[Ditemukan] ID: %d | Nama: %s | Total: %.2f kg\n", dataWarga[i].id, dataWarga[i].name, dataWarga[i].totalBerat)
			found = true
		}
	}
	if !found {
		fmt.Printf("Data warga tidak ditemukan.\n")
	}
}

func binarySearch(queryID int) {
	for i := 0; i < jumlahWarga-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahWarga; j++ {
			if dataWarga[j].id < dataWarga[minIdx].id {
				minIdx = j
			}
		}
		dataWarga[i], dataWarga[minIdx] = dataWarga[minIdx], dataWarga[i]
	}

	low := 0
	high := jumlahWarga - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if dataWarga[mid].id == queryID {
			fmt.Printf("[Ditemukan via Binary] ID: %d | Nama: %s | Total: %.2f kg\n", dataWarga[mid].id, dataWarga[mid].name, dataWarga[mid].totalBerat)
			found = true
			break
		} else if dataWarga[mid].id < queryID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Printf("Data warga tidak ditemukan.\n")
	}
}

func menuUrutWarga() {
	fmt.Printf("\n=== PENGURUTAN DATA WARGA ===\n")
	fmt.Printf("1. Urutkan dengan Ascending)\n")
	fmt.Printf("2. Urutkan dengan Descending\n")
	fmt.Printf("Pilih metode (1-2): ")
	var sub string
	fmt.Scan(&sub)

	if sub == "1" {
		selectionSort()
		fmt.Printf("Data diurutkan dengan Selection Sort secara Ascending.\n")
		showDataWarga()
	} else if sub == "2" {
		insertionSort()
		fmt.Printf("Data diurutkan dengan Insertion Sort secara Descending.\n")
		showDataWarga()
	}
}

func selectionSort() {
	for i := 0; i < jumlahWarga-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahWarga; j++ {
			if dataWarga[j].totalBerat < dataWarga[maxIdx].totalBerat {
				maxIdx = j
			}
		}
		dataWarga[i], dataWarga[maxIdx] = dataWarga[maxIdx], dataWarga[i]
	}
}

func insertionSort() {
	for i := 1; i < jumlahWarga; i++ {
		key := dataWarga[i]
		j := i - 1
		for j >= 0 && dataWarga[j].totalBerat < key.totalBerat {
			dataWarga[j+1] = dataWarga[j]
			j = j - 1
		}
		dataWarga[j+1] = key
	}
}

func tampilkanStatistik() {
	fmt.Printf("\n=== WASTE-TRACK ===\n")
	fmt.Printf("Masukkan Statistik Minggu Ke (1-4): ")
	var targetWeek int
	fmt.Scan(&targetWeek)

	var totalSemua float64 = 0

	jenisSampah := []string{}       // trash type
	beratJenisSampah := []float64{} // weight of each type of trash

	for i := 0; i < jumlahWarga; i++ {
		for j := 0; j < dataWarga[i].jumlahLog; j++ {
			tx := dataWarga[i].setoran[j]
			exist := false
			for k := 0; k < len(jenisSampah); k++ {
				if jenisSampah[k] == tx.jenis {
					exist = true
					break
				}
			}

			if !exist { // if the type of the trash not exist, add it
				jenisSampah = append(jenisSampah, tx.jenis)
				beratJenisSampah = append(beratJenisSampah, 0)
			}
			if tx.whichWeek == targetWeek {
				totalSemua += tx.berat
				for k := 0; k < len(jenisSampah); k++ {
					if tx.jenis == jenisSampah[k] {
						beratJenisSampah[k] += tx.berat
					}
				}
			}
		}
	}

	fmt.Printf("Statistik Akumulasi Sampah Minggu %d:\n", targetWeek)
	for i := 0; i < len(jenisSampah); i++ {
		fmt.Printf("- Sampah %s: %.2f\n", jenisSampah[i], beratJenisSampah[i])
	}
	fmt.Printf("TOTAL KESELURUHAN: %.2f kg\n", totalSemua)
}
