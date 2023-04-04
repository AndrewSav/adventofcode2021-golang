package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

func tryGetCookieFromFile() (string, error) {
	data, err := os.ReadFile("cookie.txt")
	return strings.Trim(string(data), "\r\n"), err
}

func GetBrowserCookie() (string, error) {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.adventofcode.com`), kooky.Name(`session`))
	for _, cookie := range cookies {
		return cookie.Value, nil
	}
	return "", fmt.Errorf("cookie not found")
}

func TryGetCookie() string {
	cookie, err := tryGetCookieFromFile()
	if err != nil {
		cookie, err2 := GetBrowserCookie()
		if err2 != nil {
			log.Fatalf("unable to find cookie: failed to get cookie from file: %v; failed to get cookie from browser: %v", err, err2)
		}
		return cookie
	}
	return cookie
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

	err = os.WriteFile(inputFile, body, normalMode)
	if err != nil {
		log.Fatal(err)
	}
}
