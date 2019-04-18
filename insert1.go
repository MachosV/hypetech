package main

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

type product struct {
	ID       string
	Name     string
	Desc     string
	Quantity int
}

func ExitOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var flag bool = false

func main() {
	colors := [13]string{
		"white",
		"black",
		"red",
		"green",
		"orange",
		"blue",
		"purple",
		"brown",
		"yellow",
		"golden",
		"silver",
		"bronze",
		"platinum",
	}
	items := [7]string{
		"football",
		"basketball",
		"baseball",
		"action-figure",
		"bicycle",
		"card",
		"comic",
	}

	desc1 := [4]string{
		"An interesting, really really interesting item.",
		"Not really interesting item.",
		"A somewhat interesting item.",
		"An extraordinary item.",
	}
	desc2 := [5]string{
		"Holds insignificant value.",
		"Holds really significant value.",
		"Somewhat significant value.",
		"Not worth your time. Low quality item.",
		"Hold a tremendous amount of value.",
	}
	desc3 := [9]string{
		"It's size is really small.",
		"The size of the item is small.",
		"Medium sized item.",
		"A large item.",
		"A really large item.",
		"Huge item.",
		"Humongous item.",
		"An item of titanic proportions.",
		"Are you fucking kidding me porportions.",
	}

	desc4 := [6]string{
		"Not shiny at all.",
		"A bit shiny.",
		"You could tell it's shiny.",
		"Yes. It shines. Yes. Precious.",
		"Oh my, this one definetely shines.",
		"Buy it now. The shine can be seen from a mile away.",
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	dbHandler := os.Getenv("DB_HANDLER")
	if dbHandler == "" {
		log.Fatalln("No db handler")
	}
	log.Println(dbHandler)
	db, err := sql.Open("mysql", dbHandler)
	ExitOnError(err)
	stmt, err := db.Prepare("INSERT into products (pserial,pname,pdesc,quantity) VALUES(?,?,?,?)")
	ExitOnError(err)
	go logger()
	for i := 0; i < 50*1000; i++ {
		id := uuid.Must(uuid.NewV4()).String()
		name := colors[r1.Int()%13] + " - " + items[r1.Int()%7]
		desc := desc1[rand.Int()%4] + " " + desc2[rand.Int()%5] + " " + desc3[rand.Int()%9] + " " + desc4[rand.Int()%6]
		qnt := rand.Int() % 120
		stmt.Exec(id, name, desc, qnt)
	}
	flag = true
}

func logger() {
	for {
		log.Println(" ")
		time.Sleep(1 * time.Second)
		if flag {
			break
		}
	}
}

/*comment*/
