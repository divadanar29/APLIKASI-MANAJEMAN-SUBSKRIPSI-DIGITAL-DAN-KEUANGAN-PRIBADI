package main
import "fmt"
// deklarasi tipe array
const NMAX int = 1000
type tabPelanggan [NMAX]langganan

// tipe bentukan data pelanggan
type langganan struct{
	nama string
	kategori string 
	biaya_langganan float64 
	tanggal_bayar string
	metode string
	status string
}

func tambahLangganan( A *tabPelanggan, n *int){ // masukkan data langganan sesuai yang diinginkan
	fmt.Print("Nama Langganan: ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Kategori: ")
	fmt.Scan(&A[*n].kategori)
	fmt.Print("Biaya langganan: ")
	fmt.Scan(&A[*n].biaya_langganan)
	fmt.Print("Tanggal bayar: ")
	fmt.Scan(&A[*n].tanggal_bayar)
	fmt.Print("Metode: ")
	fmt.Scan(&A[*n].metode)
	fmt.Print("Status: ")
	fmt.Scan(&A[*n].status)
	*n = *n + 1
}
func tampilkanLangganan(A *tabPelanggan, n *int){ // akan ditampilkan semua data pelanggan dari tambah langganan sebelumnya
	for i := 1; i < *n; i++{ // mulai indeks 1 hingga n (banyak data) 
		fmt.Println("Langganan",i)
		fmt.Println("1. Nama: ",A[i].nama)
		fmt.Println("2. Kategori: ",A[i].kategori)
		fmt.Println("3. Biaya Langganan: ",A[i].biaya_langganan)
		fmt.Println("4. Tanggal Bayar: ",A[i].tanggal_bayar)
		fmt.Println("5. Metode: ",A[i].metode)
		fmt.Println("6. Status: ",A[i].status)
		fmt.Println("---")
	}
}
func ubahLangganan(A *tabPelanggan, n *int){ // data pelanggan yang ingin diubah 
	var i int 
	fmt.Print("Langganan ke:")
	fmt.Scan(&i) // data ke berapa
	if i < *n {
		fmt.Print("Nama Langganan baru: ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Kategori baru: ")
		fmt.Scan(&A[i].kategori)
		fmt.Print("Biaya Langganan baru: ")
		fmt.Scan(&A[i].biaya_langganan)
		fmt.Print("Tanggal bayar baru: ")
		fmt.Scan(&A[i].tanggal_bayar)
		fmt.Print("Metode baru: ")
		fmt.Scan(&A[i].metode)
		fmt.Print("Status baru: ")
		fmt.Scan(&A[i].status)
	}
}
func hapusLangganan( A *tabPelanggan, n *int){ // data pelanggan yang ingin di hapus 
	var i, hapus int
	fmt.Print("Langganan yang akan dihapus: ")
	fmt.Scan(&hapus) // masukkan data ke berapa yang ingin di hapus
	*n = *n - 1 // data akan berkurang karena banyak data akan mundur 
	for i = hapus; i < *n; i++{ // mulai dari data ke "hapus" hingga n 
		A[i] = A[i+1]
	}
}
func carilayanan(A *tabPelanggan, n *int, x string ) { //sequential search
	var i int
	fmt.Print("kategori:")
	fmt.Scan(&x)
	for i = 1; i < *n; i++{
		if A[i].kategori == x {
			fmt.Println("Nama Langganan",i)
			fmt.Println("1. Nama:",A[i].nama)
			fmt.Println("2. Kategori:",A[i].kategori)
			fmt.Println("3. Biaya Langganan:",A[i].biaya_langganan)
			fmt.Println("4. Tanggal Bayar:",A[i].tanggal_bayar)
			fmt.Println("5. Metode:",A[i].metode)
			fmt.Println("6. Status:",A[i].status)
			fmt.Println("---")
		} else {
			fmt.Println("Langganan tidak ditemukan")
		}
	}
}
func totalBiaya(A *tabPelanggan, n *int) {
	var i int
	var jumlah float64 = 0
	for i = 1; i < *n; i++{
		jumlah = jumlah + A[i].biaya_langganan
	}
	fmt.Println("Total biaya bulanan: ",jumlah)
}
func biayaTermahal(A tabPelanggan, n int) int{
	var i, max int
	if n <= 1 {
		return 1
	}
	max = 1
	for i = 1; i < n; i++ {
		if A[i].biaya_langganan > A[max].biaya_langganan {
			max = i
		}
	}
	return max
}
func main(){
	var n int = 1
	var p tabPelanggan
	var x string
	var pilihan int = 0
	for pilihan != 11{
		fmt.Println("\n=== Manajemen Langganan Digital ===")
		fmt.Println("1. Tambah Langganan")
		fmt.Println("2. Tampilkan Langganan")
		fmt.Println("3. Ubah Langganan")
		fmt.Println("4. Hapus Langganan")
		fmt.Println("5. Cari layanan(Sequential)") // 
		fmt.Println("6. Urutkan Layanan (Binary Search)")// terurut 
		fmt.Println("7. Urutkan Biaya (Selection Sort)")
		fmt.Println("8. Urutkan Nama (Insertion Sort)")
		fmt.Println("9. Cek Jatuh Tempo")
		fmt.Println("10. Total Biaya Bulanan") //bisa fungsi
		fmt.Println("11. biaya termahal ")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahLangganan(&p,&n)
		} else if pilihan == 2 {
			tampilkanLangganan(&p,&n)
		} else if pilihan == 3 {
			ubahLangganan(&p,&n)
		} else if pilihan == 4{
			hapusLangganan(&p,&n)
		} else if pilihan == 5 {
			carilayanan(&p,&n,x)
		} else if pilihan == 6 {
			biayaBulanan(&p,&n)
		} else {
			max := biayaTermahal(p, n)
			fmt.Println("Langganan", max, "termahal:", p[max].biaya_langganan)
		}
	}
}