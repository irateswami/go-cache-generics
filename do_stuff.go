package genericscache

import "time"

func main() {

	// you still need to declare some kind of type when creating a container of
	// a generic type, so it needs an object to instantiate.

	// this feels like a fancy code generator and I'm 100% okay with that
	// or *stenciling*, if that's what you're into

	// NOTE that if you have only one object in the Storable interface,
	// instantiation is not needed and the compiler will automatically infer
	// the object that it's creating
	object2Cache := New[Object2](5*time.Minute, 10*time.Minute)
	object1Cache := New[Object1](5*time.Minute, 10*time.Minute)
	_ = object2Cache
	_ = object1Cache

}
