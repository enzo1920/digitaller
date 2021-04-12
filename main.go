package main

import (
    "fmt"
    "image"
    "log"
    "os"
    "path"
    "path/filepath"
    "io/ioutil"
    "github.com/jung-kurt/gofpdf"
)

func main() {
    err := PdfJpegGenerate("hello.pdf")
    if err != nil {
        panic(err)
    }
}

// Find file in folder for creation jpg
func PdfJpegGenerate(filename string) error {
    dir_to_scan := "./img"
    files, err := ioutil.ReadDir(dir_to_scan)
    if err != nil {
       log.Fatal(err)
    }

    //im, _, err := image.DecodeConfig()
    //create pdf document 
    //pdf := gofpdf.New("P", "pt ", "A4", "")
    pdf := gofpdf.NewCustom(&gofpdf.InitType{
            OrientationStr:  "P",
            UnitStr:        "pt",
            Size: gofpdf.SizeType{
                   Ht: 5000.0,
                   Wd: 4000.0,
            },
    })
    //pdf.SetFont("Arial", "B", 16)
    for _, f := range files {

        if reader, err := os.Open(filepath.Join(dir_to_scan, f.Name())); err == nil {
            defer reader.Close()
            im, _, err := image.DecodeConfig(reader)
            if err != nil {
                fmt.Fprintf(os.Stderr, "%s: %v\n", f.Name(), err)
                continue
            }
             pdf.AddPage()
             // ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
             pdf.ImageOptions(
             path.Join("./img/",f.Name()),
             0, 0,
             float64(im.Width), float64(im.Height),
             false,
             gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
             0,
             "",
               )
            fmt.Printf("%s %d %d\n", f.Name(), im.Width, im.Height)


          } else {
              fmt.Println("Impossible to open the file:", err)
            }

        fmt.Println(f)
    }

    return pdf.OutputFileAndClose(filename)
}


// GeneratePdf generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdf(filename string) error {

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    // CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
    pdf.CellFormat(190, 7, "Welcome to golangcode.com", "0", 0, "CM", false, 0, "")

    // ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
    pdf.ImageOptions(
        "avatar.jpg",
        80, 20,
        0, 0,
        false,
        gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
        0,
        "",
    )

    return pdf.OutputFileAndClose(filename)
}
