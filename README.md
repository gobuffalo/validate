# github.com/s3rj1k/validator

This package provides a framework for writing validations for Go applications. 
It does provide you with few validators, but if you need others you can easly build them.

## Installation

```bash
$ go get github.com/s3rj1k/validator
```

## Usage

Using validate is pretty easy, just define `Validate` object for a struct.

Here is a pretty simple example:

```go
package main

import (
        "fmt"

        v "github.com/s3rj1k/validator"
)

type User struct {
        Name  string
        Email string
}

func (u *User) Validate(errors *v.Errors) {
        if u.Name == "" {
                errors.Add("name", "Name must not be blank!")
        }
        if u.Email == "" {
                errors.Add("email", "Email must not be blank!")
        }
}

func main() {
        err := v.Validate(&User{})
        if err != nil {
                fmt.Println(err)
        }
}
```

In the previous example a single `Validator` for the `User` struct was used. 
To really get the benefit of using validator, one can create custom validators.

```go
package main

import (
        "fmt"
        "strings"

        v "github.com/s3rj1k/validator"
)

type User struct {
        Name  string
        Email string
}

type PresenceValidator struct {
        Field string
        Value string
}

func (v *PresenceValidator) Validate(errors *v.Errors) {
        if v.Value == "" {
                errors.Add(strings.ToLower(v.Field), fmt.Sprintf("%s must not be blank!", v.Field))
        }
}

func main() {
        u := User{Name: "", Email: ""}
        err := v.Validate(&PresenceValidator{"Email", u.Email}, &PresenceValidator{"Name", u.Name})
        if err != nil {
                fmt.Println(err)
        }
}
```

## Built-in Validators

To make it even simpler, this package has a children package with some nice built-in validators.

```go
package main

import (
        "fmt"

        "github.com/s3rj1k/validator"
        "github.com/s3rj1k/validator/validators"
)

type User struct {
        Name  string
        Email string
}


func main() {
        u := User{Name: "", Email: ""}
        err := validator.Validate(
                &validators.EmailIsPresent{Name: "Email", Field: u.Email, Message: "Mail is not in the right format."},
                &validators.StringIsPresent{Field: u.Name, Name: "Name"},
        )
        if err != nil {
                fmt.Println(err)
        }
}
```

All fields are required for each validators, except Message (every validator has a default error message).

## Todo
* update docs with fork details
* port more validators from `https://github.com/asaskevich/govalidator`
* add list of build-in validators
* add contributors list
* GoDoc?
