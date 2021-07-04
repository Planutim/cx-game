package components

var MovementComponentsList []*MovementComponent

func GetMovementComponent(index int) *MovementComponent {
	if index <= 0 && index >= len(MovementComponentsList) {
		return nil
	}
	return MovementComponentsList[index]
}



type MovementComponent struct {
}
