// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleCar() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the car's edges.

	// create car vertex with its edges.
	c := client.Car.
		Create().
		SetModel("string").
		SetRegisteredAt(time.Now()).
		SaveX(ctx)
	log.Println("car created:", c)

	// query edges.

	// Output:
}
func ExampleGroup() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the group's edges.
	u0 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u0)

	// create group vertex with its edges.
	gr := client.Group.
		Create().
		SetName("string").
		AddUsers(u0).
		SaveX(ctx)
	log.Println("group created:", gr)

	// query edges.
	u0, err = gr.QueryUsers().First(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	log.Println("users found:", u0)

	// Output:
}
func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.
	c0 := client.Car.
		Create().
		SetModel("string").
		SetRegisteredAt(time.Now()).
		SaveX(ctx)
	log.Println("car created:", c0)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetAge(1).
		SetName("string").
		AddCars(c0).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	c0, err = u.QueryCars().First(ctx)
	if err != nil {
		log.Fatalf("failed querying cars: %v", err)
	}
	log.Println("cars found:", c0)

	// Output:
}
