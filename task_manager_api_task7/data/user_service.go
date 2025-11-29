package data

import (
	"context"
	"fmt"
	"time"

	"github.com/zaahidali/task_manager_api/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
)

var JwtSecret = []byte("YOUR_SECRET_KEY")

func Register(user models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	if UserID == 1 {
		user.Role = "admin"
	}
	user.ID = UserID
	UserID++
	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Login(user models.User) (string, error) {
	filter := bson.M{"username": user.Username}

	var exist models.User
	err := UserCollection.FindOne(context.TODO(), filter).Decode(&exist)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(user.Password)) != nil {
		return "", fmt.Errorf("unvalid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       exist.ID,
		"role":     exist.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"username": exist.Username,
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Promote(id int) error {
	filter := bson.M{"id": id}

	var exist models.User
	err := UserCollection.FindOne(context.TODO(), filter).Decode(&exist)
	if err != nil {
		return err
	}

	updateFields := bson.D{}

	updateFields = append(updateFields, bson.E{Key: "role", Value: "admin"})

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, nerr := UserCollection.UpdateOne(context.TODO(), filter, update)
	if nerr != nil {
		return nerr
	}

	return nil
}
