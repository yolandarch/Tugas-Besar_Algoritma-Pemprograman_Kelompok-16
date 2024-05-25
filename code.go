package main

import (
	"fmt"
)

type Akun struct {
	username, password string
}

type Time struct {
	tanggal, bulan, tahun int
}

type Penelitian struct {
	ID              string
	Model           string
	Judul           string
	Peneliti     	[3]string
	Institusi     	string
	WaktuMulai		Time
	WaktuSelesai	Time
	WaktuPublikasi 	Time
	Metode         	string
	Summary         string
}

var akun Account
var Research []Penelitian

func main() {
	buatAkun(&akun)
	login(akun)
	menuUtama()
}

func buatAkun(A *Account) {
	fmt.Println("Create Account")
	fmt.Print("Username: ")
	fmt.Scan(&A.username)
	fmt.Print("Password: ")
	fmt.Scan(&A.password)
	fmt.Println("Akun berhasil dibuat")
}

func login(A Account) {
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

	start, _ := time.Parse("02-01-2006", startDate)
	end, _ := time.Parse("02-01-2006", endDate)
	pub, _ := time.Parse("02-01-2006", publicationDate)

	fmt.Print("Metode Penelitian: ")
	fmt.Scan(&method)

	fmt.Print("Summary Penelitian (akhiri dengan bintang): ")
	fmt.Scan(&summary)

	id := fmt.Sprintf("%dP", len(researches)+1)
	newResearch := Research{
		ID:              	id,
		Model:           	model,
		Judul:           	title,
		Peneliti:     		researchers,
		Institusi:     		institution,
		WaktuMulai:       	start,
		WaktuSelesai:       end,
		WaktuPublikasi: 	pub,
		Metode:          	method,
		Summary:         	summary,
	}

	fmt.Printf("Penelitian %s berhasil ditambahkan\n", title)
	displayResearch(newResearch)
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

func displayResearch(r Research) {
	fmt.Println("==============================================================================")
	fmt.Printf("| IdPenelitian       | %s\n", r.ID)
	fmt.Printf("| Modul              | %s\n", r.Model)
	fmt.Println("|====================|========================================================")
	fmt.Printf("| Judul Penelitian   | %s\n", r.Title)
	for i, researcher := range r.Researchers {
		fmt.Printf("| Nama Peneliti %d    | %s\n", i+1, researcher)
	}
	fmt.Printf("| Lembaga Penelitian | %s\n", r.Institution)
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
		filterAndDisplayResearch(func(r Research) bool {
			return strings.EqualFold(r.Model, model)
	})
	case 2:
		var date string
		fmt.Print("Masukkan tanggal publikasi (dd-mm-yyyy): ")
		fmt.Scan(&date)
		pubDate, _ := time.Parse("02-01-2006", date)
		filterAndDisplayResearch(func(r Research) bool {
			return r.PublicationDate.Equal(pubDate)
		})
	case 3:
		filterAndDisplayResearch(func(r Research) bool {
			return true
		})
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func filterAndDisplayResearch(filterFunc func(Research) bool) {
	filtered := []Research{}
	for _, r := range researches {
		if filterFunc(r) {
			filtered = append(filtered, r)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Tidak ada penelitian yang ditemukan")
		return
	}

	fmt.Println("Pilih Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	var pengurutanPilihan int
	fmt.Print("Pilih urutan: ")
	fmt.Scan(&pengurutanPilihan)

	sort.Slice(filtered, func(i, j int) bool {
		if pengurutanPilihan == 1 {
			return filtered[i].Title < filtered[j].Title
		}
		return filtered[i].Title > filtered[j].Title
	})

	for _, r := range filtered {
		displayResearch(r)
	}
}

func viewResearchSummary() {
	var idOrTitle string
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian: ")
	fmt.Scan(&idOrTitle)

	for _, r := range researches {
		if strings.EqualFold(r.ID, idOrTitle) || strings.EqualFold(r.Title, idOrTitle) {
			fmt.Println("==============================================================================")
			fmt.Printf("| IdPenelitian       | %s\n", r.ID)
			fmt.Printf("| Modul              | %s\n", r.Model)
			fmt.Println("|====================|========================================================")
			fmt.Printf("| Judul Penelitian   | %s\n", r.Title)
			for i, researcher := range r.Researchers {
				fmt.Printf("| Nama Peneliti %d    | %s\n", i+1, researcher)
			}
			fmt.Printf("| Summary Penelitian | %s\n", r.Summary)
			fmt.Println("==============================================================================")
			return
		}
	}
	fmt.Println("Penelitian tidak ditemukan")
}

func compareResearch() {
	var idOrTitle1, idOrTitle2 string
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian pertama: ")
	fmt.Scan(&idOrTitle1)
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian kedua: ")
	fmt.Scan(&idOrTitle2)

	var r1, r2 *Research
	for i := range researches {
		if strings.EqualFold(researches[i].ID, idOrTitle1) || strings.EqualFold(researches[i].Title, idOrTitle1) {
			r1 = &researches[i]
		}
		if strings.EqualFold(researches[i].ID, idOrTitle2) || strings.EqualFold(researches[i].Title, idOrTitle2) {
			r2 = &researches[i]
		}
	}

	if r1 == nil || r2 == nil {
		fmt.Println("Salah satu atau kedua penelitian tidak ditemukan")
		return
	}

	fmt.Println("==============================================================================")
	fmt.Printf("| Perbandingan antara %s dan %s\n", r1.Title, r2.Title)
	fmt.Println("==============================================================================")
	duration1 := r1.PublicationDate.Sub(r1.StartDate)
	duration2 := r2.PublicationDate.Sub(r2.StartDate)

	fmt.Printf("| %s: %v hari\n", r1.Title, duration1.Hours()/24)
	fmt.Printf("| %s: %v hari\n", r2.Title, duration2.Hours()/24)
	if duration1 < duration2 {
		fmt.Printf("| %s lebih cepat dipublikasikan\n", r1.Title)
	} else if duration1 > duration2 {
		fmt.Printf("| %s lebih cepat dipublikasikan\n", r2.Title)
	} else {
		fmt.Println("| Kedua penelitian memiliki durasi yang sama hingga publikasi")
	}
	fmt.Println("==============================================================================")
}

func editOrDeleteResearch() {
	var idOrTitle string
	fmt.Print("Masukkan IdPenelitian atau Judul Penelitian yang akan diubah/hapus: ")
	fmt.Scan(&idOrTitle)

	for i := range researches {
		if strings.EqualFold(researches[i].ID, idOrTitle) || strings.EqualFold(researches[i].Title, idOrTitle) {
			fmt.Println("Penelitian ditemukan. Pilih opsi:")
			fmt.Println("1. Edit Penelitian")
			fmt.Println("2. Hapus Penelitian")
			var pilihan int
			fmt.Print("Pilih opsi: ")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				editResearch(&researches[i])
				return
			} else if pilihan == 2 {
				deleteResearch(i)
				return
			} else {
				fmt.Println("Pilihan tidak valid")
				return
			}
		}
	}
	fmt.Println("Penelitian tidak ditemukan")
}

func editResearch(r *Research) {
	fmt.Println("Masukkan data baru (kosongkan untuk tidak mengubah):")
	var model, title, institution, method, summary string
	var researchers [3]string
	var startDate, endDate, publicationDate string

	fmt.Print("Model Penelitian (sekarang: " + r.Model + "): ")
	fmt.Scan(&model)
	if model != "" {
		r.Model = model
	}

	fmt.Print("Judul Penelitian (sekarang: " + r.Title + "): ")
	fmt.Scan(&title)
	if title != "" {
		r.Title = title
	}

	for j := 0; j < 3; j++ {
		fmt.Printf("Nama Peneliti %d (sekarang: %s): ", j+1, r.Researchers[j])
		fmt.Scan(&researchers[j])
		if researchers[j] != "" {
			r.Researchers[j] = researchers[j]
		}
	}

	fmt.Print("Lembaga Penelitian (sekarang: " + r.Institution + ", berhenti dengan titik): ")
	fmt.Scan(&institution)
	if institution != "" {
		institution = strings.TrimSuffix(institution, ".")
		r.Institution = institution
	}

	fmt.Print("Waktu Mulai Penelitian (sekarang: " + r.StartDate.Format("02-01-2006") + "): ")
	fmt.Scan(&startDate)
	if startDate != "" {
		start, _ := time.Parse("02-01-2006", startDate)
		r.StartDate = start
	}

	fmt.Print("Waktu Berakhir Penelitian (sekarang: " + r.EndDate.Format("02-01-2006") + "): ")
	fmt.Scan(&endDate)
	if endDate != "" {
		end, _ := time.Parse("02-01-2006", endDate)
		r.EndDate = end
	}

	fmt.Print("Waktu Publikasi Penelitian (sekarang: " + r.PublicationDate.Format("02-01-2006") + "): ")
	fmt.Scan(&publicationDate)
	if publicationDate != "" {
		pub, _ := time.Parse("02-01-2006", publicationDate)
		r.PublicationDate = pub
	}

	fmt.Print("Metode Penelitian (sekarang: " + r.Method + "): ")
	fmt.Scan(&method)
	if method != "" {
		r.Method = method
	}

	fmt.Print("Summary Penelitian (sekarang: " + r.Summary + ", akhiri dengan bintang): ")
	fmt.Scan(&summary)
	if summary != "" {
		summary = strings.TrimSuffix(summary, "*")
		r.Summary = summary
	}

	fmt.Println("Penelitian berhasil diubah")
	displayResearch(*r)
}

func deleteResearch(index int) {
	researches = append(researches[:index], researches[index+1:]...)
	fmt.Println("Penelitian berhasil dihapus")
}

