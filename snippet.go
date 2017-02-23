// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"io/ioutil"
	"golang.org/x/net/context"

	// [START imports]
//	"golang.org/x/oauth2/google"
//	"google.golang.org/api/compute/v1"
	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
//	"google.golang.org/api/option"
	// [END imports]
)
func main() {
	ctx := context.TODO()

	//client, err := google.DefaultClient(ctx, compute.ComputeScope)
	client, err := translate.NewClient(ctx)
	if err != nil {
  // Handle error.
	}
	//computeService, err := compute.New(client)
	if err != nil {
  // Handle error.
	}
    	fmt.Printf("hello ramesh\n")

//}
//func createClientWithKey() {
	// import "cloud.google.com/go/translate"
	// import "google.golang.org/api/option"
	// import "golang.org/x/text/language"
	//ctx := context.Background()
	langs, err := client.SupportedLanguages(ctx, language.English)
	for _, lang := range langs {
		fmt.Printf("%q: %s\n", lang.Tag, lang.Name)
	}
	const apiKey = "YOUR_TRANSLATE_API_KEY"
	//client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	//if err != nil {
	//	log.Fatal(err)
	//}

	resp, err := client.Translate(ctx, []string{"Hello, world!"}, language.Russian, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", resp)
	//fmt.Printf(
	//textout,err := translateText("es","The quick brown fox jumped over the lazy dog.")
	//translateTextn("The quick brown fox jumped over the lazy dog.")
	//fmt.Printf(textout)
	translateTextnf("translate.txt")
}
func translateTextnf(file string) (string, error) {
	ctx := context.Background()
	//ngs, err := client.SupportedLanguages(ctx, language.English)
	//for _, lang := range langs {
	//	fmt.Printf("%q: %s\n", lang.Tag, lang.Name)
	//}
	//lang, err := language.Parse(targetLanguage)
//	if err != nil {
//		return "", err
//	}
	textb, err := ioutil.ReadFile(file)
    	if err != nil {
     //   panic(err)
    	}
	text := string(textb)
	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()
	langs, err := client.SupportedLanguages(ctx, language.English)
	for _, lang := range langs {
		fmt.Printf("%q: %s\n", lang.Tag, lang.Name)
//		lang1, err := language.Parse(lang.Tag)
	
		resp, err := client.Translate(ctx, []string{text}, lang.Tag, nil)
		fmt.Printf(resp[0].Text)
		fmt.Printf("\n")
		if err != nil {
		return "", err
	}
	}
	if err != nil {
		return "", err
	}
	//return resp[0].Text, nil
	return "", nil
}

func translateTextn(text string) (string, error) {
	ctx := context.Background()
	//ngs, err := client.SupportedLanguages(ctx, language.English)
	//for _, lang := range langs {
	//	fmt.Printf("%q: %s\n", lang.Tag, lang.Name)
	//}
	//lang, err := language.Parse(targetLanguage)
//	if err != nil {
//		return "", err
//	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()
	langs, err := client.SupportedLanguages(ctx, language.English)
	for _, lang := range langs {
		fmt.Printf("%q: %s\n", lang.Tag, lang.Name)
//		lang1, err := language.Parse(lang.Tag)
	
		resp, err := client.Translate(ctx, []string{text}, lang.Tag, nil)
		fmt.Printf(resp[0].Text)
		fmt.Printf("\n")
		if err != nil {
		return "", err
	}
	}
	if err != nil {
		return "", err
	}
	//return resp[0].Text, nil
	return "", nil
}

func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

func detectLanguage(text string) (*translate.Detection, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	return &lang[0][0], nil
}

func listSupportedLanguages(w io.Writer, targetLanguage string) error {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	langs, err := client.SupportedLanguages(ctx, lang)
	if err != nil {
		return err
	}

	for _, lang := range langs {
		fmt.Fprintf(w, "%q: %s\n", lang.Tag, lang.Name)
	}

	return nil
}
