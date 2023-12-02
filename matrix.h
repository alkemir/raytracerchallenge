#include <stdbool.h>

#include "tuple.h"

typedef struct matrix_s {
    int rows;
    int cols;

    float* data;
} matrix;

extern matrix identityMatrix;

matrix newMatrix(int rows, int cols, float* data);

bool equalMatrix(matrix a, matrix b);

float get(matrix m, int row, int col);

matrix multiplyMatrices(matrix a, matrix b);

tuple multiplyMatrixTuple(matrix a, tuple b);