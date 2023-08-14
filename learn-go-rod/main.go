package main

import "github.com/go-rod/rod"

func main() {
	page := rod.New().MustConnect().MustPage("https://www.youtube.com/playlist?list=PLVpaYHJMoSnqKTDtbGYUzgpyEdXctTYK3")
	page.MustElement("#page-manager > ytd-browse > ytd-playlist-header-renderer > div > div.immersive-header-content.style-scope.ytd-playlist-header-renderer > div.thumbnail-and-metadata-wrapper.style-scope.ytd-playlist-header-renderer > div > div.play-menu.spaced-row.wide-screen-form.style-scope.ytd-playlist-header-renderer > ytd-button-renderer.play-button.style-scope.ytd-playlist-header-renderer > yt-button-shape > a > yt-touch-feedback-shape > div > div.yt-spec-touch-feedback-shape__fill").MustClick()
	c := make(chan struct{})
	<-c
}
