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

void test_transposeMatrix() {
    matrix m = newMatrix(
        4, 4,
        (float*)(float[4][4]){0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8});
    matrix tm = newMatrix(
        4, 4,
        (float*)(float[4][4]){0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8});

    matrix t = transpose(m);

    assert(equalMatrix(tm, t));
}

void test_transposeIdentity() {
    matrix t = transpose(identityMatrix);

    assert(equalMatrix(t, identityMatrix));
}

void test_determinant2x2() {
    matrix m = newMatrix(2, 2, (float*)(float[2][2]){1, 5, -3, 2});

    float det = determinant(m);

    assert(equalFloat(det, 17));
}

void test_submatrixOf3x3Is2x2() {
    matrix m =
        newMatrix(3, 3, (float*)(float[3][3]){1, 5, 0, -3, 2, 7, 0, 6, -3});
    matrix sm = newMatrix(2, 2, (float*)(float[2][2]){-3, 2, 0, 6});

    matrix sub = submatrix(m, 0, 2);

    assert(equalMatrix(sm, sub));
}

void test_submatrixOf4x4Is3x3() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){-6, 1, 1, 6, -8, 5, 8, 6, -1, 0,
                                               8, 2, -7, 1, -1, 1});
    matrix sm =
        newMatrix(3, 3, (float*)(float[3][3]){-6, 1, 6, -8, 8, 6, -7, -1, 1});

    matrix sub = submatrix(m, 2, 1);

    assert(equalMatrix(sm, sub));
}

void test_minor3x3() {
    matrix m =
        newMatrix(3, 3, (float*)(float[3][3]){3, 5, 0, 2, -1, -7, 6, -1, 5});
    matrix sm = submatrix(m, 1, 0);

    float subDet = determinant(sm);
    float mMinor = minor(m, 1, 0);

    assert(equalFloat(subDet, 25));
    assert(equalFloat(mMinor, 25));
}

void test_cofactor3x3() {
    matrix m =
        newMatrix(3, 3, (float*)(float[3][3]){3, 5, 0, 2, -1, -7, 6, -1, 5});

    float mMinor1 = minor(m, 0, 0);
    float mMinor2 = minor(m, 1, 0);
    float mCofactor1 = cofactor(m, 0, 0);
    float mCofactor2 = cofactor(m, 1, 0);

    assert(equalFloat(mMinor1, -12));
    assert(equalFloat(mMinor2, 25));
    assert(equalFloat(mCofactor1, -12));
    assert(equalFloat(mCofactor2, -25));
}

void test_determinant3x3() {
    matrix m =
        newMatrix(3, 3, (float*)(float[3][3]){1, 2, 6, -5, 8, -4, 2, 6, 4});

    float mCofactor1 = cofactor(m, 0, 0);
    float mCofactor2 = cofactor(m, 0, 1);
    float mCofactor3 = cofactor(m, 0, 2);
    float det = determinant(m);

    assert(equalFloat(mCofactor1, 56));
    assert(equalFloat(mCofactor2, 12));
    assert(equalFloat(mCofactor3, -46));
    assert(equalFloat(det, -196));
}

void test_determinant4x4() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){-2, -8, 3, 5, -3, 1, 7, 3, 1, 2,
                                               -9, 6, -6, 7, 7, -9});

    float mCofactor1 = cofactor(m, 0, 0);
    float mCofactor2 = cofactor(m, 0, 1);
    float mCofactor3 = cofactor(m, 0, 2);
    float mCofactor4 = cofactor(m, 0, 3);
    float det = determinant(m);

    assert(equalFloat(mCofactor1, 690));
    assert(equalFloat(mCofactor2, 447));
    assert(equalFloat(mCofactor3, 210));
    assert(equalFloat(mCofactor4, 51));
    assert(equalFloat(det, -4071));
}

void test_isInvertible_yes() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3,
                                               -7, 9, 1, 7, -6});

    assert(equalFloat(determinant(m), -2120));
    assert(isInvertible(m));
}

void test_isInvertible_no() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){-4, 2, -2, -3, 9, 6, 2, 6, 0, -5,
                                               1, -5, 0, 0, 0, 0});

    assert(equalFloat(determinant(m), 0));
    assert(!isInvertible(m));
}

void test_inverse() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){-5, 2, 6, -8, 1, -5, 1, 8, 7, 7,
                                               -6, -7, 1, -3, 7, 4});
    matrix inv_exp =
        newMatrix(4, 4,
                  (float*)(float[4][4]){0.21805, 0.45113, 0.24060, -0.04511,
                                        -0.80827, -1.45677, -0.44361, 0.52068,
                                        -0.07895, -0.22368, -0.05263, 0.19737,
                                        -0.52256, -0.81391, -0.30075, 0.30639});

    matrix m_inv = inverse(m);

    assert(equalFloat(determinant(m), 532));
    assert(equalFloat(cofactor(m, 2, 3), -160));
    assert(equalFloat(get(m_inv, 3, 2), -160.0 / 532));
    assert(equalFloat(cofactor(m, 3, 2), 105));
    assert(equalFloat(get(m_inv, 2, 3), 105.0 / 532));
    assert(equalMatrix(m_inv, inv_exp));
}

void test_inverse_second() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){8, -5, 9, 2, 7, 5, 6, 1, -6, 0,
                                               9, 6, -3, 0, -9, -4});
    matrix inv_exp =
        newMatrix(4, 4,
                  (float*)(float[4][4]){
                      -0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308,
                      0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308,
                      -0.69231, -0.69231, -0.76923, -1.92308});

    matrix m_inv = inverse(m);

    assert(equalMatrix(m_inv, inv_exp));
}

void test_inverse_third() {
    matrix m = newMatrix(4, 4,
                         (float*)(float[4][4]){9, 3, 0, 9, -5, -2, -6, -3, -4,
                                               9, 6, 4, -7, 6, 6, 2});
    matrix inv_exp =
        newMatrix(4, 4,
                  (float*)(float[4][4]){-0.04074, -0.07778, 0.14444, -0.22222,
                                        -0.07778, 0.03333, 0.36667, -0.33333,
                                        -0.02901, -0.14630, -0.10926, 0.12963,
                                        0.17778, 0.06667, -0.26667, 0.33333});

    matrix m_inv = inverse(m);

    assert(equalMatrix(m_inv, inv_exp));
}

void test_inverse_multiply() {
    matrix a = newMatrix(4, 4,
                         (float*)(float[4][4]){3, -9, 7, 3, 3, -8, 2, -9, -4, 4,
                                               4, 1, -6, 5, -1, 1});
    matrix b = newMatrix(4, 4,
                         (float*)(float[4][4]){8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5,
                                               4, 6, -2, 0, 5});
    matrix c = multiplyMatrices(a, b);

    matrix b_inv = inverse(b);

    assert(equalMatrix(multiplyMatrices(c, b_inv), a));
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
    test_transposeMatrix();
    test_transposeIdentity();
    test_submatrixOf3x3Is2x2();
    test_submatrixOf4x4Is3x3();
    test_minor3x3();
    test_cofactor3x3();
    test_determinant3x3();
    test_determinant4x4();
    test_isInvertible_yes();
    test_isInvertible_no();
    test_inverse();
    test_inverse_second();
    test_inverse_third();
    test_inverse_multiply();
    printf("Matrix module tests successful\n");
}