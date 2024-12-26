package widgets

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/views"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// SimulationView represents the simulation display, including a map and simulation objects
type SimulationView struct {
	*widget.Container
	background *ebiten.Image
	objects    []views.IViewAgent
	events     chan Events
	fontTitle  font.Face
	fontDate   font.Face
	fontMoney  font.Face
	startDate  time.Time
	step       int

	franceMoney  float32
	spainMoney   float32
	englandMoney float32
}

func NewSimulationView(events chan Events) *SimulationView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(utils.GameWidth, utils.GameHeight),
		),
	)

	// Load the background image
	fond, err := utils.NewImageFromPath(utils.MapImage)
	if err != nil {
		log.Fatal(err)
	}

	// Convert simulation objects to view objects
	viewAgents, err1 := views.ConvertAllAgentsToView([]simulation.Agent{})
	if err1 != nil {
		log.Fatal(err1)
	}

	// Load font
	fontPath := "assets/font/money.otf"
	fontData, err := utils.LoadFile(fontPath) // Implement LoadFile to read font bytes
	if err != nil {
		log.Fatal(err)
	}
	ttFont, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	// Create font faces for different sizes
	fontTitle, err := opentype.NewFace(ttFont, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	fontDate, err := opentype.NewFace(ttFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	fontMoney, err := opentype.NewFace(ttFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Set the start date to March 1, 1660
	startDate := time.Date(1660, time.March, 1, 0, 0, 0, 0, time.UTC)

	return &SimulationView{
		Container:  rootContainer,
		background: fond,
		objects:    viewAgents,
		events:     events,
		fontTitle:  fontTitle,
		fontDate:   fontDate,
		fontMoney:  fontMoney,
		startDate:  startDate,
		step:       0,
	}
}

func (s *SimulationView) Update(simu *simulation.Simulation) {
	viewAgents, err := views.ConvertAllAgentsToView(simu.Agents())
	if err != nil {
		log.Fatal(err)
	}

	s.objects = viewAgents
	s.step = simu.Parameters().Step() // Keep track of the simulation step

	// NEW: Store the current money values from the simulation
	s.franceMoney = simu.Environment().Market().France()
	s.spainMoney = simu.Environment().Market().Spain()
	s.englandMoney = simu.Environment().Market().England()
}

func (s *SimulationView) Render(screen *ebiten.Image) {
	// Draw the background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(s.background, op)

	// Draw all simulation objects
	for _, obj := range s.objects {
		obj.Draw(screen)
	}

	// Calculate the current date based on the simulation step
	currentDate := s.startDate.AddDate(0, 0, s.step)
	dateStr := fmt.Sprintf("%s %d", currentDate.Month(), currentDate.Year())

	// Draw the title "Piraaaaaates" at the top-left corner
	text.Draw(screen, "Piraaaaaates", s.fontTitle, 25, 70, color.White)

	// Draw the date just below the title
	text.Draw(screen, dateStr, s.fontDate, 25, 95, color.White)

	// Draw the money values at the top-right with flags
	moneyStrings := []string{
		fmt.Sprintf("%.2f€", s.franceMoney),
		fmt.Sprintf("%.2f€", s.spainMoney),
		fmt.Sprintf("%.2f€", s.englandMoney),
	}

	flagPaths := []string{
		"assets/flags/france.png",
		"assets/flags/spain.png",
		"assets/flags/england.png",
	}

	// Fixed flag dimensions
	flagWidth, flagHeight := 40.0, 30.0

	x := utils.GameWidth - 225 // Fixed offset for aligning text
	y := 20                    // Starting y position

	for i, moneyStr := range moneyStrings {
		// Load the flag image
		flag, err := utils.NewImageFromPath(flagPaths[i])
		if err != nil {
			log.Fatal(err)
		}

		// Draw the flag image with fixed dimensions
		flagOp := &ebiten.DrawImageOptions{}
		scaleX := flagWidth / float64(flag.Bounds().Dx())
		scaleY := flagHeight / float64(flag.Bounds().Dy())
		flagOp.GeoM.Scale(scaleX, scaleY)
		flagOp.GeoM.Translate(float64(x-50), float64(y)) // Adjust icon position
		screen.DrawImage(flag, flagOp)

		// Draw the money string next to the flag
		text.Draw(screen, moneyStr, s.fontMoney, x, y+22, color.White)
		y += 40 // Increment y position for the next line
	}
}
