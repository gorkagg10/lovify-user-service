package profile

type MusicProviderData struct {
	connected  bool
	topTracks  []Track
	topArtists []Artist
}

func (m *MusicProviderData) Connected() bool {
	return m.connected
}

type Track struct {
	name    string
	album   string
	artists []string
}

type Artist struct {
	name  string
	genre string
}
