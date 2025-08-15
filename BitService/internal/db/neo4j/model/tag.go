package model

type Tag int

func (g Tag) String() string {
	if g < Rock || g > JPop {
		return "Unknown"
	}
	return tagNames[g]
}

const (
	Rock Tag = iota
	Pop
	HipHop
	Rap
	RandB
	Jazz
	Blues
	Classical
	Electronic
	Dance
	House
	Techno
	Trance
	Dubstep
	Reggae
	Ska
	Metal
	Punk
	Folk
	Country
	Soul
	Funk
	Disco
	Gospel
	Opera
	Ambient
	Soundtrack
	World
	Latin
	KPop
	JPop
)

var tagNames = [...]string{
	"Rock",
	"Pop",
	"HipHop",
	"Rap",
	"R&B",
	"Jazz",
	"Blues",
	"Classical",
	"Electronic",
	"Dance",
	"House",
	"Techno",
	"Trance",
	"Dubstep",
	"Reggae",
	"Ska",
	"Metal",
	"Punk",
	"Folk",
	"Country",
	"Soul",
	"Funk",
	"Disco",
	"Gospel",
	"Opera",
	"Ambient",
	"Soundtrack",
	"World",
	"Latin",
	"KPop",
	"JPop",
}
