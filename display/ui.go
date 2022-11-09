package display

import (
	"fmt"
	"github.com/webview/webview"
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

func (d *Display) showInput(w webview.WebView) func() {
	return func() {
		w.Eval(fmt.Sprintf("show('%s');", d.Content))
	}
}

func (d *Display) Close() {
	d.webview.Terminate()
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
			window.showInput();
			setInterval(update, 1);
		}
	</script>
	<body onload="update()">
	<div id="app">Init</div>
	</body>
	</html>`
	d.webview = webview.New(true)
	d.webview.SetTitle("Game")
	d.webview.SetSize(width, height, webview.HintFixed)
	d.webview.SetHtml(page)
	d.webview.Bind("showInput", d.showInput(d.webview))
	return &d
}

func (d *Display) Show() {
	d.webview.Run()
}
