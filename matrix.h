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

matrix transpose(matrix m);

float determinant(matrix m);

matrix submatrix(matrix m, int row, int col);

float minor(matrix m, int row, int col);

float cofactor(matrix m, int row, int col);

void printMatrix(matrix m);