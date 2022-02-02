package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	delTrans, err := template.ParseFiles("main.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	delTrans.ExecuteTemplate(w, "main", nil)
}

func GetStorage(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres dbname=Flower sslmode=disable password = '123'"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", err)
		return
	}
	rows, err := db.Query("select * from \"Flower\".Flower")
	if err != nil {
		fmt.Fprintf(w, "<h3 align=\"center\">%s</h3>\n", err)
		return
	}

	defer rows.Close()

	Body2 := "Данные по продажам цветочного магазина \"Цветик семицветик\":"
	fmt.Fprintf(w, "<h1 align=\"center\">%s\n</h1>", Body2)
	for rows.Next() {
		var id int
		var name_fl string
		var count_fl int64
		var money_fl float64
		var date_fl time.Time
		err = rows.Scan(&id, &name_fl, &count_fl, &money_fl, &date_fl)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "<h4>"+"&emsp;"+"&emsp;"+"Идентификатор: %d"+"&emsp;"+"&emsp;"+"Наименование: %s"+"&emsp;"+"&emsp;"+"&emsp;"+"Количество: %d"+"&emsp;"+"&emsp;"+"&emsp;"+"Цена (за штуку): %v"+"&emsp;"+"&emsp;"+"&emsp;"+"Дата продажи: %s"+"&emsp;</h4>\n", id, name_fl, count_fl, money_fl, date_fl)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
		return
	}
}

func GetData(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres dbname=Flower sslmode=disable password = '123'"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", err)
		return
	}
	rows2, err := db.Query("select * from \"Flower\".\"Flower\".\"ShopFlower\"")
	if err != nil {
		fmt.Fprintf(w, "<h3 align=\"center\">%s</h3>\n", err)
		return
	}

	defer rows2.Close()

	Body2 := "Данные по складу цветочного магазина \"Цветик семицветик\":"
	fmt.Fprintf(w, "<h1 align=\"center\">%s\n</h1>", Body2)
	for rows2.Next() {
		var id int
		var name_fl string
		var count_fl int64
		var date_fl_shop time.Time
		err = rows2.Scan(&id, &name_fl, &count_fl, &date_fl_shop)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "<h4>"+"&emsp;"+"&emsp;"+"Идентификатор: %d"+"&emsp;"+"&emsp;"+"Наименование: %s"+"&emsp;"+"&emsp;"+"&emsp;"+"Количество: %d"+"&emsp;"+"&emsp;"+"&emsp;"+"Дата поступления: %s"+"&emsp;</h4>\n", id, name_fl, count_fl, date_fl_shop)
	}
	err = rows2.Err()
	if err != nil {
		panic(err)
		return
	}
}

func AddFlower(w http.ResponseWriter, r *http.Request) {
	af, err := template.ParseFiles("addFlower.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	af.ExecuteTemplate(w, "addFlower", nil)
}

func SaveTransaction(w http.ResponseWriter, r *http.Request) {
	name_fl := r.FormValue("name_fl")
	count_fl, _ := strconv.Atoi(r.FormValue("count_fl"))
	money_fl, _ := strconv.ParseFloat(r.FormValue("money_fl"), 32)
	date_fl := r.FormValue("date_fl")

	db, err := sql.Open("postgres", "user=postgres dbname=Flower sslmode=disable password = 123")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO \"Flower\".Flower (name_fl, count_fl, money_fl, date_fl) VALUES ('%s', %d, %f, '%s')", name_fl, count_fl, money_fl, date_fl)
	insert, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/storage", http.StatusSeeOther)
}

func UpFlower(w http.ResponseWriter, r *http.Request) {
	af, err := template.ParseFiles("update_fl.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	af.ExecuteTemplate(w, "update_fl", nil)
}

func UpdateDataFlower(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name_fl := r.FormValue("name")
	count_fl, _ := strconv.Atoi(r.FormValue("opf"))
	money_fl, _ := strconv.ParseFloat(r.FormValue("inn"), 32)
	date_fl := r.FormValue("kpp")

	db, err := sql.Open("postgres", "user=postgres dbname=Flower sslmode=disable password = 123")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `UPDATE "Flower".Flower SET name_fl = $2, count_fl = $3, money_fl = $4, date_fl = $5 where id = $1`
	if err != nil {
		panic(err)
	}

	res, err := db.Exec(sqlStatement, id, name_fl, count_fl, money_fl, date_fl)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	http.Redirect(w, r, "/storage", http.StatusSeeOther)
}

func DelFlower(w http.ResponseWriter, r *http.Request) {
	delTrans, err := template.ParseFiles("del_flower.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	delTrans.ExecuteTemplate(w, "del_flower", nil)
}

func DelFlowerTrans(w http.ResponseWriter, r *http.Request) {
	DTrans, _ := strconv.Atoi(r.FormValue("DTrans"))

	db, err := sql.Open("postgres", "user=postgres dbname=Flower sslmode=disable password = 123")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var rezult, _ = db.Exec("DELETE FROM \"Flower\".Flower WHERE id = $1", DTrans)
	if err != nil {
		panic(err)
	}

	fmt.Println(rezult.RowsAffected())
	http.Redirect(w, r, r.Header.Get("/data"), http.StatusFound)
}

func corStorage(w http.ResponseWriter, r *http.Request) {
	af, err := template.ParseFiles("corstorage.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	af.ExecuteTemplate(w, "corstorage", nil)
}

func CorrectStorage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name_fl := r.FormValue("name_fl")
	count_fl, _ := strconv.Atoi(r.FormValue("count_fl"))
	date_fl_shop := r.FormValue("date_fl_shop")

	db, err := sql.Open("postgres", "user=postgres dbname=Flower sslmode=disable password = 123")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `UPDATE "Flower"."Flower"."ShopFlower" SET name_fl = $2, count_fl = $3, date_fl_shop = $4 where id = $1`
	if err != nil {
		panic(err)
	}

	res, err := db.Exec(sqlStatement, id, name_fl, count_fl, date_fl_shop)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	http.Redirect(w, r, "/data", http.StatusSeeOther)
}

func addStorage(w http.ResponseWriter, r *http.Request) {
	af, err := template.ParseFiles("addstorage.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	af.ExecuteTemplate(w, "addstorage", nil)
}

func addStorageFlower(w http.ResponseWriter, r *http.Request) {
	name_fl := r.FormValue("name_fl")
	count_fl, _ := strconv.Atoi(r.FormValue("count_fl"))
	date_fl_shop := r.FormValue("date_fl_shop")

	db, err := sql.Open("postgres", "user=postgres dbname=Flower sslmode=disable password = 123")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO \"Flower\".\"Flower\".\"ShopFlower\" (name_fl, count_fl, date_fl_shop) VALUES ('%s', %d, '%s')", name_fl, count_fl, date_fl_shop)
	insert, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/data", http.StatusSeeOther)
}
func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/data", GetData)
	http.HandleFunc("/storage", GetStorage)
	http.HandleFunc("/addflower", AddFlower)
	http.HandleFunc("/savetransaction", SaveTransaction)
	http.HandleFunc("/upflower", UpFlower)
	http.HandleFunc("/updflower", UpdateDataFlower)
	http.HandleFunc("/delflower", DelFlower)
	http.HandleFunc("/delflowertrans", DelFlowerTrans)
	http.HandleFunc("/corstorage", corStorage)
	http.HandleFunc("/correctStorage", CorrectStorage)
	http.HandleFunc("/addstorage", addStorage)
	http.HandleFunc("/addstoragefl", addStorageFlower)
	http.ListenAndServe(":3050", nil)
}
