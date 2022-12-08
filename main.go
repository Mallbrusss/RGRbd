package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "ip"
	port     = 5432
	user     = "user"
	password = "pass"
	dbname   = "db"
)

func checkError(err error) { // прописываем ошибки
	if err != nil {
		panic(err)
	}
}

func connectTo() string { // connect to db

	sqlConn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	return sqlConn
}

func openDb() *sql.DB { // open db

	db, err := sql.Open("postgres", connectTo())
	checkError(err)

	return db
}

func show_table_dom() { // показываем таблицу дом
	rows, err := openDb().Query(`SELECT "этажность", "адрес", "индивидуальное_отопление" FROM "BdRealtor"."дом"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var etazhnost int
		var adres, otop string

		err = rows.Scan(&etazhnost, &adres, &otop)
		checkError(err)

		fmt.Println(etazhnost, adres, otop)
	}

	checkError(err)
}

func show_table_kvartira() { // показываем таблицу квартира
	rows, err := openDb().Query(`SELECT "площадь", "этаж", "номер" FROM "BdRealtor"."квартира"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var ploshad, etaz, number int

		err = rows.Scan(&ploshad, &etaz, &number)
		checkError(err)

		fmt.Println(ploshad, etaz, number)
	}

	checkError(err)
}

func show_table_comp() { // показываем таблицу компания
	rows, err := openDb().Query(`SELECT "название_кп", "юр_адрес", "ИНН" FROM "BdRealtor"."компания_застройщик"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {

		var name_cp, ur_adres, inN string

		err = rows.Scan(&name_cp, &ur_adres, &inN)
		checkError(err)

		fmt.Println(name_cp, ur_adres, inN)
	}

	checkError(err)
}

func show_table_obl() { // показываем таблицу объявление
	rows, err := openDb().Query(`SELECT "с_мебелью", "дата_публикации", "номер_объявления" FROM "BdRealtor"."объявление"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var nomObl int
		var mebel, dateObl string

		err = rows.Scan(&mebel, &dateObl, &nomObl)
		checkError(err)

		fmt.Println(mebel, dateObl, nomObl)
	}

	checkError(err)
}

