package sonos

// Household contains details, groups and players in a household
type Household struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Players []struct {
		Name         string   `json:"name"`
		WebsocketURL string   `json:"websocketUrl"`
		DeviceIds    []string `json:"deviceIds"`
		ID           string   `json:"id"`
		Icon         string   `json:"icon"`
	} `json:"players"`
	Groups []struct {
		PlayerIds     []string `json:"playerIds"`
		PlaybackState string   `json:"playbackState"`
		CoordinatorID string   `json:"coordinatorId"`
		ID            string   `json:"id"`
		Name          string   `json:"name"`
	} `json:"groups"`
}
