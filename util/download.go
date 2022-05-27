package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zellyn/kooky"
	"github.com/zellyn/kooky/browser/chrome"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

func chromeCookiePath() (string, error) {
	if p, set := os.LookupEnv("CHROME_PROFILE_PATH"); set {
		return filepath.Join(p, "Network", "Cookies"), nil
	}

	if runtime.GOOS == "windows" {
		localAppData, err := os.UserCacheDir()
		return filepath.Join(localAppData, "Google", "Chrome", "User Data", "Default", "Network", "Cookies"), err
	}

	return "", fmt.Errorf("chrome cookie path for GOOS %s not implemented, set CHROME_PROFILE_PATH instead", runtime.GOOS)
}

func tryGetCookieFromFile() (string, error) {
	data, err := os.ReadFile("cookie.txt")
	return strings.Trim(string(data), "\r\n"), err
}

func GetChromeCookie() (string, error) {
	cookiePath, err := chromeCookiePath()
	if err != nil {
		return "", err
	}

	cookies, err := chrome.ReadCookies(cookiePath, kooky.Valid, kooky.Name("session"), kooky.Domain(".adventofcode.com"))
	if err != nil {
		return "", err
	}

	if len(cookies) != 1 {
		return "", fmt.Errorf("session cookie not found or too many results. Got %d, want 1, ensure that you are logged in", len(cookies))
	}

	return cookies[0].Cookie.Value, nil
}

func TryGetCookie() string {
	cookie, err := tryGetCookieFromFile()
	if err != nil {
		cookie, err2 := GetChromeCookie()
		if err2 != nil {
			log.Fatalf("unable to find cookie: failed to get cookie from file: %v; failed to get cookie from chrome: %v", err, err2)
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
