package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var html = `<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Immutable demo</title>
</head>
<script type="application/javascript">
</script>
<body>
napis.jest.tu
<script src="s.js"></script>
</body>
</html>`

func main() {
	cnt := 0
	http.HandleFunc("/s.js", func(w http.ResponseWriter, r *http.Request) {
		log.Print("redirect")
		if cnt < 10 {
			http.Redirect(w, r, "/2017/10/09/f.js", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/2017/10/10/f.js", http.StatusSeeOther)
		}
	})
	http.HandleFunc("/2017/10/09/f.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public,max-age=31536000,immutable")
		fmt.Fprintf(w, "alert('file 1 %s');", time.Now())
		log.Print("file 1")
	})
	http.HandleFunc("/2017/10/10/f.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public,max-age=31536000,immutable")
		fmt.Fprintf(w, "alert('file 2 %s');", time.Now())
		log.Print("file 2")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})
	http.ListenAndServe(":8080", nil)
}
