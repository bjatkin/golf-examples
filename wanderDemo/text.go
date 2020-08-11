package main

import (
	"fmt"
	"strings"
)

var allLines = []string{
	"Sorry kid, we can't let you pass.",
	"Ya, the developer forgot to code in this area! If we let you pass, the game might crash",
	"Ending all life as we know it!",
	"The end of the universe!",
	"Armageden realized!",
	"... or you might to end up off screen. Its hard to know for sure.",
	"achoo... *sniffel*",
	"Oh it's you! Our town has waited 100 years for your arival! The prophasy has been fulfilled!",
	"You shal bring peach an justice to this land! Here young warrior take this magic sword and face your destitny!",
	"You got a used tissue...",
	"...blah blah blah wisdom and stuff... haha, I love being old :).",
	"I'm pretty sure my dad is loosing it. ...At least he seems happy.",
	"I'm so tired of standing by this stupid well. I just feel like I've been rooted to this spot forever you know.",
	"fortunately I had a coin to make a wish. Once it's granted I'm gonna see the world!",
	"I heard that guy call me stupid. There's no way I'm granting his wish now! ",
	"good luck seeing the world with now path finding ai weirdo!",
	"Man, I'd love to be a fish... Clothes are the worst.",
	"Fish get to be naked any time they want. Stupid lucky fish...",
	"Man, my land lord raised the rent again. I don't know how I'm gonna keep up with it.",
	"Well, I don't know how i'll do it next month but I managed to scrape together 18,000 kibble in time to pay my rent.",
	"Does it seem odd to you that I have to pay my rent in dog food?",
	"*bark* *bark* man home ownership is great.",
	"All these lazy millenials complain that it's impossible.",
	"In my expericne all it takes is a little hard work and a small loan of a million dollars from your dad.",
	"{sec}...\n{sec+1}...\n{sec+2}...",
	"I'm gonna be the first kid ever to count to 256!",
	"Oh good! You're here!",
	"I wanted someone to witness me finishing the count!",
	"254...\n255...",
	"1!!!\nWait what???",
	"I'm so confused...",
}

func fillTemplate(in string) string {
	ret := strings.ReplaceAll(in, "{sec}", fmt.Sprintf("%d", g.Frames()/60))
	ret = strings.ReplaceAll(ret, "{sec+1}", fmt.Sprintf("%d", (g.Frames()/60)+1))
	ret = strings.ReplaceAll(ret, "{sec+2}", fmt.Sprintf("%d", (g.Frames()/60)+2))
	return ret
}
