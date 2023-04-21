package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/godown"
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/browser/all"
)

const cookieTemplate = `{
	"Name": "session",
	"Value": "%value%",
	"Path": "/",
	"Domain": ".adventofcode.com",
	"Expires": "2050-01-01T10:10:10.317088+12:00",
	"RawExpires": "",
	"MaxAge": 0,
	"Secure": true,
	"HttpOnly": true,
	"SameSite": 0,
	"Raw": "",
	"Unparsed": null
}`

var normalMode = os.FileMode(0644)

func createDirectoryIfNotExists(path string) error {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func tryGetCookieFromFile() (string, error) {
	data, err := os.ReadFile("cookie.txt")
	return strings.Trim(string(data), "\r\n"), err
}

func GetBrowserCookie() string {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.adventofcode.com`), kooky.Name(`session`))
	for _, cookie := range cookies {
		return cookie.Value
	}
	return ""
}

func TryGetCookie() (string, error) {
	cookie, err := tryGetCookieFromFile()
	if err != nil {
		cookie = GetBrowserCookie()
		if cookie == "" {
			return "", fmt.Errorf("unable to find cookie: failed to get cookie from file: %v; failed to get cookie from browser", err)
		}
	}
	return cookie, nil
}

func DownloadInput(cookie string, day int, inputFile string) {
	if _, err := os.Stat(inputFile); err == nil {
		return
	} else if !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	var sessionToken http.Cookie
	json.Unmarshal([]byte(strings.ReplaceAll(cookieTemplate, `%value%`, cookie)), &sessionToken)
	req.AddCookie(&sessionToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("unexpected status code %s downloading %s: %s", resp.Status, url, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = createDirectoryIfNotExists(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(inputFile, body, normalMode)
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadDescriptions(cookie string) {
	for i := 1; i <= 25; i++ {
		fmt.Printf("Downloading day %d\n", i)
		url := fmt.Sprintf("https://adventofcode.com/2021/day/%d", i)
		outputFile := fmt.Sprintf("day%02d.md", i)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}
		if cookie != "" {
			var sessionToken http.Cookie
			json.Unmarshal([]byte(strings.ReplaceAll(cookieTemplate, `%value%`, cookie)), &sessionToken)
			req.AddCookie(&sessionToken)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			log.Fatalf("unexpected status code %s downloading %s: %s", resp.Status, url, body)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		startMarker := "<article"
		stopMarker := "</article>"

		data := string(body)
		err = createDirectoryIfNotExists(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		defer w.Flush()

		for {
			index := strings.Index(data, startMarker)

			if index < 0 {
				break
			}
			data = data[index:]
			index = strings.Index(data, stopMarker)
			if index < 0 {
				break
			}
			block := data[:index+len(stopMarker)]
			data = data[index+len(stopMarker):]

			err := godown.Convert(w, strings.NewReader(block), &godown.Option{})

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
