package model

import "encoding/json"

type Profile struct {
	Name string
	Age int
	Height string
	Weight string
	Income string
	Marriger string
	Xingzuo string
	// Occupation string
	// Hourse string
	// Car string
	// Hokou string
	// Education string
	// Gender string
}

func FormJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)

	return profile, err
}
