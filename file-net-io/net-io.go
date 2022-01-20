package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadfile(w http.ResponseWriter, r *http.Request) {

	// restrict upload of 5 MB files.
	r.ParseMultipartForm(5 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	fName, fheader, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer fName.Close()

	fmt.Printf("Uploaded File: %+v\n", fheader.Filename)
	fmt.Printf("File Size: %+v\n", fheader.Size)
	fmt.Printf("MIME Header: %+v\n", fheader.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp", "upload-*.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer tempFile.Close()

	// read all the contents of our uploaded file into a
	// byte array

	// fB, err := ioutil.ReadFile(fheader.Filename)
	fBytes, err := ioutil.ReadAll(fName)
	if err != nil {
		fmt.Printf(err.Error())
	}
	// write this byte array to our temporary file
	tempFile.Write(fBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setRoutes() {
	http.HandleFunc("/load", loadfile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setRoutes()
}
