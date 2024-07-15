package mongodb

import (
	"context"
	"fmt"
	entity "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface defines the contract for user repository operations.
type Interface interface {
	Add(ctx context.Context, user *entity.USER) error
	GetAll(ctx context.Context) ([]entity.USER, error)
	GetOne(ctx context.Context, email string) (*entity.USER, error)
	Delete(ctx context.Context, email string) error
}

// UserRepository implements the Interface for MongoDB operations.
type UserRepository struct {
	collection *mongo.Collection
}

// NewOfficerRepository creates a new UserRepository instance.
// It takes a MongoDB database connection and a collection name as parameters.
func NewOfficerRepository(db *mongo.Database, collectionName string) *UserRepository {
	return &UserRepository{
		collection: db.Collection(collectionName),
	}
}

// Add inserts a new user into the MongoDB collection.
// It checks for existing users with the same email before insertion.
func (o *UserRepository) Add(ctx context.Context, user *entity.USER) error {
	filter := bson.M{"email": user.Email}
	count, _ := o.collection.CountDocuments(ctx, filter)
	if count > 0 {
		return fmt.Errorf("user already exists")
	}
	_, err := o.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert: %v", err)
	}
	return nil
}

// GetAll retrieves all users from the MongoDB collection.
func (o *UserRepository) GetAll(ctx context.Context) ([]entity.USER, error) {
	cursor, err := o.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all records: %v", err)
	}
	defer cursor.Close(ctx)

	var users []entity.USER
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetOne retrieves a single user from the MongoDB collection based on the provided email.
func (o *UserRepository) GetOne(ctx context.Context, email string) (*entity.USER, error) {
	filter := bson.M{"email": email}
	var user entity.USER
	err := o.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get record: %v", err)
	}
	return &user, nil
}

// Delete removes a user from the MongoDB collection based on the provided email.
func (o *UserRepository) Delete(ctx context.Context, email string) error {
	filter := bson.M{"email": email}
	_, err := o.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete user acc")
	}
	return nil
}