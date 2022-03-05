package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (boid *Boid) calcAcceleration() Vector2D {
	upper, lower := boid.position.AddValue(viewRadius), boid.position.AddValue(-viewRadius)
	meanVelocity := Vector2D{0, 0}
	count := 0.0
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != boid.id {
				if dist := boids[otherBoidId].position.GetDistance(boid.position); dist < viewRadius {
					count++
					meanVelocity = meanVelocity.Add(boids[otherBoidId].velocity)
				}
			}
		}
	}

	acceleration := Vector2D{0, 0}
	if count > 0 {
		meanVelocity = meanVelocity.DivisionValue(count)
		acceleration = meanVelocity.Subtract(boid.velocity).MultiplyValue(adjustmentRate)
	}
	return acceleration
}

func (boid *Boid) moveOne() {
	acceleration := boid.calcAcceleration()
	boid.velocity = boid.velocity.Add(acceleration).normalize()
	next := boid.position.Add(boid.velocity)
	if next.x >= screenWidth || next.x < 0 {
		boid.velocity = Vector2D{x: -boid.velocity.x, y: boid.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		boid.velocity = Vector2D{x: boid.velocity.x, y: -boid.velocity.y}
	}
	boidMap[int(boid.position.x)][int(boid.position.y)] = -1
	boid.position = boid.position.Add(boid.velocity)
	boidMap[int(boid.position.x)][int(boid.position.y)] = boid.id
}

func (boid *Boid) start() {
	for {
		boid.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(id int) {
	boid := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       id,
	}
	boids[id] = &boid
	go boid.start()
}
