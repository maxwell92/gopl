package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	cur := time.Now()
	fmt.Println(cur.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2017-01-26T16:11:37+08:00")
	fmt.Println(t1)
	//	fmt.Println(e)
	fmt.Println(cur.Format("3--04AM"))
	form := "3--04"
	t2, err := time.Parse(form, "4--26")
	fmt.Println(t2)
	fmt.Println(err)

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "4:26")
	fmt.Println(e)
}
