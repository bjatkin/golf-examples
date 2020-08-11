package main

import (
	"fmt"
	"strings"
)

var allLines = []string{
	"Guard1: \nSorry kid, \nwe cant let you pass.",
	"Guard2: \nYa, the developer forgot \nto code in this area! \nIf we let you pass,\nthe game might crash",
	"Guard1: \nEnding all life \nas we know it!",
	"Guard2: \nThe end of the universe!",
	"Guard1: \nArmageden realized!",
	"Guard2: \n... \nor you might just end up \noff screen. Its hard to \nknow for sure.",
	"Old Man: \nachoo... *sniffel*",
	"Old Man: \nOh its you! \nOur town has waited 100 years \nfor your arival!",
	"Old Man: \nYou shal bring peace\nand justice to this land!\n",
	"Notification: \nYou got a... \nused tissue...",
	"Old Man: \n...blah blah blah \nwisdom and stuff... haha, \nI love being old :).",
	"Lady: \nIm pretty sure my dad\nis losing it. \n...At least he seems happy.",
	"Traveler: \nIm so tired of standing \nby this stupid well.",
	"Traveler: \nfortunately I had a coin to \nmake a wish in that well. \nOnce its granted Im \ngonna see the world!",
	"Well: \nI heard that guy call me \nstupid. Theres no way \nim granting his wish now!",
	"Well: \ngood luck seeing the world \nwith no path finding ai \nweirdo!",
	"Fish Boy: \nMan, id love to be \na fish... \nClothes are the worst.",
	"Fish Boy: \nFish get to be naked \nany time they want. \nStupid lucky fish...",
	"Broke Millennial: \nMan, my landlord raised \nthe rent again. I dont know \nhow Im gonna keep up \nwith it.",
	"Broke Millennial: \nWell, I dont know how \nill do it next month but",
	"Broke Millennial: \nDoes it seem odd to you that \nI have to pay my rent in \ndog food?",
	"Trust Fund Dog: \n*bark* *bark* \nMan, home ownership is great.",
	"Trust Fund Dog: \nAll these lazy \nmillenials complain that \nits impossible.",
	"Trust Fund Dog: \nIn my expericne \nall it takes is a little \nhard work...",
	"Little Kid: \n{sec}...\n{sec+1}...\n{sec+2}...",
	"Little Kid: \nIm gonna be \nthe first kid \never to count to 256!",
	"Little Kid: \nOh good! Youre here!",
	"Little Kid: \nI wanted someone to\nwitness me finishing\nthe count!",
	"Little Kid: \n253...\n254...\n255...",
	"Little Kid: \n1!!!\nWait what???",
	"Little Kid: \nIm so confused...",
	"The sign says, \nwelcome to... \nuh red... blue town... \nI guess...",
	"it seems pretty unsure of \nitself",
	"Old Man: \nThe prophecy has been \nfulfilled!",
	"Old Man: \nHere young warrior take this\nmagic sword and face\nyour destitny!",
	"Trust Fund Dog: \nand a small loan of a \nmillion dollars from your dad.",
	"Broke Millennial: \nI managed to scrape together \n18,000 kibble in time to \npay rent.",
	"Traveler: \nI just feel like Ive been \nrooted to this spot forever \nyou know.",
}

func fillTemplate(in string) string {
	ret := strings.ReplaceAll(in, "{sec}", fmt.Sprintf("%d", g.Frames()/30))
	ret = strings.ReplaceAll(ret, "{sec+1}", fmt.Sprintf("%d", (g.Frames()/30)+1))
	ret = strings.ReplaceAll(ret, "{sec+2}", fmt.Sprintf("%d", (g.Frames()/30)+2))
	return ret
}
