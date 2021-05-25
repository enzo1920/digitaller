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
     version := "0.0.3"
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
        //path_to_jpg
        fullpath_jpg := ""
        //jpg dimention  info
        width,height :=0,0

    for _, f := range files {
        fullpath_jpg = filepath.Join(dir_to_scan, f.Name())
        if reader, err := os.Open(fullpath_jpg); err == nil {
            defer reader.Close()
            pdf.AddPage()
            img, _, err := image.Decode(reader)
            if err != nil {
                      log.Fatal(err)
            }
            //получаем длину и ширину фотки, чтобы вписать в пдф
            width,height = getImageDimension(fullpath_jpg) 
            if err != nil {
                log.Fatal(err)
            }

            // теперь правильно делаем ресайз
            if height > width {
                   new_image := resize.Resize(1024, 1365, img, resize.Lanczos3)
                   err = jpeg.Encode(buf, new_image, nil)
                   if err != nil {
                        log.Fatal(err)
                   }

            }else{
                   new_image := resize.Resize(1024,768 , img, resize.Lanczos3)
                   err = jpeg.Encode(buf, new_image, nil)
                   if err != nil {
                        log.Fatal(err)
                   }
            }
            //err = jpeg.Encode(buf, new_image, nil)

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

     pdf.WritePdf(filename)
}


func getImageDimension(filepath string)   (int, int) {
    w,h :=0,0
    if file, err := os.Open(filepath); err == nil {
        defer file.Close()
        img, _, err := image.DecodeConfig(file)
        if err != nil {
                log.Fatal(err)
        }
                //fmt.Println("Width:", img.Width, "Height:", img.Height)
                w = img.Width
                h = img.Height
        }else {
              fmt.Println("Impossible to open the file:", err)
        }
    return w, h
}



