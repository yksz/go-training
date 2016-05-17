package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"

	"gopl.io/ch7/eval"
)

var env = eval.Env{"PI": math.Pi}

type Calc struct {
	Expr string
	Ans  string
}

func calculate(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return "", err
	}
	ans := fmt.Sprintf("%.6g", expr.Eval(env))
	return ans, nil
}

func main() {
	calcTemplate := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		expr := r.FormValue("expr")
		ans, err := calculate(expr)
		if err != nil {
			log.Print(err)
			ans = "ERROR"
		}
		calc := &Calc{Expr: expr, Ans: ans}
		if err := calcTemplate.Execute(w, calc); err != nil {
			log.Print(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
