package widgets

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// ControlView makes btns to choose the speed of simulation

type ControlView struct {
	simu *simulation.Simulation

	fontFace font.Face

	isPlusPressed  bool
	isMinusPressed bool

	plusBtnX, plusBtnY          int
	plusBtnWidth, plusBtnHeight int

	minusBtnX, minusBtnY          int
	minusBtnWidth, minusBtnHeight int
}

func NewControlView(simu *simulation.Simulation) *ControlView {
	// Load a font for the button text
	fontData, err := utils.LoadFile("assets/font/money.otf")
	if err != nil {
		log.Fatal("Failed to load font file:", err)
	}
	parsedFont, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal("Failed to parse font:", err)
	}
	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("Failed to create font face:", err)
	}

	margin := 25
	btnWidth := 30
	btnHeight := 30
	xStart := utils.MapWidth - margin - btnWidth*2
	yStart := utils.MapHeigt - margin - btnHeight

	return &ControlView{
		simu: simu,

		fontFace: face,

		plusBtnX:      xStart,
		plusBtnY:      yStart,
		plusBtnWidth:  btnWidth,
		plusBtnHeight: btnHeight,

		minusBtnX:      xStart + btnWidth + 10,
		minusBtnY:      yStart,
		minusBtnWidth:  btnWidth,
		minusBtnHeight: btnHeight,
	}
}

func (cv *ControlView) Update() {
	mouseX, mouseY := ebiten.CursorPosition()
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if mouseX >= cv.plusBtnX && mouseX <= cv.plusBtnX+cv.plusBtnWidth &&
		mouseY >= cv.plusBtnY && mouseY <= cv.plusBtnY+cv.plusBtnHeight {
		cv.isPlusPressed = mousePressed
		if mousePressed {
			cv.simu.Parameters().IncreaseSpeed()
		}
	} else {
		cv.isPlusPressed = false
	}

	if mouseX >= cv.minusBtnX && mouseX <= cv.minusBtnX+cv.minusBtnWidth &&
		mouseY >= cv.minusBtnY && mouseY <= cv.minusBtnY+cv.minusBtnHeight {
		cv.isMinusPressed = mousePressed
		if mousePressed {
			cv.simu.Parameters().DecreaseSpeed()
		}
	} else {
		cv.isMinusPressed = false
	}
}

func (cv *ControlView) Render(screen *ebiten.Image) {
	// SHow btns + and - to modify speed
	plusOpacity := uint8(255)
	if cv.isPlusPressed {
		plusOpacity = 128
	}
	plusRect := ebiten.NewImage(cv.plusBtnWidth, cv.plusBtnHeight)
	plusRect.Fill(color.RGBA{0, 0, 0, plusOpacity})
	plusOp := &ebiten.DrawImageOptions{}
	plusOp.GeoM.Translate(float64(cv.plusBtnX), float64(cv.plusBtnY))
	screen.DrawImage(plusRect, plusOp)

	plusLabelX := cv.plusBtnX + 12
	plusLabelY := cv.plusBtnY + 20
	text.Draw(screen, "+", cv.fontFace, plusLabelX, plusLabelY, color.White)

	minusOpacity := uint8(255)
	if cv.isMinusPressed {
		minusOpacity = 128
	}
	minusRect := ebiten.NewImage(cv.minusBtnWidth, cv.minusBtnHeight)
	minusRect.Fill(color.RGBA{0, 0, 0, minusOpacity})
	minusOp := &ebiten.DrawImageOptions{}
	minusOp.GeoM.Translate(float64(cv.minusBtnX), float64(cv.minusBtnY))
	screen.DrawImage(minusRect, minusOp)

	minusLabelX := cv.minusBtnX + 12
	minusLabelY := cv.minusBtnY + 20
	text.Draw(screen, "-", cv.fontFace, minusLabelX, minusLabelY, color.White)

	// Display current steps/s next to the buttons to speed up the simulation
	speed := int(cv.simu.Parameters().TimeSpeed())
	if speed > 1 {
		speedStr := fmt.Sprintf("%d steps/s", int(cv.simu.Parameters().TimeSpeed()))
		speedX := cv.minusBtnX + cv.minusBtnWidth - 66
		speedY := cv.minusBtnY - 15
		text.Draw(screen, speedStr, cv.fontFace, speedX, speedY, color.White)
	} else {
		speedStr := fmt.Sprintf("%d step/s", int(cv.simu.Parameters().TimeSpeed()))
		speedX := cv.minusBtnX + cv.minusBtnWidth - 61
		speedY := cv.minusBtnY - 15
		text.Draw(screen, speedStr, cv.fontFace, speedX, speedY, color.White)
	}

}
