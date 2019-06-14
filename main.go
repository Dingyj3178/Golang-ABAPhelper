package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("ABAP Helper")
	window.SetMinimumSize2(800, 800)
	layout := widgets.NewQVBoxLayout()

	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	label := widgets.NewQLabel2("copy you definition here", nil, 0)
	layout.AddWidget(label, 0, 0)

	textEdit := widgets.NewQPlainTextEdit2("", nil)
	layout.AddWidget(textEdit, 0, 0)

	button := widgets.NewQPushButton2("Format Definition", nil)
	layout.AddWidget(button, 0, 0)

	button.ConnectClicked(func(checked bool) {
		textEdit.SetPlainText(beautifyDef(textEdit.ToPlainText()))
	})

	window.SetCentralWidget(mainWidget)
	window.Show()

	app.Exec()
}
