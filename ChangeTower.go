package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"flag"
	"github.com/rapidloop/skv"
)

var wg sync.WaitGroup
var mut sync.Mutex

func Database(link string, length string, d bool, l bool){
	var old string
	store, _ := skv.Open("len.db")
	_ = store.Get(link, &old)
	if old == ""{
		if d == false{
			log.Printf("New %s\n",link)
		}
		if l == false{
			fmt.Printf("New %s\n",link)
		}else{
			fmt.Printf("%s\n",link)
		}
		_ = store.Put(link, length)
	}else{
		switch old {
			case length:
			default:
				_ = store.Delete(link)
				_ = store.Put(link, length)
				if d == false{
					log.Printf("Changed %s\n",link)
				}
				if l == false{
					fmt.Printf("Changed %s\n",link)
				}else{
					fmt.Printf("%s\n",link)
				}
		}
	}
	err := store.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func Length(link string, d bool, l bool) {
	defer wg.Done()
	resp, err := http.Get(link)
	if resp == nil {
		if err != nil {
			log.Fatal(err)
		}
		length := "0"
		mut.Lock()
		defer mut.Unlock()
		Database(link,length,d,l)
	} else {
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		length := fmt.Sprintf("%d", len(body))
		mut.Lock()
		defer mut.Unlock()
		Database(link,length,d,l)
	}
}

const (
	LinkColor = "\033[1;34m%s\033[0m"
	HelpColor = "\033[1;36m%s\033[0m"
	BannerMColor = "\033[1;33m%s\033[0m"
	BannerColor = "\033[1;31m%s\033[0m"
)

func main(){
	s := flag.Bool("s",false,"silent mode (no banner)")
	u := flag.Bool("u",false,"example of usage")
	d := flag.Bool("d",false,"dont log the data")
	w := flag.Bool("w",false,"without color (for windows user)")
	l := flag.Bool("l",false,"return links without any tag (dont effect in log file)")
	flag.Parse()
	if *s == false && *w == true{
			fmt.Printf("\n╔═╗┬ ┬┌─┐┌┐┌┌─┐┌─┐  ╔╦╗┌─┐┬ ┬┌─┐┬─┐\n║  ├─┤├─┤││││ ┬├┤    ║ │ ││││├┤ ├┬┘\n╚═╝┴ ┴┴ ┴┘└┘└─┘└─┘   ╩ └─┘└┴┘└─┘┴└─\nThe Cats : https://github.com/DC4ts\n\n")
	}else if *s == false{
		fmt.Printf(BannerColor,"\n╔═╗┬ ┬┌─┐┌┐┌┌─┐┌─┐  ╔╦╗┌─┐┬ ┬┌─┐┬─┐")
		fmt.Printf(BannerMColor,"\n║  ├─┤├─┤││││ ┬├┤    ║ │ ││││├┤ ├┬┘\n")
		fmt.Printf(BannerColor,"╚═╝┴ ┴┴ ┴┘└┘└─┘└─┘   ╩ └─┘└┴┘└─┘┴└─\n")
		fmt.Printf("The Cats : ")
		fmt.Printf(LinkColor,"https://github.com/DC4ts\n\n")
	}
	if *u == true && *w == true{
			FileName := os.Args[0]
			fmt.Printf("for list of urls:\n\tcat links.txt | %s\nfor single url:\n\techo \"https://example.com\" | %s\n",FileName,FileName)
	}else if *u == true{
			FileName := os.Args[0]
			fmt.Printf(HelpColor,"for list of urls:\n\t")
			fmt.Printf("cat links.txt | %s\n",FileName)
			fmt.Printf(HelpColor,"for single url:\n\t")
			fmt.Printf("echo \"https://example.com\" | %s\n",FileName)
	}else{
		if *d == false{
			now := time.Now()
			year := now.Year()
			month := int(now.Month())
			day := now.Day()
			hour := now.Hour()
			min := now.Minute()
			name := fmt.Sprintf("%d-%d-%d-%d-%d.md",year,month,day,hour,min)
			file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			log.SetOutput(file)
			if *s == false{
				fmt.Printf("result log file: %s\n",name)
			}
		}
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		scanner := bufio.NewScanner(os.Stdin)
		var link string
		for scanner.Scan(){
			link = scanner.Text()
			go Length(link,*d,*l)
			wg.Add(1)
		}
		wg.Wait()
	}
}

