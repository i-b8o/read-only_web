package not_found_controller

type viewModelState struct {
	Title       string
	Description string
}

func getState() viewModelState {
	return viewModelState{Title: NOT_FOUND_TITLE, Description: ""}
}

const NOT_FOUND_TITLE = `404 страница не найдена`
