package main

// Example exercises the full package.
func Example() {

	main()
	// Output:
	// investors [{i1 100 100} {i2 200 200} {i3 150 150} {i4 50 50}]
	// loans [{l1 100 50} {l2 150 150} {l3 100 50} {l4 20 20} {l5 80 80} {l6 120 100}]
	//
	// loan &{l1 100 50}
	// investor &{i1 100 100}
	// loan < investor
	// loan &{l1 100 0}
	// investor &{i1 100 50}
	//
	// loan &{l2 150 150}
	// investor &{i1 100 50}
	// loan >= investor
	// loan &{l2 150 100}
	// investor &{i1 100 0}
	//
	// loan &{l2 150 100}
	// investor &{i2 200 200}
	// loan < investor
	// loan &{l2 150 0}
	// investor &{i2 200 100}
	//
	// loan &{l3 100 50}
	// investor &{i2 200 100}
	// loan < investor
	// loan &{l3 100 0}
	// investor &{i2 200 50}
	//
	// loan &{l4 20 20}
	// investor &{i2 200 50}
	// loan < investor
	// loan &{l4 20 0}
	// investor &{i2 200 30}
	//
	// loan &{l5 80 80}
	// investor &{i2 200 30}
	// loan >= investor
	// loan &{l5 80 50}
	// investor &{i2 200 0}
	//
	// loan &{l5 80 50}
	// investor &{i3 150 150}
	// loan < investor
	// loan &{l5 80 0}
	// investor &{i3 150 100}
	//
	// loan &{l6 120 100}
	// investor &{i3 150 100}
	// loan >= investor
	// loan &{l6 120 0}
	// investor &{i3 150 0}
	//
	// investors [{i1 100 0} {i2 200 0} {i3 150 0} {i4 50 50}]
	// loans [{l1 100 0} {l2 150 0} {l3 100 0} {l4 20 0} {l5 80 0} {l6 120 0}]
}
