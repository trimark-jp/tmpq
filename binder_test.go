package tmpq

import (
	"fmt"
	"testing"
)

func TestBindParam(t *testing.T) {
	b := NewBinder(`
        SELECT * FROM ${table}
        WHERE (
            name LIKE ${name}
            OR
            display_name LIKE ${name}
        ) AND email LIKE ${email}

    `)

	b.Bind("table", "users")
	b.Bind("name", "admin")
	b.Bind("email", "%@trimark.jp")
	t.Log(b.QueryArgs())
}

func ExampleBinder() {
	b := NewBinder(`
    SELECT * FROM users
    WHERE (
        name LIKE ${name}
        OR
        display_name LIKE ${name}
    ) AND email LIKE ${email}`)

	b.Bind("name", "admin")
	b.Bind("email", "%@trimark.jp")

	query, args := b.QueryArgs()

	fmt.Println(query)
	fmt.Println()
	fmt.Println(args)

	// Output:
	// SELECT * FROM users
	//     WHERE (
	//         name LIKE $1
	//         OR
	//         display_name LIKE $1
	//     ) AND email LIKE $2
	//
	// [admin %@trimark.jp]
}
