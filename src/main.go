package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprintf(w, "Random Image Server 0.0.1! Request [Get] http://%s/rand-image", r.Host)
	})
	router.Get("/rand-image", randImageHandler)

	port := 9811
	log.Println(fmt.Sprintf("The random image server listens at %s:%d", "http://localhost", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func randImageHandler(w http.ResponseWriter, r *http.Request) {
	imgPaths, err := dirwalk("files")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	randInt := getRandomInt(0, len(imgPaths))

	imgPath := imgPaths[randInt]

	img, err := os.Open(imgPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer img.Close()

	buf := make([]byte, 512)
	img.Read(buf)
	contentType := http.DetectContentType(buf)
	img.Seek(0, 0)

	w.Header().Set("Content-Type", contentType)
	io.Copy(w, img)
}

func dirwalk(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			nexts, err := dirwalk(filepath.Join(dir, file.Name()))
			if err != nil {
				continue
			}
			paths = append(paths, nexts...)
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, nil
}

func getRandomInt(min, max int) int {
	if min > max {
		return 0
	}
	return rand.Intn(max-min) + min
}
