package update

type githubAPIJSON struct {
	App struct {
		Description string `json:"description"`
		Version     string `json:"version"`
		Endpoint    string `json:"endpoint"`
	} `json:"App"`
}
