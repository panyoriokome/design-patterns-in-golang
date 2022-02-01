package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
	// interfaceを実装すればいい
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("この操作はサポートしていません")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("この操作はサポートしていません")
}

// ISP
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

// PrinterとScannerの機能を持つ
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {

}
func (p Photocopier) Print(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func main() {
	ofp := OldFashionedPrinter{}
	ofp.Scan()
}
