package saturation

import (
	"crypto/sha256"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"

	"github.com/mrmxf/opentsg-core/config"
	examplejson "github.com/mrmxf/opentsg-widgets/exampleJson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBars(t *testing.T) {
	myImage := image.NewNRGBA64(image.Rect(0, 0, 2330, 600))
	s := saturationJSON{GridLoc: &config.Grid{Alias: "testlocation"}}
	colours := [][]string{{"red", "green", "blue"}, {"red", "blue"}, {"blue"}, {}}
	explanation := []string{"redGreenBlue", "redBlue", "blue", "defualt"}

	for i, c := range colours {
		s.Colours = c
		genErr := s.Generate(myImage)
		examplejson.SaveExampleJson(s, widgetType, explanation[i])

		f, _ := os.Open(fmt.Sprintf("./testdata/ordertest%v.png", i))

		baseVals, _ := png.Decode(f)
		// Assign the colour to the correct type of image NGRBA64 and replace the colour values
		readImage := image.NewNRGBA64(baseVals.Bounds())
		draw.Draw(readImage, readImage.Bounds(), baseVals, image.Point{0, 0}, draw.Over)

		// Make a hash of the pixels of each image
		hnormal := sha256.New()
		htest := sha256.New()
		hnormal.Write(readImage.Pix)
		htest.Write(myImage.Pix)
		// F, _ := os.Create(testF[i] + fmt.Sprintf("%v.png", i))
		// Png.Encode(f, myImage)

		Convey("Checking saturations ramps can be generated for differenent colours", t, func() {
			Convey("Comparing the generated ramp to the base test", func() {
				Convey("No error is returned and the file matches", func() {
					So(genErr, ShouldBeNil)
					So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
				})
			})
		})

	}
}
