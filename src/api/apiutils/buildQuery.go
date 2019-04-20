package apiutils

import (
	"log"
	"strings"
)

func BuildQuery(direction, pivot, offsetBegin, offsetEnd string) string {
	qb := strings.Builder{}
	if direction == "-1" { //first page
		//ok return a sorted query of the first 10
		qb.WriteString("SELECT id,pserial,pname,pdesc,quantity FROM products ")
		qb.WriteString("ORDER BY ")
		qb.WriteString(pivot)
		qb.WriteString(" LIMIT 10;")
	} else {
		qb.WriteString("SELECT id,pserial,pname,pdesc,quantity FROM ")
		if direction == "1" { //going next
			qb.WriteString("products where ")
			qb.WriteString(pivot)
			qb.WriteString(" > ")
			qb.WriteString(offsetEnd)
			qb.WriteString(" LIMIT 10;")
		} else {
			qb.WriteString("(SELECT id,pserial,pname,pdesc,quantity FROM products where ")
			qb.WriteString(pivot)
			qb.WriteString(" < ")
			qb.WriteString(offsetBegin)
			qb.WriteString(" ORDER BY ")
			qb.WriteString(pivot)
			qb.WriteString(" DESC LIMIT 10)t ORDER BY ")
			qb.WriteString(pivot)
			qb.WriteString(";")
		}
	}
	log.Println(qb.String())
	return qb.String()
}
