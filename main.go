package main

import (
    "fmt"
    "log"
    "os"
    "path"
    "path/filepath"
    "io/ioutil"
    "bytes"
    "github.com/signintech/gopdf"
    "github.com/nfnt/resize"
    "image"
    "image/jpeg"
)

func main() {
     version := "0.0.2"
     fmt.Println("Pdf jpeg creator version:"+version)
     fmt.Println("start dir is:   img")
     fmt.Println("Press any key to start!!!")
     fmt.Scanln()
     start_dir :="./img"
     folders, err := ioutil.ReadDir(start_dir)
     if err != nil {
        log.Fatal(err)
     }
     for _, f := range folders {
            fmt.Println(start_dir+"/"+f.Name())
            PdfJpegGenerate(f.Name()+".pdf", path.Join(start_dir,f.Name()))

     }


}

// Find file in folder for creation jpg
func PdfJpegGenerate(filename string, dir_to_scan string) {

    files, err := ioutil.ReadDir(dir_to_scan)
    if err != nil {
       log.Fatal(err)
    }
        pdf := gopdf.GoPdf{}
        pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 })  
        //buffer for resize jpeg
        buf := new(bytes.Buffer)

    for _, f := range files {

        if reader, err := os.Open(filepath.Join(dir_to_scan, f.Name())); err == nil {
            defer reader.Close()
            pdf.AddPage()
            img, _, err := image.Decode(reader)
            if err != nil {
                log.Fatalln(err)
            }
            new_image := resize.Resize(1024, 473, img, resize.Lanczos3)
            err = jpeg.Encode(buf, new_image, nil)

            imgH1, err := gopdf.ImageHolderByBytes(buf.Bytes())
            if err != nil {
                  log.Print(err.Error())
                  return
            }

            pdf.ImageByHolder(imgH1, 10, 10, nil)

            //clear buffer
            buf.Reset()

          } else {
              fmt.Println("Impossible to open the file:", err)
            }

        fmt.Println(f)
    }

     pdf.WritePdf(filename) //pdf.OutputFileAndClose(filename)
}

/*
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
*/