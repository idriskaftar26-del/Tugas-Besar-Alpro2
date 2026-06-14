package main

import "fmt"

type Transaksi struct {
	date  string
	jenis string
	berat float64
}

type warga struct {
	name       string
	id         int
	setoran    []Transaksi
	jumlahLog  int
	totalBerat float64
}

var (
	dataWarga   []warga
	jenisSampah []string
)

func printFeature() {
	fmt.Printf("=========== WASTE-TRACK ===========\n")
	fmt.Printf("|1. Manajemen Data Warga          |\n")
	fmt.Printf("|2. Catat Setoran Sampah          |\n")
	fmt.Printf("|3. Cari Data Warga               |\n")
	fmt.Printf("|4. Urutkan Data Warga            |\n")
	fmt.Printf("|5. Tampilkan Statistik           |\n")
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
	if len(dataWarga) == 0 {
		fmt.Printf("Data warga kosong\n")
		return
	}
	for i := 0; i < len(dataWarga); i++ {
		fmt.Printf(
			"Nama : %s\nID : %d\nJumlah Log : %d\nTotal Berat : %.2f\n\n",
			dataWarga[i].name, dataWarga[i].id, dataWarga[i].jumlahLog, dataWarga[i].totalBerat,
		)
	}
}

func findIndexByID(id int) int { // cari index dari warga berdasarkan id
	for i := 0; i < len(dataWarga); i++ {
		if dataWarga[i].id == id {
			return i
		}
	}
	return -1
}

func tambahWarga() {
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

	dataWarga = append(dataWarga, newWarga)
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

	dataWarga = append(dataWarga[:idx], dataWarga[idx+1:]...) // updating the dataWarga slice (upacking)
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

	fmt.Printf("\n=== Jenis Sampah yang Terdata ===\n")
	if len(jenisSampah) == 0 {
		fmt.Printf("Belum ada jenis sampah yang terdata\n")
	} else {
		for i, sampah := range jenisSampah {
			fmt.Printf("%d. %s\n", i+1, sampah)
		}
	}

	fmt.Printf("===============================\n")
	fmt.Printf("Pilih/Ketik Jenis Sampah (Ketik nama baru jika belum terdaftar):\n")
	var pilJenis string
	fmt.Scan(&pilJenis)

	//check if jenis sampah already exist
	exist := false
	for _, jenis := range jenisSampah {
		if jenis == pilJenis {
			exist = true
			break
		}
	}

	if !exist {
		jenisSampah = append(jenisSampah, pilJenis)
		fmt.Printf("Jenis sampah baru '%s' berhasil ditambahkan\n", pilJenis)
	}

	fmt.Printf("Masukkan Berat Sampah (kg): ")
	var berat float64
	fmt.Scan(&berat)

	var tgl, bln, thn int
	fmt.Printf("Masukkan Tanggal (dd mm yyyy): ")
	fmt.Scan(&tgl, &bln, &thn)

	tanggalFormat := fmt.Sprintf("%02d-%02d-%04d", tgl, bln, thn) // using sprintf to format the date

	newTx := Transaksi{
		date:  tanggalFormat,
		jenis: pilJenis,
		berat: berat,
	}

	dataWarga[idx].setoran = append(dataWarga[idx].setoran, newTx)
	dataWarga[idx].jumlahLog++
	dataWarga[idx].totalBerat += berat
	fmt.Printf("Setoran Sampah berhasil dicatat pada tanggal %s\n", tanggalFormat)

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
	for i := 0; i < len(dataWarga); i++ {
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
	//making slice tempData from dataWarga
	tempData := make([]warga, len(dataWarga))

	// copy it
	copy(tempData, dataWarga)

	// sorting the tempData to prevent sorting the original dataWarga
	for i := 0; i < len(tempData)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(tempData); j++ {
			if tempData[j].id < tempData[minIdx].id {
				minIdx = j
			}
		}
		tempData[i], tempData[minIdx] = tempData[minIdx], tempData[i]
	}

	// Binary Search on the tempData
	low := 0
	high := len(tempData) - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if tempData[mid].id == queryID {
			fmt.Printf("[Ditemukan via Binary] ID: %d | Nama: %s | Total: %.2f kg\n",
				tempData[mid].id, tempData[mid].name, tempData[mid].totalBerat)
			found = true
			break
		} else if tempData[mid].id < queryID {
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
	for i := 0; i < len(dataWarga)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(dataWarga); j++ {
			if dataWarga[j].totalBerat < dataWarga[minIdx].totalBerat {
				minIdx = j
			}
		}
		dataWarga[i], dataWarga[minIdx] = dataWarga[minIdx], dataWarga[i]
	}
}

func insertionSort() {
	for i := 1; i < len(dataWarga); i++ {
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
	fmt.Printf("\n=== WASTE-TRACK STATISTIK ===\n")
	fmt.Printf("1. Statistik Berdasarkan Bulan\n")
	fmt.Printf("2. Statistik Berdasarkan Tahun\n")
	fmt.Printf("Pilih filter statistik (1-2): ")
	var pilihan string
	fmt.Scan(&pilihan)

	var targetPeriode string
	var labelPeriode string

	if pilihan == "1" {
		var bln, thn int
		fmt.Printf("Masukkan Bulan Statistik (1-12): ")
		fmt.Scan(&bln)
		fmt.Printf("Masukkan Tahun Statistik (Contoh: 2026): ")
		fmt.Scan(&thn)

		// Format target: "-MM-YYYY" (e.g., "-06-2026")
		targetPeriode = fmt.Sprintf("-%02d-%04d", bln, thn)
		labelPeriode = fmt.Sprintf("Bulan %02d-%d", bln, thn)

	} else if pilihan == "2" {
		var thn int
		fmt.Printf("Masukkan Tahun Statistik (Contoh: 2026): ")
		fmt.Scan(&thn)

		// Format target: "-YYYY" (e.g., "-2026")
		targetPeriode = fmt.Sprintf("-%04d", thn)
		labelPeriode = fmt.Sprintf("Tahun %d", thn)

	} else {
		fmt.Printf("Pilihan tidak valid.\n")
		return
	}

	var totalSemua float64 = 0
	beratJenisSampah := []float64{}

	for i := 0; i < len(dataWarga); i++ {
		for j := 0; j < len(dataWarga[i].setoran); j++ {
			tx := dataWarga[i].setoran[j]

			// Logika slicing string fleksibel:
			// Jika filter Bulan: cek apakah 7 karakter terakhir cocok dengan "-06-2026"
			// Jika filter Tahun: cek apakah 5 karakter terakhir cocok dengan "-2026"
			if len(tx.date) >= len(targetPeriode) && tx.date[len(tx.date)-len(targetPeriode):] == targetPeriode {
				totalSemua += tx.berat

				for k := 0; k < len(jenisSampah); k++ {
					if tx.jenis == jenisSampah[k] {
						beratJenisSampah[k] += tx.berat
					}
				}
			}
		}
	}

	fmt.Printf("\nStatistik Akumulasi Sampah %s:\n", labelPeriode)
	if len(jenisSampah) == 0 {
		fmt.Printf("Tidak ada transaksi pada periode ini.\n")
	} else {
		for i := 0; i < len(jenisSampah); i++ {
			fmt.Printf("- Sampah %s: %.2f kg\n", jenisSampah[i], beratJenisSampah[i])
		}
	}
	fmt.Printf("TOTAL KESELURUHAN: %.2f kg\n\n", totalSemua)
}
