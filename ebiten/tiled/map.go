package tiled

/*
  __  __
 |  \/  |__ _ _ __
 | |\/| / _` | '_ \
 |_|  |_\__,_| .__/
             |_|
*/

// Map is a TMX file structure representing the map as a whole.
type Map struct {
	Version     string `xml:"title,attr"`
	Orientation string `xml:"orientation,attr"`
	// Width is the number of tiles
	Width int `xml:"width,attr"`
	// Height is the number of tiles
	Height     int `xml:"height,attr"`
	TileWidth  int `xml:"tilewidth,attr"`
	TileHeight int `xml:"tileheight,attr"`
	//Properties   []*Property    `xml:"properties>property"`
	//Tilesets     []*Tileset     `xml:"tileset"`
	//TileLayers   []*TileLayer   `xml:"layer"`
	ObjectGroups []*ObjectGroup `xml:"objectgroup"`
	//Infinite     bool           `xml:"infinite,attr"`
	//ImageLayers  []*ImageLayer  `xml:"imagelayer"`
	//
	//canvas *pixelgl.Canvas
	// dir is the directory the tmx file is located in.  This is used to access images for tilesets via a relative path.
	dir string
}
