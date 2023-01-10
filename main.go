package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"sync"

	fb "github.com/huandu/facebook"
)

var pageName = flag.String("n", "", "Facebook page name such as: scottiepippen")
var numOfWorkersPtr = flag.Int("c", 2, "the number of concurrent rename workers. default = 2")
var mutex sync.Mutex
var TOKEN string

func init() {
	TOKEN = os.Getenv("FBTOKEN")
}

func downloadWorker(destDir string, linkChan chan DLData, wg *sync.WaitGroup) {
	defer wg.Done()

	for target := range linkChan {
		var imageType string

		if strings.Contains(target.ImageURL, ".png") {
			imageType = ".png"
		} else {
			imageType = ".jpg"
		}

		resp, err := http.Get(target.ImageURL)
		
		if err != nil {
			log.Println("Http.Get\nerror: " + err.Error() + "\ntarget: " + target.ImageURL)
			
			continue
		}

		defer resp.Body.Close()

		m, _, err := image.Decode(resp.Body)
		
		if err != nil {
			log.Println("image.Decode\nerror: " + err.Error() + "\ntarget: " + target.ImageURL)
			continue
		}

		bounds := m.Bounds()
		
		if bounds.Size().X > 300 && bounds.Size().Y > 300 {
			out, err := os.Create(destDir + "/" + target.ImageID + imageType)
			if err != nil {
				log.Println("os.Create\nerror: %s", err)
				continue
			}

			defer out.Close()
			
			if imageType == ".png" {
				png.Encode(out, m)
			} else {
				jpeg.Encode(out, m, nil)
			}
		}
	}
}

func parseMapToStruct(inData interface{}, decodeStruct interface{}) {
	jret, _ := json.Marshal(inData)
	err := json.Unmarshal(jret, &decodeStruct)

	if err != nil {
		log.Fatalln(err)
	}
}

func findPhotoFromAlbum(ownerName, albumName, albumId, baseDir string, photoCount, photoOffset int) {
	photoRet := FBPhotos{}
	var queryString string

	if photoOffset > 0 {
		queryString = fmt.Sprintf("/%s/photos?limit=%d&offset=%d", albumId, photoCount, photoOffset)

	} else {
		queryString = fmt.Sprintf("/%s/photos?limit=%d", albumId, photoCount)
	}

	resPhoto := runFBGraphAPI(queryString)
	parseMapToStruct(resPhoto, &photoRet)
	dir := fmt.Sprintf("%v/%v/%v - %v", baseDir, ownerName, albumId, albumName)

	os.MkdirAll(dir, 0755)

	linkChannel := make(chan DLData)
	wg := new(sync.WaitGroup)

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go downloadWorker(dir, linkChannel, wg)
	}

	for _, value := range photoRet.Data {
		dlChan := DLData{}
		dlChan.ImageID = value.ID
		dlChan.ImageURL = value.Source
		linkChannel <- dlChan
	}
}

func runFBGraphAPI(query string) (queryResult interface{}) {
	res, err := fb.Get(query, fb.Params{
		"access_token": TOKEN,
	})

	if err != nil {
		log.Fatalln("FB connect error, err=", err.Error())
	}
	return res
}

func main() {
	flag.Parse()
	var inputPage string

	if TOKEN == "" {
		log.Fatalln("Set your FB token as environment variables 'export FBTOKEN=XXXXXXX'")
	}

	if *pageName == "" {
		log.Fatalln("You need to input -n=Name_or_ID.")
	}

	inputPage = *pageName

	user, _ := user.Current()
	baseDir := fmt.Sprintf("%v/Pictures/fb_photo_downloader", user.HomeDir)

	resUser := runFBGraphAPI("/" + inputPage + "/albums")
	userRet := FBUser{}
	parseMapToStruct(resUser, &userRet)

	resAlbums := runFBGraphAPI("/" + inputPage + "/albums")
	albumRet := FBAlbums{}
	parseMapToStruct(resAlbums, &albumRet)

	const maxCount int = 30

	userFolderName := fmt.Sprintf("[%s]%s", userRet.UserName, userRet.Name)

	for _, value := range albumRet.Data {
		fmt.Println("Starting download ["+value.Name+"]-"+value.From.Name, " total count:", value.Count)

		if value.Count > maxCount {
			currentOffset := 0

			for {
				if currentOffset > value.Count {
					break
				}

				findPhotoFromAlbum(userFolderName, value.Name, value.ID, baseDir, maxCount, currentOffset)
				currentOffset = currentOffset + maxCount
			}
		} else {
			findPhotoFromAlbum(userFolderName, value.Name, value.ID, baseDir, value.Count, 0)
		}
	}
}
