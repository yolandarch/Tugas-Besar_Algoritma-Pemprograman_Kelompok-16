package main

import (
	"fmt"
	"strings"
	"time"
)

type Akun struct {
	username, password string
}

type Penelitian struct {
	ID             string
	Model          string
	Judul          string
	Peneliti       [3]string
	Institusi      string
	WaktuMulai     time.Time
	WaktuSelesai   time.Time
	WaktuPublikasi time.Time
	Metode         string
	Summary        string
}

var akun Akun
var penelitianList [100]Penelitian
var penelitianCount int

func main() {
	buatAkun(&akun)
	login(akun)
	menuUtama()
}

func buatAkun(A *Akun) {
	fmt.Println("Create Account")
	fmt.Print("Username: ")
	fmt.Scan(&A.username)
	fmt.Print("Password: ")
	fmt.Scan(&A.password)
	fmt.Println("Akun berhasil dibuat")
}

func login(A Akun) {
	var username, password string
	fmt.Println("Login")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	if username != A.username {
		fmt.Println("Username tidak terdaftar")
		login(A)
	} else {
		fmt.Print("Password: ")
		fmt.Scan(&password)
		if password != A.password {
			fmt.Println("Password salah")
			login(A)
		} else {
			fmt.Println("Login berhasil")
		}
	}
}

func menuUtama() {
	fmt.Println("Menu:")
	fmt.Println("1. Tambahkan Penelitian")
	fmt.Println("2. Lihat Penelitian")
	fmt.Println("3. Lihat Summary Penelitian")
	fmt.Println("4. Lihat Perbandingan Penelitian")
	fmt.Println("5. Edit/Hapus Penelitian")
	fmt.Println("6. Keluar")

	var pilihan int
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		addResearch()
	} else if pilihan == 2 {
		viewResearch()
	} else if pilihan == 3 {
		viewResearchSummary()
	} else if pilihan == 4 {
		compareResearch()
	} else if pilihan == 5 {
		editOrDeleteResearch()
	} else if pilihan == 6 {
		fmt.Println("Keluar")
		return
	} else {
		fmt.Println("Pilihan tidak valid")
		menuUtama()
	}
}

func addResearch() {
	var model, title, institution, method, summary string
	var researchers [3]string
	var startDate, endDate, publicationDate string

	fmt.Println("Tambah Penelitian")
	fmt.Println("Pilih model:")
	fmt.Println("1. Machine Learning")
	fmt.Println("2. Deep Learning")
	fmt.Println("3. Neural Networks")
	fmt.Println("4. Generative Models")
	fmt.Println("5. Reinforcement Learning")
	fmt.Println("6. Supervised Learning")
	fmt.Println("7. Unsupervised Learning")
	var modelPilihan int
	fmt.Print("Model Pilihan: ")
	fmt.Scan(&modelPilihan)
	model = getModelName(modelPilihan)

	fmt.Print("Judul Penelitian: ")
	fmt.Scan(&title)
	for i := 0; i < 3; i++ {
		fmt.Printf("Nama Peneliti %d: ", i+1)
		fmt.Scan(&researchers[i])
	}
	fmt.Print("Lembaga Penelitian (berhenti dengan titik): ")
	fmt.Scan(&institution)

	fmt.Print("Waktu Mulai Penelitian (dd-mm-yyyy): ")
	fmt.Scan(&startDate)
	fmt.Print("Waktu Berakhir Penelitian (dd-mm-yyyy): ")
	fmt.Scan(&endDate)
	fmt.Print("Waktu Publikasi Penelitian (dd-mm-yyyy): ")
	fmt.Scan(&publicationDate)

	start, err := time.Parse("02-01-2006", startDate)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}
	end, err := time.Parse("02-01-2006", endDate)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}
	pub, err := time.Parse("02-01-2006", publicationDate)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}

	fmt.Print("Metode Penelitian: ")
	fmt.Scan(&method)

	fmt.Print("Summary Penelitian (akhiri dengan bintang): ")
	fmt.Scan(&summary)

	id := fmt.Sprintf("%dP", penelitianCount+1)
	newResearch := Penelitian{
		ID:             id,
		Model:          model,
		Judul:          title,
		Peneliti:       researchers,
		Institusi:      institution,
		WaktuMulai:     start,
		WaktuSelesai:   end,
		WaktuPublikasi: pub,
		Metode:         method,
		Summary:        summary,
	}

	penelitianList[penelitianCount] = newResearch
	penelitianCount++
	fmt.Printf("Penelitian %s berhasil ditambahkan\n", title)
	displayResearch(newResearch)
	menuUtama()
}

func getModelName(pilihan int) string {
	models := []string{
		"Machine Learning",
		"Deep Learning",
		"Neural Networks",
		"Generative Models",
		"Reinforcement Learning",
		"Supervised Learning",
		"Unsupervised Learning",
	}
	if pilihan >= 1 && pilihan <= 7 {
		return models[pilihan-1]
	}
	return "Unknown"
}

