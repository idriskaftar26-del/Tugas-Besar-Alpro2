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
	setoran    [100]Transaksi
	jumlahLog  int
	totalBerat float64
}

var (
	dataWarga   [100]warga
	jumlahWarga int
	jenisSampah = []string{"Sampah Organik", "Sampah Anorganik", "Sampah B3", "Sampah Residu"}
)

func printFeature() {
	fmt.Printf("=========== WASTE-TRACK ===========\n")
	fmt.Printf("|1. Manajemen Data Warga					|\n")
	fmt.Printf("|2. Catat Setoran Sampah Mingguan	|\n")
	fmt.Printf("|3. Cari Data Warga								|\n")
	fmt.Printf("|4. Urutkan Data Warga						|\n")
	fmt.Printf("|5. Tampilkan Statistik Mingguan	|\n")
	fmt.Printf("|9. Exit													|\n")
	fmt.Printf("===================================\n")
}

func evalChoice(choice string, run *bool) { // main menu
	if choice == "1" {
		menuCRUDWarga()
	} else if choice == "2" {
		catatSetoran()
	} else if choice == "3" {
		menuCariWarga()
	} else if choice == "4" {
		menuUrutWarga()
	} else if choice == "5" {
		tampilkanStatistik()
	} else if choice == "9" {
		*run = false
	}
}

func main() {
	var run bool = true
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

func menuCRUDWarga() { // pilihan sub menu
	var subRun bool = true
	for subRun {
		CRUDWarga()
		fmt.Printf("Masukkan pilihan menu warga: ")
		var subChoice string
		fmt.Scan(&subChoice)

		if subChoice == "1" {
			showDataWarga()
		} else if subChoice == "2" {
			tambahWarga()
		} else if subChoice == "3" {
			editWarga()
		} else if subChoice == "4" {
			hapusWarga()
		} else if subChoice == "7" {
			subRun = false
		}
	}
}

func showDataWarga() {
	if jumlahWarga == 0 {
		fmt.Printf("Belum ada data warga.\n")
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

	fmt.Printf("Pilih Jenis Sampah:\n")
	for i := 0; i < 4; i++ {
		fmt.Printf("%d. %s\n", i+1, jenisSampah[i])
	}
	fmt.Printf("Pilihan (1-4): ")
	var pilJenis int
	fmt.Scan(&pilJenis)
	if pilJenis < 1 || pilJenis > 4 {
		fmt.Printf("Pilihan tidak valid.\n")
		return
	}

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
		jenis:     jenisSampah[pilJenis-1],
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
	var tOrganik float64 = 0
	var tAnorganik float64 = 0
	var tB3 float64 = 0
	var tResidu float64 = 0

	for i := 0; i < jumlahWarga; i++ {
		for j := 0; j < dataWarga[i].jumlahLog; j++ {
			tx := dataWarga[i].setoran[j]
			if tx.whichWeek == targetWeek {
				totalSemua += tx.berat
				if tx.jenis == "Sampah Organik" {
					tOrganik += tx.berat
				} else if tx.jenis == "Sampah Anorganik" {
					tAnorganik += tx.berat
				} else if tx.jenis == "Sampah B3" {
					tB3 += tx.berat
				} else if tx.jenis == "Sampah Residu" {
					tResidu += tx.berat
				}
			}
		}
	}

	fmt.Printf("Statistik Akumulasi Sampah Minggu %d:\n", targetWeek)
	fmt.Printf("- Sampah Organik: %.2f kg\n", tOrganik)
	fmt.Printf("- Sampah Anorganik: %.2f kg\n", tAnorganik)
	fmt.Printf("- Sampah B3: %.2f kg\n", tB3)
	fmt.Printf("- Sampah Residu: %.2f kg\n", tResidu)
	fmt.Printf("TOTAL KESELURUHAN: %.2f kg\n", totalSemua)
	fmt.Printf("==============================\n")
}
