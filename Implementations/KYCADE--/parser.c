#include "parser.h"

token_t* tokenizeFile(char* file, char** contentsPTR, uint64_t* numTokens, uint64_t* flen) {
	FILE* dcj = fopen(file, "r");
	if (dcj == NULL) return NULL;
	fseek(dcj, 0L, SEEK_END);
	*flen = ftell(dcj);
	fseek(dcj, 0L, SEEK_SET);
	
	*contentsPTR = malloc(*flen);
	char* contents = *contentsPTR;
	token_t* tokens = malloc(0);
	
	uint8_t inString = 0;
	*numTokens = 0;
	for (uint64_t i = 0; i < *flen; i++) {
		contents[i] = fgetc(dcj);
		if (strchr("\"`", contents[i])) {
			if (i > 1) {
				if (contents[i-1] != '\\') {
					inString = !inString;
					tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
					tokens[*numTokens].token = contents[i];
					tokens[*numTokens].location = i;
					*numTokens += 1;
					printf("STR TOKEN, TOKEN NUMBER %llu\n", *numTokens);
				}
			} else {
				inString = !inString;
				tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
				tokens[*numTokens].token = contents[i];
				tokens[*numTokens].location = i;
				*numTokens += 1;
				printf("STR TOKEN, TOKEN NUMBER %llu\n", *numTokens);
			}
		} else if (strchr(" {}[]\n", contents[i]) && !inString) {
			if (i > 1) {
				if (contents[i-1] != '\\') {
					tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
					tokens[*numTokens].token = contents[i];
					tokens[*numTokens].location = i;
					*numTokens += 1;
					printf("FUNC TOKEN, TOKEN NUMBER %llu\n", *numTokens);
				}
			} else {
				tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
				tokens[*numTokens].token = contents[i];
				tokens[*numTokens].location = i;
				*numTokens += 1;
				printf("FUNC TOKEN, TOKEN NUMBER %llu\n", *numTokens);
			}
		}
	}
	fclose(dcj);
	
	return tokens;
}