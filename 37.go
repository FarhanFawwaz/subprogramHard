func printPolaBilangan(baris, kolom, n int) {
	/* I.S bilangan bulat baris, kolom, dan n
	F.S menampilkan pola string "*" yang berbentuk keliling persegi yang berukuran N × N */

	if n >= baris {
		if n >= kolom {
			if baris == 1 || baris == n || kolom == 1 || kolom == n {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
			printPolaBilangan(baris, kolom+1, n)
		} else {
			fmt.Println()
			printPolaBilangan(baris+1, 1, n)
		}
	}
}