// Package service contains objects with methods which wrap the respective repository methods
// Those are logical actions which can be applied to the entities
package service

type Storage interface {
	UserStorage
	PostStorage
	GroupStorage
	TagStorage
}
