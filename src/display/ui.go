package display

import (
	"fmt"
	"net/url"
	"github.com/zserge/webview"
)

type Display struct {
	Content  string
	Width    int
	Height   int
	webview  webview.WebView
}

func (d *Display) SetContent(content string) {
	d.Content = content
}

func (d *Display) showInput(w webview.WebView, data string) {
	w.Eval(fmt.Sprintf("show('%s');", d.Content))
}

func (d *Display) Close() {
	d.webview.Exit()
}

func CreateDisplay(width, height int) *Display {	
	d := Display{
		"",
		800,
		800,
		nil}
		
	page := `
	<!doctype html><html>
	<script language=javascript>
		function show(value){
			document.getElementById('app').innerHTML = value;
		}
		function update(){
			window.external.invoke();
			setInterval(update, 1);
		}
	</script>
	<body onload="update()">
	<div id="app">Init</div>
	</body>
	</html>`
	d.webview = webview.New(webview.Settings{
		Title:                  "Game",
		URL:                    `data:text/html,` + url.PathEscape(page),
		Width:                  width,
		Height:                 height,
		ExternalInvokeCallback: d.showInput})
	return &d
}

func (d *Display) Show() {
	d.webview.Run()
}
