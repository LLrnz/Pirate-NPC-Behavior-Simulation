package widgets

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type InfoPanelView struct {
	background      *ebiten.Image
	fontFace        font.Face
	agentInfo       string
	nationalityFlag *ebiten.Image
	nationality     string

	heartImg *ebiten.Image
	coinImg  *ebiten.Image
}

func NewInfoPanelView() *InfoPanelView {
	background, err := utils.NewImageFromPath("assets/chest_open.png")
	if err != nil {
		log.Println("Error loading background, using plain color:", err)
		background = ebiten.NewImage(100, 50)
		background.Fill(color.NRGBA{0, 100, 100, 255})
	} else {
		originalBounds := background.Bounds()
		newWidth := int(float64(originalBounds.Dx()) * 1.8)
		newHeight := int(float64(originalBounds.Dy()) * 1.8)
		resizedBackground := ebiten.NewImage(newWidth, newHeight)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.8, 1.8)
		resizedBackground.DrawImage(background, op)
		background = resizedBackground
	}

	// Load the font
	fontPath := "assets/font/money.otf"
	fontData, err := utils.LoadFile(fontPath)
	if err != nil {
		log.Fatal("Failed to load font file:", err)
	}

	ttFont, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal("Failed to parse font:", err)
	}

	fontFace, err := opentype.NewFace(ttFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("Failed to create font face:", err)
	}

	// Load heart and coin images
	heartImg, err := utils.NewImageFromPath("assets/heart.png")
	if err != nil {
		log.Println("Error loading heart image, using fallback:", err)
		heartImg = ebiten.NewImage(1, 1)
		heartImg.Fill(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	}

	coinImg, err := utils.NewImageFromPath("assets/coin.png")
	if err != nil {
		log.Println("Error loading coin image, using fallback:", err)
		coinImg = ebiten.NewImage(1, 1)
		coinImg.Fill(color.NRGBA{R: 255, G: 255, B: 0, A: 255})
	}

	return &InfoPanelView{
		background: background,
		fontFace:   fontFace,
		agentInfo:  "Cliquez sur un agent",

		heartImg: heartImg,
		coinImg:  coinImg,

		nationalityFlag: nil,
		nationality:     "",
	}
}

func (i *InfoPanelView) UpdateInfo(agent simulation.Agent) {
	if agent == nil {
		i.agentInfo = "Cliquez sur un agent"
		i.nationalityFlag = nil
		i.nationality = ""
		return
	}

	name := agent.Id()
	health := agent.Health()

	switch agent.AgtType() {
	case "Kraken":
		i.nationalityFlag = nil
		i.nationality = ""

		infoBuilder := strings.Builder{}
		infoBuilder.WriteString(fmt.Sprintf("Name: %s\n", name))
		infoBuilder.WriteString(fmt.Sprintf("Health: %d\n", health))
		i.agentInfo = infoBuilder.String()

	case "Ship":
		ship, ok := agent.(*simulation.Ship)
		if !ok {
			i.nationalityFlag = nil
			i.nationality = ""

			i.agentInfo = fmt.Sprintf("Name: %s\nHealth: %d\n", name, health)
			return
		}

		nationality := ship.Nationality()

		if nationality != "" && nationality != "Without" {
			flagPath := getFlagPath(nationality)
			flagImg, err := utils.NewImageFromPath(flagPath)
			if err != nil {
				log.Println("Error loading nationality flag:", err)
				i.nationalityFlag = nil
			} else {
				i.nationalityFlag = flagImg
			}
			i.nationality = nationality
		} else {
			i.nationalityFlag = nil
			i.nationality = ""
		}

		agentType := ship.ShipType()

		infoBuilder := strings.Builder{}
		infoBuilder.WriteString(fmt.Sprintf("%s\n", name))
		infoBuilder.WriteString(fmt.Sprintf("Vie : %d\n", health))

		switch agentType {
		case "Merchant", "Pirate", "Privateer":
			money := ship.Money()

			infoBuilder.WriteString(fmt.Sprintf("Argent : %d\n", int(money)))

			resources := ship.Ressources()

			infoBuilder.WriteString("Resources:\n")

			for resource, quantity := range resources {
				if quantity > 0 {
					infoBuilder.WriteString(fmt.Sprintf("    - %s: %d\n", resource, quantity))
				}
			}
		default:
		}

		i.agentInfo = infoBuilder.String()

	default:
		i.nationalityFlag = nil
		i.nationality = ""
		i.agentInfo = fmt.Sprintf("%s\nVie : %d\n", name, health)
	}
}

func (i *InfoPanelView) Render(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(i.background, op)

	lines := splitLines(i.agentInfo)
	lineHeight := 20
	textX := x + 45
	textY := y + 95

	if i.nationalityFlag != nil {
		flagOp := &ebiten.DrawImageOptions{}
		flagWidth, flagHeight := 40.0, 30.0
		scaleX := flagWidth / float64(i.nationalityFlag.Bounds().Dx())
		scaleY := flagHeight / float64(i.nationalityFlag.Bounds().Dy())

		flagOp.GeoM.Scale(scaleX, scaleY)
		flagOp.GeoM.Translate(float64(textX), float64(y+80))
		screen.DrawImage(i.nationalityFlag, flagOp)

		textY = y + 135
	}

	for idx, line := range lines {
		drawX := textX
		drawY := textY + (idx * lineHeight)

		text.Draw(screen, line, i.fontFace, drawX, drawY, color.White)

		if strings.HasPrefix(line, "Vie :") {
			iconOp := &ebiten.DrawImageOptions{}
			scale := 0.1
			iconOp.GeoM.Scale(scale, scale)
			iconOp.GeoM.Translate(float64(drawX+75), float64(drawY-14))
			screen.DrawImage(i.heartImg, iconOp)
		} else if strings.HasPrefix(line, "Argent :") {
			iconOp := &ebiten.DrawImageOptions{}
			scale := 0.1
			iconOp.GeoM.Scale(scale, scale)
			iconOp.GeoM.Translate(float64(drawX+75), float64(drawY-14))
			screen.DrawImage(i.coinImg, iconOp)
		}
	}
}

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func getFlagPath(nationality string) string {
	switch nationality {
	case "French":
		return "assets/flags/france.png"
	case "Spanish":
		return "assets/flags/spain.png"
	case "English":
		return "assets/flags/england.png"
	default:
		return "assets/flags/blank.png"
	}
}
