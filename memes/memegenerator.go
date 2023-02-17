package memes

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/TannerKvarfordt/imgflipgo"
)

type Generator struct {
	IMGFLIPUsername string
	IMGFLIPPassword string
	l               *log.Logger
}

func NewGenerator(username, pwd string, l *log.Logger) *Generator {
	return &Generator{username, pwd, l}
}

func (m *Generator) Generate(captions []string) (*imgflipgo.CaptionResponse, error) {
	m.l.Println(captions)
	template, err := randomTemplate(len(captions))
	if err != nil {
		m.l.Printf("Can't generate meme with %d captions\n", len(captions))
		return nil, err
	}

	boxes, err := fillCaptions(*template, captions)
	if err != nil {
		m.l.Println(err)
		return nil, err
	}
	m.l.Println(boxes)
	response, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: template.ID,
		Username:   m.IMGFLIPUsername,
		Password:   m.IMGFLIPPassword,
		TextBoxes:  boxes,
	})
	if err != nil {
		m.l.Printf("Caption request failed, err=%v\n", err)
		return nil, err
	}

	return &response, nil
}

func randomTemplate(captures int) (*imgflipgo.Meme, error) {
	memes, err := imgflipgo.GetMemes()
	var newMemes []imgflipgo.Meme
	fmt.Printf("Generating memes with %d captions\n", captures)
	for _, m := range memes {
		if m.BoxCount == uint(captures) {
			newMemes = append(newMemes, m)
		}
	}

	if err != nil {
		return nil, err
	}

	if len(newMemes) == 0 {
		return nil, fmt.Errorf("Can't create meme with %d captions", captures)
	}

	return &newMemes[rand.Intn(len(newMemes))], nil
}

func fillCaptions(template imgflipgo.Meme, captions []string) ([]imgflipgo.TextBox, error) {
	boxes := make([]imgflipgo.TextBox, template.BoxCount)
	if len(boxes) != len(captions) {
		return nil, fmt.Errorf("Can't fill captions(%d), template has different amount of boxes(%d)", len(captions), len(boxes))
	}
	for i := 0; i < len(boxes); i++ {
		boxes[i].Text = captions[i]
		boxes[i].SetColor(0xFFFFFF)   //white
		boxes[i].SetOutlineColor(0x0) //black
	}
	return boxes, nil
}
