package invoice

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/signintech/gopdf"
)

const (
	quantityColumnOffset = 360
	rateColumnOffset     = 405
	amountColumnOffset   = 480
)

func writeLogo(pdf *gopdf.GoPdf, logo string, from string) {
	if logo != "" {
		width, height := getImageDimension(logo)
		scaledWidth := 100.0
		scaledHeight := float64(height) * scaledWidth / float64(width)
		_ = pdf.Image(logo, pdf.GetX(), pdf.GetY(), &gopdf.Rect{W: scaledWidth, H: scaledHeight})
		pdf.Br(scaledHeight + 24)
	}
	pdf.SetTextColor(55, 55, 55)

	formattedFrom := strings.ReplaceAll(from, `\n`, "\n")
	fromLines := strings.Split(formattedFrom, "\n")

	for i := 0; i < len(fromLines); i++ {
		if i == 0 {
			_ = pdf.SetFont("Inter", "", 12)
			_ = pdf.Cell(nil, fromLines[i])
			pdf.Br(18)
		} else {
			_ = pdf.SetFont("Inter", "", 10)
			_ = pdf.Cell(nil, fromLines[i])
			pdf.Br(15)
		}
	}
	pdf.Br(21)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX(), pdf.GetY(), 260, pdf.GetY())
	pdf.Br(36)
}

func writeTitle(pdf *gopdf.GoPdf, title, id, date string) {
	_ = pdf.SetFont("Inter-Bold", "", 24)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Cell(nil, title)
	pdf.Br(36)
	_ = pdf.SetFont("Inter", "", 12)
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, "#")
	_ = pdf.Cell(nil, id)
	pdf.SetTextColor(150, 150, 150)
	_ = pdf.Cell(nil, "  ·  ")
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, date)
	pdf.Br(48)
}

func writeDueDate(pdf *gopdf.GoPdf, locale Locale, due string) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, locale.DueLabel)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(11)
	pdf.SetX(amountColumnOffset - 15)
	_ = pdf.Cell(nil, due)
	pdf.Br(12)
}

func writeBillTo(pdf *gopdf.GoPdf, locale Locale, to string) {
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont("Inter", "", 9)
	_ = pdf.Cell(nil, locale.ToLabel)
	pdf.Br(18)
	pdf.SetTextColor(75, 75, 75)

	formattedTo := strings.ReplaceAll(to, `\n`, "\n")
	toLines := strings.Split(formattedTo, "\n")

	for i := 0; i < len(toLines); i++ {
		if i == 0 {
			_ = pdf.SetFont("Inter", "", 15)
			_ = pdf.Cell(nil, toLines[i])
			pdf.Br(20)
		} else {
			_ = pdf.SetFont("Inter", "", 10)
			_ = pdf.Cell(nil, toLines[i])
			pdf.Br(15)
		}
	}
	pdf.Br(32)
}

func writeHeaderRow(pdf *gopdf.GoPdf, locale Locale) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, locale.ItemLabel)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, locale.QuantityLabel)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, locale.RateLabel)
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, locale.AmountLabel)
	pdf.Br(24)
}

func writeNotes(pdf *gopdf.GoPdf, locale Locale, notes string) {
	pdf.SetY(650)

	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, locale.NotesLabel)
	pdf.Br(18)
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(0, 0, 0)

	formattedNotes := strings.ReplaceAll(notes, `\n`, "\n")
	notesLines := strings.Split(formattedNotes, "\n")

	for i := 0; i < len(notesLines); i++ {
		_ = pdf.Cell(nil, notesLines[i])
		pdf.Br(15)
	}

	pdf.Br(48)
}

func writeFooter(pdf *gopdf.GoPdf, id string) {
	pdf.SetY(800)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, id)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX()+10, pdf.GetY()+6, 550, pdf.GetY()+6)
	pdf.Br(48)
}

func writeRow(pdf *gopdf.GoPdf, item string, quantity int, rate float64, currency string) {
	_ = pdf.SetFont("Inter", "", 11)
	pdf.SetTextColor(0, 0, 0)

	total := float64(quantity) * rate
	amount := strconv.FormatFloat(total, 'f', 2, 64)

	_ = pdf.Cell(nil, item)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, strconv.Itoa(quantity))
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, getCurrencySymbol(currency)+strconv.FormatFloat(rate, 'f', 2, 64))
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, getCurrencySymbol(currency)+amount)
	pdf.Br(24)
}

func writeTotals(pdf *gopdf.GoPdf, locale Locale, subtotal float64, tax float64, discount float64, currency string) {
	pdf.SetY(650) // TODO: factor out constants like these

	writeTotal(pdf, locale, locale.SubtotalLabel, subtotal, currency)
	if tax > 0 {
		writeTotal(pdf, locale, locale.TaxLabel, tax, currency)
	}
	if discount > 0 {
		writeTotal(pdf, locale, locale.DiscountLabel, discount, currency)
	}
	writeTotal(pdf, locale, locale.TotalLabel, subtotal+tax-discount, currency)
}

func writeTotal(pdf *gopdf.GoPdf, locale Locale, label string, total float64, currency string) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, label)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(12)
	pdf.SetX(amountColumnOffset - 15)
	if label == locale.TotalLabel {
		_ = pdf.SetFont("Inter-Bold", "", 11.5)
	}
	_ = pdf.Cell(nil, getCurrencySymbol(currency)+strconv.FormatFloat(total, 'f', 2, 64))
	pdf.Br(24)
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
