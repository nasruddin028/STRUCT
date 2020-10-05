package main

import "fmt"

type Data struct 
    {
    Nama, Nilai string
    Nim      int
    }
func main() {

	fmt.Println("============================")
	fmt.Println("    FORM NILAI SEMESTER 3   ")
	fmt.Println("============================")


	multiA := [10]Data {
		for jumlah := 0; jumlah < 10; jumlah++{
			fmt.Println("Nama :")
			fmt.Scan(&Data.Nama)
			fmt.Println("Masukan Nim :")
			fmt.Scan(&Data.Nilai)
			fmt.Println("Masukan Nim :")
			fmt.Scan(&Data.Nim)
			}
	}			
	
	fmt.Println("============================")
	fmt.Println("BIODATA MAHASISWA PANCASAKTI")
	
	for jumlah := 0; jumlah < 10; jumlah++ {
	 fmt.Println("===========================")
	 fmt.Printf("Nama \t: %v\n", multiA[jumlah].Nama)
	 fmt.Printf("Alamat \t: %v\n", multiA[jumlah].Nilai)
	 fmt.Printf("NIM \t: %v\n", multiA[jumlah].Nim)
	}
    
	fmt.Println("============================")
	fmt.Println(" GUNAKAN DATA DENGAN BIJAK  ")
}