func displayResearch(r Penelitian) {
	fmt.Println("==============================================================================")
	fmt.Printf("| IdPenelitian       | %s\n", r.ID)
	fmt.Printf("| Model              | %s\n", r.Model)
	fmt.Println("|====================|========================================================")
	fmt.Printf("| Judul Penelitian   | %s\n", r.Judul)
	for i, researcher := range r.Peneliti {
		fmt.Printf("| Nama Peneliti %d    | %s\n", i+1, researcher)
	}
	fmt.Printf("| Lembaga Penelitian | %s\n", r.Institusi)
	fmt.Println("==============================================================================")
}

func viewResearch() {
	fmt.Println("Lihat Penelitian Berdasarkan:")
	fmt.Println("1. Jenis Model")
	fmt.Println("2. Tanggal Publikasi")
	fmt.Println("3. Semua Penelitian")
	var pilihan int
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		var model string
		fmt.Print("Masukkan jenis model: ")
		fmt.Scan(&model)
		filterAndDisplayResearch(func(r Penelitian) bool {
			return strings.EqualFold(r.Model, model)
		})
	} else if pilihan == 2 {
		var date string
		fmt.Print("Masukkan tanggal publikasi (dd-mm-yyyy): ")
		fmt.Scan(&date)
		pubDate, err := time.Parse("02-01-2006", date)
		if err != nil {
			fmt.Println("Format tanggal salah")
			return
		}
		filterAndDisplayResearch(func(r Penelitian) bool {
			return r.WaktuPublikasi.Equal(pubDate)
		})
	} else if pilihan == 3 {
		filterAndDisplayResearch(func(r Penelitian) bool {
			return true
		})
	} else {
		fmt.Println("Pilihan tidak valid")
		menuUtama()
	}
}

func filterAndDisplayResearch(filterFunc func(Penelitian) bool) {
	var filtered [100]Penelitian
	var count int

	for i := 0; i < penelitianCount; i++ {
		if filterFunc(penelitianList[i]) {
			filtered[count] = penelitianList[i]
			count++
		}
	}

	if count == 0 {
		fmt.Println("Tidak ada penelitian yang ditemukan")
		return
	}

	fmt.Println("Pilih Urutan:")
	fmt.Println("1. ID (Binary Sort)")
	fmt.Println("2. Model (Insertion Sort)")
	fmt.Println("3. Judul (Selection Sort)")
	fmt.Println("4. Metode (Sequential Sort)")
	var pengurutanPilihan int
	fmt.Print("Pilih urutan: ")
	fmt.Scan(&pengurutanPilihan)

	if pengurutanPilihan == 1 {
		binarySortByID(filtered[:count])
	} else if pengurutanPilihan == 2 {
		insertionSortByModel(filtered[:count])
	} else if pengurutanPilihan == 3 {
		selectionSortByTitle(filtered[:count])
	} else if pengurutanPilihan == 4 {
		sequentialSortByMethod(filtered[:count])
	} else {
		fmt.Println("Pilihan tidak valid")
		return
	}

	for i := 0; i < count; i++ {
		displayResearch(filtered[i])
	}
	menuUtama()
}

func binarySortByID(penelitian []Penelitian) {
	for i := 1; i < len(penelitian); i++ {
		key := penelitian[i]
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) / 2
			if penelitian[mid].ID < key.ID {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		for j := i - 1; j >= low; j-- {
			penelitian[j+1] = penelitian[j]
		}
		penelitian[low] = key
	}
}

func insertionSortByModel(penelitian []Penelitian) {
	for i := 1; i < len(penelitian); i++ {
		key := penelitian[i]
		j := i - 1
		for j >= 0 && penelitian[j].Model > key.Model {
			penelitian[j+1] = penelitian[j]
			j--
		}
		penelitian[j+1] = key
	}
}

func selectionSortByTitle(penelitian []Penelitian) {
	for i := 0; i < len(penelitian)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(penelitian); j++ {
			if penelitian[j].Judul < penelitian[minIndex].Judul {
				minIndex = j
			}
		}
		penelitian[i], penelitian[minIndex] = penelitian[minIndex], penelitian[i]
	}
}

func sequentialSortByMethod(penelitian []Penelitian) {
	n := len(penelitian)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if penelitian[i].Metode > penelitian[j].Metode {
				penelitian[i], penelitian[j] = penelitian[j], penelitian[i]
			}
		}
	}
}

func viewResearchSummary() {
	var id string
	fmt.Print("Masukkan ID Penelitian: ")
	fmt.Scan(&id)

	for i := 0; i < penelitianCount; i++ {
		if penelitianList[i].ID == id {
			fmt.Println("Summary Penelitian:")
			fmt.Println(penelitianList[i].Summary)
			return
		}
	}
	fmt.Println("Penelitian dengan ID tersebut tidak ditemukan")
	menuUtama()
}

