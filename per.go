// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/performance"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://visiblealpha.com/`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			err := performance.SetTimeDomain(performance.SetTimeDomainTimeDomainTimeTicks).Do(ctx)
			if err != nil {
				return err
			}

			err = performance.Enable().Do(ctx)
			if err != nil {
				return err
			}

			metrics, err := performance.GetMetrics().Do(ctx)
			if err != nil {
				return err
			}
			for _, element := range metrics {
				fmt.Println(element)
			}
			return nil
		}),
	}
}
