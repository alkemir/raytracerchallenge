#include "transformations.h"

#include <assert.h>
#include <stdio.h>

#include "matrix.h"
#include "tuple.h"

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

int main() {
    test_translatePoint();
    test_translatePoint_reverse();
    test_translateVector();
    test_scalePoint();
    test_scaleVector();
    test_scaleVector_reverse();
    printf("Transformations module tests successful\n");
}