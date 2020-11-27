package tiled

type ObjectGroup struct {
	Name       string      `xml:"name,attr"`
	Color      string      `xml:"color,attr"`
	OffSetX    float64     `xml:"offsetx,attr"`
	OffSetY    float64     `xml:"offsety,attr"`
	Opacity    float32     `xml:"opacity,attr"`
	Visible    bool        `xml:"visible,attr"`
	Properties []*Property `xml:"properties>property"`
	//Objects    []*Object   `xml:"object"`
	//
	//// parentMap is the map which contains this object
	//parentMap *Map
	title *Tile
}
