package main

import "fmt"

//import "os"
type pasien struct {
	nama         string
	umur         string
	id           int
	jkelamin     string
	goldarah     string
	nomorlayanan string
	hasilmcu     [100]hasilrekap
	time         [3]waktu
	harga        int
}
type waktu struct {
	tahun int
	bulan int
	hari  int
}
type hasilrekap struct {
	guladarah    string //paket1
	hdl          string //paket1
	ldl          string //paket1
	trigliserida string //paket1
	asamurat     string //paket2
	hemoglobin   string //paket2
	leukosit     string //paket2
	kreatin      string //paket2
	hepatitis    string //paket3
	sgot         string //paket3
	sgpt         string //paket3
}

type jpas [100]pasien

var jml_pas int

func insertpasien(data *jpas, jml_pas *int) {
	var banyakpasien int
	//*jml_pas=0
	fmt.Print("Berapa Data Pasien Yang Ingin Anda Masukan? : ")
	fmt.Scan(&banyakpasien)
	i := 1
	for i <= banyakpasien {
		fmt.Print("Masukan Nama Pasien", " ", "ke-", i, " : ")
		fmt.Scan(&data[*jml_pas].nama)
		fmt.Print("Masukan Umur : ")
		fmt.Scan(&data[*jml_pas].umur)
		fmt.Print("Masukan ID : ")
		fmt.Scan(&data[*jml_pas].id)
		fmt.Print("Masukan Jenis Kelamin (L/P) : ")
		fmt.Scan(&data[*jml_pas].jkelamin)
		fmt.Print("Masukan Golongan Darah (a/b/o) : ")
		fmt.Scan(&data[*jml_pas].goldarah)
		fmt.Print("Masukan Nomor Data Layanan (1/2/3) : ")
		fmt.Scan(&data[*jml_pas].nomorlayanan)
		fmt.Print("Masukan Waktu Data Di Input (DD MM YYYY) : ")
		fmt.Scan(&data[*jml_pas].time[1].hari, &data[*jml_pas].time[1].bulan, &data[*jml_pas].time[1].tahun)
		if data[*jml_pas].nomorlayanan == "1" {
			data[*jml_pas].harga = 180000
			fmt.Print("Masukan Kadar Gula Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].guladarah)
			fmt.Print("Masukan Kadar HDL dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].hdl)
			fmt.Print("Masukan Kadar LDL dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].ldl)
			fmt.Print("Masukan kadar Trigliserida dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].trigliserida)
		} else if data[*jml_pas].nomorlayanan == "2" {
			data[*jml_pas].harga = 200000
			fmt.Print("Masukan Kadar Asam Urat dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].asamurat)
			fmt.Print("Masukan Kadar Hemoglobin dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].hemoglobin)
			fmt.Print("Masukan Kadar Leukosit dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].leukosit)
			fmt.Print("Masukan Kadar Kreatin dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].kreatin)
		} else if data[*jml_pas].nomorlayanan == "3" {
			data[*jml_pas].harga = 250000
			fmt.Print("Apakah Positif Hepatitis (positif/negatif) : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].hepatitis)
			fmt.Print("Masukan Kadar SGOT dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].sgot)
			fmt.Print("Masukan Kadar SGPT dalam Darah : ")
			fmt.Scan(&data[*jml_pas].hasilmcu[1].sgpt)
		}
		i++
		*jml_pas++

	}
}

func editdatapasien(data *jpas, jml_pas *int) {
	var id int
	var i int
	var sukses bool
	var pilihan int
	sukses = false
	fmt.Println("Masukan ID Pasien yang Mau di Edit")
	fmt.Scan(&id)
	for i < *jml_pas {
		if data[i].id == id {
			fmt.Println("ID Ditemukan")
			fmt.Println("Apa Yang Ingin Anda Edit? : ")
			fmt.Println("1. Nama Pasien")
			fmt.Println("2. Umur Pasien")
			fmt.Println("3. Jenis Kelamin Pasien")
			fmt.Println("4. Golongan Darah Pasien")
			fmt.Print("Masukan Pilihan Anda : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				fmt.Print("Masukan Nama Baru Pasien  : ")
				fmt.Scan(&data[i].nama)
				sukses = true
			} else if pilihan == 2 {
				fmt.Print("Masukan Umur Baru Pasien : ")
				fmt.Scan(&data[i].umur)
				sukses = true
			} else if pilihan == 3 {
				fmt.Print("Masukan Jenis Kelamin Baru Pasien : ")
				fmt.Scan(&data[i].jkelamin)
				sukses = true
			} else if pilihan==4{
				fmt.Print("Masukan Golongan Darah Baru Pasien : ")
				fmt.Scan(&data[i].goldarah)
			}
		}
		i = i + 1
	}
	if !sukses {
		fmt.Println("ID Tidak Ada")
	} else {
		fmt.Println("Berhasil Diedit")
	}

}
func hapuspasien(data *jpas, jml_pas *int) {
	var id int
	var i int
	var deleted string
	var idxterhapus int
	var pasiendelete pasien
	fmt.Print("Masukan ID Pasien Yang ingin Dihapus")
	fmt.Scan(&id)
	for i < *jml_pas {
		if data[i].id == id {
			data[i] = pasiendelete
			deleted = "Berhasil"
			idxterhapus = i
			*jml_pas = *jml_pas - 1

		}
		i = i + 1
	}
	if deleted != "Berhasil" {
		fmt.Print("ID Pasien Tidak Ada")
	} else {
		counter := 1
		fmt.Println("ID Pasien", id, "Berhasil Dihapus")
		for (i + counter) <= *jml_pas {
			data[idxterhapus-1] = data[idxterhapus+counter]
			counter = counter + 1
		}
	}

}

func laporanpemasukan(data *jpas, jml_pas *int) {
	var bulan int
	var i int
	var total int
	//i=0
	//total = 0
	fmt.Print("Masukan Bulan Laporan Yang ingin Dilihat : ")
	fmt.Scan(&bulan)
	for i <= *jml_pas {
		//total=0
		if data[i].time[1].bulan == bulan {
			total = data[i].harga+total
			//i = i + 1
		}
		i = i + 1
		//total=0
	}
	fmt.Println(total)

}

func carinama(data *jpas, jml_pas *int) {
	var nama string
	var i int
	i=0
	fmt.Println("Apabila Pasien Tidak Ditemukan Otomatis Akan Kembali Ke Menu Utama")
	fmt.Print("Masukan Nama Pasien Yang ingin Dilihat : ")
	fmt.Scan(&nama)
	for i < *jml_pas {
		if data[i].nama == nama {
			fmt.Print("Nomor ID Pasien Adalah                                : ")
			fmt.Println(data[i].id)
			fmt.Print("Umur Pasien Adalah                                    : ")
			fmt.Println(data[i].umur)
			fmt.Print("Jenis Kelamin Pasien Adalah (L=Laki Laki/P=Perempuan) : ")
			fmt.Println(data[i].jkelamin)
			fmt.Print("Golongan Darah Pasien Adalah                          : ")
			fmt.Println(data[i].goldarah)
			if data[i].nomorlayanan=="1"{
				fmt.Print("Kadar Gula Darah Pasien                               : ")
				fmt.Println(data[i].hasilmcu[1].guladarah)
				fmt.Print("Kadar HDL dalam Darah Pasien                          : ")
				fmt.Println(data[i].hasilmcu[1].hdl)
				fmt.Print("Kadar LDL dalam Darah Pasien                          : ")
				fmt.Println(data[i].hasilmcu[1].ldl)
				fmt.Print("Kadar Trigliserida dalam Darah Pasien                 : ")
				fmt.Println(data[i].hasilmcu[1].trigliserida)
			} else if data[i].nomorlayanan == "2" {
				fmt.Print("Kadar Asam Urat dalam Darah Pasien                    : ")
				fmt.Println(data[i].hasilmcu[1].asamurat)
				fmt.Print("Kadar Hemoglobin dalam Darah Pasien                   : ")
				fmt.Println(data[i].hasilmcu[1].hemoglobin)
				fmt.Print("Kadar Leukosit dalam Darah Pasien                     : ")
				fmt.Println(data[i].hasilmcu[1].leukosit)
				fmt.Print("Kadar Kreatin dalam Darah Pasien                      : ")
				fmt.Println(data[i].hasilmcu[1].kreatin)
			} else if data[i].nomorlayanan == "3" {
				fmt.Print("Apakah Pasien Positif Hepatitis (positif/negatif)     : ")
				fmt.Println(data[i].hasilmcu[1].hepatitis)
				fmt.Print("Kadar SGOT dalam Darah Pasien                         : ")
				fmt.Println(data[i].hasilmcu[1].sgot)
				fmt.Print("Masukan Kadar SGPT dalam Darah                        : ")
				fmt.Println(data[i].hasilmcu[1].sgpt)
		}
				

			//i = i + 1
		}
		i = i + 1
	}

	

}



func cariperiode(data *jpas, jml_pas *int) {
	var periode int
	var i int
	i=0
	fmt.Println("Apabila Pasien Tidak Ditemukan Otomatis Akan Kembali Ke Menu Utama")
	fmt.Print("Masukan Bulan Periode Daftar Pasien Yang ingin Dilihat : ")
	fmt.Scan(&periode)
	for i < *jml_pas {
		if data[i].time[1].bulan == periode {
			fmt.Print("Nama Pasien Adalah                                    : ")
			fmt.Println(data[i].nama)
			fmt.Print("Nomor ID Pasien Adalah                                : ")
			fmt.Println(data[i].id)
			fmt.Print("Umur Pasien Adalah                                    : ")
			fmt.Println(data[i].umur)
			fmt.Print("Jenis Kelamin Pasien Adalah (L=Laki Laki/P=Perempuan) : ")
			fmt.Println(data[i].jkelamin)
			fmt.Print("Golongan Darah Pasien Adalah                          : ")
			fmt.Println(data[i].goldarah)
			if data[i].nomorlayanan=="1"{
				fmt.Print("Kadar Gula Darah Pasien                               : ")
				fmt.Println(data[i].hasilmcu[1].guladarah)
				fmt.Print("Kadar HDL dalam Darah Pasien                          : ")
				fmt.Println(data[i].hasilmcu[1].hdl)
				fmt.Print("Kadar LDL dalam Darah Pasien                          : ")
				fmt.Println(data[i].hasilmcu[1].ldl)
				fmt.Print("Kadar Trigliserida dalam Darah Pasien                 : ")
				fmt.Println(data[i].hasilmcu[1].trigliserida)
			} else if data[i].nomorlayanan == "2" {
				fmt.Print("Kadar Asam Urat dalam Darah Pasien                    : ")
				fmt.Println(data[i].hasilmcu[1].asamurat)
				fmt.Print("Kadar Hemoglobin dalam Darah Pasien                   : ")
				fmt.Println(data[i].hasilmcu[1].hemoglobin)
				fmt.Print("Kadar Leukosit dalam Darah Pasien                     : ")
				fmt.Println(data[i].hasilmcu[1].leukosit)
				fmt.Print("Kadar Kreatin dalam Darah Pasien                      : ")
				fmt.Println(data[i].hasilmcu[1].kreatin)
			} else if data[i].nomorlayanan == "3" {
				fmt.Print("Apakah Pasien Positif Hepatitis (positif/negatif)     : ")
				fmt.Println(data[i].hasilmcu[1].hepatitis)
				fmt.Print("Kadar SGOT dalam Darah Pasien                         : ")
				fmt.Println(data[i].hasilmcu[1].sgot)
				fmt.Print("Masukan Kadar SGPT dalam Darah                        : ")
				fmt.Println(data[i].hasilmcu[1].sgpt)
		}
				

			//i = i + 1
		}
		i = i + 1
	}
}
//var data jpas
func sorting(data jpas,jml_pas *int){
	//var temp int
	var i int
	var j int
	var k int

	i=0
	j=0
	for  i < *jml_pas{
		min := i
		j = i
		for j < *jml_pas {
			if data[j].time[1].bulan < data[min].time[1].bulan //&& data[j].time[1].hari < data[min].time[1].hari{
				min = j

			}
			j=j+1
		}
		//temp = data[*jml_pas].time[*jml_pas].bulan
		data[i].time[1].bulan, data[min].time[1].bulan=data[min].time[1].bulan, data[i].time[1].bulan
		data[i].nama, data[min].nama=data[min].nama, data[i].nama
		data[i].time[1].hari, data[min].time[1].hari=data[min].time[1].hari, data[i].time[1].hari



		i=i+1
	}
	k=0 
	fmt.Println("Nama Pasien Terurut Berdasarkan Bulan Kedatangan")
	for k < *jml_pas {
		fmt.Println("Nama Pasien : ",(data)[k].nama)
		fmt.Println(" ","-", "Tanggal Masuk Pasien : ",(data)[k].time[1].hari)
		fmt.Println(" ","-", "Bulan Masuk Pasien :   ",(data)[k].time[1].bulan)

		k=k+1
	}
}








func main() {
	var temp string
	var data jpas
	for temp != "12" {
		fmt.Println("_____________________________________________________________")
		fmt.Println("|                                                           |")
		fmt.Println("| Menu :                                                    |")
		fmt.Println("| 1. Penambahan Data Pasien                                 |")
		fmt.Println("| 2. Edit Data Pasien                                       |")
		fmt.Println("| 3. Hapus Pasien                                           |")
		fmt.Println("| 4. Laporan Pemasukan                                      |")
		fmt.Println("| 5. Pencarian Berdasarkan Nama                             |")
		fmt.Println("| 6. Pencarian Berdasarkan Periode                          |")
		fmt.Println("| 12. Exit Program                                          |")
		fmt.Println("|___________________________________________________________|")
		fmt.Print("Tulislah angka yang akan anda pilih : ")
		fmt.Scan(&temp)
		switch temp {
		case "1":
			insertpasien(&data, &jml_pas)
		case "2":
			editdatapasien(&data, &jml_pas)
		case "3":
			hapuspasien(&data, &jml_pas)
		case "4":
			laporanpemasukan(&data, &jml_pas)
		case "5" :
			carinama(&data,&jml_pas)
		case "6" :
			cariperiode(&data, &jml_pas)
		case "7" : 
			sorting(data, &jml_pas)

		case "11":
			fmt.Println(data[1].nama)
			fmt.Println(data[2].nama)
			fmt.Println(data[3].nama)
			fmt.Println(jml_pas)
			fmt.Println(data[1].harga)
			fmt.Println(data[0].time[1].bulan)
			fmt.Println(data[1].time[1].bulan)
			fmt.Println(data[2].time[1].bulan)

		}
	}
}
