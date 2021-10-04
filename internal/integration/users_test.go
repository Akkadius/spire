package integration

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSimpleUserCreateWithRelationships(t *testing.T) {
	c := bootApp()
	c.DbConnections().SpireMigrate(true)

	var user models.User
	user.UserName = "Akkadius"
	c.DbConnections().SpireDb().Create(&user)

	var connection models.ServerDatabaseConnection
	connection.DbHost = "Something"
	connection.CreatedBy = user.ID
	c.DbConnections().SpireDb().Create(&connection)

	var userConnection models.UserServerDatabaseConnection
	userConnection.UserId = 1
	userConnection.ServerDatabaseConnectionId = connection.ID
	userConnection.CreatedBy = user.ID
	c.DbConnections().SpireDb().Create(&userConnection)

	fmt.Printf("Created connection [%v]\n", connection.ID)

	var fUser models.User
	query := c.DbConnections().SpireDb().Model(&fUser)
	for _, relationship := range fUser.Relationships() {
		query = query.Preload(relationship)
	}

	query.Find(&fUser, user.ID)

	pJson, err := json.MarshalIndent(fUser, "", " ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%+v\n", string(pJson))

	// assert
	// make sure we can see owned connections
	assert.Contains(t, string(pJson), "owned_connections")
	assert.Contains(t, string(pJson), "user_connections")
	assert.Contains(t, string(pJson), "database_connection")
}
