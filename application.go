package main


func (this *Application) Init () {
	this.Updater = new(Updater)
	this.Updater.Parent = this
	this.SetupConfig()
}

func (this *Application ) SetupConfig () {
	this.Config = *new(Config)
	this.Config.Load()
	this.Config.Parent = this
}
