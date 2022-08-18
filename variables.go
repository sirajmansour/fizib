package website

import (
	"errors"
)

const dataBasePath = "C://UsersData/"
const htmlPagesPath = "C://HtmlPages/MUXLESS/"
const defaultPPPath = "C://DefaultPP/"

// Port is the port which the server is listening on, and it's set to ":8080" as default and can be changed.
var Port = ":8080"

var (
	errNoRegisteredOrDeletedAccount = errors.New("no account registered with this gmail? Make one?")
	errLostOrDeletedData            = errors.New("lost or deleted data? Visit your actions log?")
	errAlreadyOccupiedGmail         = errors.New("an account already registered with the given gmail? Use another one or check your gmail links?")
)

var (
	firstName      Name
	lastName       Name
	nation         Name
	gendre         Gendre
	gmail          Gmail
	password       Password
	verifyPassword Password
	birthday       Birthday
)
