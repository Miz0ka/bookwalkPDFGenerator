package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//matched, err := filepath.Match("*.mp3", fi.Name())

func main() {
	var shousetsuTOC = []string{
		"cover",
		"title",
		"caution",
		"toc",
		"page",
		"colophon",
	}
	var image, image2 string

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			title := file.Name()
			fmt.Println("<<<<<<<<<< Generation du PDF : " + title + " >>>>>>>>>>>>>")
			pdf := initBooks()
			nbPage := 0
			for _, typePage := range shousetsuTOC {
				image = ""
				image2 = ""
				switch typePage {
				case "page":
					chap := 1
					for chap >= 0 {
						page := 0
						for page >= 0 {
							image = filepath.Join(".", title, typePage+"-"+fmt.Sprintf("%03d", chap)+"-"+fmt.Sprintf("%03d", page)+".jpg")
							image2 = filepath.Join(".", title, typePage+"AndPicture-"+fmt.Sprintf("%03d", chap)+"-"+fmt.Sprintf("%03d", page)+".jpg")
							if fileExists(image) {
								fmt.Println(typePage + " chap " + fmt.Sprintf("%03d", chap) + " page " + fmt.Sprintf("%03d", page) + " Add")
								addPage(image, pdf)
								//numberPage(nbPage, pdf)
								nbPage++
								page++
							} else if fileExists(image2) {
								fmt.Println(typePage + " chap " + fmt.Sprintf("%03d", chap) + " page " + fmt.Sprintf("%03d", page) + " Add")
								addPage(image2, pdf)
								nbPage++
								page++
							} else {
								//log.Println(typePage + " chap " + fmt.Sprintf("%03d", chap) + " page " + fmt.Sprintf("%03d", page) + " missing")
								if chap > 23 && page == 0 {
									chap = -10
								}
								page = -10
							}

						}
						chap++
					}
					break
				case "toc":
					chap := 1
					for chap >= 0 {
						image = filepath.Join(".", title, typePage+"-"+fmt.Sprintf("%03d", chap)+".jpg")
						if fileExists(image) {
							fmt.Println(typePage + " Add")
							addPage(image, pdf)
							chap++
							nbPage++
						} else {
							log.Println(typePage + " missing")
							chap = -1
						}
					}
					break
				default:
					image = filepath.Join(".", title, typePage+".jpg")
					if fileExists(image) {
						fmt.Println(typePage + " Add")
						addPage(image, pdf)
						nbPage++
					} else {
						log.Println(typePage + " missing")
					}
				}

				if err != nil {
					log.Fatal(err)
				}

				//addPage(,pdf)
			}
			err = pdf.OutputFileAndClose(title + ".pdf")
			fmt.Println("Un fichier PDF de " + fmt.Sprintf("%d", nbPage) + " à été generer")
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func initBooks() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A6", "")
	pdf.SetFont("Arial", "B", 16)
	return pdf
}

func addPage(image string, pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.Image(image, 0, 0, 105, 149, false, "", 0, "")
}

func numberPage(num int, pdf *gofpdf.Fpdf) {
	pdf.SetX(-50)
	pdf.SetY(-140)
	pdf.Cell(40, 10, fmt.Sprintf("%d", num))
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
