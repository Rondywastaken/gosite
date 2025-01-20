package main

import (
    "context"
    "fmt"
    "net/http"
    "log"
    "os"
    "encoding/json"
    "time"

    "github.com/lrstanley/go-ytdlp"
)

func main()  {
    //http.HandleFunc("/", handler)
    //log.Fatal(http.ListenAndServe(":8080", nil))
    exampleClient()
}

func handler(res http.ResponseWriter, req *http.Request) {
    fmt.Println(res, "Hello World!")
}

func exampleClient() {
 	ytdlp.MustInstall(context.TODO(), nil)

	dl := ytdlp.New().
		PrintJSON().
		NoProgress().
		FormatSort("res,ext:mp4:m4a").
		RecodeVideo("mp4").
		NoPlaylist().
		NoOverwrites().
		Continue().
		ProgressFunc(100*time.Millisecond, func(prog ytdlp.ProgressUpdate) {
			fmt.Printf( //nolint:forbidigo
				"%s @ %s [eta: %s] :: %s\n",
				prog.Status,
				prog.PercentString(),
				prog.ETA(),
				prog.Filename,
			)
		}).
		Output("%(extractor)s - %(title)s.%(ext)s")

	r, err := dl.Run(context.TODO(), "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("results.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")

	if err = enc.Encode(r); err != nil {
		panic(err)
	}

	log.Println("wrote results to results.json")   
}
