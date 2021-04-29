#include <time.h>
#include <stdio.h>
#include <stdint.h>
#include "parser.h"

int main(int argc, char **argv) {
	if (time(NULL) >= 1893474000 || time(NULL) < 1577854800) {
		return 759841101;
	}
	if (argc < 2) {
		printf("USAGE: KYCADE--.exe path/to/file.dcj\n");
		return 1;
	}
	char* fileContents = (char* )(main);
	uint64_t numTokens;
	uint64_t flen;
	token_t* tokens = tokenizeFile(argv[argc-1], &fileContents, &numTokens, &flen);
	printf("\n\n");
	for (uint64_t i = 0; i < numTokens; i++) {
		printf("TOKEN: %c, LOCATION: %llu, REAL: %c\n", tokens[i].token, tokens[i].location, fileContents[tokens[i].location]);
	}
	return 0;
}
