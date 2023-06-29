package model

type Transfer struct {
	ID uint
	ControllerTransfer
	Controller *Controller
}
