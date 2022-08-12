package main

import "fmt"

func Secim() {
	secim := [4]string{"Para Çekme", "Para Yatırma", "Bakiye Sorgulama", "Çıkış"}
	for idx, val := range secim {
		idx += 1
		fmt.Printf("[%v] %v\n", idx, val)
	}
}

func main() {
	fmt.Println("Bankamıza Hoş Geldiniz!")

	_sifre := int(6565)
	sifre_sayac := 3
	_bakiye := 5000.0

	fmt.Println("Lütfen şifrenizi giriniz!")

	for {

		var kullanici_sifre int
		fmt.Scanln(&kullanici_sifre)

		if kullanici_sifre != _sifre {
			sifre_sayac -= 1
			fmt.Printf("Şifrenizi yanlış girdiniz! Kalan hak: %v\n", sifre_sayac)
			if sifre_sayac <= 0 {
				fmt.Println("Çok sayıda yanlış deneme yaptınız!")
				break
			}
		} else {
			fmt.Println("Hesabınıza hoş geldiniz!\nLütfen bir seçim yapınız.")
			Secim()
			var _secim int
			fmt.Scanln(&_secim)
			// SWITCH

			switch _secim {
			case 1:
				_cekmeSayac := 3
				for {
					fmt.Printf("Bakiyeniz: %v$\n", _bakiye)
					fmt.Println("Çekmek istediğiniz tutarı giriniz.")
					var _cekilenTutar float64
					fmt.Scanln(&_cekilenTutar)

					if _cekilenTutar > _bakiye {
						_cekmeSayac -= 1
						fmt.Println("Yeterli bakiyeniz bulunmamaktadır.")
						if _cekmeSayac <= 0 {
							fmt.Println("Çok sayıda yanlış çekme işlemi yaptınız!\nÇıkışı yapıldı!")
							break
						}
					} else {
						_bakiye -= _cekilenTutar
						fmt.Printf("%v$ para çekme işleminiz onaylandı.\nGüncel bakiyeniz: %v$\n", _cekilenTutar, _bakiye)
						fmt.Println("Çıkışınız yapıldı.")
						break
					}
				}

			case 2:

				fmt.Printf("Bakiyeniz: %v$\n", _bakiye)
				fmt.Println("Yatırmak istediğiniz tutarı giriniz.")
				var _yatirilanTutar float64
				fmt.Scanln(&_yatirilanTutar)

				_bakiye += _yatirilanTutar
				fmt.Printf("%v$ para yatırma işleminiz onaylandı.\nGüncel bakiyeniz: %v$\n", _yatirilanTutar, _bakiye)
				fmt.Println("Çıkışınız yapıldı.")

			case 3:
				fmt.Printf("Güncel bakiyeniz: %v$\n", _bakiye)
				fmt.Println("Çıkışınız yapıldı.")
			case 4:
				fmt.Println("Çıkışınız yapıldı.")
			default:
				fmt.Println("Yanlış seçim yaptınız. Çıkışınız yapıldı!")
				return
			}
			break
		}
	}
}
