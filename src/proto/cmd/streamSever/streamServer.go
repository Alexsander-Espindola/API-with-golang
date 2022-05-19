package streamServer

import (
	"fmt"
	"log"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func StreamClient(stream pb.PostUser_TestStreamClient, client pb.PostUserClient) {
	if err := ui.Init(); err != nil {
		log.Fatalf("Couldn't init UI: %v", err)
	}
	defer ui.Close()

	lc := widgets.NewPlot()
	ui.Render(lc)

	uiEvents := ui.PollEvents()

	for {
		fmt.Println("Context done")
		select {
		case e := <-uiEvents: // Listen for Keyboard
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-stream.Context().Done():
			fmt.Println("All done")
			break
		}
	}
}
