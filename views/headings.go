package views

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func headingSection(level int) Node {
	return Group([]Node{
		El(fmt.Sprintf("h%d", level), Text(fmt.Sprintf("Heading %d", level))),
		P(Text("Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex. Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce sagittis in est a consequat.")),
		P(Text("Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Pellentesque condimentum sem vel nulla efficitur cursus. Ut accumsan leo eget dictum mollis. Aenean vel dictum nulla. Fusce eget quam sollicitudin, scelerisque ex sed, dignissim mauris. Sed consequat condimentum nibh, eget tristique lectus viverra id. Donec facilisis ultrices lectus ac aliquet. Mauris facilisis ante et vulputate elementum. Interdum et malesuada fames ac ante ipsum primis in faucibus.")),
	})
}

func Headings() Node {
	sections := []Node{}
	for i := 2; i <= 6; i++ {
		sections = append(sections, headingSection(i))
	}

	return Group([]Node{
		HGroup(
			H1(Text("Headings")),
			P(Text("Fill in some info here")),
		),
		Div(
			Attr("role", "document"),
			Group(sections),
		),
	})
}
