// 1st of October, 2021
// Author Asliddin Abdivasiyev
// Email space.coding.programmer@gmail.com
// This file provides configurations of library
package gocache

// Driver type
// Client selects redis or memcached
type Driver int

// Available drivers
const (
	RedisDriver     Driver = iota // redis driver
	MemcachedDriver               // memcached driver
)

// Configuration of the package
type Config struct {
	Driver     Driver   // driver type
	Addresses  []string // connection address
	Username   string   // username
	Password   string   // password
	DB         int      // database
	Expiration int      // expiration of the cache in seconds
}

// Construct new configuration
func NewConfig() *Config {
	return &Config{
		Driver:     RedisDriver,
		Addresses:  []string{"127.0.0.1:6379"},
		Username:   "",
		Password:   "",
		DB:         0,
		Expiration: 3,
	}
}

// Change default driver to your own
func (c *Config) WithDriver(driver Driver) *Config {
	c.Driver = driver
	return c
}

// Change connection addresses or set replica addresses
func (c *Config) WithAddr(addresses ...string) *Config {
	c.Addresses = addresses
	return c
}

// Change credentials
func (c *Config) WithCreds(username, password string) *Config {
	c.Username = username
	c.Password = password
	return c
}

// Change db
func (c *Config) WithDB(db int) *Config {
	c.DB = db
	return c
}

// Change expiration time
func (c *Config) WithExpiration(exp int) *Config {
	c.Expiration = exp
	return c
}
