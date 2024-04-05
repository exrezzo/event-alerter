//#!/Users/at0794/go/bin/gorun

package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"log"
	"time"
)

func main() {
	fmt.Println("COMPILATO! | color=blue")
	fmt.Println("---")
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	fmt.Println("Open website | href=https://xbarapp.com | color=red")

	playSound()
}

//go:embed chinese-beat-190047.mp3
var music embed.FS

func playSound() {
	open, err := music.Open("chinese-beat-190047.mp3")
	if err != nil {
		panic(err)
	}

	streamer, format, err := mp3.Decode(open)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	//speaker.Play(streamer)
	//time.Sleep(3 * time.Second)
}