func compareResearch() {
	var id1, id2 string
	fmt.Print("Masukkan ID Penelitian Pertama: ")
	fmt.Scan(&id1)
	fmt.Print("Masukkan ID Penelitian Kedua: ")
	fmt.Scan(&id2)

	var penelitian1, penelitian2 *Penelitian
	for i := 0; i < penelitianCount; i++ {
		if penelitianList[i].ID == id1 {
			penelitian1 = &penelitianList[i]
		}
		if penelitianList[i].ID == id2 {
			penelitian2 = &penelitianList[i]
		}
	}

	if penelitian1 == nil || penelitian2 == nil {
		fmt.Println("Salah satu ID Penelitian tidak ditemukan")
		menuUtama()
		return
	}

	fmt.Println("Perbandingan Penelitian:")
	fmt.Printf("Judul: %s vs %s\n", penelitian1.Judul, penelitian2.Judul)
	fmt.Printf("Model: %s vs %s\n", penelitian1.Model, penelitian2.Model)
	fmt.Printf("Peneliti: %v vs %v\n", penelitian1.Peneliti, penelitian2.Peneliti)
	fmt.Printf("Institusi: %s vs %s\n", penelitian1.Institusi, penelitian2.Institusi)
	fmt.Printf("Metode: %s vs %s\n", penelitian1.Metode, penelitian2.Metode)
	menuUtama()
}

func editOrDeleteResearch() {
	var id string
	fmt.Print("Masukkan ID Penelitian: ")
	fmt.Scan(&id)

	for i := 0; i < penelitianCount; i++ {
		if penelitianList[i].ID == id {
			fmt.Println("Penelitian ditemukan. Pilih aksi:")
			fmt.Println("1. Edit")
			fmt.Println("2. Hapus")
			var aksi int
			fmt.Print("Pilih aksi: ")
			fmt.Scan(&aksi)
			if aksi == 1 {
				editResearch(&penelitianList[i])
			} else if aksi == 2 {
				deleteResearch(i)
			} else {
				fmt.Println("Aksi tidak valid")
				menuUtama()
			}
			return
		}
	}
	fmt.Println("Penelitian dengan ID tersebut tidak ditemukan")
	menuUtama()
}

func editResearch(p *Penelitian) {
	var model, title, institution, method, summary string
	var researchers [3]string
	var startDate, endDate, publicationDate string

	fmt.Println("Edit Penelitian")

	fmt.Println("Pilih model:")
	fmt.Println("1. Machine Learning")
	fmt.Println("2. Deep Learning")
	fmt.Println("3. Neural Networks")
	fmt.Println("4. Generative Models")
	fmt.Println("5. Reinforcement Learning")
	fmt.Println("6. Supervised Learning")
	fmt.Println("7. Unsupervised Learning")
	var modelPilihan int
	fmt.Print("Model Pilihan: ")
	fmt.Scan(&modelPilihan)
	model = getModelName(modelPilihan)

	fmt.Print("Judul Penelitian: ")
	fmt.Scan(&title)
	for i := 0; i < 3; i++ {
		fmt.Printf("Nama Peneliti %d: ", i+1)
		fmt.Scan(&researchers[i])
	}
	fmt.Print("Lembaga Penelitian: ")
	fmt.Scan(&institution)

	fmt.Print("Waktu Mulai Penelitian (dd-mm-yyyy): ")
	fmt.Scan(&startDate)
	fmt.Print("Waktu Berakhir Penelitian (dd-mm-yyyy): ")
	fmt.Scan(&endDate)
	fmt.Print("Waktu Publikasi Penelitian (dd-mm-yyyy): ")
	fmt.Scan(&publicationDate)

	start, err := time.Parse("02-01-2006", startDate)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}
	end, err := time.Parse("02-01-2006", endDate)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}
	pub, err := time.Parse("02-01-2006", publicationDate)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}

	fmt.Print("Metode Penelitian: ")
	fmt.Scan(&method)

	fmt.Print("Summary Penelitian: ")
	fmt.Scan(&summary)

	p.Model = model
	p.Judul = title
	p.Peneliti = researchers
	p.Institusi = institution
	p.WaktuMulai = start
	p.WaktuSelesai = end
	p.WaktuPublikasi = pub
	p.Metode = method
	p.Summary = summary

	fmt.Printf("Penelitian %s berhasil diedit\n", title)
	displayResearch(*p)
	menuUtama()
}

func deleteResearch(index int) {
	for i := index; i < penelitianCount-1; i++ {
		penelitianList[i] = penelitianList[i+1]
	}
	penelitianCount--
	fmt.Println("Penelitian berhasil dihapus")
	menuUtama()
}
