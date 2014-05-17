#include <stdio.h>

int main(int argc, char const *argv[]) {
	if (argc != 2) {
		return 1;
	}

	FILE *file = fopen(argv[1], "r");
	fseek(file, 1, SEEK_END);
	long length = ftell(file) - 1;
	fclose(file);

	printf("%lu\n", length);

	return 0;
}