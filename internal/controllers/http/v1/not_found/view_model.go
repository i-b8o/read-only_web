package not_found_controller

type viewModelState struct {
	Title *string
}

func getState() viewModelState {
	title := `404 страница не найдена`
	return viewModelState{Title: &title}
}
