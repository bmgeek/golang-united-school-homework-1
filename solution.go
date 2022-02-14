package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	helloWorldMap := fmt.Sprintf("%s", "Hello :world_map:!")
	emoji.Sprint(helloWorldMap)
	return helloWorldMap
}
