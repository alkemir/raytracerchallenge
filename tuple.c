#include "tuple.h"

#include <math.h>

const float EPSILON = 0.00001;
const float POINT_W = 1.0;
const float VECTOR_W = 0.0;
const float COLOR_W = 0.0;

tuple newPoint(float x, float y, float z) {
    tuple t = {x, y, z, POINT_W};
    return t;
}

tuple newVector(float x, float y, float z) {
    tuple t = {x, y, z, VECTOR_W};
    return t;
}

tuple newColor(float r, float g, float b) {
    tuple t = {r, g, b, COLOR_W};
    return t;
}

bool isPoint(tuple t) { return t.w == 1.0; }

bool isVector(tuple t) { return t.w == 0.0; }

bool equalFloat(float a, float b) { return fabsf(a - b) < EPSILON; }

bool equalTuple(tuple a, tuple b) {
    return equalFloat(a.x, b.x) && equalFloat(a.y, b.y) &&
           equalFloat(a.z, b.z) && equalFloat(a.w, b.w);
}

tuple addTuple(tuple a, tuple b) {
    a.x += b.x;
    a.y += b.y;
    a.z += b.z;
    a.w += b.w;
    return a;
}

tuple substractTuple(tuple a, tuple b) {
    a.x -= b.x;
    a.y -= b.y;
    a.z -= b.z;
    a.w -= b.w;
    return a;
}

tuple negateTuple(tuple a) {
    a.x *= -1;
    a.y *= -1;
    a.z *= -1;
    a.w *= -1;
    return a;
}

tuple multiplyTuple(tuple a, float s) {
    a.x *= s;
    a.y *= s;
    a.z *= s;
    a.w *= s;
    return a;
}

tuple divideTuple(tuple a, float d) {
    a.x /= d;
    a.y /= d;
    a.z /= d;
    a.w /= d;
    return a;
}

float magnitude(tuple a) { return sqrtf(a.x * a.x + a.y * a.y + a.z * a.z); }

tuple normalize(tuple a) { return divideTuple(a, magnitude(a)); }

float dot(tuple a, tuple b) {
    return a.x * b.x + a.y * b.y + a.z * b.z + a.w * b.w;
}

tuple cross(tuple a, tuple b) {
    tuple x = {a.y * b.z - a.z * b.y, a.z * b.x - a.x * b.z,
               a.x * b.y - a.y * b.x};
    return x;
}

tuple hadamard(tuple a, tuple b) {
    a.x *= b.x;
    a.y *= b.y;
    a.z *= b.z;
    return a;
}