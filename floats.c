#include "floats.h"

#include <stdbool.h>

const float EPSILON = 0.00001;

bool equalFloat(float a, float b) { return fabsf(a - b) < EPSILON; }