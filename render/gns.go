package render

type GNS struct {
	AutoClose           bool        `json:"auto_close"`
	AutoOpen            bool        `json:"auto_open"`
	AutoStart           bool        `json:"auto_start"`
	DrawingGridSize     int         `json:"drawing_grid_size"`
	GridSize            int         `json:"grid_size"`
	Name                string      `json:"name"`
	ProjectID           string      `json:"project_id"`
	Revision            int         `json:"revision"`
	SceneHeight         int         `json:"scene_height"`
	SceneWidth          int         `json:"scene_width"`
	ShowGrid            bool        `json:"show_grid"`
	ShowInterfaceLabels bool        `json:"show_interface_labels"`
	ShowLayers          bool        `json:"show_layers"`
	SnapToGrid          bool        `json:"snap_to_grid"`
	Supplier            interface{} `json:"supplier"`
	Topology            struct {
		Computes []interface{} `json:"computes"`
		Drawings []interface{} `json:"drawings"`
		Links    []struct {
			Filters struct {
			} `json:"filters"`
			LinkID    string `json:"link_id"`
			LinkStyle struct {
			} `json:"link_style"`
			Nodes []struct {
				AdapterNumber int `json:"adapter_number"`
				Label         struct {
					Rotation int    `json:"rotation"`
					Style    string `json:"style"`
					Text     string `json:"text"`
					X        int    `json:"x"`
					Y        int    `json:"y"`
				} `json:"label"`
				NodeID     string `json:"node_id"`
				PortNumber int    `json:"port_number"`
			} `json:"nodes"`
			Suspend bool `json:"suspend"`
		} `json:"links"`
		Nodes []struct {
			ComputeID        string        `json:"compute_id"`
			Console          int           `json:"console"`
			ConsoleAutoStart bool          `json:"console_auto_start"`
			ConsoleType      string        `json:"console_type"`
			CustomAdapters   []interface{} `json:"custom_adapters"`
			FirstPortName    interface{}   `json:"first_port_name"`
			Height           int           `json:"height"`
			Label            struct {
				Rotation int    `json:"rotation"`
				Style    string `json:"style"`
				Text     string `json:"text"`
				X        int    `json:"x"`
				Y        int    `json:"y"`
			} `json:"label"`
			Locked          bool   `json:"locked"`
			Name            string `json:"name"`
			NodeID          string `json:"node_id"`
			NodeType        string `json:"node_type"`
			PortNameFormat  string `json:"port_name_format"`
			PortSegmentSize int    `json:"port_segment_size"`
			Properties      struct {
				PortsMapping []struct {
					Name       string `json:"name"`
					PortNumber int    `json:"port_number"`
					Type       string `json:"type"`
					Vlan       int    `json:"vlan"`
				} `json:"ports_mapping"`
			} `json:"properties"`
			Symbol     string `json:"symbol"`
			TemplateID string `json:"template_id"`
			Width      int    `json:"width"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Z          int    `json:"z"`
		} `json:"nodes"`
	} `json:"topology"`
	Type      string      `json:"type"`
	Variables interface{} `json:"variables"`
	Version   string      `json:"version"`
	Zoom      int         `json:"zoom"`
}
