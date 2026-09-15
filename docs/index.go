package docs

var Index Categories

func init() {
	Index = Categories{
		NewCategory(
			"Getting started",
			AboutPage,
			ModulesPage,
		),
		NewCategory(
			"Theming",
			ColorsPage,
			PalettePage,
		),
		NewCategory(
			"Layout",
			MarginPage,
			SpacingPage,
			AccordionPage,
			CardPage,
			TabsPage,
		),
		NewCategory(
			"Content",
			AvatarPage,
			CalloutPage,
			TypographyPage,
			ListPage,
			TablePage,
		),
		NewCategory(
			"Navigation",
			AnchorPage,
			BreadcrumbPage,
			NavlistPage,
		),
		NewCategory(
			"Forms",
			ButtonPage,
			CheckboxPage,
			InputPage,
			TextareaPage,
			RadioPage,
			RangePage,
			SelectPage,
			SwitchPage,
		),
		NewCategory(
			"Loading",
			ProgressPage,
			SkeletonPage,
		),
		NewCategory(
			"Popups",
			DialogPage,
			MenuPage,
			PopoverPage,
			NotificationPage,
		),
		NewCategory(
			"Data",
			PaginationPage,
			TagPage,
		),
	}
}
