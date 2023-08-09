package in

import "time"

type Profile struct {
	ID             int64  `json:"id"`
	WantedJobTitle string `json:"wantedJobTitle"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Address        string `json:"address"`
	PostalCode     int64  `json:"postalCode"`
	DrivingLicense string `json:"drivingLicense"`
	Nationality    string `json:"nationality"`
	PlaceOfBirth   string `json:"placeOfBirth"`
	DateOfBirth    string `json:"dateOfBirth"`
	Deleted        string `json:"deleted"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (input *Profile) ValidateInsertProfile() (err error) {

	return
}

func (input *Profile) ValidateUpdateProfile() (err error) {

	return
}

func (input *Profile) ValidateGetProfile() (err error) {

	return
}
