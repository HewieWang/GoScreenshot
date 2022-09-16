package main

import (
    "github.com/go-rod/rod"
    "os"
)

func main() {
    page := rod.New().MustConnect().MustPage(os.Args[1])
    page.MustWaitLoad().MustScreenshot(os.Args[2])
}
