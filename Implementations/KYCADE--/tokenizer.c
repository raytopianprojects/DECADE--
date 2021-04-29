#include "tokenizer.h"

token_t* tokenizeFile(char* file, char* contents, uint64_t* numTokens, uint64_t* flen) {
	FILE* dcj = fopen(file, "r");
	fseek(dcj, 0L, SEEK_END);
	*flen = ftell(dcj);
	fseek(dcj, 0L, SEEK_SET);
	
	contents = malloc(*flen);
	token_t* tokens = malloc(0);
	
	for (uint64_t i = 0; i < *flen; i++) {
		contents[i] = fgetc(dcj);
		if (strchr(" []\n", contents[i])) {
			tokens = realloc(tokens, (*numTokens+1)*sizeof(token_t));
			tokens[*numTokens].token = contents[i];
			tokens[*numTokens].location = i;
			*numTokens += 1;
			break;
		}
	}
	fclose(dcj);
	
	return tokens;
}