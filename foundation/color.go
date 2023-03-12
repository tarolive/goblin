package foundation

type BrandColor struct {
	Primary   string
	Secondary string
}

type BackgroundColor struct {
	Primary       string
	Secondary     string
	SecondaryGray string
	Tertiary      string
	TertiaryGray  string
}

type TextColor struct {
	Primary   string
	Secondary string
}

type LinkColor struct {
	Primary             string
	PrimaryBackground   string
	Secondary           string
	SecondaryBackground string
}

type AlertColor struct {
	Primary           string
	PrimaryBackground string
	Secondary         string
}

func NewBrandColor(primary string, secondary string) BrandColor {

	var (
		brandColor = BrandColor{
			Primary:   primary,
			Secondary: secondary,
		}
	)

	return brandColor
}

func NewBackgroundColor(primary string, secondary string, secondaryGray string, tertiary string, tertiaryGray string) BackgroundColor {

	var (
		backgroundColor = BackgroundColor{
			Primary:       primary,
			Secondary:     secondary,
			SecondaryGray: secondaryGray,
			Tertiary:      tertiary,
			TertiaryGray:  tertiaryGray,
		}
	)

	return backgroundColor
}

func NewTextColor(primary string, secondary string) TextColor {

	var (
		textColor = TextColor{
			Primary:   primary,
			Secondary: secondary,
		}
	)

	return textColor
}

func NewLinkColor(primary string, primaryBackground string, secondary string, secondaryBackground string) LinkColor {

	var (
		linkColor = LinkColor{
			Primary:             primary,
			PrimaryBackground:   primaryBackground,
			Secondary:           secondary,
			SecondaryBackground: secondaryBackground,
		}
	)

	return linkColor
}

func NewAlertColor(primary string, primaryBackground string, secondary string) AlertColor {

	var (
		alertColor = AlertColor{
			Primary:           primary,
			PrimaryBackground: primaryBackground,
			Secondary:         secondary,
		}
	)

	return alertColor
}
