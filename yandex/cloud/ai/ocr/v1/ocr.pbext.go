// Code generated by protoc-gen-goext. DO NOT EDIT.

package ocr

func (m *Polygon) SetVertices(v []*Vertex) {
	m.Vertices = v
}

func (m *Vertex) SetX(v int64) {
	m.X = v
}

func (m *Vertex) SetY(v int64) {
	m.Y = v
}

func (m *TextAnnotation) SetWidth(v int64) {
	m.Width = v
}

func (m *TextAnnotation) SetHeight(v int64) {
	m.Height = v
}

func (m *TextAnnotation) SetBlocks(v []*Block) {
	m.Blocks = v
}

func (m *TextAnnotation) SetEntities(v []*Entity) {
	m.Entities = v
}

func (m *TextAnnotation) SetTables(v []*Table) {
	m.Tables = v
}

func (m *TextAnnotation) SetFullText(v string) {
	m.FullText = v
}

func (m *Entity) SetName(v string) {
	m.Name = v
}

func (m *Entity) SetText(v string) {
	m.Text = v
}

func (m *Block) SetBoundingBox(v *Polygon) {
	m.BoundingBox = v
}

func (m *Block) SetLines(v []*Line) {
	m.Lines = v
}

func (m *Block) SetLanguages(v []*Block_DetectedLanguage) {
	m.Languages = v
}

func (m *Block) SetTextSegments(v []*TextSegments) {
	m.TextSegments = v
}

func (m *Block_DetectedLanguage) SetLanguageCode(v string) {
	m.LanguageCode = v
}

func (m *Line) SetBoundingBox(v *Polygon) {
	m.BoundingBox = v
}

func (m *Line) SetText(v string) {
	m.Text = v
}

func (m *Line) SetWords(v []*Word) {
	m.Words = v
}

func (m *Line) SetTextSegments(v []*TextSegments) {
	m.TextSegments = v
}

func (m *Word) SetBoundingBox(v *Polygon) {
	m.BoundingBox = v
}

func (m *Word) SetText(v string) {
	m.Text = v
}

func (m *Word) SetEntityIndex(v int64) {
	m.EntityIndex = v
}

func (m *Word) SetTextSegments(v []*TextSegments) {
	m.TextSegments = v
}

func (m *TextSegments) SetStartIndex(v int64) {
	m.StartIndex = v
}

func (m *TextSegments) SetLength(v int64) {
	m.Length = v
}

func (m *Table) SetBoundingBox(v *Polygon) {
	m.BoundingBox = v
}

func (m *Table) SetRowCount(v int64) {
	m.RowCount = v
}

func (m *Table) SetColumnCount(v int64) {
	m.ColumnCount = v
}

func (m *Table) SetCells(v []*TableCell) {
	m.Cells = v
}

func (m *TableCell) SetBoundingBox(v *Polygon) {
	m.BoundingBox = v
}

func (m *TableCell) SetRowIndex(v int64) {
	m.RowIndex = v
}

func (m *TableCell) SetColumnIndex(v int64) {
	m.ColumnIndex = v
}

func (m *TableCell) SetColumnSpan(v int64) {
	m.ColumnSpan = v
}

func (m *TableCell) SetRowSpan(v int64) {
	m.RowSpan = v
}

func (m *TableCell) SetText(v string) {
	m.Text = v
}

func (m *TableCell) SetTextSegments(v []*TextSegments) {
	m.TextSegments = v
}
