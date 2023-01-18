package main

func main() {

	// mLog := log.NewHelper(logger.NewLogger(config))

	app := App()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
