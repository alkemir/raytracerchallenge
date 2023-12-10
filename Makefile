CC = clang

all : test

.PHONY : clean test all

tuple_test.exe : tuple_test.o floats.o
	$(CC) -o tuple_test.exe tuple_test.o tuple.o floats.o

matrix_test.exe : matrix_test.o matrix.o floats.o tuple.o
	$(CC) -o matrix_test.exe matrix_test.o matrix.o floats.o tuple.o

canvas_test.exe : canvas_test.o canvas.o tuple.o floats.o
	$(CC) -o canvas_test.exe canvas_test.o canvas.o tuple.o floats.o

floats.o : floats.h
tuple.o : tuple.h floats.h
matrix.o : floats.h matrix.h
canvas.o : canvas.h tuple.h


test : matrix_test.exe tuple_test.exe canvas_test.exe
	matrix_test.exe
	tuple_test.exe
	canvas_test.exe


clean :
	rm *.o
	rm *.exe