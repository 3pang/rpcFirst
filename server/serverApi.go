package server

type Item struct {
	Title string
	Body  string
}

//type API int

type API struct {
	Database []Item
}

func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = a.Database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range a.Database {
		if val.Title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	a.Database = append(a.Database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(item Item, reply *Item) error {
	var changed Item

	for idx, val := range a.Database {
		if val.Title == item.Title {
			a.Database[idx] = Item{item.Title, item.Body}
			changed = a.Database[idx]
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range a.Database {
		if val.Title == item.Title && val.Body == item.Body {
			a.Database = append(a.Database[:idx], a.Database[idx+1:]...)
			del = item
			break
		}
	}

	*reply = del
	return nil
}
