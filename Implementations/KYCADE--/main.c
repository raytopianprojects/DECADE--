#include <time.h>
#include <stdio.h>
#include <stdint.h>
#include "parser.h"

int main(int argc, char **argv) {
	if (time(NULL) >= 1893474000 || time(NULL) < 1577854800) {
		return 759841101;
	}
	if (argc < 2) {
		printf("USAGE: KYCADE--.exe path/to/file.dcj");
	}
	char* fileContents = (char* )(main);
	uint64_t numTokens;
	uint64_t flen;
	tokenizeFile(argv[argc-1], fileContents, &numTokens, &flen);
	return 0;
}
