#include "matrix.h"

#include <assert.h>
#include <stdio.h>

#include "floats.h"
#include "tuple.h"

void test_matrixSubscriptable4x4() {
    matrix m =
        newMatrix(4, 4,
                  (float*)(float[4][4]){1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10,
                                        11, 12, 13.5, 14.5, 15.5, 16.5});

    assert(equalFloat(get(m, 0, 0), 1));
    assert(equalFloat(get(m, 0, 3), 4));
    assert(equalFloat(get(m, 1, 0), 5.5));
    assert(equalFloat(get(m, 1, 2), 7.5));
    assert(equalFloat(get(m, 2, 2), 11));
    assert(equalFloat(get(m, 3, 0), 13.5));
    assert(equalFloat(get(m, 3, 2), 15.5));
}

void test_matrixSubscriptable2x2() {
    matrix m = newMatrix(2, 2, (float*)(float[2][2]){-3, 5, 1, -2});

    assert(equalFloat(get(m, 0, 0), -3));
    assert(equalFloat(get(m, 0, 1), 5));
    assert(equalFloat(get(m, 1, 0), 1));
    assert(equalFloat(get(m, 1, 1), -2));
}

void test_matrixSubscriptable3x3() {
    matrix m =
        newMatrix(3, 3, (float*)(float[3][3]){-3, 5, 0, 1, -2, -7, 0, 1, 1});

    assert(equalFloat(get(m, 0, 0), -3));
    assert(equalFloat(get(m, 1, 1), -2));
    assert(equalFloat(get(m, 2, 2), 1));
}

void test_compareMatricesSame() {
    matrix m1 = newMatrix(
        4, 4,
        (float*)(float[4][4]){1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2});
    matrix m2 = newMatrix(
        4, 4,
        (float*)(float[4][4]){1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2});

    assert(equalMatrix(m1, m2));
}

void test_compareMatricesDifferent() {
    matrix m1 = newMatrix(
        4, 4,
        (float*)(float[4][4]){1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2});
    matrix m2 = newMatrix(
        4, 4,
        (float*)(float[4][4]){2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1});

    assert(!equalMatrix(m1, m2));
}

void test_multiplyMatrices() {
    matrix m1 = newMatrix(
        4, 4,
        (float*)(float[4][4]){1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2});
    matrix m2 = newMatrix(4, 4,
                          (float*)(float[4][4]){-2, 1, 2, 3, 3, 2, 1, -1, 4, 3,
                                                6, 5, 1, 2, 7, 8});

    matrix p =
        newMatrix(4, 4,
                  (float*)(float[4][4]){20, 22, 50, 48, 44, 54, 114, 108, 40,
                                        58, 110, 102, 16, 26, 46, 42});

    assert(equalMatrix(multiplyMatrices(m1, m2), p));
}

void test_multiplyMatrixByTuple() {
    matrix m = newMatrix(
        4, 4,
        (float*)(float[4][4]){1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1});
    tuple t = newTuple(1, 2, 3, 1);

    tuple p = newTuple(18, 24, 33, 1);

    assert(equalTuple(multiplyMatrixTuple(m, t), p));
}

void test_multiplyIdentityByMatrix() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8,
                                               16, 4, 8, 16, 32});

    assert(equalMatrix(multiplyMatrices(m, identityMatrix), m));
}

void test_multiplyIdentityByTuple() {
    tuple t = newTuple(1, 2, 3, 4);

    assert(equalTuple(multiplyMatrixTuple(identityMatrix, t), t));
}

int main() {
    test_matrixSubscriptable4x4();
    test_matrixSubscriptable2x2();
    test_matrixSubscriptable3x3();
    test_compareMatricesSame();
    test_compareMatricesDifferent();
    test_multiplyMatrices();
    test_multiplyMatrixByTuple();
    test_multiplyIdentityByMatrix();
    test_multiplyIdentityByTuple();
    printf("Matrix module tests successful\n");
}