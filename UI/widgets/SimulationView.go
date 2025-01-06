package widgets

import (
	"fmt"
	"image/color"
	"log"
	"math"
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

// Simulation panel shows the simulation, the agents, and combine the infopanel and controlview

type SimulationView struct {
	*widget.Container

	fontTitle font.Face
	fontDate  font.Face
	fontMoney font.Face

	background *ebiten.Image
	objects    []views.IViewAgent

	controlView   *ControlView
	infoPanelView *InfoPanelView
	events        chan Events

	startDate     time.Time
	step          int
	franceMoney   float32
	spainMoney    float32
	englandMoney  float32
	selectedAgent simulation.Agent
}

func NewSimulationView(simu *simulation.Simulation, events chan Events, infoPanel *InfoPanelView) *SimulationView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(utils.GameWidth, utils.GameHeight),
		),
	)

	controlView := NewControlView(simu)

	fond, err := utils.NewImageFromPath(utils.MapImage)
	if err != nil {
		log.Println(err)
	}

	viewAgents, _, err1 := views.ConvertAllAgentsToView([]simulation.Agent{})
	if err1 != nil {
		log.Println(err1)
	}

	fontData, err := utils.LoadFile("assets/font/money.otf")
	if err != nil {
		log.Println("Failed to load font :", err)
	}

	parsedOT, err := opentype.Parse(fontData)
	if err != nil {
		log.Println("Failed to parse font :", err)
	}

	fontTitle, err := opentype.NewFace(parsedOT, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Println(err)
	}

	fontDate, err := opentype.NewFace(parsedOT, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Println(err)
	}

	fontMoney, err := opentype.NewFace(parsedOT, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Println(err)
	}

	startDate := time.Date(1716, time.January, 1, 0, 0, 0, 0, time.UTC)

	return &SimulationView{
		Container:     rootContainer,
		background:    fond,
		objects:       viewAgents,
		events:        events,
		fontTitle:     fontTitle,
		fontDate:      fontDate,
		fontMoney:     fontMoney,
		startDate:     startDate,
		step:          0,
		controlView:   controlView,
		infoPanelView: infoPanel,
	}
}

func isMouseOver(mouseX, mouseY int, obj views.IViewAgent) bool {
	pos := obj.Position()
	spriteWidth := 16
	spriteHeight := 16

	if mouseX >= pos.X && mouseX <= (pos.X+spriteWidth) &&
		mouseY >= pos.Y && mouseY <= (pos.Y+spriteHeight) {
		return true
	}
	return false
}

func (s *SimulationView) Update(simu *simulation.Simulation) {
	s.controlView.Update()

	viewAgents, agentMapping, err := views.ConvertAllAgentsToView(simu.Agents())
	if err != nil {
		log.Println(err)
	}

	s.objects = viewAgents
	s.step = simu.Parameters().Step()
	s.franceMoney = simu.Environment().Market().France()
	s.spainMoney = simu.Environment().Market().Spain()
	s.englandMoney = simu.Environment().Market().England()

	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, obj := range s.objects {
			if isMouseOver(x, y, obj) {
				originalAgent := agentMapping[obj]

				s.selectedAgent = originalAgent

				s.infoPanelView.UpdateInfo(originalAgent)
				break
			}
		}
	}

	s.infoPanelView.UpdateInfo(s.selectedAgent)

}

func (s *SimulationView) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(s.background, op)

	for _, obj := range s.objects {
		obj.Draw(screen)
	}

	currentDate := s.startDate.AddDate(0, 0, s.step)
	dateStr := fmt.Sprintf("%s %d", currentDate.Month(), currentDate.Year())

	text.Draw(screen, "Piraaaaaates", s.fontTitle, 25, 70, color.White)

	text.Draw(screen, dateStr, s.fontDate, 25, 95, color.White)

	moneyValues := []int{
		int(math.Floor(float64(s.franceMoney))),
		int(math.Floor(float64(s.spainMoney))),
		int(math.Floor(float64(s.englandMoney))),
	}
	flagPaths := []string{
		"assets/flags/france.png",
		"assets/flags/spain.png",
		"assets/flags/england.png",
	}

	coinImg, err := utils.NewImageFromPath("assets/coin.png")
	if err != nil {
		log.Println(err)
	}

	flagWidth, flagHeight := 40.0, 30.0
	coinWidth, coinHeight := 20.0, 20.0

	x := utils.GameWidth - 100
	y := 20

	for i, val := range moneyValues {
		flag, err := utils.NewImageFromPath(flagPaths[i])
		if err != nil {
			log.Println(err)
		}

		flagOp := &ebiten.DrawImageOptions{}
		scaleFlagX := flagWidth / float64(flag.Bounds().Dx())
		scaleFlagY := flagHeight / float64(flag.Bounds().Dy())
		flagOp.GeoM.Scale(scaleFlagX, scaleFlagY)
		flagOp.GeoM.Translate(float64(x-50), float64(y))
		screen.DrawImage(flag, flagOp)

		moneyStr := fmt.Sprintf("%d", val)

		text.Draw(screen, moneyStr, s.fontMoney, x, y+22, color.White)

		bound := text.BoundString(s.fontMoney, moneyStr)
		textWidth := bound.Dx()

		coinOp := &ebiten.DrawImageOptions{}
		scaleCoinX := coinWidth / float64(coinImg.Bounds().Dx())
		scaleCoinY := coinHeight / float64(coinImg.Bounds().Dy())
		coinOp.GeoM.Scale(scaleCoinX, scaleCoinY)
		coinOp.GeoM.Translate(
			float64(x+textWidth+5),
			float64(y+6),
		)
		screen.DrawImage(coinImg, coinOp)

		y += 40
	}

	s.controlView.Render(screen)

	panelX := 25
	panelY := utils.GameHeight - 295
	s.infoPanelView.Render(screen, panelX, panelY)
}
