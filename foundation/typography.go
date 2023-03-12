package foundation

type HeadlineTypography struct {
	XS  string
	SM  string
	MD  string
	LG  string
	XL  string
	XXL string
}

type BlockquoteTypography struct {
	SM string
	LG string
}

type TitleTypography struct {
	Card   string
	Layout string
}

func NewHeadlineTypography(xs string, sm string, md string, lg string, xl string, xxl string) HeadlineTypography {

	var (
		headlineTypography = HeadlineTypography{
			XS:  xs,
			SM:  sm,
			MD:  md,
			LG:  lg,
			XL:  xl,
			XXL: xxl,
		}
	)

	return headlineTypography
}

func NewBlockquoteTypography(sm string, lg string) BlockquoteTypography {

	var (
		blockquoteTypography = BlockquoteTypography{
			SM: sm,
			LG: lg,
		}
	)

	return blockquoteTypography
}

func NewTitleTypography(card string, layout string) TitleTypography {

	var (
		titleTypography = TitleTypography{
			Card:   card,
			Layout: layout,
		}
	)

	return titleTypography
}
