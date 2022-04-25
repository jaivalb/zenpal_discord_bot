package avatar

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	AvatarFile string
	AvatarURL  string
)

func init() {

	//flag.StringVar(&AvatarFile, "f", "", "Avatar File Name")
	flag.StringVar(&AvatarURL, "u", "", "URL to the Avatar Image")
	flag.Parse()

	if AvatarFile == "" && AvatarURL == "" {
		flag.Usage()
		os.Exit(1)
	}

}

func main() {
	var base64img string
	var contentType string

	if AvatarURL != "" {
		resp, err := http.Get(AvatarURL)
		if err != nil {
			fmt.Println("Error retrieving the file, ", err)
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		img, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading the response, ", err)
			return
		}

		contentType = http.DetectContentType(img)
		base64img = base64.StdEncoding.EncodeToString(img)
	}
	if AvatarFile != "" {
		img, err := ioutil.ReadFile(AvatarFile)
		if err != nil {
			fmt.Println(err)
		}

		contentType = http.DetectContentType(img)
		base64img = base64.StdEncoding.EncodeToString(img)
	}

	avatar := fmt.Sprintf("data:%s;base64,%s", contentType, base64img)
	_ = avatar

}
