package main

import "athena_service/cmd"

func main() {
	root := cmd.Root()

	if err := root.Execute(); err != nil {
		panic(err)
	}

}
