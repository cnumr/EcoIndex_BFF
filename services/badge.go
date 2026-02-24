package services

import (
	"io"
	"net/http"

	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/helper"
)

func GetColor(grade string) string {
	if color, ok := gradeColorMap[grade]; ok {
		return color
	}

	return "lightgrey"
}

func GetBadgeSvg(grade string, theme string) string {
	url := config.ENV.CDNUrl + "@" + config.ENV.BadgeVersion + "/assets/svg/" + theme + "/" + grade + ".svg"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	badgeSvg, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return helper.MinifyString("image/svg+xml", string(badgeSvg))
}