func show_table_prod() { // показываем таблицу продавец
	rows, err := openDb().Query(`SELECT "с-н_паспорта", "ФИО", "телефон", "e-mail" FROM "BdRealtor"."продавец"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var pasport int
		var fio, telefon, eMail string

		err = rows.Scan(&pasport, &fio, &telefon, &eMail)
		checkError(err)

		fmt.Println(pasport, fio, telefon, eMail)
	}

	checkError(err)
}

func delete_key_dom() { // удаляем из таблицы дом
	var check string
	fmt.Print("какую запись удалить? введите адрес дома:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdRealtor"."дом" where "адрес"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_kvartira() { // удаляем из таблицы квартира
	var check int
	fmt.Print("какую запись удалить? введите номер квартиры:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdRealtor"."квартира" where "номер"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_comp() { // удаляем из таблицы компания
	var check string
	fmt.Print("какую запись удалить? введите ИНН компании:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdRealtor"."компания_застройщик" where "ИНН"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_obl() { // удаляем из таблицы объявление
	var check int
	fmt.Print("какую запись удалить? номер объявления:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdRealtor"."объявление" where "номер_объявления"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_prod() { // удаляем из таблицы продавец
	var check int
	fmt.Print("какую запись удалить? введите номер паспорта:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdRealtor"."продавец" where "с-н_паспорта"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func update_table_kvartira() { // обновляем запись в таблице квартира
	var ploshad, etaz, number int

	fmt.Print("Введите номер квартиры для обновления:\n")
	fmt.Fscan(os.Stdin, &number)

	fmt.Print("Введите новую площадь:\n")
	fmt.Fscan(os.Stdin, &ploshad)

	fmt.Print("введите новый этаж: \n")
	fmt.Fscan(os.Stdin, &etaz)

	updateStm := `update "BdRealtor"."квартира" set "площадь"=$1, "этаж"=$2 where "номер"=$3`
	_, e := openDb().Exec(updateStm, &ploshad, &etaz, &number)
	checkError(e)
}

func update_table_dom() { // обновляем запись в таблице дом
	var adres, otop string
	var etazhnost int

	fmt.Print("Введите адрес для обновления:\n")
	fmt.Fscan(os.Stdin, &adres)

	fmt.Print("Введите новую этажность:\n")
	fmt.Fscan(os.Stdin, &etazhnost)

	fmt.Print("есть ли отопление? да/нет: \n")
	fmt.Fscan(os.Stdin, &otop)

	updateStm := `update "BdRealtor"."дом" set "этажность"=$1, "индивидуальное_отопление"=$2 where "адрес"=$3`
	_, e := openDb().Exec(updateStm, &etazhnost, &otop, &adres)
	checkError(e)
}

func update_table_company() { // обновляем запись в таблице компания
	var nameCp, urAdres, inN string

	fmt.Print("Введите ИНН для обновления:\n")
	fmt.Fscan(os.Stdin, &inN)

	fmt.Print("Введите новое название компании:\n")
	fmt.Fscan(os.Stdin, &nameCp)

	fmt.Print("введите новыый юр-адрес: \n")
	fmt.Fscan(os.Stdin, &urAdres)

	updateStm := `update "BdRealtor"."компания_застройщик" set "название_кп"=$1, "юр_адрес"=$2 where "ИНН"=$3`
	_, e := openDb().Exec(updateStm, &nameCp, &urAdres, &inN)
	checkError(e)
}

func update_table_oble() { // обновляем запись в таблице объявление
	var mebel, date string
	var nomObl int

	fmt.Print("Введите номер объявления:\n")
	fmt.Fscan(os.Stdin, &nomObl)

	fmt.Print("Введите дату публикации:\n")
	fmt.Fscan(os.Stdin, &date)

	fmt.Print("квартира с мебелью?: \n")
	fmt.Fscan(os.Stdin, &mebel)

	updateStm := `update "BdRealtor"."объявление" set "с_мебелью"=$1 where "дата_публикации"=$2, "номер_объявления"=$3`
	_, e := openDb().Exec(updateStm, &mebel, &date, &nomObl)
	checkError(e)
}

func update_table_prodavec() { // обновляем запись в таблице продавец
	var fio, eMail string
	var telefon, pasport int

	fmt.Print("Введите номер паспорта:\n")
	fmt.Fscan(os.Stdin, &pasport)

	fmt.Print("Введите новый номер телефона:\n")
	fmt.Fscan(os.Stdin, &telefon)

	fmt.Print("Введите новую почту: \n")
	fmt.Fscan(os.Stdin, &eMail)

	fmt.Print("Введите новые ФИО: \n")
	fmt.Fscan(os.Stdin, &fio)

	updateStm := `update "BdRealtor"."продавец" set "ФИО"=$1, "телефон"=$2, "e-mail"=$3 where "с-н_паспорта"=$4`
	_, e := openDb().Exec(updateStm, &fio, &telefon, &eMail, &pasport)
	checkError(e)
}

func add_znach_to_dom() { // insert function enter value to dom table
	var adres, otop string
	var etazhnost int
	fmt.Print("Введите адрес:\n")
	fmt.Fscan(os.Stdin, &adres)

	fmt.Print("Введите этажность:\n")
	fmt.Fscan(os.Stdin, &etazhnost)

	fmt.Print("есть ли отопление? да/нет: \n")
	fmt.Fscan(os.Stdin, &otop)

	insertToDyn := `insert into "BdRealtor"."дом"("этажность","адрес","индивидуальное_отопление") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &etazhnost, &adres, &otop)
	checkError(e)

}

func add_znach_to_kvartira() { // insert function enter value to kvartira table
	var ploshad, etaz, nomer int
	fmt.Print("Введите площадь квартиры: \n")
	fmt.Fscan(os.Stdin, &ploshad)

	fmt.Print("Введите этаж квартиры: \n")
	fmt.Fscan(os.Stdin, &etaz)

	fmt.Print("Введите номер квартиры: \n")
	fmt.Fscan(os.Stdin, &nomer)

	insertToDyn := `insert into "BdRealtor"."квартира"("площадь","этаж","номер") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &ploshad, &etaz, &nomer)
	checkError(e)

}

func add_znach_to_company() { // insert function enter value to company table
	var name_cp, ur_adres, inN string
	fmt.Print("Введите название компании: \n")
	fmt.Fscan(os.Stdin, &name_cp)

	fmt.Print("Введите юр.адрес: \n")
	fmt.Fscan(os.Stdin, &ur_adres)

	fmt.Print("Введите Инн: \n")
	fmt.Fscan(os.Stdin, &inN)

	insertToDyn := `insert into "BdRealtor"."компания_застройщик"("название_кп","юр_адрес","ИНН") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &name_cp, &ur_adres, &inN)
	checkError(e)

}

