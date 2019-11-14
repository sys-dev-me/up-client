package main


func (this *Module) isEnabled () bool{
	 
	return this.Config.ServiceEnabled 

}

func (this *Module) isNetwork () bool {
	return this.Config.ServiceNetwork
}
