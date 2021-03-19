package main

func _canVisitAllRooms(rooms [][]int, room int, routes map[int]bool) map[int]bool {
	_, ok := routes[room]
	if ok {
		return routes
	}
	routes[room] = true

	nextKeys := rooms[room]
	for _, k := range nextKeys {
		routes = _canVisitAllRooms(rooms, k, routes)
	}
	return routes
}

func canVisitAllRooms(rooms [][]int) bool {
	routes := _canVisitAllRooms(rooms, 0, map[int]bool{})
	return len(routes) == len(rooms)
}
