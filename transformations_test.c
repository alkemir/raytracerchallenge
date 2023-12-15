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

int main() {
    test_translatePoint();
    test_translatePoint_reverse();
    printf("Transformations module tests successful\n");
}