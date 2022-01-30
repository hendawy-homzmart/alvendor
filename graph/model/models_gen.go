// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Abilities struct {
	ID      string `json:"id"`
	Action  string `json:"action"`
	Subject string `json:"subject"`
}

type Address struct {
	ID           string       `json:"id"`
	FloorNo      int          `json:"floor_no"`
	AppartmentNo int          `json:"appartment_no"`
	LandMark     string       `json:"land_mark"`
	Street       string       `json:"street"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	Country      string       `json:"country"`
	Coordinates  *Coordinates `json:"coordinates"`
}

type Admin struct {
	ID       string    `json:"id"`
	UserInfo *User     `json:"user_info"`
	UserID   string    `json:"user_id"`
	Manages  []*Seller `json:"manages"`
}

type Coordinates struct {
	Lat int `json:"lat"`
	Lng int `json:"lng"`
}

type Customer struct {
	ID        string     `json:"id"`
	UserInfo  *User      `json:"user_info"`
	UserID    string     `json:"user_id"`
	Addresses []*Address `json:"addresses"`
}

type Role struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Ability []*Abilities `json:"ability"`
}

type Seller struct {
	ID        string     `json:"id"`
	UserInfo  *User      `json:"user_info"`
	UserID    string     `json:"user_id"`
	Addresses []*Address `json:"addresses"`
	Shops     []*Shops   `json:"shops"`
}

type Shops struct {
	ID      string        `json:"id"`
	Name    *BusinessName `json:"name"`
	Address *Address      `json:"address"`
}

type User struct {
	ID        string     `json:"id"`
	FirstName *FirstName `json:"first_name"`
	LastName  *LastName  `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Mobile    *int       `json:"mobile"`
	Role      *Role      `json:"role"`
	IsActive  *bool      `json:"isActive"`
}

type BusinessName struct {
	Ar *string `json:"ar"`
	En *string `json:"en"`
}

type FirstName struct {
	Ar *string `json:"ar"`
	En *string `json:"en"`
}

type LastName struct {
	Ar *string `json:"ar"`
	En *string `json:"en"`
}

type NewAbilities struct {
	ID      string `json:"id"`
	Action  string `json:"action"`
	Subject string `json:"subject"`
}

type NewAddress struct {
	ID           string          `json:"id"`
	FloorNo      int             `json:"floor_no"`
	AppartmentNo int             `json:"appartment_no"`
	LandMark     string          `json:"land_mark"`
	Street       string          `json:"street"`
	City         string          `json:"city"`
	State        string          `json:"state"`
	Country      string          `json:"country"`
	Coordinates  *NewCoordinates `json:"coordinates"`
}

type NewAdmin struct {
	ID       string     `json:"id"`
	UserInfo *NewUser   `json:"user_info"`
	UserID   string     `json:"user_id"`
	Manages  *NewSeller `json:"manages"`
}

type NewBusinessName struct {
	Ar *string `json:"ar"`
	En *string `json:"en"`
}

type NewCoordinates struct {
	Lat int `json:"lat"`
	Lng int `json:"lng"`
}

type NewCustomer struct {
	ID        string      `json:"id"`
	UserInfo  *NewUser    `json:"user_info"`
	UserID    string      `json:"user_id"`
	Addresses *NewAddress `json:"addresses"`
}

type NewFirstName struct {
	Ar *string `json:"ar"`
	En *string `json:"en"`
}

type NewLastName struct {
	Ar *string `json:"ar"`
	En *string `json:"en"`
}

type NewRole struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Ability *NewAbilities `json:"ability"`
}

type NewSeller struct {
	ID        string      `json:"id"`
	UserInfo  *NewUser    `json:"user_info"`
	UserID    string      `json:"user_id"`
	Addresses *NewAddress `json:"addresses"`
	Shops     *NewShops   `json:"shops"`
}

type NewShops struct {
	ID      string           `json:"id"`
	Name    *NewBusinessName `json:"name"`
	Address *NewAddress      `json:"address"`
}

type NewUser struct {
	ID        string        `json:"id"`
	FirstName *NewFirstName `json:"first_name"`
	LastName  *NewLastName  `json:"last_name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Mobile    *int          `json:"mobile"`
	Role      *NewRole      `json:"role"`
	IsActive  *bool         `json:"isActive"`
}
