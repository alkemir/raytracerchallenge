#include "matrix.h"

#include <malloc.h>
#include <stdbool.h>
#include <stdio.h>

#include "floats.h"

matrix identityMatrix = {
    4, 4,
    (float*)(float[4][4]){1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}};

matrix newMatrix(int rows, int cols, float* data) {
    float* dataCpy = malloc(rows * cols * sizeof(float));

    if (data != NULL) {
        for (int y = 0; y < rows; y++) {
            for (int x = 0; x < cols; x++) {
                dataCpy[x + y * cols] = data[x + y * cols];
            }
        }
    }

    matrix m = {rows, cols, dataCpy};
    return m;
}

bool equalMatrix(matrix a, matrix b) {
    if (a.rows != b.rows) {
        return false;
    }

    if (a.cols != b.cols) {
        return false;
    }

    for (int y = 0; y < a.rows; y++) {
        for (int x = 0; x < a.cols; x++) {
            if (!equalFloat(a.data[x + y * a.cols], b.data[x + y * a.cols])) {
                return false;
            }
        }
    }

    return true;
}

float get(matrix m, int row, int col) { return m.data[row * m.cols + col]; }

matrix multiplyMatrices(matrix a, matrix b) {
    float* dstData = malloc(a.rows * b.cols * sizeof(float));

    for (int y = 0; y < a.rows; y++) {
        for (int x = 0; x < b.cols; x++) {
            dstData[x + y * b.cols] = 0;
            for (int i = 0; i < a.cols; i++) {
                dstData[x + y * a.cols] +=
                    a.data[i + y * a.cols] * b.data[x + i * a.cols];
            }
        }
    }

    matrix r = {a.rows, b.cols, dstData};
    return r;
}

tuple multiplyMatrixTuple(matrix m, tuple t) {
    tuple r;

    r.x = m.data[0 + 0 * m.cols] * t.x + m.data[1 + 0 * m.cols] * t.y +
          m.data[2 + 0 * m.cols] * t.z + m.data[3 + 0 * m.cols] * t.w;

    r.y = m.data[0 + 1 * m.cols] * t.x + m.data[1 + 1 * m.cols] * t.y +
          m.data[2 + 1 * m.cols] * t.z + m.data[3 + 1 * m.cols] * t.w;

    r.z = m.data[0 + 2 * m.cols] * t.x + m.data[1 + 2 * m.cols] * t.y +
          m.data[2 + 2 * m.cols] * t.z + m.data[3 + 2 * m.cols] * t.w;

    r.w = m.data[0 + 3 * m.cols] * t.x + m.data[1 + 3 * m.cols] * t.y +
          m.data[2 + 3 * m.cols] * t.z + m.data[3 + 3 * m.cols] * t.w;

    return r;
}

matrix transpose(matrix m) {
    matrix d = newMatrix(m.rows, m.cols, NULL);
    for (int y = 0; y < m.rows; y++) {
        for (int x = 0; x < m.cols; x++) {
            d.data[y + x * m.cols] = m.data[x + y * m.cols];
        }
    }

    return d;
}

float determinant(matrix m) {
    if (m.cols == 2 && m.rows == 2) {
        return m.data[0] * m.data[3] - m.data[1] * m.data[2];
    }

    float det = 0;
    for (int y = 0; y < m.cols; y++) {
        det += get(m, 0, y) * cofactor(m, 0, y);
    }
    return det;
}

matrix submatrix(matrix m, int row, int col) {
    matrix s = newMatrix(m.rows - 1, m.cols - 1, NULL);

    int yOff = 0;
    for (int y = 0; y < m.rows; y++) {
        if (y == row) {
            yOff = 1;
            continue;
        }

        int xOff = 0;
        for (int x = 0; x < m.cols; x++) {
            if (x == col) {
                xOff = 1;
                continue;
            }

            s.data[x - xOff + (y - yOff) * s.cols] = m.data[x + y * m.cols];
        }
    }

    return s;
}

float minor(matrix m, int row, int col) {
    return determinant(submatrix(m, row, col));
}

float cofactor(matrix m, int row, int col) {
    float factor = 1;
    if ((row + col) % 2 == 1) {
        factor = -1;
    }

    return factor * minor(m, row, col);
}

void printMatrix(matrix m) {
    for (int y = 0; y < m.rows; y++) {
        for (int x = 0; x < m.cols; x++) {
            printf("%f ", m.data[x + y * m.cols]);
        }
        printf("\n");
    }
}