package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	helloWorldMap := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloWorldMap)
	return helloWorldMap
}
