#include "intersection.h"

intersection newIntersection(float t, sphere *object) {
    intersection i = {t, object};
    return i;
}