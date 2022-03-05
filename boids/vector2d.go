package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

func (vector Vector2D) AddValue(value float64) Vector2D {
	return Vector2D{x: vector.x + value, y: vector.y + value}
}

func (vector Vector2D) MultiplyValue(value float64) Vector2D {
	return Vector2D{x: vector.x * value, y: vector.y * value}
}

func (vector Vector2D) DivisionValue(value float64) Vector2D {
	return Vector2D{x: vector.x / value, y: vector.y / value}
}

func (vector Vector2D) normalize() Vector2D {
	factor := math.Sqrt(math.Pow(vector.x, 2) + math.Pow(vector.y, 2))
	return Vector2D{x: vector.x / factor, y: vector.y / factor}
}

func (v1 Vector2D) GetDistance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
