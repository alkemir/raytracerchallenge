#include "transformations.h"

#include <assert.h>
#include <math.h>
#include <stdio.h>

#include "matrix.h"
#include "tuple.h"

#define _USE_MATH_DEFINES

// Remove this when I figure out how to link against the library correctly

#define M_PI_2 1.57079632679489661923   // pi/2
#define M_PI_4 0.785398163397448309616  // pi/4
#define M_SQRT2 1.41421356237309504880  // sqrt(2)

void test_translatePoint() {
    matrix t = translation(5, -3, 2);
    tuple p = newPoint(-3, 4, 5);

    tuple r = multiplyMatrixTuple(t, p);

    assert(equalTuple(r, newPoint(2, 1, 7)));
}

void test_translatePoint_reverse() {
    matrix t = translation(5, -3, 2);
    matrix t_inv = inverse(t);
    tuple p = newPoint(-3, 4, 5);

    tuple r = multiplyMatrixTuple(t_inv, p);

    assert(equalTuple(r, newPoint(-8, 7, 3)));
}

void test_translateVector() {
    matrix t = translation(5, -3, 2);
    tuple v = newVector(-3, 4, 5);

    tuple r = multiplyMatrixTuple(t, v);

    assert(equalTuple(r, v));
}

void test_scalePoint() {
    matrix s = scaling(2, 3, 4);
    tuple p = newPoint(-4, 6, 8);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(-8, 18, 32)));
}

void test_scaleVector() {
    matrix s = scaling(2, 3, 4);
    tuple v = newVector(-4, 6, 8);

    tuple r = multiplyMatrixTuple(s, v);

    assert(equalTuple(r, newVector(-8, 18, 32)));
}

void test_scaleVector_reverse() {
    matrix s = scaling(2, 3, 4);
    matrix s_inv = inverse(s);
    tuple v = newVector(-4, 6, 8);

    tuple r = multiplyMatrixTuple(s_inv, v);

    assert(equalTuple(r, newVector(-2, 2, 2)));
}

void test_reflectionIsNegativeScaling() {
    matrix s = scaling(-1, 1, 1);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(-2, 3, 4)));
}

void test_rotationX() {
    matrix half_quarter = rotateX(M_PI_4);
    matrix full_quarter = rotateX(M_PI_2);
    tuple p = newPoint(0, 1, 0);

    tuple r1 = multiplyMatrixTuple(half_quarter, p);
    tuple r2 = multiplyMatrixTuple(full_quarter, p);

    assert(equalTuple(r1, newPoint(0, M_SQRT2 / 2, M_SQRT2 / 2)));
    assert(equalTuple(r2, newPoint(0, 0, 1)));
}

void test_rotationX_reverse() {
    matrix half_quarter = rotateX(M_PI_4);
    matrix r_inv = inverse(half_quarter);
    tuple p = newPoint(0, 1, 0);

    tuple r1 = multiplyMatrixTuple(r_inv, p);

    assert(equalTuple(r1, newPoint(0, M_SQRT2 / 2, -1 * M_SQRT2 / 2)));
}

void test_rotationY() {
    matrix half_quarter = rotateY(M_PI_4);
    matrix full_quarter = rotateY(M_PI_2);
    tuple p = newPoint(0, 0, 1);

    tuple r1 = multiplyMatrixTuple(half_quarter, p);
    tuple r2 = multiplyMatrixTuple(full_quarter, p);

    assert(equalTuple(r1, newPoint(M_SQRT2 / 2, 0, M_SQRT2 / 2)));
    assert(equalTuple(r2, newPoint(1, 0, 0)));
}

void test_rotationZ() {
    matrix half_quarter = rotateZ(M_PI_4);
    matrix full_quarter = rotateZ(M_PI_2);
    tuple p = newPoint(0, 1, 0);

    tuple r1 = multiplyMatrixTuple(half_quarter, p);
    tuple r2 = multiplyMatrixTuple(full_quarter, p);

    assert(equalTuple(r1, newPoint(-1 * M_SQRT2 / 2, M_SQRT2 / 2, 0)));
    assert(equalTuple(r2, newPoint(-1, 0, 0)));
}

int main() {
    test_translatePoint();
    test_translatePoint_reverse();
    test_translateVector();
    test_scalePoint();
    test_scaleVector();
    test_scaleVector_reverse();
    test_reflectionIsNegativeScaling();
    test_rotationX();
    test_rotationX_reverse();
    test_rotationY();
    test_rotationZ();
    printf("Transformations module tests successful\n");
}