package main

func main() {
	var warriors = Warriors{}
	warriors = append(warriors, NewMRSatan())
	warriors = append(warriors, NewGoku())

	executeWithoutIPS(warriors)
}
