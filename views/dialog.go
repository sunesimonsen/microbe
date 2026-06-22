package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DialogView() Node {
	dialogText := "Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex. Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce sagittis in est a consequat."

	return docpage(
		HGroup(
			H1(Text("Dialog")),
			P(Text("A temporary, overlapping user interface window that sits on top of a main application screen. It interrupts the user to demand attention, display critical information, or require a specific decision before they can return to the main application.")),
		),
		NewExample(
			"Example",
			Nodes(
				Class("grid"),
				Button(
					Class("outline"),
					Attr("command", "show-modal"),
					Attr("commandfor", "example-dialog"),
					Text("Open dialog"),
				),
				Dialog(
					ID("example-dialog"),
					Attr("closedby", "any"),
					Header(
						Text("Header"),
						Button(
							Rel("prev"),
							Aria("label", "Close"),
							Attr("commandfor", "example-dialog"),
							Attr("command", "close"),
							TabIndex("1"),
						),
					),
					Section(
						P(Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla nulla arcu, pellentesque quis sollicitudin at, egestas et leo. Sed placerat vel odio id pharetra. Nunc euismod eget erat id ultricies. Vivamus a lacus lacus. Vivamus malesuada justo nulla, ut fringilla nisi pellentesque vel. Morbi ligula lacus, lacinia eget volutpat vitae, vestibulum ac massa. Interdum et malesuada fames ac ante ipsum primis in faucibus. Praesent viverra, nisl vitae pulvinar aliquam, ex enim scelerisque leo, a pellentesque sem ante ac velit. Phasellus euismod mauris vulputate ultrices pretium. Nulla aliquet justo id arcu tincidunt gravida.")),
						P(Text("Quisque a magna rutrum, placerat lorem sit amet, finibus elit. Vivamus a facilisis tortor. Ut eget ultricies ante. Nullam convallis feugiat felis. Aenean gravida congue neque, at dapibus neque blandit sed. Donec venenatis ullamcorper enim, ut dictum tortor ultrices in. In hac habitasse platea dictumst. Aenean fringilla condimentum orci, sit amet egestas elit fringilla sed. Phasellus sed dui vel risus bibendum rhoncus vitae quis ex. Fusce ut ornare ex. Proin ipsum ex, varius non elementum et, hendrerit id lacus. Aliquam erat volutpat. Morbi id dictum libero. Sed mattis sem quis urna egestas, vel dapibus leo tincidunt. Vivamus suscipit, augue sed volutpat lobortis, tellus dui fringilla risus, efficitur facilisis arcu quam et est. Mauris pulvinar nec nisi varius iaculis.")),
						P(Text("Etiam tempor in ligula eu euismod. Ut in leo euismod, scelerisque metus eu, vehicula quam. Quisque lacinia nulla non gravida facilisis. Nam porttitor diam eget quam malesuada, sed vestibulum purus molestie. Vivamus ultricies, tortor ac sagittis feugiat, dui ex volutpat eros, in consectetur est arcu at velit. Praesent vel dui vitae ex varius ornare vel nec nisl. Quisque egestas, sem eget condimentum venenatis, sapien lorem vestibulum mi, cursus mollis tortor lorem sit amet ipsum. Curabitur eu odio quis leo ullamcorper gravida. Donec eget augue vitae orci tincidunt luctus. Morbi auctor dictum interdum.")),
						P(Text("Pellentesque sit amet eros at turpis luctus venenatis. Sed ut magna et mauris pulvinar molestie. Sed varius massa sit amet tempor rhoncus. In lobortis pellentesque leo eu elementum. Cras faucibus leo et magna pharetra scelerisque. Nam ut tortor nunc. Ut in enim rutrum lacus volutpat finibus. Donec vel ipsum varius, placerat nulla in, vestibulum leo. Curabitur tincidunt egestas suscipit. Vivamus ac lobortis magna.")),
						P(Text("Sed enim magna, lacinia sed massa vel, vestibulum gravida lectus. Curabitur eu purus tortor. Sed vitae tortor volutpat, placerat urna vel, facilisis ligula. Nunc sodales faucibus turpis, quis ornare sem aliquet vitae. Mauris eu pulvinar metus. Sed sit amet ex nisl. Maecenas vulputate ipsum sit amet lorem accumsan, ut egestas tellus vulputate. Morbi a luctus ex.")),
						P(Text("Duis tristique accumsan libero, ac efficitur massa pretium et. Suspendisse id justo in massa condimentum eleifend. Sed neque lorem, tempor a purus et, tincidunt placerat enim. Nullam eget velit lobortis, aliquet tortor a, dictum orci. Aenean id consequat mi. Proin semper blandit est, sit amet aliquet tellus rhoncus vitae. Ut facilisis tempor maximus. Cras laoreet vitae eros in suscipit. Vestibulum laoreet id tortor ut pretium. Maecenas condimentum suscipit turpis sit amet accumsan. Vivamus quis odio sed ante pulvinar pharetra in id neque.")),
						P(Text("Suspendisse tristique tellus ac nibh placerat rhoncus. Morbi posuere fringilla nibh, pellentesque gravida felis aliquet a. Quisque ac nisi eget velit imperdiet congue id sit amet eros. Integer non varius nunc. Phasellus est urna, fermentum quis semper non, viverra sed orci. Vestibulum convallis in dolor luctus consequat. Sed ultrices, nunc ac hendrerit pulvinar, massa nunc porttitor augue, a eleifend mauris odio et felis. Suspendisse eu lacinia purus. Curabitur consectetur consectetur augue eget aliquam. Etiam dictum libero sit amet dui consectetur lobortis. Duis sit amet augue rutrum, efficitur mauris nec, rutrum elit.")),
						P(Text("Vivamus rutrum tortor erat, eget posuere ex mattis ac. Sed vitae nisi eu nibh tempor maximus id eu mauris. Nulla facilisi. Sed nec porta magna. Pellentesque elementum convallis cursus. Donec vitae elit eget tellus dignissim ullamcorper eu eu odio. Curabitur dapibus vehicula libero vitae semper. Curabitur malesuada neque ut nibh egestas, ut sodales nibh bibendum.")),
						P(Text("Suspendisse venenatis semper sapien sed gravida. In lobortis, sem id lacinia bibendum, ipsum turpis sollicitudin velit, at commodo tortor lectus sit amet nisl. Pellentesque facilisis at metus a placerat. Aenean vitae diam interdum velit faucibus dapibus sagittis ut lectus. Ut nec urna ut sem ultrices eleifend eu eu odio. Cras et lorem at turpis varius sodales ac a odio. Etiam ut imperdiet purus. Quisque pellentesque elit nisi, sit amet consectetur felis ultrices at. Pellentesque vel sodales sem. Donec massa erat, luctus sit amet justo vel, porta mollis nulla. Nulla dignissim tortor et porta ultricies. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean leo tellus, dictum vel varius eu, malesuada eu nunc. Phasellus non massa neque. Nunc in quam et libero porttitor ultrices eget quis ante. Curabitur in diam eget nibh blandit viverra.")),
						P(Text("Donec a varius velit, quis porta eros. Nullam auctor ultricies elit, ut sodales arcu aliquet vehicula. Phasellus quis eros sit amet sapien hendrerit maximus. In interdum metus lorem, id tincidunt ligula elementum vitae. Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis vehicula, lorem non vulputate accumsan, neque justo tincidunt augue, non congue massa neque et velit. Praesent cursus ut velit eget efficitur. Pellentesque finibus accumsan lacus, id dapibus erat feugiat fringilla. Maecenas vestibulum sapien vitae lectus semper laoreet. Maecenas varius vel sapien ac convallis. Nam dapibus eget lectus nec cursus. Praesent dignissim augue vel ipsum vestibulum, ac sodales nibh tincidunt. Aliquam eget pretium tortor. Fusce eleifend nisi et mi pulvinar ornare.")),
					),
					Footer(
						Class("actions"),
						Button(
							Class("outline"),
							Attr("command", "close"),
							Attr("commandfor", "example-dialog"),
							Text("Decline"),
						),
						Button(
							Class("solid"),
							Attr("command", "close"),
							Attr("commandfor", "example-dialog"),
							Text("Accept"),
						),
					),
				),
			),
		),
		NewExample(
			"Small",
			Nodes(
				Class("grid"),
				Button(
					Class("outline"),
					Attr("command", "show-modal"),
					Attr("commandfor", "example-small-dialog"),
					Text("Open dialog"),
				),
				Dialog(
					ID("example-small-dialog"),
					Class("small"),
					Attr("closedby", "any"),
					Header(
						Text("Header"),
						Button(
							Rel("prev"),
							Aria("label", "Close"),
							Attr("commandfor", "example-small-dialog"),
							Attr("command", "close"),
							TabIndex("1"),
						),
					),
					Section(P(Text(dialogText))),
					Footer(
						Class("actions"),
						Button(
							Class("outline"),
							Attr("command", "close"),
							Attr("commandfor", "example-small-dialog"),
							Text("Decline"),
						),
						Button(
							Class("solid"),
							Attr("command", "close"),
							Attr("commandfor", "example-small-dialog"),
							Text("Accept"),
						),
					),
				),
			),
		),
	)
}
