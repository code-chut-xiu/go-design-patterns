package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != carWheels {
		t.Errorf("Expeted number of wheels is %[1]d, actual is %[2]d", carWheels, car.Wheels)
	}

	if car.Seats != carSeats {
		t.Errorf("Expeted number of seats is %[1]d, actual is %[2]d", carSeats, car.Seats)
	}

	if car.Structure != carStructure {
		t.Errorf("Expeted structure is %[1]s, actual is %[2]s", carStructure, car.Structure)
	}
}
