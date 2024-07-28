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
    matrix half_quarter = rotationX(M_PI_4);
    matrix full_quarter = rotationX(M_PI_2);
    tuple p = newPoint(0, 1, 0);

    tuple r1 = multiplyMatrixTuple(half_quarter, p);
    tuple r2 = multiplyMatrixTuple(full_quarter, p);

    assert(equalTuple(r1, newPoint(0, M_SQRT2 / 2, M_SQRT2 / 2)));
    assert(equalTuple(r2, newPoint(0, 0, 1)));
}

void test_rotationX_reverse() {
    matrix half_quarter = rotationX(M_PI_4);
    matrix r_inv = inverse(half_quarter);
    tuple p = newPoint(0, 1, 0);

    tuple r1 = multiplyMatrixTuple(r_inv, p);

    assert(equalTuple(r1, newPoint(0, M_SQRT2 / 2, -1 * M_SQRT2 / 2)));
}

void test_rotationY() {
    matrix half_quarter = rotationY(M_PI_4);
    matrix full_quarter = rotationY(M_PI_2);
    tuple p = newPoint(0, 0, 1);

    tuple r1 = multiplyMatrixTuple(half_quarter, p);
    tuple r2 = multiplyMatrixTuple(full_quarter, p);

    assert(equalTuple(r1, newPoint(M_SQRT2 / 2, 0, M_SQRT2 / 2)));
    assert(equalTuple(r2, newPoint(1, 0, 0)));
}

void test_rotationZ() {
    matrix half_quarter = rotationZ(M_PI_4);
    matrix full_quarter = rotationZ(M_PI_2);
    tuple p = newPoint(0, 1, 0);

    tuple r1 = multiplyMatrixTuple(half_quarter, p);
    tuple r2 = multiplyMatrixTuple(full_quarter, p);

    assert(equalTuple(r1, newPoint(-1 * M_SQRT2 / 2, M_SQRT2 / 2, 0)));
    assert(equalTuple(r2, newPoint(-1, 0, 0)));
}

void test_shearing_xtoy() {
    matrix s = shearing(1, 0, 0, 0, 0, 0);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(5, 3, 4)));
}

void test_shearing_xtoz() {
    matrix s = shearing(0, 1, 0, 0, 0, 0);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(6, 3, 4)));
}

void test_shearing_ytox() {
    matrix s = shearing(0, 0, 1, 0, 0, 0);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(2, 5, 4)));
}

void test_shearing_ytoz() {
    matrix s = shearing(0, 0, 0, 1, 0, 0);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(2, 7, 4)));
}

void test_shearing_ztox() {
    matrix s = shearing(0, 0, 0, 0, 1, 0);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(2, 3, 6)));
}

void test_shearing_ztoy() {
    matrix s = shearing(0, 0, 0, 0, 0, 1);
    tuple p = newPoint(2, 3, 4);

    tuple r = multiplyMatrixTuple(s, p);

    assert(equalTuple(r, newPoint(2, 3, 7)));
}

void test_individual_transformations() {
    tuple p = newPoint(1, 0, 1);
    matrix A = rotationX(M_PI_2);
    matrix B = scaling(5, 5, 5);
    matrix C = translation(10, 5, 7);

    tuple p2 = multiplyMatrixTuple(A, p);
    tuple p3 = multiplyMatrixTuple(B, p2);
    tuple p4 = multiplyMatrixTuple(C, p3);

    assert(equalTuple(p2, newPoint(1, -1, 0)));
    assert(equalTuple(p3, newPoint(5, -5, 0)));
    assert(equalTuple(p4, newPoint(15, 0, 7)));
}

void test_chained_transformations() {
    tuple p = newPoint(1, 0, 1);
    matrix A = rotationX(M_PI_2);
    matrix B = scaling(5, 5, 5);
    matrix C = translation(10, 5, 7);

    matrix T = multiplyMatrices(C, multiplyMatrices(B, A));
    tuple r = multiplyMatrixTuple(T, p);

    assert(equalTuple(r, newPoint(15, 0, 7)));
}

void test_fluent_transformations() {
    tuple p = newPoint(1, 0, 1);

    matrix T =
        translate(10, 5, 7, scale(5, 5, 5, rotateX(M_PI_2, identityMatrix)));
    tuple r = multiplyMatrixTuple(T, p);

    assert(equalTuple(r, newPoint(15, 0, 7)));
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
    test_shearing_xtoy();
    test_shearing_xtoz();
    test_shearing_ytox();
    test_shearing_ytoz();
    test_shearing_ztox();
    test_shearing_ztoy();
    test_individual_transformations();
    test_chained_transformations();
    test_fluent_transformations();
    printf("Transformations module tests successful\n");
}