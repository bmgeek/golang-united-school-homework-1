package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	helloWorldMap := emoji.Sprint("Hello :world_map:!")
	return helloWorldMap
}
