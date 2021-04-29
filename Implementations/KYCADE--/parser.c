#include "parser.h"

token_t* tokenizeFile(char* file, char* contents, uint64_t* numTokens, uint64_t* flen) {
	FILE* dcj = fopen(file, "r");
	if (dcj == NULL) return NULL;
	fseek(dcj, 0L, SEEK_END);
	*flen = ftell(dcj);
	fseek(dcj, 0L, SEEK_SET);
	
	contents = malloc(*flen);
	token_t* tokens = malloc(0);
	
	uint8_t inString = 0;
	for (uint64_t i = 0; i < *flen; i++) {
		contents[i] = fgetc(dcj);
		if (strchr("\"", contents[i])) {
			if (i > 1) {
				if (contents[i-1] != '\\') {
					inString = !inString;
					tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
					tokens[*numTokens].token = contents[i];
					tokens[*numTokens].location = i;
					*numTokens += 1;
				}
			}
		} else if (strchr(" {}[]\n", contents[i]) && !inString) {
			if (i > 1) {
				if (contents[i-1] != '\\') {
					tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
					tokens[*numTokens].token = contents[i];
					tokens[*numTokens].location = i;
					*numTokens += 1;
					break;
				}
			}
		}
	}
	fclose(dcj);
	
	return tokens;
}
