package main

import "fmt"

type Author struct {
	Name      string
	VIP       bool
	Icon      string
	Signature string
	Followers int
}

type Video struct {
	views         int
	barrages      int
	likes         int
	collections   int
	coins         int
	transmissions int
}

type RVideo struct {
	authorName string
	videoName string
}

type RelevantRecommendations struct {
	Recommendations []Video
}

type Website struct {
	Author
	Video
	RelevantRecommendations
}

func (a *Video) giveLike() {
	a.likes++
}

func (a *Video) giveCollections() {
	a.collections++
}

func (a *Video) giveCoins() {
	a.coins++
}

func (a *Video) transmit() {
	a.transmissions++
}

//goland:noinspection SpellCheckingInspection
func (a *Video) giveSanlian() {
	a.giveLike()
	a.giveCollections()
	a.giveCoins()
	a.transmit()
}

func  releaseVideo(a string,b string) RVideo{
	s:=RVideo{}
	s.authorName=a
	s.videoName=b
	return s
}

func main() {
	web1:= Website{
		Author: Author{
			"asd",
			true,
			"asd",
			"asd",
			100,
		},
		Video: Video{
			100,
			100,
			100,
			100,
			100,
			100,
		},
		RelevantRecommendations: RelevantRecommendations{
			Recommendations: []Video{
				{1,1,1,1,1,1},
				{1,1,1,1,1,1},
				{1,1,1,1,1,1},
				{1,1,1,1,1,1},
				{1,1,1,1,1,1},
				{1,1,1,1,1,1},
			},
		},
	}

	fmt.Println(web1)
}
