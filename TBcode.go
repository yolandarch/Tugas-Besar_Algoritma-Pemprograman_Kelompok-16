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
	ID              string
	Model           string
	Judul           string
	Peneliti        [3]string
	Institusi       string
	WaktuMulai      time.Time
	WaktuSelesai    time.Time
	WaktuPublikasi  time.Time
	Metode          string
	Summary         string
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
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambahkan Penelitian")
		fmt.Println("2. Lihat Penelitian")
		fmt.Println("3. Lihat Summary Penelitian")
		fmt.Println("4. Lihat Perbandingan Penelitian")
		fmt.Println("5. Edit/Hapus Penelitian")
		fmt.Println("6. Kembali")

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
			fmt.Println("Kembali ke menu utama")
			return
		} else {
			fmt.Println("Pilihan tidak valid")
		}
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
	}
	menuUtama()
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
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	var pengurutanPilihan int
	fmt.Print("Pilih urutan: ")
	fmt.Scan(&pengurutanPilihan)

	if pengurutanPilihan == 1 {
		for i := 0; i < count-1; i++ {
			for j := 0; j < count-i-1; j++ {
				if filtered[j].Judul > filtered[j+1].Judul {
					filtered[j], filtered[j+1] = filtered[j+1], filtered[j]
				}
			}
		}
	} else {
		for i := 0; i < count-1; i++ {
			for j := 0; j < count-i-1; j++ {
				if filtered[j].Judul < filtered[j+1].Judul {
					filtered[j], filtered[j+1] = filtered[j+1], filtered[j]
				}
			}
		}
	}

	for i := 0; i < count; i++ {
		displayResearch(filtered[i])
	}
}

func viewResearchSummary() {
	var idOrTitle string
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian: ")
	fmt.Scan(&idOrTitle)

	for i := 0; i < penelitianCount; i++ {
		if strings.EqualFold(penelitianList[i].ID, idOrTitle) || strings.EqualFold(penelitianList[i].Judul, idOrTitle) {
			fmt.Println("==============================================================================")
			fmt.Printf("| IdPenelitian       | %s\n", penelitianList[i].ID)
			fmt.Printf("| Model              | %s\n", penelitianList[i].Model)
			fmt.Println("|====================|========================================================")
			fmt.Printf("| Judul Penelitian   | %s\n", penelitianList[i].Judul)
			for j, researcher := range penelitianList[i].Peneliti {
				fmt.Printf("| Nama Peneliti %d    | %s\n", j+1, researcher)
			}
			fmt.Printf("| Summary Penelitian | %s\n", penelitianList[i].Summary)
			fmt.Println("==============================================================================")
			menuUtama()
			return
		}
	}
	fmt.Println("Penelitian tidak ditemukan")
	menuUtama()
}

func compareResearch() {
	var idOrTitle1, idOrTitle2 string
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian pertama: ")
	fmt.Scan(&idOrTitle1)
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian kedua: ")
	fmt.Scan(&idOrTitle2)

	var r1, r2 *Penelitian
	for i := 0; i < penelitianCount; i++ {
		if strings.EqualFold(penelitianList[i].ID, idOrTitle1) || strings.EqualFold(penelitianList[i].Judul, idOrTitle1) {
			r1 = &penelitianList[i]
		}
		if strings.EqualFold(penelitianList[i].ID, idOrTitle2) || strings.EqualFold(penelitianList[i].Judul, idOrTitle2) {
			r2 = &penelitianList[i]
		}
	}

	if r1 == nil || r2 == nil {
		fmt.Println("Salah satu atau kedua penelitian tidak ditemukan")
		menuUtama()
		return
	}

	fmt.Println("==============================================================================")
	fmt.Printf("| Perbandingan antara %s dan %s\n", r1.Judul, r2.Judul)
	fmt.Println("==============================================================================")
	duration1 := r1.WaktuPublikasi.Sub(r1.WaktuMulai)
	duration2 := r2.WaktuPublikasi.Sub(r2.WaktuMulai)

	fmt.Printf("| %s: %v hari\n", r1.Judul, duration1.Hours()/24)
	fmt.Printf("| %s: %v hari\n", r2.Judul, duration2.Hours()/24)
	if duration1 < duration2 {
		fmt.Printf("| %s lebih cepat dipublikasikan\n", r1.Judul)
	} else if duration1 > duration2 {
		fmt.Printf("| %s lebih cepat dipublikasikan\n", r2.Judul)
	} else {
		fmt.Println("| Kedua penelitian memiliki durasi yang sama hingga publikasi")
	}
	fmt.Println("==============================================================================")
	menuUtama()
}

