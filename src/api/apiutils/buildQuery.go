package apiutils

import "strings"

func BuildQuery(qb *strings.Builder, direction, pivot, offsetBegin, offsetEnd string) {
	if direction == "-1" { //first page
		//ok return a sorted query of the first 10
		qb.WriteString("ORDER BY ")
		qb.WriteString(pivot)
		qb.WriteString(" LIMIT 10;")
	} else {
		qb.WriteString("WHERE ")
		qb.WriteString(pivot)
		if direction == "1" { //going next
			qb.WriteString(" > ")
			qb.WriteString(offsetEnd)
		} else {
			qb.WriteString(" < ")
			qb.WriteString(offsetBegin)
			qb.WriteString(" and ")
			qb.WriteString(pivot)
			qb.WriteString(" > ")
			qb.WriteString(offsetEnd)
		}
		qb.WriteString(" LIMIT 10;")
	}
}
