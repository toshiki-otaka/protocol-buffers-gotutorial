// Code generated by protoc-gen-counter. DO NOT EDIT.

package tutorialpb

func (x *Person) Count() int {
	return len(x.Phones)
}

func (x *AddressBook) Count() int {
	return len(x.People)
}
