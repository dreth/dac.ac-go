package components

import "fmt"

// STANDALONE TEXT SIZES
var RtFooterClasses string = "text-xxs xs:text-xs sm:text-sm md:text-sm lg:text-base xl:text-base"
var RtSClasses string = "text-xxs xs:text-xs sm:text-base md:text-lg xl:text-xl"
var RtClasses string = "text-xs xs:text-sm sm:text-base md:text-lg xl:text-xl"
var RtHClasses string = "text-sm xs:text-base sm:text-lg md:text-xl xl:text-2xl"

// TEXT SIZES
func RtFooter(classes string) string {
	return fmt.Sprintf("%s %s", RtFooterClasses, classes)
}

func RtS(classes string) string {
	return fmt.Sprintf("%s %s", RtSClasses, classes)
}

func Rt(classes string) string {
	return fmt.Sprintf("%s %s", RtClasses, classes)
}

func RtH(classes string) string {
	return fmt.Sprintf("%s %s", RtHClasses, classes)
}

// CLASSES FOR ARTICLES
// H1
var H1Classes string = "text-lg xs:text-xl sm:text-2xl md:text-3xl lg:text-4xl xl:text-5xl text-h1-h3-article-color text-center pb-2"

// H2
var H2Classes string = "text-base xs:text-lg sm:text-xl md:text-2xl lg:text-3xl xl:text-4xl text-h2-article-color pb-4"

// H3
var H3Classes string = "text-sm xs:text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl text-h1-h3-article-color py-3 pt-2"

// H4
var H4Classes string = "text-sm xs:text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl text-h1-h3-article-color py-3 pt-2"

// H5
var H5Classes string = "text-sm xs:text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl text-h1-h3-article-color py-3 pt-2"

// A
var AClasses string = Rt("text-links hover:text-links-hover")

// P, Span
var PClasses string = Rt("text-text pb-2 pt-2")

// UL
var UlClasses string = Rt("list-disc pb-2 pl-4 pr-4 pt-2")

// LI, OL
var LiClasses string = Rt("pb-2")
var OlClasses string = Rt("list-decimal pb-2 pl-2 pr-2 pt-2")

// Pre
var PreClasses string = RtFooter("text-light bg-code p-4 rounded-md py-5 my-4 bg-code-block overflow-x-auto border border-quote border-l-1 border-t-0 border-b-0 border-r-0")

// HR
var HrClasses string = Rt("my-4 sm:my-6")

// Blockquote
var BlockquoteClasses string = Rt("text-text bg-quote p-4 py-3 my-4 bg-quote-block word-wrap-break italic font-medium rounded-md border border-quote border-l-4 border-t-0 border-b-0 border-r-0 bg-quote-block")

// IMG
var ImgClassesOutsideFigure string = "w-10/12 sm:w-3/5 lg:w-1/2 h-auto rounded-md mx-auto items-center justify-center py-4 sm:py-5"
var ImgClassesSmallerOutsideFigure string = "w-5/12 sm:w-4/12 lg:w-3/12 h-auto rounded-md mx-auto items-center justify-center py-4 sm:py-5"
var ImgClassesInsideFigure string = "w-10/12 sm:w-3/5 lg:w-1/2 h-auto rounded-md mx-auto items-center justify-center"
var ImgClassesSmallerInsideFigure string = "w-5/12 sm:w-4/12 lg:w-3/12 h-auto rounded-md mx-auto items-center justify-center"

// FIGURE
var FigureClasses string = "flex flex-col items-center justify-center py-4 sm:py-5"

// FIGCAPTION
var FigcaptionClasses string = RtFooterClasses
