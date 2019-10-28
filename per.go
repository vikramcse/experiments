// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/performance"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, getPerformanceMetrics(`https://codingwithvikram.com/`)); err != nil {
		log.Fatal(err)
	}

}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func getPerformanceMetrics(urlstr string) chromedp.Tasks {
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
				if element.Name == "TaskDuration" {
					v := fmt.Sprintf("%f", element.Value)
					fmt.Println(element.Name, "->", v)
				}
			}
			return nil
		}),
	}
}
