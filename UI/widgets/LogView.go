package widgets

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

//The LogView is the view at the right of screen that can be used to add agents to the simulation

type LogView struct {
	*widget.Container

	fontText font.Face

	pirateImg    *ebiten.Image
	krakenImg    *ebiten.Image
	merchantImg  *ebiten.Image
	armyImg      *ebiten.Image
	privateerImg *ebiten.Image

	simu *simulation.Simulation
}

// NewLogView creates a LogView
func NewLogView(simu *simulation.Simulation) *LogView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(color.RGBA{100, 0, 100, 255}),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(130, 100),
		),
	)

	fontData, err := utils.LoadFile("assets/font/money.otf")
	if err != nil {
		log.Println("No font file :", err)
	}

	var fontText font.Face
	if fontData != nil {
		parsedOT, err := opentype.Parse(fontData)
		if err != nil {
			log.Println("Problem parse font :", err)
		} else {
			fontText, err = opentype.NewFace(parsedOT, &opentype.FaceOptions{
				Size:    16,
				DPI:     72,
				Hinting: font.HintingFull,
			})
			if err != nil {
				log.Println("Problem create font :", err)
			}
		}
	}

	// Sprites to show for each agent
	pirateImg := loadImage("assets/pirateship.png")
	krakenImg := loadImage("assets/kraken.png")
	merchantImg := loadImage("assets/merchantship.png")
	armyImg := loadImage("assets/armyship.png")
	privateerImg := loadImage("assets/privateership.png")

	// Create and return the LogView struct
	return &LogView{
		Container:    rootContainer,
		fontText:     fontText,
		pirateImg:    pirateImg,
		krakenImg:    krakenImg,
		merchantImg:  merchantImg,
		armyImg:      armyImg,
		privateerImg: privateerImg,
		simu:         simu,
	}
}

// mustLoadImage is a helper that loads an Ebiten image from path or returns a fallback.
func loadImage(path string) *ebiten.Image {
	img, err := utils.NewImageFromPath(path)
	if err != nil {
		log.Printf("Problem load img %s : %v", path, err)
		// Create fake image so it doesn't crash
		fallback := ebiten.NewImage(2, 2)
		fallback.Fill(color.RGBA{255, 0, 255, 255})
		return fallback
	}
	return img
}

func (l *LogView) Render(screen *ebiten.Image) {
	// Agent counts
	agentCount := l.simu.CountAgents()

	// Part 1: No nationality (Pirate et Kraken)
	pirateCount := agentCount["Pirate"]
	krakenCount := agentCount["Kraken"]

	// Part 2: France
	merchantCountFrench := agentCount["Merchant[French]"]
	privateerCountFrench := agentCount["Privateer[French]"]
	armyCountFrench := agentCount["Army[French]"]

	// Part 3: England
	merchantCountEnglish := agentCount["Merchant[English]"]
	privateerCountEnglish := agentCount["Privateer[English]"]
	armyCountEnglish := agentCount["Army[English]"]

	// Part 4: Spain
	merchantCountSpanish := agentCount["Merchant[Spanish]"]
	privateerCountSpanish := agentCount["Privateer[Spanish]"]
	armyCountSpanish := agentCount["Army[Spanish]"]

	offsetX := float64(l.GetWidget().Rect.Min.X)
	offsetY := float64(l.GetWidget().Rect.Min.Y) + 10
	sectionSpacing := 50.0
	cellWidth := 50.0
	cellHeight := 90.0

	// Part 1
	y := offsetY + 30
	text.Draw(screen, "Sans nationalité", l.fontText, int(offsetX), int(y), color.White)
	y += 20
	l.drawAgentCell(screen, l.pirateImg, pirateCount, offsetX, y, func() {
		l.simu.AddPirate()
	})
	l.drawAgentCell(screen, l.krakenImg, krakenCount, offsetX+cellWidth, y, func() {
		l.simu.AddKraken()
	})

	// Part 2
	y += cellHeight + sectionSpacing
	text.Draw(screen, "France", l.fontText, int(offsetX), int(y), color.White)
	y += 20
	l.drawAgentCell(screen, l.merchantImg, merchantCountFrench, offsetX, y, func() {
		l.simu.AddMerchant(constants.France)
	})
	l.drawAgentCell(screen, l.privateerImg, privateerCountFrench, offsetX+cellWidth, y, func() {
		l.simu.AddPrivateer(constants.France)
	})
	l.drawAgentCell(screen, l.armyImg, armyCountFrench, offsetX+2*cellWidth, y, func() {
		l.simu.AddArmy(constants.France)
	})

	// Part 3
	y += cellHeight + sectionSpacing
	text.Draw(screen, "Angleterre", l.fontText, int(offsetX), int(y), color.White)
	y += 20
	l.drawAgentCell(screen, l.merchantImg, merchantCountEnglish, offsetX, y, func() {
		l.simu.AddMerchant(constants.England)
	})
	l.drawAgentCell(screen, l.privateerImg, privateerCountEnglish, offsetX+cellWidth, y, func() {
		l.simu.AddPrivateer(constants.England)
	})
	l.drawAgentCell(screen, l.armyImg, armyCountEnglish, offsetX+2*cellWidth, y, func() {
		l.simu.AddArmy(constants.England)
	})

	// Part 4
	y += cellHeight + sectionSpacing
	text.Draw(screen, "Espagne", l.fontText, int(offsetX), int(y), color.White)
	y += 20
	l.drawAgentCell(screen, l.merchantImg, merchantCountSpanish, offsetX, y, func() {
		l.simu.AddMerchant(constants.Spain)
	})
	l.drawAgentCell(screen, l.privateerImg, privateerCountSpanish, offsetX+cellWidth, y, func() {
		l.simu.AddPrivateer(constants.Spain)
	})
	l.drawAgentCell(screen, l.armyImg, armyCountSpanish, offsetX+2*cellWidth, y, func() {
		l.simu.AddArmy(constants.Spain)
	})
}

func (l *LogView) drawAgentCell(
	screen *ebiten.Image,
	agentImg *ebiten.Image,
	count int,
	x, y float64,
	addAgent func(),
) {
	// Image of agent
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(agentImg, op)

	// Number of agents of this type
	labelStr := fmt.Sprintf("%d", count)
	labelX := int(x + 2)
	labelY := int(y + 40)
	text.Draw(screen, labelStr, l.fontText, labelX, labelY, color.White)

	// Btn to add agent of this type
	plusW, plusH := 20.0, 20.0
	plusX := float64(labelX)
	plusY := float64(labelY + 10)
	plusOp := &ebiten.DrawImageOptions{}
	plusOp.GeoM.Translate(plusX, plusY)
	plusImg := ebiten.NewImage(int(plusW), int(plusH))
	plusImg.Fill(color.RGBA{0, 128, 0, 255})
	screen.DrawImage(plusImg, plusOp)
	plusTextX := int(plusX + 6)
	plusTextY := int(plusY + 15)
	text.Draw(screen, "+", l.fontText, plusTextX, plusTextY, color.White)

	// If click, add agent
	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if float64(mx) >= plusX && float64(mx) <= plusX+plusW &&
			float64(my) >= plusY && float64(my) <= plusY+plusH {
			addAgent()
		}
	}
}
