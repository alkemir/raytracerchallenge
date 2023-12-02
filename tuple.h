#include <stdbool.h>

#pragma once

// TODO: Make this struct opaque when I figure out how to do it
typedef struct tuple_s {
    float x;
    float y;
    float z;
    float w;
} tuple;

tuple newPoint(float x, float y, float z);

tuple newVector(float x, float y, float z);

tuple newColor(float r, float g, float b);

bool isPoint(tuple t);

bool isVector(tuple t);

bool equalFloat(float a, float b);

bool equalTuple(tuple a, tuple b);

tuple addTuple(tuple a, tuple b);

tuple substractTuple(tuple a, tuple b);

tuple negateTuple(tuple a);

tuple multiplyTuple(tuple a, float s);

tuple divideTuple(tuple a, float d);

float magnitude(tuple a);

tuple normalize(tuple a);

float dot(tuple a, tuple b);

tuple cross(tuple a, tuple b);

tuple hadamard(tuple a, tuple b);