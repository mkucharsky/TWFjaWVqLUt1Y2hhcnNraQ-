package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func (app *application) getDataFromURL(id int64, url string, ) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()

	r, err := client.Get(url)

	if err != nil {
		app.responses.Insert(id, nil, 5, time.Now())
		app.errorLog.Println(err)
		return
	}
	defer r.Body.Close()

	duration := time.Since(start).Seconds()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		app.responses.Insert(id, nil, 5, time.Now())
		app.errorLog.Println(err)
		return
	}

	str := string(body)

	app.responses.Insert(id, &str, duration, time.Now())
}

func (app *application) GetUrlsToWorker() {
	data, err := app.urls.Get()

	if err != nil {
		app.errorLog.Println(err)
	}

	for _, d := range data {
		app.worker.NewJob(d.ID, d.Interval).Run(app.getDataFromURL, d.ID, d.URL)
	}
}
