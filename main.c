#include <assert.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>

typedef struct tuple {
    float x;
    float y;
    float z;
    float w;
} tuple;

tuple newPoint(float x, float y, float z) {
    tuple t = {x, y, z, 1.0};
    return t;
}

bool isPoint(tuple t) { return t.w == 1.0; }

tuple newVector(float x, float y, float z) {
    tuple t = {x, y, z, 0.0};
    return t;
}

bool isVector(tuple t) { return t.w == 0.0; }

bool equalFloat(float a, float b) { return fabsf(a - b) < 0.00001; }

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

void test_tupleIsPoint() {
    tuple t = newPoint(4.3, -4.2, 3.1);

    assert(equalFloat(t.x, 4.3));
    assert(equalFloat(t.y, -4.2));
    assert(equalFloat(t.z, 3.1));
    assert(equalFloat(t.w, 1.0));
    assert(isPoint(t));
    assert(!isVector(t));
}

void test_tupleIsVector() {
    tuple t = newVector(4.3, -4.2, 3.1);

    assert(equalFloat(t.x, 4.3));
    assert(equalFloat(t.y, -4.2));
    assert(equalFloat(t.z, 3.1));
    assert(equalFloat(t.w, 0.0));
    assert(!isPoint(t));
    assert(isVector(t));
}

void test_addTuples() {
    tuple p = newPoint(3, -2, 5);
    tuple v = newVector(-2, 3, 1);

    assert(equalTuple(addTuple(p, v), newPoint(1, 1, 6)));
}

void test_substractPointFromPoint() {
    tuple p1 = newPoint(3, 2, 1);
    tuple p2 = newPoint(5, 6, 7);

    assert(equalTuple(substractTuple(p1, p2), newVector(-2, -4, -6)));
}

void test_substractVectorFromPoint() {
    tuple p = newPoint(3, 2, 1);
    tuple v = newVector(5, 6, 7);

    assert(equalTuple(substractTuple(p, v), newPoint(-2, -4, -6)));
}

void test_substractVectorFromVector() {
    tuple v1 = newVector(3, 2, 1);
    tuple v2 = newVector(5, 6, 7);

    assert(equalTuple(substractTuple(v1, v2), newVector(-2, -4, -6)));
}

void test_substractVectorFromZero() {
    tuple v = newVector(1, -2, 3);
    tuple zero = newVector(0, 0, 0);

    assert(equalTuple(substractTuple(zero, v), newVector(-1, 2, -3)));
}

void test_negateTuple() {
    tuple t = {1, -2, 3, -4};
    tuple nt = {-1, 2, -3, 4};

    assert(equalTuple(negateTuple(t), nt));
}

void test_multiplyTupleByScalar() {
    tuple t = {1, -2, 3, -4};
    tuple p = {3.5, -7, 10.5, -14};

    assert(equalTuple(multiplyTuple(t, 3.5), p));
}

void test_multiplyTupleByFraction() {
    tuple t = {1, -2, 3, -4};
    tuple p = {0.5, -1, 1.5, -2};

    assert(equalTuple(multiplyTuple(t, 0.5), p));
}

void test_divideTupleByScalar() {
    tuple t = {1, -2, 3, -4};
    tuple d = {0.5, -1, 1.5, -2};

    assert(equalTuple(divideTuple(t, 2), d));
}

void test_magnitude() {
    tuple v1 = newVector(1, 0, 0);
    tuple v2 = newVector(0, 1, 0);
    tuple v3 = newVector(0, 0, 1);
    tuple v4 = newVector(1, 2, 3);
    tuple v5 = newVector(-1, -2, -3);

    assert(equalFloat(magnitude(v1), 1.0));
    assert(equalFloat(magnitude(v2), 1.0));
    assert(equalFloat(magnitude(v3), 1.0));
    assert(equalFloat(magnitude(v4), sqrtf(14)));
    assert(equalFloat(magnitude(v5), sqrtf(14)));
}

void test_normalize() {
    tuple v1 = newVector(4, 0, 0);
    tuple n1 = newVector(1, 0, 0);
    tuple v2 = newVector(1, 2, 3);
    tuple n2 = divideTuple(newVector(1, 2, 3), sqrtf(14));

    assert(equalTuple(normalize(v1), n1));
    assert(equalTuple(normalize(v2), n2));
    assert(equalFloat(magnitude(normalize(v2)), 1.0));
}

void test_dotProduct() {
    tuple v1 = newVector(1, 2, 3);
    tuple v2 = newVector(2, 3, 4);

    assert(equalFloat(dot(v1, v2), 20));
}

void test_crossProduct() {
    tuple v1 = newVector(1, 2, 3);
    tuple v2 = newVector(2, 3, 4);
    tuple vx = newVector(-1, 2, -1);

    assert(equalTuple(cross(v1, v2), vx));
    assert(equalTuple(cross(v2, v1), negateTuple(vx)));
}

int main() {
    test_tupleIsPoint();
    test_tupleIsVector();
    test_addTuples();
    test_substractPointFromPoint();
    test_substractVectorFromPoint();
    test_substractVectorFromVector();
    test_substractVectorFromZero();
    test_negateTuple();
    test_multiplyTupleByScalar();
    test_multiplyTupleByFraction();
    test_divideTupleByScalar();
    test_magnitude();
    test_normalize();
    test_dotProduct();
    test_crossProduct();
}