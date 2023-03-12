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

func NewRedBrandColor() BrandColor {

	var (
		redBrandColor = NewBrandColor("#EE0000", "#D40000")
	)

	return redBrandColor
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

func NewLightBackgroundColor() BackgroundColor {

	var (
		lightBackgroundColor = NewBackgroundColor("", "", "", "", "")
	)

	return lightBackgroundColor
}