func editOrDeleteResearch() {
	var idOrTitle string
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian yang akan diubah/hapus: ")
	fmt.Scan(&idOrTitle)

	for i := 0; i < penelitianCount; i++ {
		if strings.EqualFold(penelitianList[i].ID, idOrTitle) || strings.EqualFold(penelitianList[i].Judul, idOrTitle) {
			fmt.Println("Penelitian ditemukan. Pilih opsi:")
			fmt.Println("1. Edit Penelitian")
			fmt.Println("2. Hapus Penelitian")
			var pilihan int
			fmt.Print("Pilih opsi: ")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				editResearch(&penelitianList[i])
				return
			} else if pilihan == 2 {
				deleteResearch(i)
				return
			} else {
				fmt.Println("Pilihan tidak valid")
				menuUtama()
				return
			}
		}
	}
	fmt.Println("Penelitian tidak ditemukan")
	menuUtama()
}

func editResearch(r *Penelitian) {
	fmt.Println("Masukkan data baru (kosongkan untuk tidak mengubah):")
	var model, title, institution, method, summary string
	var researchers [3]string
	var startDate, endDate, publicationDate string

	fmt.Print("Model Penelitian (sekarang: " + r.Model + "): ")
	fmt.Scan(&model)
	if model != "" {
		r.Model = model
	}

	fmt.Print("Judul Penelitian (sekarang: " + r.Judul + "): ")
	fmt.Scan(&title)
	if title != "" {
		r.Judul = title
	}

	for j := 0; j < 3; j++ {
		fmt.Printf("Nama Peneliti %d (sekarang: %s): ", j+1, r.Peneliti[j])
		fmt.Scan(&researchers[j])
		if researchers[j] != "" {
			r.Peneliti[j] = researchers[j]
		}
	}

	fmt.Print("Lembaga Penelitian (sekarang: " + r.Institusi + ", berhenti dengan titik): ")
	fmt.Scan(&institution)
	if institution != "" {
		institution = strings.TrimSuffix(institution, ".")
		r.Institusi = institution
	}

	fmt.Print("Waktu Mulai Penelitian (sekarang: " + r.WaktuMulai.Format("02-01-2006") + "): ")
	fmt.Scan(&startDate)
	if startDate != "" {
		start, err := time.Parse("02-01-2006", startDate)
		if err == nil {
			r.WaktuMulai = start
		}
	}

	fmt.Print("Waktu Berakhir Penelitian (sekarang: " + r.WaktuSelesai.Format("02-01-2006") + "): ")
	fmt.Scan(&endDate)
	if endDate != "" {
		end, err := time.Parse("02-01-2006", endDate)
		if err == nil {
			r.WaktuSelesai = end
		}
	}

	fmt.Print("Waktu Publikasi Penelitian (sekarang: " + r.WaktuPublikasi.Format("02-01-2006") + "): ")
	fmt.Scan(&publicationDate)
	if publicationDate != "" {
		pub, err := time.Parse("02-01-2006", publicationDate)
		if err == nil {
			r.WaktuPublikasi = pub
		}
	}

	fmt.Print("Metode Penelitian (sekarang: " + r.Metode + "): ")
	fmt.Scan(&method)
	if method != "" {
		r.Metode = method
	}

	fmt.Print("Summary Penelitian (sekarang: " + r.Summary + ", akhiri dengan bintang): ")
	fmt.Scan(&summary)
	if summary != "" {
		summary = strings.TrimSuffix(summary, "*")
		r.Summary = summary
	}

	fmt.Println("Penelitian berhasil diubah")
	displayResearch(*r)
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
