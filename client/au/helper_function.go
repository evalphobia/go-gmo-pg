package au

func trimCommodityUnderLimit(commodity string) string {
	const maxCommodityLen = 24
	commodityRune := []rune(commodity)
	if len(commodityRune) > maxCommodityLen {
		commodity = string(commodityRune[:maxCommodityLen])
	}
	return commodity
}
