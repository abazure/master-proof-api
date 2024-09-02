package config

import (
	"context"
	"firebase.google.com/go/v4/auth"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitializeFirebase() *firebase.App {
	ctx := context.Background()

	opt := option.WithCredentialsFile("proof-master-firebase-adminsdk-91rvv-9df75f40f1.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	return app
}

func FirebaseAuthInitialize(app *firebase.App) *auth.Client {
	ctx := context.Background()
	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}
	return firebaseAuth
}