func add_znach_to_oble() { // insert function enter value to oble table
	var mebel, date string
	var number_obl int
	fmt.Print("квартира с мебелью?: \n")
	fmt.Fscan(os.Stdin, &mebel)

	fmt.Print("дата объявления: \n")
	fmt.Fscan(os.Stdin, &date)

	fmt.Print("номер объявления: \n")
	fmt.Fscan(os.Stdin, &number_obl)

	insertToDyn := `insert into "BdRealtor"."объявление"("с_мебелью","дата_публикации","номер_объявления") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &mebel, &date, &number_obl)
	checkError(e)

}

func add_key_to_prodavec() { // insert function enter value to prodavec table
	var e_mail, fio string
	var number_tel, pasport int
	fmt.Print("серия номер паспорта: \n")
	fmt.Fscan(os.Stdin, &pasport)

	fmt.Print("ФИО: \n")
	fmt.Fscan(os.Stdin, &fio)

	fmt.Print("телефон: \n")
	fmt.Fscan(os.Stdin, &number_tel)

	fmt.Print("ваша почта: \n")
	fmt.Fscan(os.Stdin, &e_mail)

	insertToDyn := `insert into "BdRealtor"."продавец"("с-н_паспорта","ФИО","телефон","e-mail") values($1, $2, $3, $4)`
	_, e := openDb().Exec(insertToDyn, &pasport, &fio, &number_tel, &e_mail)
	checkError(e)

}

func add_switch_case() { // функция выбора таблицы для добавления записи
	var vibor string
	fmt.Print("выберите дейсвтие: add_home -  чтобы добавить данные в таблицу дом\n add_kvartira - добавить данные в таблицу квартира\n add_company добавить данные в таблицу квартира\n add_oble - добавить данные в таблицу объявление\n add_prod - добавить данные в таблицу продавец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "add_home":
		add_znach_to_dom()
	case "add_kvartira":
		add_znach_to_kvartira()
	case "add_company":
		add_znach_to_company()
	case "add_oble":
		add_znach_to_oble()
	case "add_prod":
		add_key_to_prodavec()
	}
}

func update_switch_case() { // функция выбора таблицы для обеовления
	var vibor string
	fmt.Print("выберите дейсвтие: updateHome -  чтобы обновить данные в таблице дом\n updateKvartira - обновить данные в таблице квартира\n updateCompany обновить данные в таблице компания\n updateOble - обновить данные в таблице объявление\n updateProd - обновить данные в таблице продавец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "updateHome":
		update_table_dom()
	case "updateKvartira":
		update_table_kvartira()
	case "updateCompany":
		update_table_company()
	case "updateOble":
		update_table_oble()
	case "updateProd":
		update_table_prodavec()
	}
}

func delete_switch_case() { // функция выбора таблицы для удаления записи
	var vibor string
	fmt.Print("выберите дейсвтие: delDom -  чтобы удалить данные из таблицы дом\n delKvartira - удалить данные из таблицы квартира\n delCompany удалить данные из таблицы компания\n delOble - удалить данные из таблицы объявление\n delProd - удалить данные из таблицы продавец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "delDom":
		delete_key_dom()
	case "delKvartira":
		delete_key_kvartira()
	case "delCompany":
		delete_key_comp()
	case "delOble":
		delete_key_obl()
	case "delProd":
		delete_key_prod()
	}
}

func show_switch_case() { // функция выбора таблицы для выводы данных
	var vibor string
	fmt.Print("выберите дейсвтие: showDom -  чтобы показать данные из таблицы дом\n showKvartira - показать данные из таблицы квартира\n showCompany показать данные из таблицы компания\n showOble - показать данные из таблицы объявление\n showProd - показать данные из таблицы продавец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "showDom":
		show_table_dom()
	case "showKvartira":
		show_table_kvartira()
	case "showCompany":
		show_table_comp()
	case "showOble":
		show_table_obl()
	case "showProd":
		show_table_prod()
	}
}

func main() {
	var v1 string
	//close db
	defer openDb().Close()
	//check db
	err := openDb().Ping()
	checkError(err)

	fmt.Print("Что вы хотите сделать?\n Чтобы добавить значения в таблицу введите addTable\n Чтобы обновить запись введите updateTable\n Чтобы удалить запись из таблицы введите deleteFromTable\n Чтобы показать данные в таблице введите showTable\n ")
	fmt.Scanf("%s\n", &v1)

	switch v1 {
	case "addTable":
		add_switch_case()
	case "updateTable":
		update_switch_case()
	case "deleteFromTable":
		delete_switch_case()
	case "showTable":
		show_switch_case()
	}
}
