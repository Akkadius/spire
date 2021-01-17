package contexts

import (
	"errors"
	"strconv"
)

type ConnectionCreateContext struct {
	userId            uint
	connectionName    string
	dbName            string
	dbHost            string
	dbPort            string
	dbUsername        string
	dbPassword        string
	contentDbName     string
	contentDbHost     string
	contentDbPort     string
	contentDbUsername string
	contentDbPassword string
	createdFromIp     string
}

func NewConnectionCreateContext(
	userId uint,
	connectionName string,
	dbName string,
	dbHost string,
	dbPort string,
	dbUsername string,
	dbPassword string,
) (*ConnectionCreateContext, error) {
	ctx := &ConnectionCreateContext{
		userId:         userId,
		connectionName: connectionName,
		dbName:         dbName,
		dbHost:         dbHost,
		dbPort:         dbPort,
		dbUsername:     dbUsername,
		dbPassword:     dbPassword,
	}

	return ctx, validate(ctx)
}

func validate(ctx *ConnectionCreateContext) error {
	if len(ctx.connectionName) == 0 {
		return errors.New("database connection name cannot be empty")
	}
	if len(ctx.dbHost) == 0 {
		return errors.New("database hostname cannot be empty")
	}
	if len(ctx.dbName) == 0 {
		return errors.New("database name cannot be empty")
	}
	if len(ctx.dbUsername) == 0 {
		return errors.New("database username cannot be empty")
	}
	if len(ctx.dbPassword) == 0 {
		return errors.New("database password cannot be empty")
	}
	if len(ctx.dbPort) == 0 {
		return errors.New("database port cannot be empty")
	}
	portNum, _ := strconv.ParseInt(ctx.dbPort, 10, 32)
	if portNum <= 0 || portNum >= 65535 {
		return errors.New("database port must be a valid port")
	}

	return nil
}

func (c ConnectionCreateContext) DbPassword() string {
	return c.dbPassword
}

func (c ConnectionCreateContext) DbUsername() string {
	return c.dbUsername
}

func (c ConnectionCreateContext) DbPort() string {
	return c.dbPort
}

func (c ConnectionCreateContext) DbHost() string {
	return c.dbHost
}

func (c ConnectionCreateContext) DbName() string {
	return c.dbName
}

func (c ConnectionCreateContext) ConnectionName() string {
	return c.connectionName
}

func (c ConnectionCreateContext) UserId() uint {
	return c.userId
}

func (c *ConnectionCreateContext) CreatedFromIp() string {
	return c.createdFromIp
}

func (c *ConnectionCreateContext) SetCreatedFromIp(createdFromIp string) {
	c.createdFromIp = createdFromIp
}

func (c *ConnectionCreateContext) ContentDbPassword() string {
	return c.contentDbPassword
}

func (c *ConnectionCreateContext) SetContentDbPassword(contentDbPassword string) {
	c.contentDbPassword = contentDbPassword
}

func (c *ConnectionCreateContext) ContentDbUsername() string {
	return c.contentDbUsername
}

func (c *ConnectionCreateContext) SetContentDbUsername(contentDbUsername string) {
	c.contentDbUsername = contentDbUsername
}

func (c *ConnectionCreateContext) ContentDbPort() string {
	return c.contentDbPort
}

func (c *ConnectionCreateContext) SetContentDbPort(contentDbPort string) {
	c.contentDbPort = contentDbPort
}

func (c *ConnectionCreateContext) ContentDbHost() string {
	return c.contentDbHost
}

func (c *ConnectionCreateContext) SetContentDbHost(contentDbHost string) {
	c.contentDbHost = contentDbHost
}

func (c *ConnectionCreateContext) ContentDbName() string {
	return c.contentDbName
}

func (c *ConnectionCreateContext) SetContentDbName(contentDbName string) {
	c.contentDbName = contentDbName
}
