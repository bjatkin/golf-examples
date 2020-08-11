package main

type interaction struct {
	location    vec2
	lines       []int
	currentLine int
	cache       string
}

func (i *interaction) reset() {
	i.cache = ""
	i.currentLine = 0
}

func (i *interaction) next() bool {
	i.cache = ""
	i.currentLine++
	return i.currentLine < len(i.lines)
}

func (i *interaction) line() string {
	if i.cache == "" {
		line := allLines[i.lines[i.currentLine]]
		i.cache = fillTemplate(line)
	}
	return i.cache
}

var guardInter = interaction{
	lines: []int{0, 1, 2, 3, 4, 5},
}

var oldManInter = interaction{
	lines: []int{6, 7, 8, 9},
}

var ladyInter = interaction{
	lines: []int{11},
}

var wellManInter = interaction{
	lines: []int{12, 13},
}

var talkingWellInter = interaction{
	lines: []int{14, 15},
}

var fishKidInter = interaction{
	lines: []int{16, 17},
}

var rentGuyInter = interaction{
	lines: []int{18},
}

var talkingDogInter = interaction{
	lines: []int{21, 22, 23},
}

var countingKidInter = interaction{
	lines: []int{24, 25},
}
