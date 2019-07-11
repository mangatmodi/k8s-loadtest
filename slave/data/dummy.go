//Package data holds all the functions required to deal with data layer
package data

//GetData returns empty interface, implement it further
func GetData() (interface{}, error) {
	return []interface{}{Empty{}}, nil
}

type Empty struct{}
