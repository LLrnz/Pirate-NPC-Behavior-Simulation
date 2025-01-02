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

	fond, err := utils.NewImageFromPath(utils.MapImage)
	if err != nil {
		log.Fatal(err)
	}

	viewAgents, err1 := views.ConvertAllAgentsToView([]simulation.Agent{})
	if err1 != nil {
		log.Fatal(err1)
	}

	fontPath := "assets/font/money.otf"
	fontData, err := utils.LoadFile(fontPath)
	if err != nil {
		log.Fatal(err)
	}
	ttFont, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

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

	startDate := time.Date(1716, time.January, 1, 0, 0, 0, 0, time.UTC)

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
	s.step = simu.Parameters().Step()

	s.franceMoney = simu.Environment().Market().France()
	s.spainMoney = simu.Environment().Market().Spain()
	s.englandMoney = simu.Environment().Market().England()
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
		log.Fatal(err)
	}

	flagWidth, flagHeight := 40.0, 30.0
	coinWidth, coinHeight := 20.0, 20.0

	x := utils.GameWidth - 225
	y := 20

	for i, val := range moneyValues {
		flag, err := utils.NewImageFromPath(flagPaths[i])
		if err != nil {
			log.Fatal(err)
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
}
